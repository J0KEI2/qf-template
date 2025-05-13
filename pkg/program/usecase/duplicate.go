package usecase

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/AlekSi/pointer"
	"github.com/google/uuid"

	helpers "github.com/zercle/gofiber-helpers"
	permissionConstant "github.com/zercle/kku-qf-services/pkg/constant/permission"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	permissionQuery "github.com/zercle/kku-qf-services/pkg/models/query-model/permission"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	rapQuery "github.com/zercle/kku-qf-services/pkg/models/query-model/role-and-permission"
	"gorm.io/gorm"
)

func (u programUsecase) DuplicateProgram(userUID uuid.UUID, request dto.ProgramDuplicateRequestDto) (newProgramId *uuid.UUID, err error) {
	newProgramID := uuid.New()
	if err = helper.ExecuteTransaction(u.CommonRepository, u.CreateNewProgram(newProgramID)); err != nil {
		return nil, err
	}

	programQuery := query.ProgramMainQueryEntity{
		ID: &request.ProgramMainID,
	}

	err = u.CommonRepository.GetFirst(&programQuery,
		"ProgramGeneralDetail",
		"ProgramGeneralDetail.ProgramMajor",
		"ProgramGeneralDetail.ProgramMajor.ProgramPlanDetail",
		"ProgramGeneralDetail.ProgramMajor.ProgramPlanDetail.ProgramSubPlan",
		"ProgramPolicyAndStrategic",
		"ProgramOwner",
		"ProgramLecturer",
		"ProgramThesisLecturer",
		"ProgramQualityAssurance",
		"ProgramSystemAndMechanic",
		"ProgramPermission")
	if err != nil {
		return nil, err
	}

	mapOldSubPlanToNewOne := map[uint]uint{}
	mapOldPlanToNewOne := map[uint]uint{}
	mapOldYearToNewOne := map[uint]uint{}
	mapOldCourseToNewOne := map[uint]uint{}
	mapOldYLOKsecToNewOne := map[uint]uint{}
	mapOldMapKsecPloIdToNewOne := map[uint]uint{}
	mapOldPloIdToNewOne := map[uint]uint{}
	maoOldKsaToNewOne := map[uint]uint{}

	if err = helper.ExecuteTransaction(u.CommonRepository, u.DuplicateGeneralDetail(programQuery, newProgramID, mapOldSubPlanToNewOne, mapOldPlanToNewOne)); err != nil {
		return nil, err
	}

	if err = helper.ExecuteTransaction(u.CommonRepository, u.DuplicateYearAndSemester(programQuery, newProgramID, mapOldSubPlanToNewOne, mapOldYearToNewOne)); err != nil {
		return nil, err
	}

	if err = helper.ExecuteTransaction(u.CommonRepository, u.DuplicateCourseStruct(programQuery, newProgramID, mapOldSubPlanToNewOne, mapOldYearToNewOne, mapOldCourseToNewOne)); err != nil {
		return nil, err
	}

	if err = helper.ExecuteTransaction(u.CommonRepository, u.DuplicateSubMainData(programQuery, newProgramID)); err != nil {
		return nil, err
	}

	if err = helper.ExecuteTransaction(u.CommonRepository, u.DuplicatePlo(programQuery, mapOldSubPlanToNewOne, mapOldMapKsecPloIdToNewOne, mapOldPloIdToNewOne)); err != nil {
		return nil, err
	}

	if err = helper.ExecuteTransaction(u.CommonRepository, u.DuplicateYLOKsec(programQuery, newProgramID, mapOldSubPlanToNewOne, mapOldYearToNewOne, mapOldYLOKsecToNewOne, mapOldMapKsecPloIdToNewOne, mapOldPloIdToNewOne)); err != nil {
		return nil, err
	}

	if err = helper.ExecuteTransaction(u.CommonRepository, u.DuplicatePlanAndEvaluate(newProgramID, mapOldSubPlanToNewOne, programQuery)); err != nil {
		return nil, err
	}

	// if err = helper.ExecuteTransaction(u.CommonRepository, u.DuplicateCLO(mapOldPlanToNewOne, programQuery)); err != nil {
	// 	return nil, err
	// }

	if err = helper.ExecuteTransaction(u.CommonRepository, u.DuplicateKsaDetail(mapOldSubPlanToNewOne, programQuery, maoOldKsaToNewOne)); err != nil {
		return nil, err
	}

	if err = helper.ExecuteTransaction(u.CommonRepository, u.DuplicateMapCurMapResp(mapOldSubPlanToNewOne, programQuery, mapOldCourseToNewOne, mapOldPloIdToNewOne, maoOldKsaToNewOne)); err != nil {
		return nil, err
	}

	if err = helper.ExecuteTransaction(u.CommonRepository, u.DuplicateMapCurMapKsa(mapOldSubPlanToNewOne, programQuery, mapOldCourseToNewOne, mapOldPloIdToNewOne, maoOldKsaToNewOne)); err != nil {
		return nil, err
	}

	if err = helper.ExecuteTransaction(u.CommonRepository, u.DuplicateApprovals(newProgramID)); err != nil {
		return nil, err
	}

	if err = helper.ExecuteTransaction(u.CommonRepository, u.CreateProgramOwnerPermission(userUID, newProgramID), u.CreateProgramLecturerPermission(userUID, newProgramID)); err != nil {
		return nil, err
	}

	return &newProgramID, nil
}

func (u programUsecase) CreateNewProgram(newProgramID uuid.UUID) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		newProgram := query.ProgramMainQueryEntity{
			ID: &newProgramID,
		}

		err := u.CommonRepository.Create(tx, &newProgram)
		if err != nil {
			return err
		}
		return nil
	}
}

