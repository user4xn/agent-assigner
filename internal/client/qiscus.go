package client

import (
	"agent-assigner/internal/dto"
	"agent-assigner/pkg/helper"
	"agent-assigner/pkg/util"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type QiscusClientInterface interface {
	CountUnservedChat(adminToken string) (dto.ResponseAPICountUnserved, error)
	FetchUnservedRoom(body dto.BodyAPIChatRoom) (dto.ResponseAPIChatRoom, error)
}

type QiscusClient struct {
	Http      *http.Client
	BaseUrl   string
	AppIDCode string
	AppSecret string
}

func NewQiscusClient() *QiscusClient {
	return &QiscusClient{
		Http:      &http.Client{},
		BaseUrl:   util.GetEnv("QISCUS_BASE_URL", "https://omnichannel.qiscus.com"),
		AppIDCode: util.GetEnv("QISCUS_APP_ID_CODE", "fallback"),
		AppSecret: util.GetEnv("QISCUS_APP_SECRET", "fallback"),
	}
}

func (c *QiscusClient) CountUnservedChat(adminToken string) (dto.ResponseAPICountUnserved, error) {
	source := c.BaseUrl + "/api/v2/admin/service/total_unserved"
	headersMap := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": adminToken,
		"Qiscus-app-id": c.AppIDCode,
	}

	res, err := helper.GetRequest(c.Http, source, nil, headersMap)
	if err != nil {
		return dto.ResponseAPICountUnserved{}, err
	}
	defer helper.ClientClose(res)

	if res.StatusCode != 200 {
		if res.StatusCode == 401 {
			return dto.ResponseAPICountUnserved{}, fmt.Errorf("unauthorized")
		}

		err := fmt.Errorf("error count unserved chat status code: %d, body: %s", res.StatusCode, util.ResponseBodyToString(res))
		log.Print(err)
		return dto.ResponseAPICountUnserved{}, err
	}

	var commonResponse dto.CommonResponse
	if err := json.NewDecoder(res.Body).Decode(&commonResponse); err != nil {
		log.Println(err)
		return dto.ResponseAPICountUnserved{}, err
	}

	byteData, err := json.Marshal(commonResponse.Data)
	if err != nil {
		log.Println(err)
		return dto.ResponseAPICountUnserved{}, err
	}

	var response dto.ResponseAPICountUnserved
	if err := json.Unmarshal(byteData, &response); err != nil {
		log.Println(err)
		return dto.ResponseAPICountUnserved{}, err
	}

	return response, nil
}

func (c *QiscusClient) FetchUnservedRoom(body dto.BodyAPIChatRoom) (dto.ResponseAPIChatRoom, error) {
	source := c.BaseUrl + "/api/v2/customer_rooms"
	headersMap := map[string]string{
		"Content-Type":      "application/json",
		"Qiscus-App-Id":     c.AppIDCode,
		"Qiscus-Secret-Key": c.AppSecret,
	}

	reqJson, err := json.Marshal(body)
	if err != nil {
		return dto.ResponseAPIChatRoom{}, err
	}

	reqBody := strings.NewReader(string(reqJson))

	res, err := helper.PostRequest(c.Http, source, reqBody, headersMap)
	if err != nil {
		return dto.ResponseAPIChatRoom{}, err
	}
	defer helper.ClientClose(res)

	if res.StatusCode != 200 {
		if res.StatusCode == 401 {
			return dto.ResponseAPIChatRoom{}, fmt.Errorf("unauthorized while fetching unserved room")
		}

		err := fmt.Errorf("error unserved chat room status code: %d, body: %s", res.StatusCode, util.ResponseBodyToString(res))
		log.Print(err)
		return dto.ResponseAPIChatRoom{}, err
	}

	var commonResponse dto.CommonResponse
	if err := json.NewDecoder(res.Body).Decode(&commonResponse); err != nil {
		log.Println(err)
		return dto.ResponseAPIChatRoom{}, err
	}

	byteData, err := json.Marshal(commonResponse.Data)
	if err != nil {
		log.Println(err)
		return dto.ResponseAPIChatRoom{}, err
	}

	var response dto.ResponseAPIChatRoom
	if err := json.Unmarshal(byteData, &response); err != nil {
		log.Println(err)
		return dto.ResponseAPIChatRoom{}, err
	}

	return response, nil
}
