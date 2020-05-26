package veego

import "github.com/mattb2401/veego/validation"

type Controller struct {
	validation.Validator
}

func NewController() *Controller {
	return &Controller{}
}