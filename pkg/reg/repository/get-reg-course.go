package repository

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/AlekSi/pointer"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
)

func (repo regRepository) GetRegCourseByCourseCode(regToken *string, courseCode string) ([]dto.RegCourseResponse, error) {

	args := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(args) // Release resources after use

	args.Add("course_code", courseCode)

	url := viper.GetString("reg.url.base") + viper.GetString("reg.url.path.courses") + "?" + args.String()
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()
	req.Header.SetMethod("GET")
	req.Header.Set("Authorization", "Bearer "+pointer.GetString(regToken))
	req.SetRequestURI(url)
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
	result := dto.RegGetCourseResponse{}
	err := json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Course) <= 0 {
		return nil, errors.New("course not found")
	}

	return result.Course, nil

}
