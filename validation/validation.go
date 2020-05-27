package validation

import (
	"fmt"
	"github.com/fatih/structs"
	"reflect"
	"strconv"
	"strings"
)

var errortmpl = map[string]string{
	"required": "%s is a required field",
	"length":   "%s can not be empty or zero",
	"max":      "%s's maximum required number of characters are %d characters",
	"min":      "%s's minimum required number of characters are %d characters",
	//todo
	"alpha":         "%s must be valid alpha characters",
	"numeric":       "%s must be valid numeric characters",
	"alpha_numeric": "%s must be valid alpha or numeric characters",
	"string":        "%s must be a string",
}

type Validator struct{}

func (v *Validator) Validate(args map[string]string, params interface{}) error {
	for key, value := range args {
		rules := strings.Split(value, "|")
		if len(rules) >= 2 {
			for _, rule := range rules {
				switch {
				case strings.Contains(rule, "required"):
					if err := v.checkRequired(key, params); err != nil {
						return err
					}
					continue
				case strings.Contains(rule, "max"):
					arg := strings.Split(rule, ":")
					length, _ := strconv.Atoi(arg[1])
					if err := v.checkMaxLength(key, length, params); err != nil {
						return err
					}
					continue
				case strings.Contains(rule, "min"):
					arg := strings.Split(rule, ":")
					length, _ := strconv.Atoi(arg[1])
					if err := v.checkMinLength(key, length, params); err != nil {
						return err
					}
					continue
				case strings.Contains(value, "string"):
					if err := v.checkIfString(key, params); err != nil {
						return err
					}
					continue
				default:
					return fmt.Errorf("invalid rule provided")
				}
			}
		} else {
			switch {
			case strings.Contains(value, "required"):
				if err := v.checkRequired(key, params); err != nil {
					return err
				}
				continue
			case strings.Contains(value, "max"):
				arg := strings.Split(value, ":")
				rqlen, _ := strconv.Atoi(arg[1])
				if err := v.checkMaxLength(key, rqlen, params); err != nil {
					return err
				}
				continue
			case strings.Contains(value, "min"):
				arg := strings.Split(value, ":")
				rqlen, _ := strconv.Atoi(arg[1])
				if err := v.checkMinLength(key, rqlen, params); err != nil {
					return err
				}
				continue
			case strings.Contains(value, "string"):
				if err := v.checkIfString(key, params); err != nil {
					return err
				}
				continue
			default:
				return fmt.Errorf("invalid rule provided")
			}
		}
	}
	return nil
}

func (v *Validator) checkRequired(key string, params interface{}) error {
	switch reflect.ValueOf(params).Kind() {
	case reflect.Map:
		for _, v := range params.(map[string]interface{}) {
			if _, ok := params.(map[string]interface{})[key]; ok {
				switch v.(type) {
				case int:
					if v.(int) <= 0 {
						return fmt.Errorf(errortmpl["length"], key)
					}
				case string:
					if len(v.(string)) <= 0 {
						return fmt.Errorf(errortmpl["required"], key)
					}
				}
			} else {
				return fmt.Errorf(errortmpl["required"], key)
			}
		}
	case reflect.Struct:
		params = structs.Map(params)
		for _, v := range params.(map[string]interface{}) {
			if _, ok := params.(map[string]interface{})[strings.Title(key)]; ok {
				switch v.(type) {
				case int:
					if v.(int) <= 0 {
						return fmt.Errorf(errortmpl["length"], key)
					}
				case string:
					if len(v.(string)) <= 0 {
						return fmt.Errorf(errortmpl["required"], key)
					}
				}
			} else {
				return fmt.Errorf(errortmpl["required"], key)
			}
		}
	default:
		return fmt.Errorf("invalid interface provided: %v", reflect.TypeOf(params))
	}
	return nil
}

func (v *Validator) checkMaxLength(key string, rqlen int, params interface{}) error {
	switch reflect.ValueOf(params).Kind() {
	case reflect.Map:
		for _, v := range params.(map[string]interface{}) {
			if _, ok := params.(map[string]interface{})[key]; ok {
				switch v.(type) {
				case string:
					if len(v.(string)) > rqlen {
						return fmt.Errorf(errortmpl["max"], key, rqlen)
					}
				default:
					return fmt.Errorf(errortmpl["string"], key)
				}
				return nil
			}
		}
	case reflect.Struct:
		params = structs.Map(params)
		for _, v := range params.(map[string]interface{}) {
			if _, ok := params.(map[string]interface{})[strings.Title(key)]; ok {
				switch v.(type) {
				case string:
					if len(v.(string)) > rqlen {
						return fmt.Errorf(errortmpl["max"], key, rqlen)
					}
				default:
					return fmt.Errorf(errortmpl["string"], key)
				}
				return nil
			}
		}
	default:
		return fmt.Errorf("invalid interface provided: %v", reflect.TypeOf(params))
	}
	return nil
}

func (v *Validator) checkMinLength(key string, rqlen int, params interface{}) error {
	switch reflect.ValueOf(params).Kind() {
	case reflect.Map:
		for _, v := range params.(map[string]interface{}) {
			if _, ok := params.(map[string]interface{})[key]; ok {
				switch v.(type) {
				case string:
					if len(v.(string)) <= rqlen {
						return fmt.Errorf(errortmpl["min"], key, rqlen)
					}
				default:
					return fmt.Errorf(errortmpl["string"], key)
				}
				return nil
			}
		}
		return nil
	case reflect.Struct:
		params = structs.Map(params)
		for _, v := range params.(map[string]interface{}) {
			if _, ok := params.(map[string]interface{})[strings.Title(key)]; ok {
				switch v.(type) {
				case string:
					if len(v.(string)) <= rqlen {
						return fmt.Errorf(errortmpl["min"], key, rqlen)
					}
				default:
					return fmt.Errorf(errortmpl["string"], key)
				}
				return nil
			}
		}
	default:
		return fmt.Errorf("invalid interface provided: %v", reflect.TypeOf(params))
	}
	return nil
}

func (v *Validator) checkIfString(key string, params interface{}) error {
	switch reflect.ValueOf(params).Kind() {
	case reflect.Map:
		for _, v := range params.(map[string]interface{}) {
			if _, ok := params.(map[string]interface{})[strings.Title(key)]; ok {
				switch v.(type) {
				case string:
					return nil
				default:
					return fmt.Errorf(errortmpl["string"], key)
				}
				return nil
			}
		}
	case reflect.Struct:
		params = structs.Map(params)
		for _, v := range params.(map[string]interface{}) {
			if _, ok := params.(map[string]interface{})[key]; ok {
				switch v.(type) {
				case string:
					return nil
				default:
					return fmt.Errorf(errortmpl["string"], key)
				}
				return nil
			}
		}
	}
	return nil
}
