package main

import (
	// "fmt"
	// "errors"
	"time"
)

type Example struct {
	Name  string
	Email string
	Phone string
	time  time.Time
}
type Examples []Example

func (exp *Examples) add(name string, email string, phone string) {
	example := Example{
		Name:  name,
		Email: email,
		Phone: phone,
		time:  time.Now(),
	}
	*exp = append(*exp, example)
}
