package util

import (
	"agent-assigner/internal/dto"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

// utility function for get env data
func GetEnv(key string, fallback string) string {
	a, _ := godotenv.Read()
	var (
		val     string
		isExist bool
	)
	val, isExist = a[key]
	if !isExist {
		val = fallback
	}
	return val
}

// help for convert response body to string
func ResponseBodyToString(resp *http.Response) string {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(body)
}

// utility function for api response
func APIResponse(status string, code int, message string, data interface{}) dto.Response {
	meta := dto.Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := dto.Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

// utility function for json response
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(data)
}
