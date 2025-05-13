package dto

type HREducationsResponseDto struct {
	EducationLevel   *string `json:"educationlevel"`
	Qualification    *string `json:"qualification"`
	Department       *string `json:"major"`
	InstituteName    *string `json:"institute"`
	InstituteCountry *string `json:"country"`
	GraduateYear     *int    `json:"successyear"`
}

type GetRequestSignInParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ReponseSignInModel struct {
	Status string            `json:"status"`
	Data   DataResponseModel `json:"data"`
}

type DataResponseModel struct {
	ID           string   `json:"id"`
	Username     string   `json:"username"`
	Email        string   `json:"email"`
	AccessToken  string   `json:"accessToken"`
	RefreshToken string   `json:"refreshToken"`
	Roles        []string `json:"roles"`
}

type ErrNotFoundModel struct {
	Status string            `json:"status"`
	Data   ErrorNotFoundData `json:"data"`
}

type ErrorNotFoundData struct {
	Message string `json:"message"`
}

type GetEducationsResponseDto struct {
	Status string                  `json:"status"`
	Data   HRPaginationResponseDto `json:"data"`
}

type HRPaginationResponseDto struct {
	TotalItems  int                 `json:"totalItems"`
	Items       []HREducationDetail `json:"items"`
	TotalPages  int                 `json:"totalPages"`
	CurrentPage int                 `json:"currentPage"`
}

type HREducationDetail struct {
	ID             int    `json:"id"`
	SuccessYear    string `json:"successyear"`
	EducationLevel string `json:"educationlevel"`
	Qualification  string `json:"qualification"`
	Major          string `json:"major"`
	Institute      string `json:"institute"`
	Country        string `json:"country"`
}