func (u programUsecase) DuplicateGeneralDetail(programQuery query.ProgramMainQueryEntity, newProgramID uuid.UUID, mapOldSubPlanToNewOne map[uint]uint, mapOldPlanToNewOne map[uint]uint) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		err = u.duplicateGeneralDetail(tx, newProgramID, programQuery.ProgramGeneralDetail, mapOldSubPlanToNewOne, mapOldPlanToNewOne)
		if err != nil {
			return err
		}

		return
	}
}

func (u programUsecase) duplicateGeneralDetail(tx *gorm.DB, newProgramID uuid.UUID, programGeneralDetail *query.ProgramGeneralDetailQueryEntity, mapOldSubPlanToNewOne map[uint]uint, mapOldPlanToNewOne map[uint]uint) (err error) {
	programNameTh := fmt.Sprintf("%v (%v)", *programGeneralDetail.ProgramNameTH, "สำเนา")
	programNameEn := fmt.Sprintf("%v (%v)", *programGeneralDetail.ProgramNameEN, "Copied")
	newProgramGD := query.ProgramGeneralDetailQueryEntity{
		UniversityName:        programGeneralDetail.UniversityName,
		FacultyID:             programGeneralDetail.FacultyID,
		ProgramNameTH:         &programNameTh,
		ProgramNameEN:         &programNameEn,
		BranchNameTH:          programGeneralDetail.BranchNameTH,
		BranchNameEN:          programGeneralDetail.BranchNameEN,
		DegreeNameTH:          programGeneralDetail.DegreeNameTH,
		DegreeNameEN:          programGeneralDetail.DegreeNameEN,
		ProgramMajorTypeID:    programGeneralDetail.ProgramMajorTypeID,
		ProgramMajorType:      programGeneralDetail.ProgramMajorType,
		ProgramDegreeTypeID:   programGeneralDetail.ProgramDegreeTypeID,
		ProgramDegreeType:     programGeneralDetail.ProgramDegreeType,
		OverallCredit:         programGeneralDetail.OverallCredit,
		NumberOfYear:          programGeneralDetail.NumberOfYear,
		ProgramLanguageID:     programGeneralDetail.ProgramLanguageID,
		ProgramLanguage:       programGeneralDetail.ProgramLanguage,
		Admission:             programGeneralDetail.Admission,
		MOU:                   programGeneralDetail.MOU,
		MOUFilepath:           programGeneralDetail.MOUFilepath,
		ProgramTypeID:         programGeneralDetail.ProgramTypeID,
		ProgramType:           programGeneralDetail.ProgramType,
		ProgramYearID:         programGeneralDetail.ProgramYearID,
		ProgramYear:           programGeneralDetail.ProgramYear,
		Semester:              programGeneralDetail.Semester,
		SemesterYear:          programGeneralDetail.SemesterYear,
		BoardApproval:         programGeneralDetail.BoardApproval,
		BoardApprovalDate:     programGeneralDetail.BoardApprovalDate,
		AcademicCouncil:       programGeneralDetail.AcademicCouncil,
		AcademicCouncilDate:   programGeneralDetail.AcademicCouncilDate,
		UniversityCouncil:     programGeneralDetail.UniversityCouncil,
		UniversityCouncilDate: programGeneralDetail.UniversityCouncilDate,
		IsSamePlanMajor:       programGeneralDetail.IsSamePlanMajor,
	}

	err = u.CommonRepository.Create(tx, &newProgramGD)
	if err != nil {
		return err
	}

	mainQuery := query.ProgramMainQueryEntity{
		ID: &newProgramID,
	}

	mainUpdate := query.ProgramMainQueryEntity{
		ProgramGeneralDetailID: newProgramGD.ID,
	}

	err = u.CommonRepository.Update(tx, &mainQuery, &mainUpdate)

	if err != nil {
		return err
	}

	for _, majorItem := range programGeneralDetail.ProgramMajor {
		majorBody := query.ProgramMajorQueryEntity{
			ProgramGeneralDetailID: newProgramGD.ID,
			Name:                   majorItem.Name,
		}
		u.CommonRepository.Create(tx, &majorBody)
		for _, planItem := range majorItem.ProgramPlanDetail {
			planBody := query.ProgramPlanDetailQueryEntity{
				ProgramMajorID: majorBody.ID,
				PlanName:       planItem.PlanName,
				CreditRulesID:  planItem.CreditRulesID,
				CreditRules:    planItem.CreditRules,
				Credit:         planItem.Credit,
				IsSplitPlan:    planItem.IsSplitPlan,
				IsActive:       planItem.IsActive,
			}
			u.CommonRepository.Create(tx, &planBody)
			oldPlan := pointer.GetUint(planItem.ID)
			mapOldPlanToNewOne[oldPlan] = pointer.GetUint(planBody.ID)
			for _, subPlanItem := range planItem.ProgramSubPlan {
				subPlanBody := query.ProgramSubPlanQueryEntity{
					ProgramPlanDetailID: planBody.ID,
					SubPlanName:         subPlanItem.SubPlanName,
					CreditRulesID:       subPlanItem.CreditRulesID,
					CreditRules:         subPlanItem.CreditRules,
					Credit:              subPlanItem.Credit,
				}
				u.CommonRepository.Create(tx, &subPlanBody)
				oldSubPlan := pointer.GetUint(subPlanItem.ID)
				mapOldSubPlanToNewOne[oldSubPlan] = pointer.GetUint(subPlanBody.ID)
			}
		}
	}
	return nil
}

