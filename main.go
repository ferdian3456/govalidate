package main

import (
	"encoding/json"
	"fmt"
	"github.com/gookit/validate"
	"net/http"
)

type User struct {
	Username    string `validate:"required|minLen:3" json:"username"` // if you wanted to print in lowercase
	Email       string `validate:"required|email|maxLen:20"`
	Gender      string `validate:"in:Male,Female"`
	Description string `validate:"minLen:10"`
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
		Description: "",
		Age:         17,
	}

	// before set the config
	testValidator(&user)

	validate.Config(func(opt *validate.GlobalOption) {
		opt.StopOnError = false
		//opt.SkipOnEmpty = false // default is true
		fmt.Println(opt.SkipOnEmpty)
		// SkipOnEmpty vs Omitempty
		// SkipOnEmpty = global
		// Omitempty = directly to struct field
	})

	// after set the config
	testValidator(&user)

	v := validate.Struct(user)

	ValidateAndPrint(v)
}

func testValidator(user *User) {
	v := validate.Struct(user)

	ValidateAndPrint(v)
}

func ValidateAndPrint(v *validate.Validation) {
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
