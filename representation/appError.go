package representation

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (s *AppError) Error() string {
	return s.Message
}
