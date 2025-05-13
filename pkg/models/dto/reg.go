package dto

import "encoding/json"

type RegErrModel struct {
	Status RegErrStatusModel `json:"status"`
}

type RegErrStatusModel struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

type GetRegRequestSignInParam struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	GrantType    string `json:"grant_type"`
	Scope        string `json:"scope"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type RegSignInResponse struct {
	TokenType    string `json:"token_type"`
	ExpiredIn    uint   `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RegGetCourseResponse struct {
	Status  RegStatusResponse   `json:"status"`
	Message string              `json:"message"`
	Course  []RegCourseResponse `json:"course"`
}

type RegStatusResponse struct {
	Code uint   `json:"code"`
	Text string `json:"text"`
}

type RegCourseResponse struct {
	CourseID        *string      `json:"course_id"`
	CourseCode      *string      `json:"course_code"`
	RevisionCode    *string      `json:"revision_code"`
	CourseName      *string      `json:"course_name"`
	CourseNameEng   *string      `json:"course_name_eng"`
	FacultyID       *string      `json:"faculty_id"`
	DepartmentID    *string      `json:"department_id"`
	FacultyName     *string      `json:"faculty_name"`
	DepartmentName  *string      `json:"department_name"`
	CourseGroup     *string      `json:"course_group"`
	GradeMode       *string      `json:"grade_mode"`
	CreditTotal     *string      `json:"credit_total"`
	Credit1         *string      `json:"credit1"`
	Credit2         *string      `json:"credit2"`
	Credit3         *string      `json:"credit3"`
	Period1         *json.Number `json:"period1"`
	Period2         *json.Number `json:"period2"`
	Period3         *json.Number `json:"period3"`
	Description1    *string      `json:"description1"`
	DescriptionEng1 *string      `json:"descriptioneng1"`
	Description2    *string      `json:"description2"`
	DescriptionEng2 *string      `json:"descriptioneng2"`
	Description3    *string      `json:"description3"`
	DescriptionEng3 *string      `json:"descriptioneng3"`
}
