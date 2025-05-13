package repository

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
)

func (repo hrRepository) GetLecturerBySSN(HRKey, ssn string) (*dto.GetSingleEmployeeResponseDto, error) {
	url := viper.GetString("hr.url.base") + viper.GetString("hr.url.path.employees") + "/" + ssn
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()
	req.Header.SetMethod("GET")
	req.Header.SetContentType("application/json")
	req.Header.Set("hr-access-token", HRKey)
	req.SetRequestURI(url)
	res := fasthttp.AcquireResponse()

	if err := fasthttp.Do(req, res); err != nil {
		return nil, err
	}

	responseBody := res.Body()
	statusCode := res.StatusCode()
	if statusCode != 200 {
		errorModel := dto.ErrNotFoundModel{}
		err := json.Unmarshal(responseBody, &errorModel)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf(errorModel.Data.Message)
	}

	result := dto.GetSingleEmployeeResponseDto{}
	err := json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil

}
