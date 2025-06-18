package dto

type (
	Response struct {
		Meta Meta `json:"meta"`
		Data any  `json:"data"`
	}

	Meta struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
		Status  string `json:"status"`
	}

	ServerInfo struct {
		ServiceName string `json:"service_name"`
		Version     string `json:"version"`
	}
)