func (u programUsecase) DuplicateYearAndSemester(programQuery query.ProgramMainQueryEntity, newProgramID uuid.UUID, mapOldSubPlanToNewOne, mapOldYearToNewOne map[uint]uint) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		for _, majorItem := range programQuery.ProgramGeneralDetail.ProgramMajor {
			for _, planDetailItem := range majorItem.ProgramPlanDetail {
				for _, subPlanItem := range planDetailItem.ProgramSubPlan {
					oldSubPlan := pointer.GetUint(subPlanItem.ID)
					newSubPlan := mapOldSubPlanToNewOne[oldSubPlan]
					yearAndSemesterStatement := query.ProgramYearAndSemesterQueryEntity{
						ProgramSubPlanID: subPlanItem.ID,
					}

					yearAndSemesterResult := []query.ProgramYearAndSemesterQueryEntity{}

					u.CommonRepository.GetList(yearAndSemesterStatement, &yearAndSemesterResult, nil)

					for _, yearAndSemester := range yearAndSemesterResult {
						newYearAndSemester := query.ProgramYearAndSemesterQueryEntity{
							ProgramSubPlanID: &newSubPlan,
							CourseDetail:     yearAndSemester.CourseDetail,
							Year:             yearAndSemester.Year,
							Semester:         yearAndSemester.Semester,
						}
						u.CommonRepository.Create(tx, &newYearAndSemester)
						mapOldYearToNewOne[*yearAndSemester.ID] = pointer.GetUint(newYearAndSemester.ID)
					}
				}
			}
		}

		return
	}
}

func (u programUsecase) DuplicateCourseStruct(programQuery query.ProgramMainQueryEntity, newProgramID uuid.UUID, mapOldSubPlanToNewOne, mapOldYearToNewOne, mapOldCourseToNewOne map[uint]uint) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		for _, majorItem := range programQuery.ProgramGeneralDetail.ProgramMajor {
			for _, planDetailItem := range majorItem.ProgramPlanDetail {
				for _, subPlanItem := range planDetailItem.ProgramSubPlan {

					oldSubPlan := pointer.GetUint(subPlanItem.ID)
					newSubPlan := mapOldSubPlanToNewOne[oldSubPlan]

					structureStatement := query.ProgramStructureDetailQueryEntity{
						ProgramSubPlanID: &oldSubPlan,
						ParentID:         pointer.ToUint(0),
					}

					structureResult := []query.ProgramStructureDetailQueryEntity{}

					u.CommonRepository.GetListWithNilSearch(&structureStatement, &structureResult, nil, "CourseDetail", "Children")

					u.recursiveStruct(tx, structureResult, pointer.ToUint(0), nil, newSubPlan, oldSubPlan, mapOldYearToNewOne, mapOldCourseToNewOne)
				}
			}
		}

		return nil
	}
}

func (u programUsecase) recursiveStruct(tx *gorm.DB, structures []query.ProgramStructureDetailQueryEntity, oldParentID, parentId *uint, newSubPlan, oldSubPlan uint, mapOldYearToNewOne map[uint]uint, mapOldCourseToNewOne map[uint]uint) {
	structureStatement := query.ProgramStructureDetailQueryEntity{
		ProgramSubPlanID: &oldSubPlan,
		ParentID:         oldParentID,
	}

	structureResult := []query.ProgramStructureDetailQueryEntity{}

	u.CommonRepository.GetListWithNilSearch(&structureStatement, &structureResult, nil, "CourseDetail", "Children")

	for _, structure := range structureResult {
		newStructure := query.ProgramStructureDetailQueryEntity{
			ProgramSubPlanID: &newSubPlan,
			Name:             structure.Name,
			Order:            structure.Order,
			ParentID:         parentId,
			Qualification:    structure.Qualification,
			StructureCredit:  structure.StructureCredit,
		}
		u.CommonRepository.Create(tx, &newStructure)
		u.recursiveStruct(tx, structure.Children, structure.ID, newStructure.ID, newSubPlan, oldSubPlan, mapOldYearToNewOne, mapOldCourseToNewOne)

		for _, courseDetail := range structure.CourseDetail {
			newCourse := query.ProgramCourseDetailQueryEntity{
				ProgramSubPlanID:   &newSubPlan,
				ProgramStructureID: newStructure.ID,
				CourseSource:       courseDetail.CourseSource,
				REGKkuKey:          courseDetail.REGKkuKey,
				CourseKey:          courseDetail.CourseKey,
				CourseTypeID:       courseDetail.CourseTypeID,
				CourseType:         courseDetail.CourseType,
				CourseCode:         courseDetail.CourseCode,
				CourseYear:         courseDetail.CourseYear,
				CourseNameTH:       courseDetail.CourseNameTH,
				CourseNameEN:       courseDetail.CourseNameEN,
				CourseCredit:       courseDetail.CourseCredit,
				Credit1:            courseDetail.Credit1,
				Credit2:            courseDetail.Credit2,
				Credit3:            courseDetail.Credit3,
				IsCreditCalc:       courseDetail.IsCreditCalc,
			}

			oldYearAndSemester := pointer.GetUint(courseDetail.YearAndSemesterID)
			newYearAndSemesterId := mapOldYearToNewOne[oldYearAndSemester]
			if newYearAndSemesterId != 0 {
				newCourse.YearAndSemesterID = pointer.ToUint(newYearAndSemesterId)
			}

			u.CommonRepository.Create(tx, &newCourse)
			oldCourseDetail := pointer.GetUint(courseDetail.ID)
			mapOldCourseToNewOne[oldCourseDetail] = pointer.GetUint(newCourse.ID)

		}
	}
}

func (u programUsecase) DuplicateSubMainData(programQuery query.ProgramMainQueryEntity, newProgramID uuid.UUID) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		err = u.duplicatePolicyAndStrategic(tx, newProgramID, programQuery.ProgramPolicyAndStrategic)
		if err != nil {
			return err
		}

		err = u.duplicateOwner(tx, newProgramID, programQuery.ProgramOwner)
		if err != nil {
			return err
		}

		err = u.duplicateLecturer(tx, newProgramID, programQuery.ProgramLecturer)
		if err != nil {
			return err
		}

		err = u.duplicateThesisLecturer(tx, newProgramID, programQuery.ProgramThesisLecturer)
		if err != nil {
			return err
		}

		err = u.duplicateQualityAssurance(tx, newProgramID, programQuery.ProgramQualityAssurance)
		if err != nil {
			return err
		}

		err = u.duplicateSystemAndMechanic(tx, newProgramID, programQuery.ProgramSystemAndMechanic)
		if err != nil {
			return err
		}

		err = u.duplicateProgramPermission(tx, newProgramID, programQuery.ProgramPermission)
		if err != nil {
			return err
		}

		err = u.duplicateReference(tx, programQuery, &newProgramID)
		if err != nil {
			return err
		}

		return nil
	}
}

