package delivery

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	helpers "github.com/zercle/gofiber-helpers"
	oauth_models "github.com/zercle/kku-oauth2-saml/pkg/models"
	"github.com/zercle/kku-qf-services/internal/datasources"
)

// GetUserSSO from KKU's oauth
func (h *authHandler) GetUserSSO(accessToken string) (userSSO oauth_models.SsoUser, err error) {
	if h.fasthttpClient == nil {
		h.fasthttpClient = datasources.NewFastHTTPClient(true)
	}

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()

	req.SetRequestURI(viper.GetString("oauth.prd_url") + "/api/v1/user")
	req.Header.SetMethod(fasthttp.MethodGet)
	req.Header.Set("Authorization", "Bearer "+accessToken)

	err = h.fasthttpClient.Do(req, resp)
	if err != nil {
		log.Printf("GetUserDetail: %+v \n%+v", req.Header.String(), req.Body())
		err = fmt.Errorf("GetUserDetail: %+v", err)
		return
	}
	if resp.StatusCode() != http.StatusOK {
		log.Printf("GetUserDetail: %+v \n%+v", req.Header.String(), req.Body())
		log.Printf("GetUserDetail: %+v \n%+v", resp.Header.String(), req.Body())
		err = fmt.Errorf("GetUserDetail: %+v", resp.StatusCode())
		return
	}

	err = json.NewDecoder(bytes.NewReader(resp.Body())).Decode(&userSSO)

	if err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
		return
	}

	return
}
