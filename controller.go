package veego

type Controller struct {
}

type BaseController interface {
	Validate(args map[string]string, params map[string]interface{}) error
}
