package register_public_domain_test

import (
	"errors"
	"github.com/attestify/go-domain/gateway"
	"github.com/attestify/go-domain/usecase/register_public_domain"
	"github.com/attestify/go-domain/usecase/register_public_domain/public_domain"
	"testing"
)

func setup(t *testing.T) {
	t.Parallel()
}

// Given a new Register Public Domain use case is instantiated, when a user id of "1541815603606036480" is provided,
// and when a registered name of "attestify.io" is provided, then the RegisterPublicDomain use case should instantiate
// without any errors.
func Test_Instantiate_RegisterPublicDomain_Successfully(t *testing.T) {
	setup(t)

	// Assemble
	var registrationGateway register_public_domain.RegistrationGateway = NewRegistrationGatewayMock(false)
	var identityGateway gateway.IdentityGateway = NewIdentityGatewayMock(1541815603606036480)

	// Act
	_, err := register_public_domain.New(identityGateway, registrationGateway)

	// Assert
	if err != nil {
		t.Errorf("An error was returned when no error was expected.\n Error: %s ", err.Error())
	}
}

func Test_ExecuteRequest_For_RegisterPublicDomain_Successfully(t *testing.T) {
	setup(t)
	t.Errorf("Test Not Implemented")
	//// Assemble
	//newDomainRequest, err = public_domain_request.New(1541815603606036480, "attestify.io")
	//// todo - test for error
	//
	//// Act
	//usecase, err = New(identityGateway, registrationGateway)
	//err = usecase.MakeRequest(*newDomainRequest)
	//// todo - test for error
	//
	//// Assert
	//// todo - test for newDomainRequest.DomainId != 0

}


/** Sad Path **/

// todo - provide description
func Test_Nil_IdentityGateway(t *testing.T) {
	setup(t)

	// Assemble
	var registrationGateway register_public_domain.RegistrationGateway = NewRegistrationGatewayMock(false)
	var identityGateway gateway.IdentityGateway = nil

	// Act
	_, err := register_public_domain.New(identityGateway, registrationGateway)

	// Assert
	if err == nil {
		t.Fatalf("an error is expected, although no error was thrown")
	}

	expected := "the provided identity gateway is nil, please provide a valid instance of an identity gateway"
	actual := err.Error()
	if expected != actual {
		t.Errorf("The exptected error was not returned. \n Actual: %s \n Expected: %s", actual, expected)
	}

}

// todo - test with nil RegistrationGateway
func Test_Nil_RegistrationGateway(t *testing.T) {
	setup(t)

	// Assemble
	var registrationGateway register_public_domain.RegistrationGateway = nil
	var identityGateway gateway.IdentityGateway = NewIdentityGatewayMock(1541815603606036480)

	// Act
	_, err := register_public_domain.New(identityGateway, registrationGateway)

	// Assert
	if err == nil {
		t.Fatalf("an error is expected, although no error was thrown")
	}

	expected := "the provided RegistrationGateway is nil, please provide a valid instance of a RegistrationGateway"
	actual := err.Error()
	if expected != actual {
		t.Errorf("The exptected error was not returned. \n Actual: %s \n Expected: %s", actual, expected)
	}

}

/** Mocks **/

/** IdentityGatewayMock **/
type IdentityGatewayMock struct {
	expectedId int64
}

func NewIdentityGatewayMock(expectedId int64) IdentityGatewayMock {
	return IdentityGatewayMock{
		expectedId: expectedId,
	}
}

func (gateway IdentityGatewayMock) GenerateId() int64 {
	return gateway.expectedId
}


/** RegistrationGatewayMock **/
type RegistrationGatewayMock struct {
	returnError bool
	wasCalled bool

}

func NewRegistrationGatewayMock(returnError bool) RegistrationGatewayMock{
	return RegistrationGatewayMock{
		returnError: returnError,
		wasCalled: false,
	}
}

func (gateway RegistrationGatewayMock) RegisterPublicDomain(publicDomain public_domain.PublicDomain) error {
	gateway.wasCalled = true
	if gateway.returnError {
		return errors.New("error with Registration Gateway")
	} else {
		return nil
	}
}

