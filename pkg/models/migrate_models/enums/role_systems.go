package enums

// Define a custom type for days of the week
type SystemPermissions int

// Define constants for the days of the week
const (
	ViewMyProfile SystemPermissions = iota
	UpdateMyProfile
	ListAllLecture
	CreateLecture
	UpdateLecture
	DeleteLecture
	Status
)

// ToString method to convert the enum value to a string
func (sysPerm SystemPermissions) ToString() string {
	switch sysPerm {
	case ViewMyProfile:
		return "view_my_profile"
	case UpdateMyProfile:
		return "update_my_profile"
	case ListAllLecture:
		return "list_all_lecture"
	case CreateLecture:
		return "create_lecture"
	case UpdateLecture:
		return "update_lecture"
	case DeleteLecture:
		return "delete_lecture"
	case Status:
		return "status"
	default:
		return "Invalid permissions"
	}
}
