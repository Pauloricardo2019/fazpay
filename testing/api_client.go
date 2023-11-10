package testing

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"kickoff/dto"
	"net/http"
	"net/http/httptest"
)

type APIResponse struct {
	StatusCode int
	Data       []byte
}
type ApiClient struct {
	bearerToken string
}

func NewApiClient() *ApiClient {
	return &ApiClient{}
}

func (a *ApiClient) Login(request *dto.LoginRequest) (*dto.LoginResponse, *dto.Error) {
	loginResponse := &dto.LoginResponse{}
	errorResponse := &dto.Error{}

	apiResponse, err := a.doRequest("POST", "/v1/auth/login", request, nil)
	if err != nil {
		errorResponse.Message = err.Error()
		return nil, errorResponse
	}

	err = json.Unmarshal(apiResponse.Data, loginResponse)
	if err != nil {
		errorResponse.Message = err.Error()
		return nil, errorResponse
	}

	return loginResponse, nil
}

func (a *ApiClient) SetBearerToken(value string) {
	a.bearerToken = value
}

func (a *ApiClient) CreateUser(createUserRequest *dto.CreateUserRequest) (*dto.CreateResponse, *dto.Error) {
	createResponse := &dto.CreateResponse{}
	errorResponse := &dto.Error{}

	apiResponse, err := a.doRequest("POST", "/v1/user/", createUserRequest, nil)
	if err != nil {
		errorResponse.Message = err.Error()
		return nil, errorResponse
	}

	err = json.Unmarshal(apiResponse.Data, createResponse)
	if err != nil {
		errorResponse.Message = err.Error()
		return nil, errorResponse
	}

	return createResponse, nil
}

func (a *ApiClient) GetUserByID(id uint64) (*dto.GetUserResponse, *dto.Error) {
	getUserResponse := &dto.GetUserResponse{}
	errorResponse := &dto.Error{}

	apiResponse, err := a.doRequest("GET", fmt.Sprintf("/v1/user/%d", id), nil, nil)
	if err != nil {
		errorResponse.Message = err.Error()
		return nil, errorResponse
	}

	err = json.Unmarshal(apiResponse.Data, getUserResponse)
	if err != nil {
		errorResponse.Message = err.Error()
		return nil, errorResponse
	}

	return getUserResponse, nil
}

func (a *ApiClient) UpdateUser(updateUserRequest *dto.UpdateUserRequest, id uint64) *dto.Error {
	errorResponse := &dto.Error{}

	_, err := a.doRequest("PUT", fmt.Sprintf("/v1/user/%d", id), updateUserRequest, nil)
	if err != nil {
		errorResponse.Message = err.Error()
		return errorResponse
	}

	return nil
}

func (a *ApiClient) DeleteUser(id uint64) *dto.Error {
	errorResponse := &dto.Error{}

	_, err := a.doRequest("DELETE", fmt.Sprintf("/v1/user/%d", id), nil, nil)
	if err != nil {
		errorResponse.Message = err.Error()
		return errorResponse
	}

	return nil
}

func (a *ApiClient) CreateClient(client *dto.CreateClientRequest) (*dto.CreateResponse, *dto.Error) {
	createResponse := &dto.CreateResponse{}
	errorResponse := &dto.Error{}

	apiResponse, err := a.doRequest("POST", "/v1/client/", client, nil)
	if err != nil {
		errorResponse.Message = err.Error()
		return nil, errorResponse
	}

	err = json.Unmarshal(apiResponse.Data, createResponse)
	if err != nil {
		errorResponse.Message = err.Error()
		return nil, errorResponse
	}

	return createResponse, nil
}

func (a *ApiClient) GetClientByID(id uint64) (*dto.GetClientResponse, *dto.Error) {
	getClientResponse := &dto.GetClientResponse{}
	errorResponse := &dto.Error{}

	apiResponse, err := a.doRequest("GET", fmt.Sprintf("/v1/client/%d", id), nil, nil)
	if err != nil {
		errorResponse.Message = err.Error()
		return nil, errorResponse
	}

	err = json.Unmarshal(apiResponse.Data, getClientResponse)
	if err != nil {
		errorResponse.Message = err.Error()
		return nil, errorResponse
	}

	return getClientResponse, nil
}

func (a *ApiClient) GetClientList() (*dto.GetAllClientsResponse, *dto.Error) {
	getListClientResponse := &dto.GetAllClientsResponse{}
	errorResponse := &dto.Error{}

	queries := map[string]string{
		"limit":  "10",
		"offset": "0",
	}

	apiResponse, err := a.doRequest("GET", "/v1/client/", nil, &queries)
	if err != nil {
		errorResponse.Message = err.Error()
		return nil, errorResponse
	}

	err = json.Unmarshal(apiResponse.Data, getListClientResponse)
	if err != nil {
		errorResponse.Message = err.Error()
		return nil, errorResponse
	}

	return getListClientResponse, nil
}

func (a *ApiClient) GetByIDClient(id uint64) (*dto.GetClientResponse, *dto.Error) {
	getClientResponse := &dto.GetClientResponse{}
	errorResponse := &dto.Error{}

	apiResponse, err := a.doRequest("GET", fmt.Sprintf("/v1/client/%d", id), nil, nil)
	if err != nil {
		errorResponse.Message = err.Error()
		return nil, errorResponse
	}

	err = json.Unmarshal(apiResponse.Data, getClientResponse)
	if err != nil {
		errorResponse.Message = err.Error()
		return nil, errorResponse
	}

	return getClientResponse, nil
}

func (a *ApiClient) DeleteClient(id uint64) *dto.Error {
	errorResponse := &dto.Error{}

	_, err := a.doRequest("DELETE", fmt.Sprintf("/v1/client/%d", id), nil, nil)
	if err != nil {
		errorResponse.Message = err.Error()
		return errorResponse
	}

	return nil
}

func (a *ApiClient) doRequest(method string, url string, objRequest any, queries *map[string]string) (*APIResponse, error) {
	server, err := GetServer()
	if err != nil {
		return nil, err
	}

	jsonValue, _ := json.Marshal(objRequest)

	req, _ := http.NewRequest(method, url, bytes.NewBuffer(jsonValue))
	if a.bearerToken != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.bearerToken))
	}

	if queries != nil {
		q := req.URL.Query()
		for key, value := range *queries {
			q.Add(key, value)
		}
	}

	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	apiResponse := &APIResponse{
		StatusCode: w.Code,
		Data:       w.Body.Bytes(),
	}

	if apiResponse.StatusCode < 200 || apiResponse.StatusCode > 299 {
		errorMessage := &dto.Error{}
		err := json.Unmarshal(apiResponse.Data, errorMessage)
		if err != nil {
			return nil, err
		}

		return nil, errors.New(errorMessage.Message)

	}

	return apiResponse, nil
}
