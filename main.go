package main

import (
	"encoding/json"
	"fmt"
	"github.com/gookit/validate"
	"net/http"
)

type User struct {
	Username    string `validate:"required|minLen:3"`
	Email       string `validate:"required|email|maxLen:20"`
	Gender      string `validate:"in:Male,Female"`
	Description string `validate:"minLen:10"` // optional field
	Age         int    `validate:"required|min:18"`
}

type Response struct {
	Code   int                 `json:"code"`
	Status string              `json:"status"`
	Error  map[string][]string `json:"error,omitempty"`
}

func main() {
	user := User{
		Username:    "a",
		Email:       "",
		Gender:      "Mal",
		Description: "123455",
		Age:         17,
	}

	validate.Config(func(opt *validate.GlobalOption) {
		opt.StopOnError = false
		opt.SkipOnEmpty = true
	})

	// opt.SkipOnEmpty is useful when field is optional
	// when the field is filled, then it will do the validation like min len

	v := validate.Struct(user)

	errorMap := make(map[string][]string)

	if v.Validate() == false {
		for field, rules := range v.Errors {
			for _, msg := range rules {
				errorMap[field] = append(errorMap[field], msg)
			}
		}
	}

	response := Response{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Error:  errorMap,
	}

	jsonBytes, _ := json.MarshalIndent(response, "", "  ")
	fmt.Println(string(jsonBytes))
}
