package utils

type ResultArray struct {
	Status string	`json:"status"`
	Data []map[string]interface{} `json:"data"`
}
type ResultObject struct {
	Status string	`json:"status"`
	Data any `json:"data"`
}