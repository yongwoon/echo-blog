package serializer

type (
	successEntry struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}

	successSerializer struct {
		Result *successEntry `json:"result"`
	}
)

// NewSuccessResponse return success message
func NewSuccessResponse() *successSerializer {
	s := new(successEntry)
	s.Code = "200"
	s.Message = "success"

	return &successSerializer{s}
}
