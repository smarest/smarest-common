package exception

const (
	CodeNotFound         = 1001
	CodeValueInvalid     = 1002
	CodeSignatureInvalid = 1003
	CodeUnknown          = 1004
	CodeSystemError      = 9000
)

var statusText = map[int]string{
	CodeNotFound:         "Not Found",
	CodeValueInvalid:     "Invalid Value",
	CodeSignatureInvalid: "Invalid Signature",
	CodeSystemError:      "System Error.",
}

type Error struct {
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	RootCause    string `json:"-"`
}

func GetError(code int) *Error {
	return &Error{ErrorCode: code, ErrorMessage: statusText[code]}
}

func CreateError(code int, message string) *Error {
	return &Error{ErrorCode: code, ErrorMessage: message}
}

func CreateErrorWithRootCause(code int, message string, err error) *Error {
	return &Error{ErrorCode: code, ErrorMessage: message, RootCause: err.Error()}
}

func GetErrorMessage(code int) string {
	return statusText[code]
}