func (u programUsecase) DuplicatePlo(programQuery query.ProgramMainQueryEntity, mapOldSubPlanToNewOne map[uint]uint, mapOldMapKsecPloIdToNewOne map[uint]uint, mapOldPloIdToNewOne map[uint]uint) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		for _, majorItem := range programQuery.ProgramGeneralDetail.ProgramMajor {
			for _, planDetailItem := range majorItem.ProgramPlanDetail {
				for _, subPlanItem := range planDetailItem.ProgramSubPlan {

					mapOldKsecIdWithNewOne := map[uint]uint{}

					programPLOFormat := query.ProgramPloFormatQueryEntity{
						ProgramSubPlanID: subPlanItem.ID,
					}

					err = u.CommonRepository.GetFirst(&programPLOFormat)
					if err == gorm.ErrRecordNotFound {
						continue
					}

					oldSubPlan := pointer.GetUint(programPLOFormat.ProgramSubPlanID)
					newSubPlan := pointer.ToUint(mapOldSubPlanToNewOne[oldSubPlan])
					newPloFormat := query.ProgramPloFormatQueryEntity{
						ProgramSubPlanID: newSubPlan,
						PLOFormat:        programPLOFormat.PLOFormat,
					}

					u.CommonRepository.Create(tx, &newPloFormat)

					programKsecQuery := query.ProgramKsecDetailQueryEntity{
						ProgramSubPlanID: subPlanItem.ID,
					}

					programKsec := []query.ProgramKsecDetailQueryEntity{}

					if err = u.CommonRepository.GetList(&programKsecQuery, &programKsec, nil); err != nil {
						return
					}

					for _, ksec := range programKsec {
						newKsec := query.ProgramKsecDetailQueryEntity{
							ProgramSubPlanID: newSubPlan,
							Type:             ksec.Type,
							Order:            ksec.Order,
							Detail:           ksec.Detail,
							IsChecked:        ksec.IsChecked,
						}
						oldId := pointer.GetUint(ksec.ID)
						u.CommonRepository.Create(tx, &newKsec)
						mapOldKsecIdWithNewOne[oldId] = pointer.GetUint(newKsec.ID)
					}

					programPloQuery := query.ProgramPloQueryEntity{
						ProgramPloFormatID: programPLOFormat.ID,
						ParentID:           pointer.ToUint(0),
					}

					programPloResult := []query.ProgramPloQueryEntity{}

					if err = u.CommonRepository.GetListWithNilSearch(&programPloQuery, &programPloResult, nil, "ProgramMapPloWithKsecQueryEntity", "LearningSolution", "LearningEvaluation"); err != nil {
						return
					}

					//Todo : recursive func
					u.recursiveDuplicatePlo(tx, pointer.ToUint(0), nil, programPloResult, programPLOFormat.ID, newPloFormat.ID, mapOldMapKsecPloIdToNewOne, mapOldKsecIdWithNewOne, mapOldPloIdToNewOne)
				}
			}
		}

		return nil
	}
}

func (u programUsecase) recursiveDuplicatePlo(tx *gorm.DB, oldParentId, parentID *uint, programPloResult []query.ProgramPloQueryEntity, oldFormatID, programPLOFormatId *uint, mapOldMapKsecPloIdToNewOne map[uint]uint, mapOldKsecIdWithNewOne map[uint]uint, mapOldPloIdToNewOne map[uint]uint) error {
	programPloQuery := query.ProgramPloQueryEntity{
		ProgramPloFormatID: oldFormatID,
		ParentID:           oldParentId,
	}

	programPloResultList := []query.ProgramPloQueryEntity{}

	if err := u.CommonRepository.GetListWithNilSearch(&programPloQuery, &programPloResultList, nil, "ProgramMapPloWithKsecQueryEntity", "LearningSolution", "LearningEvaluation"); err != nil {
		return err
	}

	for _, programPlo := range programPloResultList {
		newPlo := query.ProgramPloQueryEntity{
			ProgramPloFormatID: programPLOFormatId,
			Order:              programPlo.Order,
			ParentID:           parentID,
			PLOPrefix:          programPlo.PLOPrefix,
			PLODetail:          programPlo.PLODetail,
		}
		u.CommonRepository.Create(tx, &newPlo)

		oldPloId := pointer.GetUint(programPlo.ID)

		mapOldPloIdToNewOne[oldPloId] = pointer.GetUint(newPlo.ID)
		u.recursiveDuplicatePlo(tx, programPlo.ID, newPlo.ID, programPlo.Children, programPlo.ProgramPloFormatID, programPLOFormatId, mapOldMapKsecPloIdToNewOne, mapOldKsecIdWithNewOne, mapOldPloIdToNewOne)

		for _, learningSolution := range programPlo.LearningSolution {
			newLearningSolution := query.ProgramPLOLearningSolutionQueryEntity{
				PloID:  newPlo.ID,
				Key:    learningSolution.Key,
				Detail: learningSolution.Detail,
				Order:  learningSolution.Order,
			}
			u.CommonRepository.Create(tx, &newLearningSolution)
		}

		for _, learningEvaluation := range programPlo.LearningEvaluation {
			newLearningEvaluation := query.ProgramPLOLearningEvaluationQueryEntity{
				PloID:  newPlo.ID,
				Key:    learningEvaluation.Key,
				Detail: learningEvaluation.Detail,
				Order:  learningEvaluation.Order,
			}
			u.CommonRepository.Create(tx, &newLearningEvaluation)
		}

		for _, mapPloWithKsec := range programPlo.ProgramMapPloWithKsecQueryEntity {
			ksecId := pointer.GetUint(mapPloWithKsec.KsecID)
			newMapPloWithKsec := query.ProgramMapPloWithKsecQueryEntity{
				PloID:  newPlo.ID,
				KsecID: pointer.ToUint(mapOldKsecIdWithNewOne[ksecId]),
			}
			u.CommonRepository.Create(tx, &newMapPloWithKsec)
			oldId := pointer.GetUint(mapPloWithKsec.ID)
			mapOldMapKsecPloIdToNewOne[oldId] = pointer.GetUint(newMapPloWithKsec.ID)
		}
	}
	return nil
}

