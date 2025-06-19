package dto

type (
	// common response struct
	Response struct {
		Meta Meta `json:"meta"`
		Data any  `json:"data"`
	}

	Meta struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
		Status  string `json:"status"`
	}

	// server info struct
	ServerInfo struct {
		ServiceName string `json:"service_name"`
		Version     string `json:"version"`
	}
)
