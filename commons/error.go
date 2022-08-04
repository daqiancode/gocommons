package commons

type Result struct {
	State       int               `json:"state"`
	Data        interface{}       `json:"data,omitempty"`
	Message     string            `json:"message,omitempty"`
	FieldErrors map[string]string `json:"fieldErrors,omitempty"`
	App         string            `json:"app,omitempty"`
}

// Error in service layer
type ServiceError struct {
	Message string
}

func (s *ServiceError) Error() string {
	return s.Message
}

type Page struct {
	PageIndex int         `json:"pageIndex"`
	PageSize  int         `json:"pageSize"`
	Total     int         `json:"total"`
	Items     interface{} `json:"items"`
}
