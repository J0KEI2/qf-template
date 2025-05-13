package approvals

import (
	"log"
	"sort"

	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	constant "github.com/zercle/kku-qf-services/pkg/constant/common"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	"github.com/zercle/kku-qf-services/pkg/models/query-model"
)

func (u *programApprovalUsecase) GetUniversityApprovalByProgramUID(programUID *uuid.UUID, userApprovalLevel uint) (result *dto.GetApprovalUniversityDto, err error) {
	result, err = u.ProgramApprovalRepository.GetUniversityApprovalByProgramUID(*programUID)
	if err != nil {
		log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
		return nil, err
	}

	var fileList []query.MapFilesSystemQueryEntity
	fileQuery := query.MapFilesSystemQueryEntity{
		ApprovalID: &result.ID,
	}
	if err = u.CommonRepository.GetList(&fileQuery, &fileList, nil, "FileSystem"); err != nil {
		return nil, err
	}

	var approvalFileList []dto.ApprovalFileList
	for _, file := range fileList {
		approvalFileList = append(approvalFileList, dto.ApprovalFileList{
			FileID:   file.FileSystem.ID,
			FileName: file.FileSystem.FileName,
		})
	}

	result.UniversityFileList = approvalFileList
	result.IsEditable = u.CheckIfApprovalEditable(uint(constant.FACULTY_APPROVAL_LEVEL), userApprovalLevel)

	sort.SliceStable(result.Submissions, func(i, j int) bool {
		return result.Submissions[i].ID < result.Submissions[j].ID
	})

	return result, nil
}
