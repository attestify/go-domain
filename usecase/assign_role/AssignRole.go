package assign_role

import (
	"github.com/attestify/go-kernel/error/internal_error"
	"github.com/attestify/go-kernel/identity/id"
)

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

func (usecase AssignRole) Assign(userId int64, entityId int64, entity string) error {

	err := usecase.gateway.RecordAssignment(id.New(userId), id.New(entityId), entity)
	if err != nil {
		return err
	}
	return nil
}
