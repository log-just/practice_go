package util

// ErrCode : 에러코드으~~
// type ErrCode string

// ErrResponse : 에러
type ErrResponse struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

const (
	// ErrCodeParam : 에러
	ErrCodeParam = "PARAM"
)
