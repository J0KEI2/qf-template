package constant

const (
	FORM_DATA_TIME_FORMAT                   string = "Mon Jan 2 2006 15:04:05 MST-0700"
	APPROVAL_ATTRIBUTE_CURRICULUM_COMMITTEE string = "curriculum_committee"
	APPROVAL_ATTRIBUTE_ACADEMIC             string = "academic"
	APPROVAL_ATTRIBUTE_UNIVERSITY           string = "university"
	// 2024-08-28T21:31:21.522912Z
// FACULTY_LEVEL            = "ระดับคณะ/หน่วยงาน"
// CURRICULUM_COMMITEE       = "ระดับคณะกรรมการกลั่นกรอง"
// ACADEMIC_COUNCIL         = "สภาวิชาการ"
// UNIVERSITY_COUNCIL_LEVEL = "สภามหาวิทยาลัย"
// CHECO                    = "CHECO"

// STATUS_CURRENT  = "กำลังดำเนินการ"
// STATUS_PENDING  = "กำลังดำเนินการตรวจสอบ"
// STATUS_REJECT   = "ไม่ผ่านการตรวจสอบ/รอการแก้ไข"
// STATUS_APPROVED = "ผ่านการตรวจสอบ"
// STATUS_WAITING  = "รอดำเนินการตรวจสอบ"
)

// var APPROVAL_LEVEL_LIST = [...]string{FACULTY_LEVEL, CURRICULUM_COMMITEE, ACADEMIC_COUNCIL, UNIVERSITY_COUNCIL_LEVEL, CHECO}
// var STATUS_LIST = [...]string{STATUS_CURRENT, STATUS_PENDING, STATUS_REJECT, STATUS_APPROVED, STATUS_WAITING}

// APPROVAL_STATUS - Custom type to hold value ranging from 1-5
type APPROVAL_STATUS uint

// PROGRAM_APPROVAL_LEVEL - Represent Program Approval Level
type PROGRAM_APPROVAL_LEVEL uint
type CHECO_STATUS uint
type STATUS uint

// Declare related constants for each APPROVAL_STATUS starting with index 1
const (
	FACULTY_APPROVAL_STATUS             APPROVAL_STATUS = iota + 1 // EnumIndex = 1
	CURRICULUM_COMMITEE_APPROVAL_STATUS                            // EnumIndex = 2
	ACADEMIC_COUNCIL_APPROVAL_STATUS                               // EnumIndex = 3
	UNIVERSITY_COUNCIL_APPROVAL_STATUS                             // EnumIndex = 4
	CHECO_APPROVAL_STATUS                                          // EnumIndex = 5
)

// Declare related constants for each PROGRAM_APPROVAL_LEVEL starting with index 1
const (
	FACULTY_APPROVAL_LEVEL              PROGRAM_APPROVAL_LEVEL = iota + 1 // EnumIndex = 1
	COMMITTEE_SELECTOR_APPROVAL_LEVEL                                     // EnumIndex = 2
	CURRICULUM_COMMITEE_APPROVAL_LEVEL                                    // EnumIndex = 3
	CURRICULUM_COMMITEES_APPROVAL_LEVEL                                   // EnumIndex = 4
	ACADEMIC_COUNCIL_APPROVAL_LEVEL                                       // EnumIndex = 5
	UNIVERSITY_COUNCIL_APPROVAL_LEVEL                                     // EnumIndex = 6
	CHECO_APPROVAL_LEVEL                                                  // EnumIndex = 7
	ADMIN_APPROVAL_LEVEL                                                  // EnumIndex = 8
)

// Declare related constants for each approval_status starting with index 1
const (
	STATUS_CURRENT  STATUS = iota + 1 // EnumIndex = 1
	STATUS_PENDING                    // EnumIndex = 2
	STATUS_REJECT                     // EnumIndex = 3
	STATUS_APPROVED                   // EnumIndex = 4
	STATUS_WAITING                    // EnumIndex = 5
)

const (
	W  CHECO_STATUS = iota + 1 // EnumIndex = 1
	W1                         // EnumIndex = 2
	S                          // EnumIndex = 3
	E                          // EnumIndex = 4
	A1                         // EnumIndex = 5
	A2                         // EnumIndex = 6
	A3                         // EnumIndex = 7
	A4                         // EnumIndex = 8
	P                          // EnumIndex = 9
)

// func (s APPROVAL_LEVEL) List() [5]string {
// 	return [...]string{"ระดับคณะ/หน่วยงาน", "ระดับคณะกรรมการกลั่นกรอง", "สภาวิชาการ", "สภามหาวิทยาลัย", "CHECO"}
// }

func (s APPROVAL_STATUS) String() string {
	return [...]string{"คณะ/หน่วยงาน", "คณะกรรมการกลั่นกรองหลักสูตร", "สภาวิชาการ", "สภามหาวิทยาลัย", "กระทรวงการอุดมศึกษา วิทยาศาสตร์ วิจัยและนวัตกรรม"}[s-1]
}
func (s APPROVAL_STATUS) StringEn() string {
	return [...]string{"Faculty/Institution", "Curriculum Committee", "Academic Council", "University Council", "CHECO"}[s-1]
}
func (s APPROVAL_STATUS) EnumIndex() uint {
	return uint(s)
}

// var CHECOs = [...]string{
// 	W:  "W (รอส่ง)",
// 	W1: "W1 (ส่งไประดับมหาวิทยาลัย)",
// 	S:  "S (ส่งไป สปอว. แล้ว)",
// 	E:  "E (ส่งให้มหาวิทยาลัยแก้ไข)",
// 	A1: "A1 (หัวหน้าฝ่าย)",
// 	A2: "A2 (ผู้อำนวยการกลุ่ม)",
// 	A3: "A3 (ผู้อำนวยการสำนัก/กอง)",
// 	A4: "A4 (ปลัดกระทรวงฯ)",
// 	P:  "P (พิจารณาความสอดคล้องและออกรหัสหลักสูตรเรียบร้อย)",
// }

// var keywords map[string]CHECO_STATUS

// func InitCHECOConstants() {
// 	keywords = make(map[string]CHECO_STATUS, P-(W+1))
// 	for i := W + 1; i < P; i++ {
// 		keywords[CHECOs[i]] = i
// 	}
// }

//	func CHECOLookup(ident string) *CHECO_STATUS {
//		if checo, is_keyword := keywords[ident]; is_keyword {
//			return &checo
//		}
//		return nil
//	}

// returns string value of checo status
func (c CHECO_STATUS) String() string {
	return [...]string{"W (รอส่ง)", "W1 (ส่งไประดับมหาวิทยาลัย)", "S (ส่งไป สปอว. แล้ว)", "E (ส่งให้มหาวิทยาลัยแก้ไข)", "A1 (หัวหน้าฝ่าย)", "A2 (ผู้อำนวยการกลุ่ม)", "A3 (ผู้อำนวยการสำนัก/กอง)", "A4 (ปลัดกระทรวงฯ)", "P (พิจารณาความสอดคล้องและออกรหัสหลักสูตรเรียบร้อย)"}[c-1]
}
func (c CHECO_STATUS) EnumIndex() uint {
	return uint(c)
}

func (s STATUS) String() string {
	return [...]string{"กำลังดำเนินการ", "กำลังดำเนินการตรวจสอบ", "ไม่ผ่านการตรวจสอบ/รอการแก้ไข", "ผ่านการตรวจสอบ", "รอดำเนินการตรวจสอบ"}[s-1]
}

func (s STATUS) EnumIndex() uint {
	return uint(s)
}
