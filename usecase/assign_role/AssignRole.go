package assign_role

type AssignRole struct {
	gateway AssignRoleGateway
}

func New(gateway AssignRoleGateway) (AssignRole, error) {
	return AssignRole{
		gateway: gateway,
	}, nil
}
