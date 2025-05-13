package repository

import (
	"github.com/valyala/fasthttp"
	"github.com/zercle/kku-qf-services/pkg/domain"
)

type hrRepository struct {
	FastHttp *fasthttp.Client
}

func NewHRRepository(fastHttp *fasthttp.Client) domain.HRRepository {
	return &hrRepository{
		FastHttp: fastHttp,
	}
}
