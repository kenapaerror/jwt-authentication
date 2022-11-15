package web

type UsersUpdateRequest struct {
	Id        string `validate:"required" json:"id"`
	FirstName string `validate:"required,min=1,max=25" json:"first_name"`
	LastName  string `validate:"required,min=1,max=50" json:"last_name"`
}
