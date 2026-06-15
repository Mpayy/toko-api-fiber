package model

type WebResponse[T any] struct {
	Data   T      `json:"data,omitempty"`
	Errors T      `json:"errors,omitempty"`
}