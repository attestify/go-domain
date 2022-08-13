package assign_role_test

import (
	"errors"
	"github.com/attestify/go-domain/usecase/assign_role"
	"github.com/attestify/go-kernel/error/internal_error"
	"testing"
)

func setup(t *testing.T) {
	t.Parallel()
}

/** Happy Path **/

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

/** Sad Path **/

// Given an nil instance of a AssignRoleGateway is provided
// When the AssignRole use case is instantiated
// Then an InternalError with the text "the provided AssignRoleGateway is nil, please provide a valid instance of an AssignRoleGateway" should be returned
func Test_Instantiate_AssignRole_With_Nil_AssignRoleGateway(t *testing.T) {
	setup(t)
	// Assemble
	var assignRoleGateway assign_role.AssignRoleGateway = nil

	// Act
	_, err := assign_role.New(assignRoleGateway)

	// Assert
	if err == nil {
		t.Fatalf("An error was expected, although no error was returned.")
	}

	if !errors.As(err, &internal_error.InternalError{}) {
		t.Errorf("did not get the epected error of type InternalError")
	}

	actualMessage := err.Error()
	expectedMessage := "the provided AssignRoleGateway is nil, please provide a valid instance of an AssignRoleGateway"
	if expectedMessage != actualMessage {
		t.Errorf("The returned error message was not expected: \n Expected: %s \n Actual %s", expectedMessage, actualMessage)
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