package util

// ErrCode : 에러코드으~~
// type ErrCode string

// ErrResponse : 에러요청 형식은 고정
type ErrResponse struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

// ErrParamBodyWithVerbose : 요
func ErrParamBodyWithVerbose(err error) *ErrResponse {
	return &ErrResponse{
		Code:   ErrCodeParam,
		Detail: err.Error(),
	}
}

// ErrParamBody : 에러
func ErrParamBody(err error) *ErrResponse {
	return &ErrResponse{
		Code: ErrCodeParam,
	}
}

const (
	ErrCodeParam = "PARAM"	// 파라미터 에러
)
