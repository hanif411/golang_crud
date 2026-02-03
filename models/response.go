package models

type BaseResponse struct {
	Status  int
	Message string
	Data    interface{}
}
