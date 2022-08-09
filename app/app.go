package app

import (
	"fmt"
	"mpbe/auth"
	"mpbe/domain"
	"mpbe/errs"
	"mpbe/logger"
	"mpbe/service"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func sanityCheck() {
	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_USER",
		"DB_PASSWD",
		"DB_ADDR",
		"DB_PORT",
		"DB_NAME",
	}

	for _, envKey := range envProps {
		if os.Getenv(envKey) == "" {
			logger.Fatal(fmt.Sprintf("environment variable %s not defined. terminating application...", envKey))
		}
	}

	logger.Info("environment variables loaded...")

}

func Start() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error loading .env file")
	}

	sanityCheck()

	dbClient := getClientDB()

	// wiring
	// * setup repository
	employeeRepositoryDB := domain.NewEmployeeRepositoryDB(dbClient)
	userRepositoryDB := domain.NewUserRepositoryDB(dbClient)
	// * setup service
	employeeService := service.NewEmployeeService(&employeeRepositoryDB)
	userService := service.NewUserService(&userRepositoryDB)
	authService := auth.NewAuthService()
	// * setup handler
	che := EmployeeHandlers{employeeService}
	chu := UserHandlers{userService, authService}

	router := gin.Default()

	router.GET("/employees", authMiddleware(authService, userService), che.getAllEmployee)
	router.GET("/employees/:id", che.getEmployeeID)
	router.POST("/employees", che.createEmployee)
	router.DELETE("/employees/:id", che.deleteEmployee)
	router.PUT("/employees/:id", che.updateEmployee)
	router.POST("/users", chu.createUser)
	router.POST("/login", chu.loginUser)

	router.Run(":9090")

}

func getClientDB() *gorm.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWD")
	dbAddress := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbAddress, dbPort, dbName)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info("success connect to database...")

	return db
}

func authMiddleware(authService auth.AuthService, userService service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			appErr := errs.NewBadRequestError("invalid token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, appErr)
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer", "", -1)
		fmt.Println("TOken string", tokenString)

		userId, ok, err := authService.ValidateToken(tokenString)
		fmt.Println("userid", userId, "ok", ok, "Err", err)
		if err != nil && userId == "" && !ok {
			appErr := errs.NewBadRequestError("invalid token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, appErr)
		} else {
			idUser, _ := strconv.Atoi(userId)
			fmt.Println("ID USER", idUser)
			fmt.Println("ID USER KRESNA", userId)
			user, err := userService.UserByID(idUser)
			if err != nil {
				appErr := errs.NewBadRequestError("invalid token")
				c.AbortWithStatusJSON(http.StatusUnauthorized, appErr)
			}
			c.Set("currentUser", user)
		}

	}
}