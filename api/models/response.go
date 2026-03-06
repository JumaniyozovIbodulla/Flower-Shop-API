package models

type Response struct {
	StatusCode  int
	Description string
	Data        interface{}
}

type CreateResponse struct {
	ID string `json:"id"`
}

type GetRequest struct {
	ID string `json:"id"`
}

type DeleteRequest struct {
	ID string `json:"id"`
}

type InternalServerError struct {
	Code    string
	Message string
}

type BadRequestError struct {
	Code    string
	Message string
}

type SeccessRequest struct {
	Code    string
	Message string
}

type ValidationError struct {
	Code        string
	Message     string
	UserMessage string
}

