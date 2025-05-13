package repository

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/url"

	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models/entity"
)

func (r *authRepository) GetThaidDataFromDopa(code string) (ThaidDataResponse entity.ThaidTokenPostResponseEntity, responseError *helpers.ResponseError) {
	body := url.Values{}
	body.Set("code", code)
	body.Set("grant_type", "authorization_code")
	body.Set("redirect_uri", viper.GetString("thaid.callback"))
	encodedData := body.Encode()

	username := viper.GetString("thaid.client_id")
	password := viper.GetString("thaid.client_secret")
	credential := username + ":" + password
	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(credential))

	postArgs := &fasthttp.Args{}
	postArgs.Add("Content-Type", "")
	postArgs.Add("Authorization", auth)

	httpRequest := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(httpRequest)
	httpRequest.SetRequestURI("https://imauth.bora.dopa.go.th/api/v2/oauth2/token/")
	httpRequest.Header.SetMethod("POST")
	httpRequest.Header.SetContentType("application/x-www-form-urlencoded")
	httpRequest.Header.Set("Authorization", auth)
	httpRequest.SetBody([]byte(encodedData))

	httpResponse := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(httpResponse)

	err := fasthttp.Do(httpRequest, httpResponse)
	if err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		responseError = &helpers.ResponseError{
			Code:    fasthttp.StatusInternalServerError,
			Title:   fasthttp.StatusMessage(fasthttp.StatusInternalServerError),
			Message: err.Error(),
		}
		return
	}

	if err := json.Unmarshal(httpResponse.Body(), &ThaidDataResponse); err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		responseError = &helpers.ResponseError{
			Code:    fasthttp.StatusInternalServerError,
			Title:   fasthttp.StatusMessage(fasthttp.StatusInternalServerError),
			Message: err.Error(),
		}
		return
	}

	return ThaidDataResponse, nil
}
