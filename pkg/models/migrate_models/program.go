package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProgramMain struct {
	ID                          uuid.UUID                 `gorm:"column:id;primaryKey;type:uuid" json:"id"`
	ProgramGeneralDetail        ProgramGeneralDetail      `gorm:"foreignKey:ProgramGeneralDetailID;references:ID"`
	ProgramGeneralDetailID      uint                      `gorm:"column:program_general_detail_id;type:bigint" json:"program_general_detail_id"`
	ProgramPolicyAndStrategic   ProgramPolicyAndStrategic `gorm:"foreignKey:ProgramPolicyAndStrategicID;references:ID"`
	ProgramPolicyAndStrategicID uint                      `gorm:"column:program_policy_and_strategic_id;type:bigint" json:"program_policy_and_strategic_id"`
	ProgramPlanAndEvaluate      ProgramPlanAndEvaluate    `gorm:"foreignKey:ProgramPlanAndEvaluateID;references:ID"`
	ProgramPlanAndEvaluateID    uint                      `gorm:"column:program_plan_and_evaluate_id;type:bigint" json:"program_plan_and_evaluate_id"`
	ProgramOwner                []ProgramOwner            `gorm:"foreignKey:ProgramMainUID;references:ID"`
	ProgramLecturer             []ProgramLecturer         `gorm:"foreignKey:ProgramMainUID;references:ID"`
	ProgramThesisLecturer       []ProgramThesisLecturer   `gorm:"foreignKey:ProgramMainUID;references:ID"`
	ProgramQualityAssurance     ProgramQualityAssurance   `gorm:"foreignKey:ProgramQualityAssuranceID;references:ID"`
	ProgramQualityAssuranceID   uint                      `gorm:"column:program_quality_assurance_id;type:bigint" json:"program_quality_assurance_id"`
	ProgramSystemAndMechanic    ProgramSystemAndMechanic  `gorm:"foreignKey:ProgramSystemAndMechanicID;references:ID"`
	ProgramSystemAndMechanicID  uint                      `gorm:"column:program_system_and_mechanic_id;type:bigint" json:"program_system_and_mechanic_id"`
	MapProgramRoles             []MapProgramsRoles        `gorm:"foreignKey:ProgramID;references:ID"`
	CreatedAt                   time.Time                 `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt                   time.Time                 `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt                   gorm.DeletedAt            `gorm:"index;column:deleted_at"`
	// ProgramPermission           []PermissionProgram       `gorm:"foreignKey:ID;references:ProgramUID"`
}

// TableName specifies the table name for the Program struct
func (ProgramMain) TableName() string {
	return "program_main"
}

