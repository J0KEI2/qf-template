package models

const (
	Lecturer = "อาจารย์"
)

type LecturerSigninResponse struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
}

type Data struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type EmployeeDetail struct {
	Status       string       `json:"status"`
	EmployeeData EmployeeData `json:"data"`
}

type EmployeeData struct {
	Title        string `json:"title"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	TitleEng     string `json:"titleEng"`
	FirstNameEng string `json:"firstnameEng"`
	LastNameEng  string `json:"lastnameEng"`
	PositionType string `json:"PositionType"`
	StatusList   string `json:"positionlist"`
	WorkLine     string `json:"workline"`
	Position     string `json:"position"`
	Faculty      string `json:"faculty"`
	Division     string `json:"division"`
	Job          string `json:"job"`
	Unit         string `json:"unit"`
	PutDay       string `json:"putday"`
	Email        string `json:"email"`
}

type CronUpdateLecturer struct {
	HRKey string `json:"hrKey,omitempty"`
	UID   string `json:"uid"`
}

type LecturerFetchWithNameRequestModel struct {
	FullName *string
}

type LecturerFetchWithNameQueryModel struct {
	FirstName *string
	LastName  *string
	Page      *int
	Size      *int
}

type LecturerFetchWithNameResponseModel struct {
	Status string                                 `json:"status"`
	Data   LecturerFetchWithNameEmployeeListModel `json:"data"`
}

type LecturerFetchWithNameEmployeeListModel struct {
	TotalItems  int                                  `json:"totalItems"`
	Items       []LecturerFetchWithNameEmployeeModel `json:"items"`
	TotalPages  int                                  `json:"totalPages"`
	CurrentPage int                                  `json:"currentPage"`
}

type LecturerFetchWithNameEmployeeModel struct {
	Title        string  `json:"title"`
	FirstName    string  `json:"firstname"`
	LastName     string  `json:"lastname"`
	TitleEN      string  `json:"titleEng"`
	FirstNameEN  string  `json:"firstnameEng"`
	LastNameEN   string  `json:"lastnameEng"`
	PositionType string  `json:"positionType"`
	StatusList   string  `json:"statuslist"`
	WorkLine     string  `json:"workline"`
	Position     string  `json:"position"`
	Faculty      string  `json:"faculty"`
	Division     string  `json:"division"`
	Job          string  `json:"job"`
	Unit         string  `json:"unit"`
	Putday       string  `json:"putday"`
	Email        *string `json:"email"`
}
