package model

type Erro struct {
	Error error `json:"error"`
	Mensagem string `json:"mensagem"`
	Code int `json:"code"`
}
