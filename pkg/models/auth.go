package models

type ThaidLoginRequestDto struct {
	Code  string `query:"code"`
	State string `query:"state"`
}

type OauthCallback struct {
	Code string `json:"code"`
}
