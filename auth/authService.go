package auth

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthService interface {
	GenerateToken(int) (string, error)
	ValidateToken(string) (string, bool, error)
}

type JwtService struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// var SECRET_KEY = os.Getenv("SECRET_KEY")

func NewAuthService() *JwtService {
	return &JwtService{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
}

func (j *JwtService) GenerateToken(userID int) (string, error) {
	// fmt.Println("secret key generate token", SECRET_KEY)
	claim := NewAuthService()
	claim.UserID = strconv.Itoa(userID)
	fmt.Println("USER ID CUYYY: ", userID)
	fmt.Println("USER ID CUYYY: ", claim.UserID)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	fmt.Println("TOKEN", token)

	theToken, err := token.SignedString([]byte("INIYA"))
	fmt.Println("The TOKEN", theToken)
	if err != nil {
		fmt.Println("err :", err)
		return theToken, err
	}
	return theToken, nil
}

func (j *JwtService) ValidateToken(encodedToken string) (string, bool, error) {
	fmt.Println("TOKEN ENCODED", encodedToken)
	// fmt.Println("secret key validate token", SECRET_KEY)
	theToken, err := jwt.ParseWithClaims(encodedToken, &JwtService{}, func(token *jwt.Token) (interface{}, error) {
		// newToken, ok := token.Method.(*jwt.SigningMethodHMAC)
		// fmt.Println("NEWTOKEN", newToken)
		// fmt.Println("OKE OK OK OK OK OK", ok)
		// fmt.Println("ENCODED TOKEN", encodedToken)
		// if !ok {
		// 	fmt.Println("TOKEN INVALID !!!!")
		// 	return nil, errors.New("INVALID TOKEN CUY")
		// }
		// fmt.Println("TEST GOLANG", []byte(SECRET_KEY))
		// fmt.Println("new token", newToken)
		return []byte("INIYA"), nil
	})

	fmt.Println("tokenvalid", theToken.Valid)
	fmt.Println("TOKEN AUTH: ", theToken)
	fmt.Println("TEST CUYYYYY: ", theToken.Claims)
	if err != nil {
		fmt.Println("OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO", err.Error())
		// fmt.Println("YYYYYYYYYYYYYYYYYYY", theToken.Claims.(*JwtService))
		return "0", false, err
	}

	if claim, ok := theToken.Claims.(*JwtService); ok && theToken.Valid {
		fmt.Println("TEST AUTH: ", claim.UserID)
		// newID := strconv.Itoa(claim.UserID)
		return claim.UserID, true, nil
	} else {
		// claim.UserID := strconv.Itoa(claim.UserID)
		return claim.UserID, false, nil
	}

}
