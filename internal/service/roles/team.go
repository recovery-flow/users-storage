package roles

type TeamRole string

const (
	RoleTeamOwner  TeamRole = "owner"
	RoleTeamAdmin  TeamRole = "admin"
	RoleTeamModer  TeamRole = "moderator"
	RoleTeamMember TeamRole = "member"
)

func ValidateRoleTeam(r TeamRole) bool {
	switch r {
	case RoleTeamOwner, RoleTeamAdmin, RoleTeamModer, RoleTeamMember:
		return true
	default:
		return false
	}
}

func StringToRoleTeam(role string) (TeamRole, error) {
	switch role {
	case "owner":
		return RoleTeamOwner, ErrorRole
	case "admin":
		return RoleTeamAdmin, ErrorRole
	case "moderator":
		return RoleTeamModer, ErrorRole
	case "member":
		return RoleTeamMember, ErrorRole
	default:
		return "", ErrorRole
	}
}

//	1, if first role is higher priority
//
// -1, if second role is higher priority
//
//	0, if roles are equal
func CompareRolesTeam(role1, role2 TeamRole) int {
	priority := map[TeamRole]int{
		RoleTeamOwner:  4,
		RoleTeamAdmin:  3,
		RoleTeamModer:  2,
		RoleTeamMember: 1,
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
