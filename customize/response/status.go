package response

type Status struct {
	Code    int32 `json:"code"`
	Message string `json:"message"`
}

func NewResponseStatus(code int32, message string) *Status {
	var status = Status{}
	status.Code = code
	status.Message = message

	return &status
}

var (
	SUCCESS      = NewResponseStatus(0, "success") //success
	SYSTEM_ERROR = NewResponseStatus(100, "system error")
)