func (u programUsecase) DuplicateYLOKsec(programQuery query.ProgramMainQueryEntity, newProgramID uuid.UUID, mapOldSubPlanToNewOne map[uint]uint, mapOldYearToNewOne map[uint]uint, mapOldYLOKsecToNewOne, mapOldMapKsecPloIdToNewOne, mapOldPloIdToNewOne map[uint]uint) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		for _, majorItem := range programQuery.ProgramGeneralDetail.ProgramMajor {
			for _, planDetailItem := range majorItem.ProgramPlanDetail {
				for _, subPlanItem := range planDetailItem.ProgramSubPlan {
					yloKsecQuery := query.ProgramYloKsecQueryEntity{
						ProgramSubPlanID: subPlanItem.ID,
					}

					yloKsecList := []query.ProgramYloKsecQueryEntity{}
					if err = u.CommonRepository.GetList(&yloKsecQuery, &yloKsecList, nil, "ProgramYearAndSemester"); err != nil {
						return err
					}

					for _, yloKsecItem := range yloKsecList {
						oldYearID := pointer.GetUint(yloKsecItem.ProgramYearAndSemesterID)
						newYearID := mapOldYearToNewOne[oldYearID]
						oldSubplan := pointer.GetUint(subPlanItem.ID)
						newSubPlan := mapOldSubPlanToNewOne[oldSubplan]
						yloKsecBody := query.ProgramYloKsecQueryEntity{
							ProgramSubPlanID:         pointer.ToUint(newSubPlan),
							ProgramYearAndSemesterID: pointer.ToUint(newYearID),
							Knowledge:                yloKsecItem.Knowledge,
							Skill:                    yloKsecItem.Skill,
							Ethic:                    yloKsecItem.Ethic,
							Character:                yloKsecItem.Character,
						}

						if err = u.CommonRepository.Create(tx, &yloKsecBody); err != nil {
							return err
						}

						oldYloKsec := pointer.GetUint(yloKsecItem.ID)
						mapOldYLOKsecToNewOne[oldYloKsec] = pointer.GetUint(yloKsecBody.ID)

						yloWithKsecQuery := query.ProgramYloWithKsecQueryEntity{
							ProgramYearAndSemesterID: yloKsecItem.ProgramYearAndSemesterID,
						}

						yloWithKsecList := []query.ProgramYloWithKsecQueryEntity{}
						if err = u.CommonRepository.GetList(&yloWithKsecQuery, &yloWithKsecList, nil); err != nil {
							return err
						}

						for _, yloWithKsecItem := range yloWithKsecList {
							oldMapPloWithKsec := pointer.GetUint(yloWithKsecItem.ProgramMapPloWithKsecID)
							newMapPloWithKsec := mapOldMapKsecPloIdToNewOne[oldMapPloWithKsec]
							yloWithKsecBody := query.ProgramYloWithKsecQueryEntity{
								ProgramYearAndSemesterID: pointer.ToUint(newYearID),
								ProgramMapPloWithKsecID:  pointer.ToUint(newMapPloWithKsec),
								Remark:                   yloWithKsecItem.Remark,
								IsChecked:                yloWithKsecItem.IsChecked,
							}

							if err = u.CommonRepository.Create(tx, &yloWithKsecBody); err != nil {
								return err
							}
						}

						yloWithPloQuery := query.ProgramYloWithPloQueryEntity{
							ProgramYearAndSemesterID: yloKsecItem.ProgramYearAndSemesterID,
						}

						yloWithPloList := []query.ProgramYloWithPloQueryEntity{}
						if err = u.CommonRepository.GetList(&yloWithPloQuery, &yloWithPloList, nil); err != nil {
							return err
						}

						for _, yloWithPloItem := range yloWithPloList {
							oldPloId := pointer.GetUint(yloWithPloItem.ProgramPloID)
							newPloId := mapOldPloIdToNewOne[oldPloId]
							yloWithPloBody := query.ProgramYloWithPloQueryEntity{
								ProgramYearAndSemesterID: pointer.ToUint(newYearID),
								ProgramPloID:             pointer.ToUint(newPloId),
								Remark:                   yloWithPloItem.Remark,
								IsChecked:                yloWithPloItem.IsChecked,
							}

							if err = u.CommonRepository.Create(tx, &yloWithPloBody); err != nil {
								return err
							}
						}
					}
				}
			}
		}

		return nil
	}
}

