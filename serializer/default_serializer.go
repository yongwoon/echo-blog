package serializer

type (
	successEntry struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}

	SuccessSerializer struct {
		Result *successEntry `json:"result"`
	}
)

// NewSuccessResponse return success message
func NewSuccessResponse() *SuccessSerializer {
	s := new(successEntry)
	s.Code = "200"
	s.Message = "success"

	return &SuccessSerializer{s}
}
