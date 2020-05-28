package validation

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestValidation_ValidParameters(t *testing.T) {
	params := map[string]interface{}{
		"username": "veego",
		"email":    "veego@email.com",
		"password": "supasecret",
	}
	validator := Validator{}
	if err := validator.Validate(map[string]string{
		"username": "required|max:20|min:4",
		"email":    "required|string",
		"password": "max:40",
	}, params); err != nil {
		t.Errorf("didnt expect any errors but we got %v", err.Error())
	}
}

func TestValidation_ValidStructParameters(t *testing.T) {

	type params struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Msisdn string `json:"msisdn"`
	}
	customer := params {
		FirstName:        "Matt",
		LastName:         "Sebuuma",
		Msisdn:           "256777000000",
	}
	var p params
	bt := []byte(fmt.Sprintf(`{"firstName": "%s", "lastName": "%s", "msisdn": "%s"}`, customer.FirstName, customer.LastName, customer.Msisdn))
	err := json.Unmarshal(bt, &p)
	if err != nil {
		t.Errorf("didnt expect any errors but we got %v", err.Error())
	}
	validator := Validator{}
	if err := validator.Validate(map[string]string{
		"firstName": "required",
		"lastName": "required",
		"msisdn": "required|max:12|min:12",
	}, p); err != nil {
		t.Errorf("didnt expect any errors but we got %v", err.Error())
	}
}

func TestValidation_InvalidParameters(t *testing.T) {
	params := map[string]interface{}{
		"email": "veego@email.com",
	}
	validator := Validator{}
	if err := validator.Validate(map[string]string{
		"username": "required",
		"emai_l":   "required",
	}, params); err == nil {
		t.Errorf("expected errors but we didnt get any")
	}
}

func TestValidation_EmptyStringParameters(t *testing.T) {
	params := map[string]interface{}{
		"email": "",
	}
	validator := Validator{}
	if err := validator.Validate(map[string]string{
		"email": "required",
		"phone": "required",
	}, params); err == nil {
		t.Errorf("expected errors but we didnt get any")
	}
}

func TestValidation_EmptyIntParameters(t *testing.T) {
	params := map[string]interface{}{
		"phone": 0,
	}
	validator := Validator{}
	if err := validator.Validate(map[string]string{
		"phone": "required",
	}, params); err == nil {
		t.Errorf("expected errors but we didnt get any: %v", err.Error())
	}
}

func TestValidation_InvalidRequiredParameters(t *testing.T) {
	params := map[string]interface{}{
		"email": "veego@email.com",
	}
	validator := Validator{}
	if err := validator.Validate(map[string]string{
		"emai_l": "required|max:100",
	}, params); err == nil {
		t.Errorf("expected errors but we didnt get any")
	}
}

func TestValidation_InvalidNestedMaxParameters(t *testing.T) {
	params := map[string]interface{}{
		"email": "veego@email.com",
	}
	validator := Validator{}
	if err := validator.Validate(map[string]string{
		"email": "required|max:3",
	}, params); err == nil {
		t.Errorf("expected errors but we didnt get any")
	}
}

func TestValidation_InvalidNestedMinParameters(t *testing.T) {
	params := map[string]interface{}{
		"email": "ve",
	}
	validator := Validator{}
	if err := validator.Validate(map[string]string{
		"email": "required|min:3",
	}, params); err == nil {
		t.Errorf("expected errors but we didnt get any")
	}
}

func TestValidation_InvalidRule(t *testing.T) {
	params := map[string]interface{}{
		"email": "veego@email.com",
	}
	validator := Validator{}
	if err := validator.Validate(map[string]string{
		"email": "unique",
	}, params); err == nil {
		t.Errorf("expected errors but we didnt get any")
	}
}

func TestValidation_InvalidNestedRule(t *testing.T) {
	params := map[string]interface{}{
		"email": "veego@email.com",
	}
	validator := Validator{}
	if err := validator.Validate(map[string]string{
		"email": "required|unique",
	}, params); err == nil {
		t.Errorf("expected errors but we didnt get any")
	}
}

func TestValidation_InvalidMaxRule(t *testing.T) {
	params := map[string]interface{}{
		"email": "veego1@email.com",
	}
	validator := Validator{}
	if err := validator.Validate(map[string]string{
		"email": "max:3",
	}, params); err == nil {
		t.Errorf("expected errors but we didnt get any")
	}
}

func TestValidation_InvalidMinRule(t *testing.T) {
	params := map[string]interface{}{
		"email": "ve",
	}
	validator := Validator{}
	if err := validator.Validate(map[string]string{
		"email": "min:3",
	}, params); err == nil {
		t.Errorf("expected errors but we didnt get any")
	}
}

func TestValidation_InvalidMaxParameterType(t *testing.T) {
	params := map[string]interface{}{
		"email": 0,
	}
	validator := Validator{}
	if err := validator.Validate(map[string]string{
		"email": "max:10",
	}, params); err == nil {
		t.Errorf("expected errors but we didnt get any")
	}
}

func TestValidation_InvalidMinParameterType(t *testing.T) {
	params := map[string]interface{}{
		"email": 0,
	}
	validator := Validator{}
	if err := validator.Validate(map[string]string{
		"email": "min:10",
	}, params); err == nil {
		t.Errorf("expected errors but we didnt get any")
	}
}