func (u programUsecase) DuplicatePlanAndEvaluate(newProgramID uuid.UUID, mapOldSubPlanToNewOne map[uint]uint, programQuery query.ProgramMainQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		for _, majorItem := range programQuery.ProgramGeneralDetail.ProgramMajor {
			for _, planItem := range majorItem.ProgramPlanDetail {
				for _, subPlanItem := range planItem.ProgramSubPlan {
					oldSubPlan := pointer.GetUint(subPlanItem.ID)
					newSubPlan := pointer.ToUint(mapOldSubPlanToNewOne[oldSubPlan])

					programPlanAndEvaQuery := query.ProgramPlanAndEvaluateQueryEntity{
						ProgramSubPlanID: subPlanItem.ID,
					}

					programPlanAndEva := []query.ProgramPlanAndEvaluateQueryEntity{}

					if err = u.CommonRepository.GetList(&programPlanAndEvaQuery, &programPlanAndEva, nil); err != nil {
						return
					}

					for _, planAndEvaItem := range programPlanAndEva {
						programPlanAndEvaBody := query.ProgramPlanAndEvaluateQueryEntity{
							ProgramSubPlanID:      newSubPlan,
							StudentCharacteristic: planAndEvaItem.StudentCharacteristic,
							ReceiveStudentPlan:    planAndEvaItem.ReceiveStudentPlan,
							ProgramIncome:         planAndEvaItem.ProgramIncome,
							ProgramOutcome:        planAndEvaItem.ProgramOutcome,
							AcademicEvaluation:    planAndEvaItem.AcademicEvaluation,
							GraduationCriteria:    planAndEvaItem.GraduationCriteria,
						}
						u.CommonRepository.Create(tx, &programPlanAndEvaBody)

						mainQuery := query.ProgramMainQueryEntity{
							ID: &newProgramID,
						}

						mainUpdate := query.ProgramMainQueryEntity{
							ProgramPlanAndEvaluateID: programPlanAndEvaBody.ID,
						}

						err = u.CommonRepository.Update(tx, &mainQuery, &mainUpdate)

						if err != nil {
							return err
						}
					}
				}
			}
		}

		return nil
	}
}

func (u programUsecase) DuplicateMapCurMapResp(mapOldSubPlanToNewOne map[uint]uint, programQuery query.ProgramMainQueryEntity, mapOldCourseToNewOne, mapOldPloIdToNewOne, maoOldKsaToNewOne map[uint]uint) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		for _, majorItem := range programQuery.ProgramGeneralDetail.ProgramMajor {
			for _, planItem := range majorItem.ProgramPlanDetail {
				for _, subPlanItem := range planItem.ProgramSubPlan {
					oldSubPlan := pointer.GetUint(subPlanItem.ID)
					newSubPlan := pointer.ToUint(mapOldSubPlanToNewOne[oldSubPlan])

					programMapCurMapRespQuery := query.ProgramMapCurMapRespQueryEntity{
						ProgramSubPlanID: subPlanItem.ID,
					}

					programMapCurMapResp := []query.ProgramMapCurMapRespQueryEntity{}

					if err = u.CommonRepository.GetList(&programMapCurMapRespQuery, &programMapCurMapResp, nil); err != nil {
						return
					}

					for _, mapCurMapRespItem := range programMapCurMapResp {
						oldCourseDetail := pointer.GetUint(mapCurMapRespItem.ProgramCourseDetailID)
						newCourseDetail := mapOldCourseToNewOne[oldCourseDetail]
						oldPloId := pointer.GetUint(mapCurMapRespItem.ProgramPloID)
						newPloId := mapOldPloIdToNewOne[oldPloId]
						mapCurMapRespBody := query.ProgramMapCurMapRespQueryEntity{
							ProgramSubPlanID:      newSubPlan,
							ProgramCourseDetailID: pointer.ToUint(newCourseDetail),
							ProgramPloID:          pointer.ToUint(newPloId),
							Status:                mapCurMapRespItem.Status,
						}
						if err = u.CommonRepository.Create(tx, &mapCurMapRespBody); err != nil {
							return
						}
					}
				}
			}
		}

		return nil
	}
}

func (u programUsecase) DuplicateMapCurMapKsa(mapOldSubPlanToNewOne map[uint]uint, programQuery query.ProgramMainQueryEntity, mapOldCourseToNewOne, mapOldPloIdToNewOne, maoOldKsaToNewOne map[uint]uint) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		var ksaIDList []int
		var newKsaIDList []uint

		for _, majorItem := range programQuery.ProgramGeneralDetail.ProgramMajor {
			for _, planItem := range majorItem.ProgramPlanDetail {
				for _, subPlanItem := range planItem.ProgramSubPlan {
					oldSubPlan := pointer.GetUint(subPlanItem.ID)
					newSubPlan := pointer.ToUint(mapOldSubPlanToNewOne[oldSubPlan])

					programMapCurMapKsaQuery := query.ProgramMapCurMapKsaQueryEntity{
						ProgramSubPlanID: subPlanItem.ID,
					}

					programMapCurMapKsa := []query.ProgramMapCurMapKsaQueryEntity{}

					if err = u.CommonRepository.GetList(&programMapCurMapKsaQuery, &programMapCurMapKsa, nil); err != nil {
						return
					}

					for _, mapCurMapKsaItem := range programMapCurMapKsa {
						err = json.Unmarshal([]byte(*mapCurMapKsaItem.KsaID), &ksaIDList)
						if err != nil {
							return
						}

						for _, ksaID := range ksaIDList {
							oldKsaId := uint(ksaID)
							newKsaId := maoOldKsaToNewOne[oldKsaId]
							newKsaIDList = append(newKsaIDList, newKsaId)
						}

						oldCourseDetail := pointer.GetUint(mapCurMapKsaItem.ProgramCourseDetailID)
						newCourseDetail := mapOldCourseToNewOne[oldCourseDetail]
						oldPloId := pointer.GetUint(mapCurMapKsaItem.ProgramPloID)
						newPloId := mapOldPloIdToNewOne[oldPloId]
						ksaIdListStr := fmt.Sprintf("%v", newKsaIDList)
						mapCurMapRespBody := query.ProgramMapCurMapKsaQueryEntity{
							ProgramSubPlanID:      newSubPlan,
							ProgramCourseDetailID: pointer.ToUint(newCourseDetail),
							ProgramPloID:          pointer.ToUint(newPloId),
							KsaID:                 pointer.ToString(ksaIdListStr),
						}
						if err = u.CommonRepository.Create(tx, &mapCurMapRespBody); err != nil {
							return
						}
					}
				}
			}
		}

		return nil
	}
}

