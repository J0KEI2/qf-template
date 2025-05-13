package repository

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
)


func (repo hrRepository) GetHrToken() (*string, error) {
	url := viper.GetString("hr.url.base") + viper.GetString("hr.url.path.sign_in")
	requestParam := dto.GetRequestSignInParam{
		Username: viper.GetString("hr.auth.username"),
		Password: viper.GetString("hr.auth.password"),
	}

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()
	jsonRequestParam, _ := json.Marshal(requestParam)
	req.SetBody(jsonRequestParam)
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")
	req.SetRequestURI(url)
	res := fasthttp.AcquireResponse()

	if err := fasthttp.Do(req, res); err != nil {

		return nil, err
	}

	responseBody := res.Body()
	statusCode := res.StatusCode()
	errorModel := dto.ErrNotFoundModel{}
	if statusCode == 401 {
		err := json.Unmarshal(responseBody, &errorModel)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf(errorModel.Data.Message)
	}

	if statusCode == 404 {
		err := json.Unmarshal(responseBody, &errorModel)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf(errorModel.Data.Message)
	}
	result := dto.ReponseSignInModel{}
	err := json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	return &result.Data.AccessToken, nil

}
