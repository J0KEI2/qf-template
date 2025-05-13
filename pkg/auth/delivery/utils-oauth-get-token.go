package delivery

import (
	"fmt"
	"log"
	"net/url"

	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/internal/datasources"
	"github.com/zercle/kku-qf-services/pkg/utils"
)

// GetOauthToken from KKU's oauth
func (h *authHandler) GetOauthToken(code string) (accessToken []byte, refreshToken []byte, err error) {
	if h.fasthttpClient == nil {
		h.fasthttpClient = datasources.NewFastHTTPClient(true)
	}

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()

	paramVal := url.Values{}
	paramVal.Set("client_id", viper.GetString("oauth.id"))
	paramVal.Set("client_secret", viper.GetString("oauth.secret"))
	paramVal.Set("grant_type", "authorization_code")
	paramVal.Set("code", code)

	req.SetRequestURI(viper.GetString("oauth.prd_url") + "/token")
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentType("application/x-www-form-urlencoded")
	req.SetBodyString(paramVal.Encode())

	if err = h.fasthttpClient.Do(req, resp); err != nil {
		log.Printf("%+v\nGetOauthToken: %+v \n%+v", helpers.WhereAmI(), req.Header.String(), req.Body())
		err = fmt.Errorf("GetOauthToken: %+v", err)
		return
	}

	if resp.StatusCode() != fasthttp.StatusOK {
		log.Printf("%+v\nGetOauthTokenResp: %+v \n%+s", helpers.WhereAmI(), resp.Header.String(), resp.Body())
		err = fmt.Errorf("%+v\nGetOauthToken: %+v", helpers.WhereAmI(), resp.StatusCode())
		return
	}

	jsonParser := utils.JsonParserPool.Get()
	defer utils.JsonParserPool.Put(jsonParser)

	vals, err := jsonParser.ParseBytes(resp.Body())
	if err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		return
	}

	accessToken = vals.GetStringBytes("access_token")
	refreshToken = vals.GetStringBytes("refresh_token")

	return
}