func (u programUsecase) DuplicateKsaDetail(mapOldSubPlanToNewOne map[uint]uint, programQuery query.ProgramMainQueryEntity, maoOldKsaToNewOne map[uint]uint) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		for _, majorItem := range programQuery.ProgramGeneralDetail.ProgramMajor {
			for _, planItem := range majorItem.ProgramPlanDetail {
				for _, subPlanItem := range planItem.ProgramSubPlan {
					oldSubPlan := pointer.GetUint(subPlanItem.ID)
					newSubPlan := pointer.ToUint(mapOldSubPlanToNewOne[oldSubPlan])

					programKsaDetailQuery := query.ProgramKsaDetailQueryEntity{
						ProgramSubPlanID: subPlanItem.ID,
					}

					programKsaDetails := []query.ProgramKsaDetailQueryEntity{}

					if err = u.CommonRepository.GetList(&programKsaDetailQuery, &programKsaDetails, nil); err != nil {
						return
					}

					for _, ksaDetail := range programKsaDetails {
						ksaDetailBody := query.ProgramKsaDetailQueryEntity{
							ProgramSubPlanID: newSubPlan,
							KsaType:          ksaDetail.KsaType,
							Order:            ksaDetail.Order,
							ShortCode:        ksaDetail.ShortCode,
							KsaDetail:        ksaDetail.KsaDetail,
						}

						if err = u.CommonRepository.Create(tx, &ksaDetailBody); err != nil {
							return
						}

						oldKsaDetail := pointer.GetUint(ksaDetail.ID)
						maoOldKsaToNewOne[oldKsaDetail] = pointer.GetUint(ksaDetailBody.ID)
					}

				}
			}
		}

		return nil
	}
}

func (u programUsecase) duplicateReference(tx *gorm.DB, programQuery query.ProgramMainQueryEntity, newProgramID *uuid.UUID) (err error) {
	programReferenceQuery := query.ProgramReferenceQueryEntity{
		ProgramID: programQuery.ID,
	}

	programRefer := []query.ProgramReferenceQueryEntity{}

	if err = u.CommonRepository.GetList(&programReferenceQuery, &programRefer, nil); err != nil {
		return err
	}

	for _, referItem := range programRefer {
		referItem.ID = nil
		referItem.ProgramID = newProgramID
		referItem.CreatedAt = nil
		referItem.UpdatedAt = nil
		u.CommonRepository.Create(tx, &referItem)
	}

	return nil
}

func (u programUsecase) duplicatePolicyAndStrategic(tx *gorm.DB, newProgramID uuid.UUID, programPolicyAndStrategic *query.ProgramPolicyAndStrategicQueryEntity) (err error) {
	pasData := query.ProgramPolicyAndStrategicQueryEntity{
		ProgramPhilosophy: programPolicyAndStrategic.ProgramPhilosophy,
		ProgramObjective:  programPolicyAndStrategic.ProgramObjective,
		ProgramPolicy:     programPolicyAndStrategic.ProgramPolicy,
		ProgramStrategic:  programPolicyAndStrategic.ProgramStrategic,
		ProgramRisk:       programPolicyAndStrategic.ProgramRisk,
		ProgramFeedback:   programPolicyAndStrategic.ProgramFeedback,
	}

	err = u.CommonRepository.Create(tx, &pasData)
	if err != nil {
		return err
	}

	mainQuery := query.ProgramMainQueryEntity{
		ID: &newProgramID,
	}

	mainUpdate := query.ProgramMainQueryEntity{
		ProgramPolicyAndStrategicID: pasData.ID,
	}

	err = u.CommonRepository.Update(tx, &mainQuery, &mainUpdate)

	if err != nil {
		return err
	}

	return nil
}

func (u programUsecase) duplicateOwner(tx *gorm.DB, newProgramID uuid.UUID, programOwner []query.ProgramOwnerQueryEntity) (err error) {
	for _, ownerItem := range programOwner {
		ownerQuery := query.ProgramOwnerQueryEntity{
			ProgramMainUID: &newProgramID,
			OwnerID:        ownerItem.OwnerID,
			Owner:          ownerItem.Owner,
		}

		err = u.CommonRepository.Create(tx, &ownerQuery)
		if err != nil {
			return err
		}
	}

	return nil
}

func (u programUsecase) duplicateLecturer(tx *gorm.DB, newProgramID uuid.UUID, programLecturer []query.ProgramLecturerQueryEntity) (err error) {
	for _, lecturerItem := range programLecturer {
		lecturerQuery := query.ProgramLecturerQueryEntity{
			ProgramMainUID: &newProgramID,
			LecturerID:     lecturerItem.LecturerID,
			Lecturer:       lecturerItem.Lecturer,
		}

		err = u.CommonRepository.Create(tx, &lecturerQuery)
		if err != nil {
			return err
		}
	}

	return nil
}

func (u programUsecase) duplicateThesisLecturer(tx *gorm.DB, newProgramID uuid.UUID, programThesisLecturer []query.ProgramThesisLecturerQueryEntity) (err error) {
	for _, thesisItem := range programThesisLecturer {
		thesisQuery := query.ProgramThesisLecturerQueryEntity{
			ProgramMainUID:   &newProgramID,
			ThesisLecturerID: thesisItem.ThesisLecturerID,
			ThesisLecturer:   thesisItem.ThesisLecturer,
		}

		err = u.CommonRepository.Create(tx, &thesisQuery)
		if err != nil {
			return err
		}
	}

	return nil
}

