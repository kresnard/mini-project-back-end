package dto

type EmployeeResponse struct {
	ID         string `json:"id" db:"employee_id"`
	Name       string `json:"name" db:"name"`
	Gender     string `json:"gender" db:"gender"`
	Birthday   string `json:"birthday" db:"birthday"`
	Department string `json:"department" db:"department"`
	Status     string `json:"status" db:"status"`
}
