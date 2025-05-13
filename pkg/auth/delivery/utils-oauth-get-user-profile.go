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

// GetUserProfile from KKU's oauth
func (h *authHandler) GetUserProfile(accessToken string) (userProfile oauth_models.UserProfile, err error) {
	if h.fasthttpClient == nil {
		h.fasthttpClient = datasources.NewFastHTTPClient(true)
	}

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()

	req.SetRequestURI(viper.GetString("oauth.prd_url") + "/api/v2/user")
	req.Header.SetMethod(fasthttp.MethodGet)
	req.Header.Set("Authorization", "Bearer "+accessToken)

	if err = h.fasthttpClient.Do(req, resp); err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		log.Printf("GetUserDetail: %+v \n%+v", req.Header.String(), req.Body())
		err = fmt.Errorf("GetUserDetail: %+v", err)
		return
	}

	if resp.StatusCode() != http.StatusOK {
		log.Printf("GetUserDetail: %+v \n%+s", req.Header.String(), req.Body())
		log.Printf("GetUserDetail: %+v \n%+s", resp.Header.String(), resp.Body())
		err = fmt.Errorf("GetUserDetail: %+v", resp.StatusCode())
		return
	}

	if err = json.NewDecoder(bytes.NewReader(resp.Body())).Decode(&userProfile); err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		log.Printf("GetUserDetail: %+v", err.Error())
		return
	}

	return
}
