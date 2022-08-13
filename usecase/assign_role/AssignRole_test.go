package assign_role_test

import (
	"github.com/attestify/go-domain/usecase/assign_role"
	"testing"
)

func setup(t *testing.T) {
	t.Parallel()
}


// Given an instance of a AssignRoleGateway is provided
// When the AssignRole use case is instantiated
// Then there should be no error
func Test_Instantiate_AssignRole_Successfully(t *testing.T) {
	setup(t)
	assignRoleGateway := NewAssignRoleGatewayMock()
	_, err := assign_role.New(assignRoleGateway)
	if err != nil {
		t.Errorf("An error was returned when no error was expected: \n %s", err.Error())
	}
}

// todo -- AssignRole Invocation
// Given the user id of "[x]" is provided,
// and the entity if of "[y]" is provided,
// and the entity of "TestEntity" is provided
// Then when the AssigneRole usecase is


type AssignRoleGatewayMock struct {
}

func NewAssignRoleGatewayMock() AssignRoleGatewayMock {
	return AssignRoleGatewayMock{}
}