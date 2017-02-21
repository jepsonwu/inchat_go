package response

type Status struct {
	Code    int32 `json:"code"`
	Message string `json:"message"`
}

func NewResponseStatus(name string, code int32, message string) *Status {
	var status = Status{}
	status.Code = code
	status.Message = message

	StatusMap[name] = &status
	return &status
}

func GetStatusByMap(name string) *Status {
	return StatusMap[name]
}

var (
	StatusMap = make(map[string]*Status)

	SUCCESS = NewResponseStatus("success", 0, "success") //success
	SYSTEM_ERROR = NewResponseStatus("system_error", 100, "system error")

	LAST_ID_INVALID = NewResponseStatus("last_id_invalid", 101, "last_id invalid")
	PAGE_INVALID = NewResponseStatus("page_invalid", 101, "page invalid")
	PER_PAGE_INVALID = NewResponseStatus("per_page_invalid", 101, "per_page invalid")

	PEEL_ID_INVALID = NewResponseStatus("peel_id_invalid", 101, "peel_id invalid")
)
