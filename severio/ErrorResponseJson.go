package severio

type ErrorResponseJson struct {
	Message string `json:"message"`
	Code    uint16 `json:"code"`
}
