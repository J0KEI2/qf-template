package repository

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"github.com/zercle/kku-qf-services/pkg/models"
)

func (r *userDetailRepository) GetSingleEmployee(hrKey, accessToken string) ([]byte, error) {
	hrDomain := viper.GetString("hr.domain")

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()

	signInpath := fmt.Sprintf("v1/employees/%v", hrKey)
	requestURL := fmt.Sprintf("%v%v", hrDomain, signInpath)
	req.Header.SetMethod("GET")
	req.Header.Set("hr-access-token", accessToken)
	req.SetRequestURI(requestURL)
	res := fasthttp.AcquireResponse()

	if err := fasthttp.Do(req, res); err != nil {
		return nil, err
	}

	responseBody := res.Body()
	statusCode := res.StatusCode()

	if statusCode == 401 {
		unauthorizedError := models.UnauthorizedError()
		return nil, unauthorizedError
	}

	if statusCode == 403 {
		forbiddenError := models.ForbiddenError()
		return nil, forbiddenError
	}

	if statusCode == 404 {
		notFoundError := models.NotFoundError()
		return nil, notFoundError
	}

	if statusCode == 404 {
		internalServerError := models.InternalServerError()
		return nil, internalServerError
	}

	return responseBody, nil
}
