# Go Validation Example with `gookit/validate`

A simple implementation using the [`github.com/gookit/validate`](https://github.com/gookit/validate) library in Go. This library provides a more convenient and expressive way to validate JSON, maps, and struct data compared to `go-playground/validator`.

## Why `gookit/validate`?

This validation library offers several advantages:

- ✅ **Built-in error messages** out of the box — no need to define your own formatting logic.
- ✅ **Supports structs and maps** — works even without struct tags.
- ✅ **Rich rule syntax** — e.g. `required|minLen:3|email` directly in tags or dynamically at runtime.
- ✅ **Custom error messages** — define per-rule or per-field custom messages easily.
- ✅ **Global configuration options** — such as `StopOnError` and `SkipOnEmpty`.

## When to Use

This library is ideal for:

- Building REST APIs that return structured and readable error responses.
- Quickly validating input without boilerplate code.
- Use cases where dynamic validation rules or map-based input needs checking.

## Example

```go
type User struct {
	Username    string `validate:"required|minLen:3"`
	Email       string `validate:"required|email|maxLen:20"`
	Gender      string `validate:"in:Male,Female"`
	Description string `validate:"minLen:10"` // optional field
	Age         int    `validate:"required|min:18"`
}
```

## Response
```cmd
{
  "code": 200,
  "status": "OK",
  "error": {
    "Age": [
      "Age min value is 18"
    ],
    "Description": [
      "Description min length is 10"
    ],
    "Email": [
      "Email is required to not be empty"
    ],
    "Gender": [
      "Gender value must be in the enum [Male Female]"
    ],
    "Username": [
      "Username min length is 3"
    ]
  }
}

```
