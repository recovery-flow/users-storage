package roles

type UserRole string

const (
	RoleUserAdmin  UserRole = "user_admin"
	RoleUserGov    UserRole = "user_gov"
	RoleUserVerify UserRole = "user_verify"
	RoleUserSimple UserRole = "user"
)

func ValidateRoleUser(r UserRole) bool {
	switch r {
	case RoleUserGov, RoleUserAdmin, RoleUserVerify, RoleUserSimple:
		return true
	default:
		return false
	}
}

func StringToRoleUser(role string) (UserRole, error) {
	switch role {
	case "admin":
		return RoleUserAdmin, ErrorRole
	case "user_gov":
		return RoleUserGov, ErrorRole
	case "user_verify":
		return RoleUserVerify, ErrorRole
	case "user":
		return RoleUserSimple, ErrorRole
	default:
		return "", ErrorRole
	}
}

//	1, if first role is higher priority
//
// -1, if second role is higher priority
//
//	0, if roles are equal
func CompareRolesUser(role1, role2 UserRole) int {
	priority := map[UserRole]int{
		RoleUserAdmin:  4,
		RoleUserGov:    3,
		RoleUserVerify: 2,
		RoleUserSimple: 1,
	}

	p1, ok1 := priority[role1]
	p2, ok2 := priority[role2]

	if !ok1 || !ok2 {
		return -1
	}

	if p1 > p2 {
		return 1
	} else if p1 < p2 {
		return -1
	}
	return 0
}
