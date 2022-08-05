package input

type EmployeeInput struct {
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	Birthday   string `json:"birthday"`
	Department string `json:"department"`
	Status     int    `json:"status"`
}
