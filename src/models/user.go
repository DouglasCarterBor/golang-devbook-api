package models

import "time"

type User struct {
	ID uint64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Nick string `json:"nick,omitempty"`
	Email string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	CreatedIn time.Time `json:"created_in"`
}