func (u programUsecase) duplicateQualityAssurance(tx *gorm.DB, newProgramID uuid.UUID, programQualityAssurance *query.ProgramQualityAssuranceQueryEntity) (err error) {
	qaQuery := query.ProgramQualityAssuranceQueryEntity{
		IsHescCheck:      programQualityAssurance.IsHescCheck,
		HescDescription:  programQualityAssurance.HescDescription,
		IsAunQaCheck:     programQualityAssurance.IsAunQaCheck,
		AunQaDescription: programQualityAssurance.AunQaDescription,
		IsAbetCheck:      programQualityAssurance.IsAbetCheck,
		AbetDescription:  programQualityAssurance.AbetDescription,
		IsWfmeCheck:      programQualityAssurance.IsWfmeCheck,
		WfmeDescription:  programQualityAssurance.WfmeDescription,
		IsAacsbCheck:     programQualityAssurance.IsAacsbCheck,
		AacsbDescription: programQualityAssurance.AacsbDescription,
	}

	err = u.CommonRepository.Create(tx, &qaQuery)
	if err != nil {
		return err
	}

	mainQuery := query.ProgramMainQueryEntity{
		ID: &newProgramID,
	}

	mainUpdate := query.ProgramMainQueryEntity{
		ProgramQualityAssuranceID: qaQuery.ID,
	}

	err = u.CommonRepository.Update(tx, &mainQuery, &mainUpdate)

	if err != nil {
		return err
	}

	return nil
}

func (u programUsecase) duplicateSystemAndMechanic(tx *gorm.DB, newProgramID uuid.UUID, programSystemAndMechanic *query.ProgramSystemAndMechanicQueryEntity) (err error) {
	samQuery := query.ProgramSystemAndMechanicQueryEntity{
		CourseExpectedAttribute: programSystemAndMechanic.CourseExpectedAttribute,
		CourseImprovingPlan:     programSystemAndMechanic.CourseImprovingPlan,
		CoursePolicies:          programSystemAndMechanic.CoursePolicies,
		CourseRisk:              programSystemAndMechanic.CourseRisk,
		CourseStrategies:        programSystemAndMechanic.CourseStrategies,
		CourseStudentComment:    programSystemAndMechanic.CourseStudentComment,
		MainContentAndStructure: programSystemAndMechanic.MainContentAndStructure,
	}

	err = u.CommonRepository.Create(tx, &samQuery)
	if err != nil {
		return err
	}

	mainQuery := query.ProgramMainQueryEntity{
		ID: &newProgramID,
	}

	mainUpdate := query.ProgramMainQueryEntity{
		ProgramSystemAndMechanicID: samQuery.ID,
	}

	err = u.CommonRepository.Update(tx, &mainQuery, &mainUpdate)

	if err != nil {
		return err
	}

	return nil
}

func (u programUsecase) duplicateProgramPermission(tx *gorm.DB, newProgramID uuid.UUID, permissionProgram []permissionQuery.PermissionProgramQueryEntity) (err error) {
	for _, permissionItem := range permissionProgram {
		thesisQuery := permissionQuery.PermissionProgramQueryEntity{
			UserUID:       permissionItem.UserUID,
			ProgramUID:    &newProgramID,
			Accessibility: permissionItem.Accessibility,
		}

		err = u.CommonRepository.Create(tx, &thesisQuery)
		if err != nil {
			return err
		}
	}

	return nil
}

func (u programUsecase) DuplicateCLO(mapOldPlanToNewOne map[uint]uint, programQuery query.ProgramMainQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		for _, majorItem := range programQuery.ProgramGeneralDetail.ProgramMajor {
			for _, planItem := range majorItem.ProgramPlanDetail {
				oldPlan := pointer.GetUint(planItem.ID)
				newPlan := pointer.ToUint(mapOldPlanToNewOne[oldPlan])

				programCLOQuery := query.ProgramCLOQueryEntity{
					PlanID: planItem.ID,
				}

				programCLO := []query.ProgramCLOQueryEntity{}

				if err = u.CommonRepository.GetList(&programCLOQuery, &programCLO, nil); err != nil {
					return
				}

				for _, cloItem := range programCLO {
					cloItem.ID = nil
					cloItem.PlanID = newPlan
					cloItem.CreatedAt = nil
					cloItem.UpdatedAt = nil
					u.CommonRepository.Create(tx, &cloItem)
				}

			}
		}

		return nil
	}
}

func (u programUsecase) DuplicateApprovals(newProgramID uuid.UUID) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		if err := u.InitProgramApprovals(newProgramID, tx); err != nil {
			log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
			return err
		}
		return
	}
}

func (u programUsecase) CreateProgramOwnerPermission(userUID uuid.UUID, newProgramID uuid.UUID) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		role := rapQuery.RoleQueryEntity{
			RoleNameEN: pointer.ToString(permissionConstant.ROLE_PROGRAM_OWNER),
		}

		u.CommonRepository.GetFirst(&role)

		programPermission := rapQuery.MapProgramsRolesQueryEntity{
			UserID:    &userUID,
			ProgramID: &newProgramID,
			RoleID:    role.ID,
		}
		if err = u.CommonRepository.Create(tx, &programPermission); err != nil {
			log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
			return err
		}
		return
	}
}

func (u programUsecase) CreateProgramLecturerPermission(userUID uuid.UUID, newProgramID uuid.UUID) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		role := rapQuery.RoleQueryEntity{
			RoleNameEN: pointer.ToString(permissionConstant.ROLE_PROGRAM_LECTURER),
		}

		u.CommonRepository.GetFirst(&role)

		programPermission := rapQuery.MapProgramsRolesQueryEntity{
			UserID:    &userUID,
			ProgramID: &newProgramID,
			RoleID:    role.ID,
		}
		if err = u.CommonRepository.Create(tx, &programPermission); err != nil {
			log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
			return err
		}
		return
	}
}
