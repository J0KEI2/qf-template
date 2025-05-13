package repository

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"github.com/zercle/kku-qf-services/pkg/models"
)

func (repo *userDetailRepository) GetAllEmployees(accessToken string, criteria *models.LecturerFetchWithNameQueryModel) (*models.LecturerFetchWithNameResponseModel, error) {
	endpoint := viper.GetString("hr.get_all_employees_url")

	params := url.Values{}

	if criteria.FirstName != nil {
		params.Set("firstname", *criteria.FirstName)
	}

	if criteria.LastName != nil {
		params.Set("lastname", *criteria.LastName)
	}

	endpoint += "?" + params.Encode()

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(endpoint)
	req.Header.Set("hr-access-token", accessToken)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := fasthttp.Do(req, resp); err != nil {
		err = fmt.Errorf("error: %s", err)
		return nil, err
	}

	employees := models.LecturerFetchWithNameResponseModel{}

	err := json.Unmarshal(resp.Body(), &employees)

	if err != nil {
		err = fmt.Errorf("error: %s", err)
		return nil, err
	}

	return &employees, err
}
