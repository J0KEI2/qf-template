package repository

import (
	"github.com/valyala/fasthttp"
	"github.com/zercle/kku-qf-services/pkg/domain"
)

type regRepository struct {
	FastHttp *fasthttp.Client
}

func NewHRRepository(fastHttp *fasthttp.Client) domain.RegRepository {
	return &regRepository{
		FastHttp: fastHttp,
	}
}
