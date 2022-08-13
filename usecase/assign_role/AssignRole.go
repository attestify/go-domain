package assign_role

import "github.com/attestify/go-kernel/error/internal_error"

type AssignRole struct {
	gateway AssignRoleGateway
}

func New(gateway AssignRoleGateway) (AssignRole, error) {

	if gateway == nil {
		return AssignRole{}, internal_error.New("the provided AssignRoleGateway is nil, please provide a valid instance of an AssignRoleGateway")
	}

	return AssignRole{
		gateway: gateway,
	}, nil
}
