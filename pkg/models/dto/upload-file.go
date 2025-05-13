package dto

type UploadFileResponseDto struct {
	FileName            string `json:"file_name"`
	FileNameForDownload string `json:"file_name_for_download"`
}
