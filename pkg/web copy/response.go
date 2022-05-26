package web

type Response struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func NewResponse(codeStatus int, message interface{}, err string) Response {
	if codeStatus < 300 {
		return Response{codeStatus, message, ""}
	} else {
		return Response{codeStatus, nil, err}
	}
}
