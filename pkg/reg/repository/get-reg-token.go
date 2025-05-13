package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"

	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
)

func (repo regRepository) GetRegToken() (*string, error) {
	url := viper.GetString("reg.url.base") + viper.GetString("reg.url.path.sign_in")
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	// Add text field
	writer.WriteField("username", viper.GetString("reg.auth.username"))
	writer.WriteField("password", viper.GetString("reg.auth.password"))
	writer.WriteField("grant_type", "password")
	writer.WriteField("scope", "*")
	writer.WriteField("client_id", viper.GetString("reg.auth.client_id"))
	writer.WriteField("client_secret", viper.GetString("reg.auth.client_secret"))

	writer.Close()

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()
	req.Header.SetMethod("POST")
	req.Header.SetContentType(writer.FormDataContentType())
	req.SetRequestURI(url)
	req.SetBody(body.Bytes())
	res := fasthttp.AcquireResponse()

	if err := fasthttp.Do(req, res); err != nil {
		return nil, err
	}

	responseBody := res.Body()
	statusCode := res.StatusCode()
	errorModel := dto.RegErrModel{}
	if statusCode == 401 {
		err := json.Unmarshal(responseBody, &errorModel)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf(errorModel.Status.Text)
	}

	if statusCode == 404 {
		err := json.Unmarshal(responseBody, &errorModel)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf(errorModel.Status.Text)
	}
	result := dto.RegSignInResponse{}
	err := json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	return &result.AccessToken, nil

}
