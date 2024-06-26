package forms

// here u put the validation rules for any form that is not a modelform
type LoginForm struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6"`
}
