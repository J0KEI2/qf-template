package repository

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
)

func (repo hrRepository) GetLecturerPagination(HRKey, firstname, lastname string, options models.PaginationOptions) (*dto.GetEmployeesResponseDto, error) {
	url := viper.GetString("hr.url.base") + viper.GetString("hr.url.path.employees")
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()
	log.Println(firstname, lastname)
	req.Header.SetMethod("GET")
	req.Header.SetContentType("application/json")
	req.Header.Set("hr-access-token", HRKey)
	req.SetRequestURI(url)
	req.URI().QueryArgs().Add("firstname", firstname)
	req.URI().QueryArgs().Add("lastname", lastname)
	req.URI().QueryArgs().Add("page", fmt.Sprintf("%d", options.GetPage()))
	req.URI().QueryArgs().Add("size", fmt.Sprintf("%d", options.GetLimit()))
	res := fasthttp.AcquireResponse()

	if err := fasthttp.Do(req, res); err != nil {
		return nil, err
	}

	responseBody := res.Body()
	log.Println(string(responseBody))
	statusCode := res.StatusCode()
	if statusCode != 200 {
		errorModel := dto.ErrNotFoundModel{}
		err := json.Unmarshal(responseBody, &errorModel)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf(errorModel.Data.Message)
	}

	result := dto.GetEmployeesResponseDto{}
	err := json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil

}
