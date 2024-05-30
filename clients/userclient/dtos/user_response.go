package dtos

type UserResponse struct {
	Id               int32  `json:"id"`
	Name             string `json:"name"`
	Age              int32  `json:"age"`
	Email            string `json:"email"`
	CurrentAddress   string `json:"current_address"`
	PermanentAddress string `json:"permanent_address"`
}
