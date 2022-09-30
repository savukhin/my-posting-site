package dto

type DefaultResponse struct {
	Has_err bool   `json:"has_error"`
	Msg     string `json:"msg"`
}

type DefaultJWTResponse struct {
	DefaultResponse
	Token string `json:"token"`
}