type ProgramGeneralDetail struct {
	ID                    uint           `gorm:"column:id;type:bigint;primaryKey;autoIncrement" json:"id"`
	UniversityName        string         `gorm:"column:university_name;type:varchar(255)" json:"university_name"`
	FacultyID             uint           `gorm:"column:faculty_id;type:uint" json:"faculty_id"`
	Faculty               Faculty        `gorm:"foreignKey:FacultyID;references:ID"`
	ProgramNameTH         string         `gorm:"column:program_name_th;type:varchar(255)" json:"program_name_th"`
	ProgramCode           string         `gorm:"column:program_code;type:varchar(63)" json:"program_code"`
	ProgramNameEN         string         `gorm:"column:program_name_en;type:varchar(255)" json:"program_name_en"`
	BranchNameTH          string         `gorm:"column:branch_name_th;type:varchar(255)" json:"branch_name_th"`
	BranchNameEN          string         `gorm:"column:branch_name_en;type:varchar(255)" json:"branch_name_en"`
	DegreeNameTH          string         `gorm:"column:degree_name_th;type:varchar(255)" json:"degree_name_th"`
	DegreeNameEN          string         `gorm:"column:degree_name_en;type:varchar(255)" json:"degree_name_en"`
	DegreeNameShortenTH   string         `gorm:"column:degree_name_shorten_th;type:varchar(255)" json:"degree_name_shorten_th"`
	DegreeNameShortenEN   string         `gorm:"column:degree_name_shorten_en;type:varchar(255)" json:"degree_name_shorten_en"`
	OverallCredit         int            `gorm:"column:overall_credit;type:int" json:"overall_credit"`
	ProgramMajorTypeID    int            `gorm:"column:program_major_type_id;type:int" json:"program_major_type_id"`
	ProgramMajorType      string         `gorm:"column:program_major_type;type:varchar(50)" json:"program_major_type"`
	ProgramDegreeTypeID   int            `gorm:"column:program_degree_type_id;type:int" json:"program_degree_type_id"`
	ProgramDegreeType     string         `gorm:"column:program_degree_type;type:varchar(50)" json:"program_degree_type"`
	NumberOfYear          int            `gorm:"column:number_of_year;type:int" json:"number_of_year"`
	ProgramLanguageID     int            `gorm:"column:program_language_id;type:int" json:"program_language_id"`
	ProgramLanguage       string         `gorm:"column:program_language;type:varchar(100)" json:"program_language"`
	Admission             string         `gorm:"column:admission;type:varchar(100)" json:"admission"`
	MOU                   string         `gorm:"column:mou;type:varchar(100)" json:"mou"`
	MOUFilepath           string         `gorm:"column:mou_filepath;type:varchar(255)" json:"mou_filepath"`
	ProgramTypeID         int            `gorm:"column:program_type_id;type:int" json:"program_type_id"`
	ProgramType           string         `gorm:"column:program_type;type:varchar(100)" json:"program_type"`
	ProgramYearID         int            `gorm:"column:program_year_id;type:int" json:"program_year_id"`
	ProgramYear           uint           `gorm:"column:program_year;type:uint" json:"program_year"`
	Semester              int            `gorm:"column:semester;type:int" json:"semester"`
	SemesterYear          int            `gorm:"column:semester_year;type:int" json:"semester_year"`
	BoardApproval         int            `gorm:"column:board_approval;type:int" json:"board_approval"`
	BoardApprovalDate     time.Time      `gorm:"column:board_approval_date;type:date" json:"board_approval_date"`
	AcademicCouncil       int            `gorm:"column:academic_council;type:int" json:"academic_council"`
	AcademicCouncilDate   time.Time      `gorm:"column:academic_council_date;type:date" json:"academic_council_date"`
	UniversityCouncil     int            `gorm:"column:university_council;type:int" json:"university_council"`
	UniversityCouncilDate time.Time      `gorm:"column:university_council_date;type:date" json:"university_council_date"`
	IsSamePlanMajor       bool           `gorm:"column:is_same_plan_major;type:boolean" json:"is_same_plan_major"`
	ProgramMajor          []ProgramMajor `gorm:"foreignKey:ProgramGeneralDetailID;references:ID"`
	IsNationalProgram     bool           `gorm:"column:is_national_program;type:boolean" json:"is_national_program"`
	IsEnglishProgram      bool           `gorm:"column:is_english_program;type:boolean" json:"is_english_program"`
	IsOther               bool           `gorm:"column:is_other;type:boolean" json:"is_other"`
	OtherName             string         `gorm:"column:other_name;type:varchar" json:"other_name"`
	ProgramAdjustFrom     string         `gorm:"column:program_adjust_from;type:varchar" json:"program_adjust_from"`
	ProgramAdjustYear     int            `gorm:"column:program_adjust_year;type:int" json:"program_adjust_year"`
	CreatedAt             time.Time      `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt             time.Time      `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt             gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

// TableName sets the table name for the Catagory1 model
func (ProgramGeneralDetail) TableName() string {
	return "program_general_detail"
}

type ProgramMajor struct {
	ID                     uint                `gorm:"column:id;type:bigint;primaryKey;autoIncrement" json:"id"`
	ProgramGeneralDetailID uint                `gorm:"column:program_general_detail_id;type:uint"`
	Name                   string              `gorm:"column:name;type:varchar(255)" json:"major"`
	ProgramPlanDetail      []ProgramPlanDetail `gorm:"foreignKey:ProgramMajorID;references:ID"`
	CreatedAt              time.Time           `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt              time.Time           `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt              gorm.DeletedAt      `gorm:"index;column:deleted_at"`
}

// TableName sets the table name for the MajorAndPlan model
func (q *ProgramMajor) TableName() string {
	return "program_major"
}

type ProgramPlanDetail struct {
	ID             uint             `gorm:"column:id;type:bigint;primaryKey;autoIncrement" json:"id"`
	ProgramMajorID uint             `gorm:"column:program_major_id;type:uint"`
	PlanName       string           `gorm:"column:plan_name;type:string" json:"plan_name"`
	CreditRulesID  int              `gorm:"column:credit_rules_id;type:int" json:"credit_rules_id"`
	CreditRules    string           `gorm:"column:credit_rules;type:varchar(255)" json:"credit_rules"`
	Credit         uint16           `gorm:"column:credit;type:smallint" json:"credit"`
	IsSplitPlan    bool             `gorm:"column:is_split_plan;type:boolean" json:"is_split_plan"`
	IsActive       bool             `gorm:"column:is_active;type:boolean" json:"is_active"`
	SubPlan        []ProgramSubPlan `gorm:"foreignKey:ProgramPlanDetailID;references:ID"`
	CreatedAt      time.Time        `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt      time.Time        `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt      gorm.DeletedAt   `gorm:"index;column:deleted_at"`
}

// TableName sets the table name for the Plan model
func (q *ProgramPlanDetail) TableName() string {
	return "program_plan_detail"
}

type ProgramSubPlan struct {
	ID                  uint                     `gorm:"column:id;type:bigint;primaryKey;autoIncrement" json:"id"`
	ProgramPlanDetailID uint                     `gorm:"column:program_plan_detail_id;type:uint"`
	SubPlanName         string                   `gorm:"column:sub_plan_name;type:string" json:"sub_plan_name"`
	CreditRulesID       int                      `gorm:"column:credit_rules_id;type:int" json:"credit_rules_id"`
	CreditRules         string                   `gorm:"column:credit_rules;type:varchar(255)" json:"credit_rules"`
	Credit              uint16                   `gorm:"column:credit;type:smallint" json:"credit"`
	CourseDetail        []ProgramCourseDetail    `gorm:"foreignKey:ProgramSubPlanID;references:ID"` //has child ref
	StructureDetail     []ProgramStructureDetail `gorm:"foreignKey:ProgramSubPlanID;references:ID"` //has child ref
	CreatedAt           time.Time                `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt           time.Time                `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt           gorm.DeletedAt           `gorm:"index;column:deleted_at"`
}

// TableName sets the table name for the Plan model
func (q *ProgramSubPlan) TableName() string {
	return "program_sub_plan"
}

type ProgramPolicyAndStrategic struct {
	ID                uint           `gorm:"column:id;type:bigint;primaryKey;autoIncrement" json:"id"`
	ProgramPhilosophy string         `gorm:"column:program_philosophy;type:varchar" json:"program_philosophy"`
	ProgramObjective  string         `gorm:"column:program_objective;type:varchar" json:"program_objective"`
	ProgramPolicy     string         `gorm:"column:program_policy;type:varchar" json:"program_policy"`
	ProgramStrategic  string         `gorm:"column:program_strategic;type:varchar" json:"program_strategic"`
	ProgramRisk       string         `gorm:"column:program_risk;type:varchar" json:"program_risk"`
	ProgramFeedback   string         `gorm:"column:program_feedback;type:varchar" json:"program_feedback"`
	CreatedAt         time.Time      `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt         time.Time      `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt         gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

func (ProgramPolicyAndStrategic) TableName() string {
	return "program_policy_and_strategic"
}

type ProgramStructureDetail struct {
	ID               uint                     `gorm:"column:id;primaryKey;autoIncrement;type:uint;index"`
	ProgramSubPlanID uint                     `gorm:"column:program_sub_plan_id;type:uint"`
	ProgramSubPlan   ProgramSubPlan           `gorm:"foreignKey:ProgramSubPlanID;references:ID"`
	Name             string                   `gorm:"column:name;type:varchar;size:255"`
	Order            uint                     `gorm:"column:order;default:99;type:uint"`
	ParentID         *uint                    `gorm:"column:parent_id;type:uint;default:null"`
	Children         []ProgramStructureDetail `gorm:"foreignKey:ParentID;references:ID"`
	CourseDetail     []ProgramCourseDetail    `gorm:"foreignKey:ProgramStructureID;references:ID"` //has child ref
	Qualification    string                   `gorm:"column:qualification;type:varchar;size:255"`
	StructureCredit  uint                     `gorm:"column:structure_credit;type:uint"`
	CreatedAt        time.Time                `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt        time.Time                `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt        gorm.DeletedAt           `gorm:"index;column:deleted_at"`
}

func (ProgramStructureDetail) TableName() string {
	return "program_structure_detail"
}

type ProgramCourseDetail struct {
	ID                  uint           `gorm:"column:id;primaryKey;autoIncrement;index;type:uint"`
	YearAndSemesterID   uint           `gorm:"column:year_and_semester_id;type:uint"` // has many on parent
	ProgramStructureID  uint           `gorm:"column:program_structure_id;type:uint"` // has many on parent
	ProgramSubPlanID    uint           `gorm:"column:program_sub_plan_id;type:uint"`  // has many on parent
	CourseSource        string         `gorm:"column:course_source;"`                 // REG || Course
	REGKkuKey           *string        `gorm:"column:reg_kku_key;"`
	CourseKey           *uuid.UUID     `gorm:"column:course_key;"`
	CourseTypeID        int            `gorm:"column:course_type_id;type:int;"`
	CourseType          string         `gorm:"column:course_type;type:varchar;size:255"`
	CourseCode          string         `gorm:"column:course_code;type:varchar;size:255"`
	CourseYear          string         `gorm:"column:course_year;type:varchar;size:255"`
	CourseNameTH        string         `gorm:"column:course_name_th;type:varchar;size:255"`
	CourseNameEN        string         `gorm:"column:course_name_en;type:varchar;size:255"`
	Version             string         `gorm:"column:course_version;type:varchar;size:255"`
	CourseCredit        uint           `gorm:"column:course_credit;type:uint"`
	Credit1             uint           `gorm:"column:credit_1;type:int;not null;default:0"`
	Credit2             uint           `gorm:"column:credit_2;type:int;not null;default:0"`
	Credit3             uint           `gorm:"column:credit_3;type:int;not null;default:0"`
	CourseConditionTH   string         `gorm:"column:course_condition_th;type:varchar;size:255"`
	CourseConditionEN   string         `gorm:"column:course_condition_en;type:varchar;size:255"`
	CourseDescriptionEN string         `gorm:"column:course_description_th;type:varchar;size:255"`
	CourseDescriptionTH string         `gorm:"column:course_description_en;type:varchar;size:255"`
	CourseObjective     string         `gorm:"column:course_objective;type:varchar;size:255"`
	IsCreditCalc        bool           `gorm:"column:is_credit_calc;type:bool"`
	IsEditedCourse      bool           `gorm:"column:is_edited_course;type:bool;default:false;"`
	IsNewCourse         bool           `gorm:"column:is_new_course;type:bool;default:false;"`
	CreatedAt           time.Time      `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt           time.Time      `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt           gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

func (ProgramCourseDetail) TableName() string {
	return "program_course_detail"
}

type ProgramYearAndSemester struct {
	ID               uint                  `gorm:"primaryKey;autoIncrement;type:uint;index;"`
	ProgramSubPlan   ProgramSubPlan        `gorm:"foreignKey:ProgramSubPlanID;references:ID"`
	ProgramSubPlanID uint                  `gorm:"column:program_sub_plan_id;type:uint"`
	CourseDetail     []ProgramCourseDetail `gorm:"foreignKey:YearAndSemesterID;references:ID"` //has child ref
	Year             string                `gorm:"column:year;type:varchar;size:20;not null"`
	Semester         string                `gorm:"column:semester;type:varchar;size:20;not null"`
	CreatedAt        time.Time             `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt        time.Time             `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt        gorm.DeletedAt        `gorm:"index;column:deleted_at"`
}

func (ProgramYearAndSemester) TableName() string {
	return "program_year_and_semester"
}

type ProgramPLOFormat struct {
	ID               uint           `gorm:"primaryKey;autoIncrement"`
	ProgramSubPlanID uint           `gorm:"column:program_sub_plan_id;type:uint"`
	ProgramSubPlan   ProgramSubPlan `gorm:"foreignKey:ProgramSubPlanID;references:ID"`
	PLOFormat        string         `gorm:"column:plo_format;type:text"`
	ProgramPLO       []ProgramPlo   `gorm:"foreignKey:ProgramPloFormatID;references:ID"`
	CreatedAt        time.Time      `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt        time.Time      `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt        gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

func (ProgramPLOFormat) TableName() string {
	return "program_plo_format"
}

type ProgramPlo struct {
	ID                    uint                           `gorm:"column:id;primaryKey;autoIncrement"`
	ProgramPloFormatID    uint                           `gorm:"column:program_plo_format_id;type:uint"`
	Order                 uint                           `gorm:"column:order;default:99;type:uint"`
	ParentID              *uint                          `gorm:"column:parent_id;type:uint;default:null"`
	Children              []ProgramPlo                   `gorm:"foreignKey:ParentID;references:ID"`
	PLOPrefix             string                         `gorm:"column:plo_prefix;type:varchar;size:255"`
	PLODetail             string                         `gorm:"column:plo_detail;type:varchar;size:255"`
	LearningSolution      []ProgramPLOLearningSolution   `gorm:"foreignKey:PloID;references:ID"`
	LearningEvaluation    []ProgramPLOLearningEvaluation `gorm:"foreignKey:PloID;references:ID"`
	ProgramMapPloWithKsec []ProgramMapPloWithKsec        `gorm:"foreignKey:PloID;references:ID"`
	CreatedAt             time.Time                      `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt             time.Time                      `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt             gorm.DeletedAt                 `gorm:"index;column:deleted_at"`
}

func (ProgramPlo) TableName() string {
	return "program_plo"
}

type ProgramKsecDetail struct {
	ID                    uint                    `gorm:"column:id;primaryKey;autoIncrement"`
	ProgramSubPlanID      uint                    `gorm:"column:program_sub_plan_id;type:uint"`
	ProgramSubPlan        ProgramSubPlan          `gorm:"foreignKey:ProgramSubPlanID;references:ID"`
	Type                  string                  `gorm:"column:type;type:varchar;size:2;index"`
	Order                 uint                    `gorm:"column:order;default:99;type:uint;index"`
	Detail                string                  `gorm:"column:detail;type:text"`
	ProgramMapPloWithKsec []ProgramMapPloWithKsec `gorm:"foreignKey:KsecID;references:ID"`
	CreatedAt             time.Time               `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt             time.Time               `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt             gorm.DeletedAt          `gorm:"index;column:deleted_at"`
}

func (ProgramKsecDetail) TableName() string {
	return "program_ksec_detail"
}

type ProgramMapPloWithKsec struct {
	ID        uint              `gorm:"column:id;primaryKey;autoIncrement"`
	PloID     uint              `gorm:"column:plo_id;type:uint;not null"`
	KsecID    uint              `gorm:"column:ksec_id;type:uint;not null"`
	PLO       ProgramPlo        `gorm:"foreignKey:PloID;references:ID"`
	KSEC      ProgramKsecDetail `gorm:"foreignKey:KsecID;references:ID"`
	CreatedAt time.Time         `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt time.Time         `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt gorm.DeletedAt    `gorm:"index;column:deleted_at"`
}

func (ProgramMapPloWithKsec) TableName() string {
	return "program_map_plo_with_ksec"
}

type ProgramPLOLearningSolution struct {
	ID        uint           `gorm:"column:id;primaryKey;autoIncrement;index"`
	PloID     uint           `gorm:"column:plo_id;type:uint;not null;index"`
	Key       string         `gorm:"column:key;type:varchar;size:255;"`
	Detail    string         `gorm:"column:detail;type:varchar;size:255;"`
	Order     uint           `gorm:"column:order;default:99;type:uint"`
	CreatedAt time.Time      `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt time.Time      `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

func (ProgramPLOLearningSolution) TableName() string {
	return "program_plo_learning_solution"
}

type ProgramPLOLearningEvaluation struct {
	ID        uint           `gorm:"column:id;primaryKey;autoIncrement;index"`
	PloID     uint           `gorm:"column:plo_id;type:uint;not null;index"`
	Key       string         `gorm:"column:key;type:varchar;size:255;"`
	Detail    string         `gorm:"column:detail;type:varchar;size:255;"`
	Order     uint           `gorm:"column:order;default:99;type:uint"`
	CreatedAt time.Time      `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt time.Time      `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

func (ProgramPLOLearningEvaluation) TableName() string {
	return "program_plo_learning_evaluation"
}

type ProgramCompetency struct {
	ID                 uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	ProgramMainID      uuid.UUID      `gorm:"column:program_main_id;type:uuid" json:"program_main_id"`
	ProgramMain        ProgramMain    `gorm:"foreignKey:ProgramMainID;references:ID"`
	Order              int            `gorm:"column:order;default:99;type:int" json:"order"`
	SpecificCompetency string         `gorm:"column:specific_competency;type:varchar" json:"specific_competency"`
	GenericCompetency  string         `gorm:"column:generic_competency;type:varchar" json:"generic_competency"`
	CreatedAt          time.Time      `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt          time.Time      `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt          gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

// TableName sets the table name for the Catagory6 model
func (ProgramCompetency) TableName() string {
	return "program_competency"
}

type ProgramYloKsec struct {
	ID                       uint                   `gorm:"primaryKey;autoIncrement" json:"id"`
	ProgramSubPlanID         uint                   `gorm:"column:program_sub_plan_id;type:uint"`
	ProgramSubPlan           ProgramSubPlan         `gorm:"foreignKey:ProgramSubPlanID;references:ID"`
	ProgramYearAndSemesterID uint                   `gorm:"column:program_year_and_semester_id;type:uint"`
	ProgramYearAndSemester   ProgramYearAndSemester `gorm:"foreignKey:ProgramYearAndSemesterID;references:ID"`
	Knowledge                string                 `gorm:"column:knowledge;type:varchar;size:255;"`
	Skill                    string                 `gorm:"column:skill;type:varchar;size:255;"`
	Ethic                    string                 `gorm:"column:ethic;type:varchar;size:255;"`
	Character                string                 `gorm:"column:character;type:varchar;size:255;"`
	CreatedAt                time.Time              `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt                time.Time              `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt                gorm.DeletedAt         `gorm:"index;column:deleted_at"`
}

// TableName sets the table name for the YLO model
func (ProgramYloKsec) TableName() string {
	return "program_ylo_ksec"
}

type ProgramYloWithKsec struct {
	ID                       uint                   `gorm:"primaryKey;autoIncrement" json:"id"`
	ProgramYearAndSemesterID uint                   `gorm:"column:program_year_and_semester_id;type:uint"`
	ProgramYearAndSemester   ProgramYearAndSemester `gorm:"foreignKey:ProgramYearAndSemesterID;references:ID"`
	ProgramMapPloWithKsecID  uint                   `gorm:"column:program_map_plo_with_ksec_id;type:varchar;size:255;"`
	ProgramMapPloWithKsec    ProgramMapPloWithKsec  `gorm:"foreignKey:ProgramMapPloWithKsecID;references:ID"`
	Remark                   string                 `gorm:"column:remark;type:varchar;size:255;"`
	IsChecked                bool                   `gorm:"column:is_checked;type:boolean;default:true;" json:"is_checked"`
	CreatedAt                time.Time              `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt                time.Time              `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt                gorm.DeletedAt         `gorm:"index;column:deleted_at"`
}

func (ProgramYloWithKsec) TableName() string {
	return "program_ylo_with_ksec"
}

type ProgramYloWithPlo struct {
	ID                       uint                   `gorm:"primaryKey;autoIncrement" json:"id"`
	ProgramYearAndSemesterID uint                   `gorm:"column:program_year_and_semester_id;type:uint"`
	ProgramYearAndSemester   ProgramYearAndSemester `gorm:"foreignKey:ProgramYearAndSemesterID;references:ID"`
	ProgramPloID             uint                   `gorm:"column:program_plo_id;type:varchar;size:255;"`
	ProgramPlo               ProgramPlo             `gorm:"foreignKey:ProgramPloID;references:ID"`
	Remark                   string                 `gorm:"column:remark;type:varchar;size:255;"`
	IsChecked                bool                   `gorm:"column:is_checked;type:boolean;default:true" json:"is_checked"`
	CreatedAt                time.Time              `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt                time.Time              `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt                gorm.DeletedAt         `gorm:"index;column:deleted_at"`
}

func (ProgramYloWithPlo) TableName() string {
	return "program_ylo_with_plo"
}

type ProgramOwner struct {
	ID             uint            `gorm:"column:id;primaryKey;autoIncrement"`
	ProgramMainUID uuid.UUID       `gorm:"column:program_main_uid;type:uuid"`
	OwnerID        uuid.UUID       `gorm:"column:owner_id;type:uuid"`
	Owner          EmployeeDetails `gorm:"foreignKey:UID;references:OwnerID"`
	CreatedAt      time.Time       `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt      time.Time       `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt      gorm.DeletedAt  `gorm:"index;column:deleted_at"`
}

func (ProgramOwner) TableName() string {
	return "program_owner"
}

type ProgramThesisLecturer struct {
	ID               uint            `gorm:"column:id;primaryKey;autoIncrement"`
	ProgramMainUID   uuid.UUID       `gorm:"column:program_main_uid;type:uuid"`
	ThesisLecturerID uuid.UUID       `gorm:"column:thesis_lecturer_id;type:uuid"`
	ThesisLecturer   EmployeeDetails `gorm:"foreignKey:UID;references:ThesisLecturerID"`
	CreatedAt        time.Time       `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt        time.Time       `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt        gorm.DeletedAt  `gorm:"index;column:deleted_at"`
}

func (ProgramThesisLecturer) TableName() string {
	return "program_thesis_lecturer"
}

type ProgramLecturer struct {
	ID             uint            `gorm:"column:id;primaryKey;autoIncrement"`
	ProgramMainUID uuid.UUID       `gorm:"column:program_main_uid;type:uuid"`
	LecturerID     uuid.UUID       `gorm:"column:lecturer_id;type:uuid"`
	Lecturer       EmployeeDetails `gorm:"foreignKey:UID;references:LecturerID"`
	CreatedAt      time.Time       `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt      time.Time       `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt      gorm.DeletedAt  `gorm:"index;column:deleted_at"`
}

func (ProgramLecturer) TableName() string {
	return "program_lecturer"
}

type ProgramPlanAndEvaluate struct {
	ID                                uint           `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ProgramSubPlan                    ProgramSubPlan `gorm:"foreignKey:ProgramSubPlanID;references:ID"`
	ProgramSubPlanID                  uint           `gorm:"column:program_sub_plan_id;type:uint"`
	StudentCharacteristic             string         `gorm:"column:student_characteristic;type:varchar" json:"student_characteristic"`
	ReceiveStudentPlan                string         `gorm:"column:receive_student_plan;type:varchar" json:"receive_student_plan"`
	ProgramIncome                     string         `gorm:"column:program_income;type:varchar" json:"program_income"`
	ProgramOutcome                    string         `gorm:"column:program_outcome;type:varchar" json:"program_outcome"`
	AcademicEvaluation                string         `gorm:"column:academic_evaluation;type:varchar" json:"academic_evaluation"`
	GraduationCriteria                string         `gorm:"column:graduation_criteria;type:varchar" json:"graduation_criteria"`
	ProgramUniversityTransferStandard string         `gorm:"column:program_university_transfer_standard;type:varchar" json:"program_university_transfer_standard"`
	ProgramPreparation                string         `gorm:"column:program_preparation;type:varchar" json:"program_preparation"`
	CreatedAt                         time.Time      `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt                         time.Time      `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt                         gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

// TableName specifies the table name for the ProgramPlanAndEvaluate struct
func (ProgramPlanAndEvaluate) TableName() string {
	return "program_plan_and_evaluate"
}

type ProgramQualityAssurance struct {
	ID               uint           `gorm:"column:id;primaryKey;autoIncrement"`
	IsHescCheck      bool           `gorm:"column:is_hesc_check;type:boolean" json:"is_hesc_check"`
	HescDescription  string         `gorm:"column:hesc_description;type:text" json:"hesc_description"`
	IsAunQaCheck     bool           `gorm:"column:is_aun_qa_check;type:boolean" json:"is_aun_qa_check"`
	AunQaDescription string         `gorm:"column:aun_qa_description;type:text" json:"aun_qa_description"`
	IsAbetCheck      bool           `gorm:"column:is_abet_check;type:boolean" json:"is_abet_check"`
	AbetDescription  string         `gorm:"column:abet_description;type:text" json:"abet_description"`
	IsWfmeCheck      bool           `gorm:"column:is_wfme_check;type:boolean" json:"is_wfme_check"`
	WfmeDescription  string         `gorm:"column:wfme_description;type:text" json:"wfme_description"`
	IsAacsbCheck     bool           `gorm:"column:is_aacsb_check;type:boolean" json:"is_aacsb_check"`
	AacsbDescription string         `gorm:"column:aacsb_description;type:text" json:"aacsb_description"`
	CreatedAt        time.Time      `gorm:"column:created_at;type:timestamp;default:current_timestamp"`
	UpdatedAt        time.Time      `gorm:"column:updated_at;type:timestamp;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt        gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

func (ProgramQualityAssurance) TableName() string {
	return "program_quality_assurance"
}

type ProgramSystemAndMechanic struct {
	ID                      uint           `gorm:"column:id;primaryKey;autoIncrement"`
	CoursePolicies          string         `gorm:"column:course_policies;type:text"`
	CourseStrategies        string         `gorm:"column:course_strategies;type:text"`
	CourseRisk              string         `gorm:"column:course_risk;type:text"`
	CourseStudentComment    string         `gorm:"column:course_student_comment;type:text"`
	CourseExpectedAttribute string         `gorm:"column:course_expected_attribute;type:text"`
	MainContentAndStructure string         `gorm:"column:main_content_and_structure;type:text"`
	CourseImprovingPlan     string         `gorm:"column:course_improving_plan;type:text"`
	CreatedAt               time.Time      `gorm:"column:created_at;type:timestamp;default:current_timestamp"`
	UpdatedAt               time.Time      `gorm:"column:updated_at;type:timestamp;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt               gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

func (ProgramSystemAndMechanic) TableName() string {
	return "program_system_and_mechanic"
}

type ProgramReference struct {
	ID                   uint            `gorm:"column:id;primaryKey;autoIncrement"`
	ProgramID            uuid.UUID       `gorm:"column:program_id;type:uuid;not null"`
	ReferenceName        string          `gorm:"column:reference_name;type:varchar;size:255"`
	ReferenceDescription string          `gorm:"column:reference_description;type:text"`
	ReferenceFileName    string          `gorm:"column:reference_file_name;type:varchar;size:255"`
	ReferenceFilePath    string          `gorm:"column:reference_file_path;type:varchar;size:255"`
	ReferenceTypeID      int             `gorm:"column:reference_type_id;type:integer"`
	ReferenceOption      ReferenceOption `gorm:"foreignKey:ReferenceTypeID;references:ID"`
	CreatedAt            time.Time       `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt            time.Time       `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt            gorm.DeletedAt  `gorm:"index;column:deleted_at"`
}

func (ProgramReference) TableName() string {
	return "program_references"
}

type CLO struct {
	ID              uint           `gorm:"column:id;primaryKey;autoIncrement"`
	PlanID          uint           `gorm:"column:plan_id;type:bigint;autoIncrement" json:"plan_id"`
	CLOGeneralData  string         `gorm:"column:clo_general_data;type:varchar" json:"clo_general_data"`
	CLOSpecificData string         `gorm:"column:clo_specific_data;type:varchar" json:"clo_specific_data"`
	CreatedAt       time.Time      `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt       time.Time      `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt       gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

func (CLO) TableName() string {
	return "program_clo"
}

type ProgramMapCurMapResp struct {
	ID                    uint                `gorm:"column:id;primaryKey;autoIncrement"`
	ProgramSubPlan        ProgramSubPlan      `gorm:"foreignKey:ProgramSubPlanID;references:ID"`
	ProgramSubPlanID      uint                `gorm:"column:program_sub_plan_id;type:uint"`
	ProgramCourseDetailID uint                `gorm:"column:program_course_detail_id;type:varchar;size:255;"`
	ProgramCourseDetail   ProgramCourseDetail `gorm:"foreignKey:ProgramCourseDetailID;references:ID"` //has child ref
	ProgramPloID          uint                `gorm:"column:program_plo_id;type:varchar;size:255;"`
	ProgramPlo            ProgramPlo          `gorm:"foreignKey:ProgramPloID;references:ID"`
	Status                int                 `gorm:"column:status;default:0;type:int" json:"status"`
	CreatedAt             time.Time           `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt             time.Time           `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt             gorm.DeletedAt      `gorm:"index;column:deleted_at"`
}

func (ProgramMapCurMapResp) TableName() string {
	return "program_map_curmap_resp"
}

type ProgramMapCurMapKsa struct {
	ID                    uint                `gorm:"column:id;primaryKey;autoIncrement"`
	ProgramSubPlan        ProgramSubPlan      `gorm:"foreignKey:ProgramSubPlanID;references:ID"`
	ProgramSubPlanID      uint                `gorm:"column:program_sub_plan_id;type:uint"`
	ProgramCourseDetailID uint                `gorm:"column:program_course_detail_id;type:varchar;size:255;"`
	ProgramCourseDetail   ProgramCourseDetail `gorm:"foreignKey:ProgramCourseDetailID;references:ID"` //has child ref
	ProgramPloID          uint                `gorm:"column:program_plo_id;type:varchar;size:255;"`
	ProgramPlo            ProgramPlo          `gorm:"foreignKey:ProgramPloID;references:ID"`
	KsaID                 string              `gorm:"column:ksa_id;type:varchar" json:"ksa_id"`
	CreatedAt             time.Time           `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt             time.Time           `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt             gorm.DeletedAt      `gorm:"index;column:deleted_at"`
}

func (ProgramMapCurMapKsa) TableName() string {
	return "program_map_curmap_ksa"
}

type ProgramKsaDetail struct {
	ID               uint           `gorm:"column:id;primaryKey;autoIncrement"`
	ProgramSubPlan   ProgramSubPlan `gorm:"foreignKey:ProgramSubPlanID;references:ID"`
	ProgramSubPlanID uint           `gorm:"column:program_sub_plan_id;type:uint"`
	KsaType          string         `gorm:"column:ksa_type;type:varchar;size:255;"`
	Order            uint           `gorm:"column:order;default:99;type:uint"`
	ShortCode        string         `gorm:"column:short_code;type:varchar" json:"short_code"`
	KsaDetail        string         `gorm:"column:ksa_detail;type:varchar" json:"ksa_detail"`
	CreatedAt        time.Time      `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt        time.Time      `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt        gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

func (ProgramKsaDetail) TableName() string {
	return "program_ksa_detail"
}
