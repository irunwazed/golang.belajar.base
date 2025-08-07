package model

// keluaran request respon
type Response[T any] struct {
	Data T `json:"data"`
}

type ResponseMessage struct {
	Message string
	Errors  string
}
