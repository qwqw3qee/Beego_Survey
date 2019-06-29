package severio

type SuccessResponseJson struct {
	Message string      `json:"message"`
	Code    uint16      `json:"code"`
	Data    interface{} `json:"data"`
}
