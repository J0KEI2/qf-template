package entity

type ThaidLoginRequestDto struct {
	Code string `query:"code"`
}

type ThaidTokenPostEntity struct {
	GrantType   string `json:"grant_type"`
	Code        string `json:"code"`
	RedirectURI string `json:"redirect_uri"`
}

type ThaidTokenPostResponseEntity struct {
	ExpireIn         int64  `json:"expire_in"`
	AccessToken      string `json:"access_token"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	Scope            string `json:"scope"`
	Pid              string `json:"pid"`
	Name             string `json:"name"`
	NameEn           string `json:"name_en"`
	Birthdate        string `json:"birthdate"`
	Address          string `json:"address"`
	GivenName        string `json:"given_name"`
	MiddleName       string `json:"middle_name"`
	FamilyName       string `json:"family_name"`
	GivenNameEn      string `json:"given_name_en"`
	MiddlenameEn     string `json:"middlename_en"`
	FamilyNameEn     string `json:"family_name_en"`
	Gender           string `json:"gender"`
	SmartcardCode    string `json:"smartcard_code"`
	Title            string `json:"title"`
	TitleEn          string `json:"title_en"`
	Ial              string `json:"ial"`
	DateOfIssuance   string `json:"date_of_issuance"`
	DateOfExpiry     string `json:"date_of_expiry"`
	Openid           string `json:"openid"`
	Error            string `query:"error"`
	ErrorDescription string `query:"error_description"`
	ErrorURI         string `query:"error_uri"`
}
