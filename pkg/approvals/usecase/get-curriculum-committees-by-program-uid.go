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

func (u *programApprovalUsecase) GetCurriculumCommitteesApprovalByProgramUID(programUID *uuid.UUID, userApprovalLevel uint) (result *dto.GetApprovalCurriculumCommitteeDto, err error) {
	approval, committees, err := u.ProgramApprovalRepository.GetCurriculumCommitteesApprovalByProgramUID(*programUID)
	if err != nil {
		log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
		return nil, err
	}

	var fileList []query.MapFilesSystemQueryEntity
	fileQuery := query.MapFilesSystemQueryEntity{
		ApprovalID: &approval.ID,
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

	sort.SliceStable(committees, func(i, j int) bool {
		return committees[i].Email < committees[j].Email
	})

	result = &dto.GetApprovalCurriculumCommitteeDto{
		ID:                          approval.ID,
		ApprovalStatusLevel:         approval.ApprovalStatusLevel,
		NameTh:                      approval.NameTh,
		NameEn:                      approval.NameEn,
		ProgramUID:                  programUID.String(),
		IsCurrent:                   approval.IsCurrent,
		IsApproved:                  approval.IsApproved,
		IsRejected:                  approval.IsRejected,
		IsEditable:                  u.CheckIfApprovalEditable(uint(constant.CURRICULUM_COMMITEES_APPROVAL_LEVEL), userApprovalLevel),
		IsCommitteeSelected:         approval.IsCommitteeSelected,
		Committees:                  committees,
		CurriculumCommitteeFileList: approvalFileList,
		CreatedAt:                   approval.CreatedAt,
		UpdatedAt:                   approval.UpdatedAt,
	}

	return result, nil
}
