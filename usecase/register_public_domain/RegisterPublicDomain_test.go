package register_public_domain_test

import (
	"errors"
	"github.com/attestify/go-domain/gateway"
	"github.com/attestify/go-domain/usecase/register_public_domain"
	"github.com/attestify/go-domain/usecase/register_public_domain/public_domain_request"
	"strings"
	"testing"
)

func setup(t *testing.T) {
	t.Parallel()
}

// todo - Provide custom errors - ValidationError, InternalError

/** Happy Path **/

// Given a new RegisterPublicDomain use case is instantiated, when a user id of "1541815603606036480" is provided,
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

// Given a user with Id "1541815603606036480" requests to register the public domain of "attestify.io"
// and Given the IdentityGateway provides a value of "1541815603606036481" for a domain Id
// then the RegisterPublicDomain usecase should execute without error
// and mutate the PublicDomainRequest domain id to be "1541815603606036481"
func Test_ExecuteRequest_For_RegisterPublicDomain_Successfully(t *testing.T) {
	setup(t)
	// Assemble
	var expectedUserId int64 = 1541815603606036480
	var expectedDomainId int64 = 1541815603606036481
	var registrationGateway register_public_domain.RegistrationGateway = NewRegistrationGatewayMock(false)
	var identityGateway gateway.IdentityGateway = NewIdentityGatewayMock(expectedDomainId)
	request := public_domain_request.New(expectedUserId, "attestify.io")

	// Act
	usecase, err := register_public_domain.New(identityGateway, registrationGateway)
	if err != nil {
		t.Fatalf("An error was returned when instantiating the RegisterPublicDomain use case. No error was expected."+
			"\n Error: %s ", err.Error())
	}
	err = usecase.Register(&request)
	if err != nil {
		t.Fatalf("An error was returned when invoking RegisterPublicDomain.Register(...). No error was expected."+
			"\n Error: %s ", err.Error())
	}

	// Assert
	actualDomainId := request.DomainId()
	if expectedDomainId != actualDomainId {
		t.Errorf("The PublicDomainRequest.DomainId() did not return the expecgted value. "+
			"/n Expdected: %d /n Actual: %d", expectedDomainId, actualDomainId)
	}

}

/** Sad Path **/

// Given there is a nil IdentityGateway implementation provided to the RegisterPublicDomain constructor
// and given the RegistrationGateway is not a nil implementation
// when the RegisterPublicDomain use case is invoked
// then an error should be returned with the text,
// "the provided identity gateway is nil, please provide a valid instance of an identity gateway"
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

// Given there is a nil RegistrationGateway implementation provided to the RegisterPublicDomain constructor
// and given the IdentityGateway is not a nil implementation
// when the RegisterPublicDomain use case is invoked
// then an error should be returned with the text,
// "the provided RegistrationGateway is nil, please provide a valid instance of a RegistrationGateway"
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

// Given a user with Id "1541815603606036480" requests to register the public domain of "attestify.io"
// and Given the IdentityGateway returns and error when invoked
// then the RegisterPublicDomain usecase should return an error that starts with "error invoking the IdentityGateway:"
func Test_IdentityGateway_Returns_Error(t *testing.T) {
	setup(t)
	// Assemble
	var expectedUserId int64 = 1541815603606036480
	var registrationGateway register_public_domain.RegistrationGateway = NewRegistrationGatewayMock(false)
	var identityGateway gateway.IdentityGateway = NewIdentityGatewayMockError()
	request := public_domain_request.New(expectedUserId, "attestify.io")

	// Act
	usecase, err := register_public_domain.New(identityGateway, registrationGateway)
	if err != nil {
		t.Fatalf("An error was returned when instantiating the RegisterPublicDomain use case. No error was expected."+
			"\n Error: %s ", err.Error())
	}
	err = usecase.Register(&request)

	// Assert
	if err == nil {
		t.Fatal("an error is expected, although no error was thrown")
	}

	containsError := strings.Contains(err.Error(), "error invoking the IdentityGateway:")
	if !containsError {
		t.Errorf("expcted the error to start with 'error invoking the IdentityGateway:' and it does not.")
	}

}

// Given a user with Id "1541815603606036480" requests to register the public domain of "attestify.io"
// and Given the RegistrationGateway returns and error when invoked
// then the RegisterPublicDomain usecase should return an error that starts with "error invoking the RegistrationGateway:"
func Test_RegistrationGateway_Returns_Error(t *testing.T) {
	setup(t)
	// Assemble
	var expectedUserId int64 = 1541815603606036480
	var expectedDomainId int64 = 1541815603606036481
	var registrationGateway register_public_domain.RegistrationGateway = NewRegistrationGatewayMock(true)
	var identityGateway gateway.IdentityGateway = NewIdentityGatewayMock(expectedDomainId)
	request := public_domain_request.New(expectedUserId, "attestify.io")

	usecase, err := register_public_domain.New(identityGateway, registrationGateway)
	if err != nil {
		t.Fatalf("An error was returned when instantiating the RegisterPublicDomain use case. No error was expected."+
			"\n Error: %s ", err.Error())
	}

	// Act
	err = usecase.Register(&request)

	// Assert
	if err == nil {
		t.Fatal("an error is expected, although no error was thrown")
	}

	containsError := strings.Contains(err.Error(), "error invoking the RegistrationGateway:")
	if !containsError {
		t.Errorf("expcted the error to start with 'error invoking the RegistrationGateway:' and it does not.")
	}

}

func Test_ExecuteRequest_For_RegisterPublicDomain_With_Bad_DomainName(t *testing.T) {
	setup(t)
	// Assemble
	var expectedUserId int64 = 1541815603606036480
	var expectedDomainId int64 = 1541815603606036481
	var registrationGateway register_public_domain.RegistrationGateway = NewRegistrationGatewayMock(false)
	var identityGateway gateway.IdentityGateway = NewIdentityGatewayMock(expectedDomainId)
	request := public_domain_request.New(expectedUserId, "attestify.io-1")

	usecase, err := register_public_domain.New(identityGateway, registrationGateway)
	if err != nil {
		t.Fatalf("An error was returned when instantiating the RegisterPublicDomain use case. No error was expected."+
			"\n Error: %s ", err.Error())
	}

	// Act
	err = usecase.Register(&request)
	if err == nil {
		t.Fatalf("An error was expected from RegisterPublicDomain.Register(...), but none was returned.")
	}

	// Assert
	expectedError := "error creating the PublicDomain entity:"
	containsError := strings.Contains(err.Error(), expectedError)
	if !containsError {
		t.Errorf("the retuned error does not start with the expected string: \n Actual %s \n Expected %s",
			err.Error(), expectedError)
	}

}

/** Mocks **/

/** IdentityGatewayMock **/
type IdentityGatewayMock struct {
	expectedId  int64
	returnError bool
}

func NewIdentityGatewayMock(expectedId int64) IdentityGatewayMock {
	return IdentityGatewayMock{
		expectedId:  expectedId,
		returnError: false,
	}
}

func NewIdentityGatewayMockError() IdentityGatewayMock {
	return IdentityGatewayMock{
		expectedId:  0,
		returnError: true,
	}
}

func (gateway IdentityGatewayMock) GenerateId() (int64, error) {
	if gateway.returnError {
		return 0, errors.New("mock error from identity gateway")
	} else {
		return gateway.expectedId, nil
	}

}

/** RegistrationGatewayMock **/
type RegistrationGatewayMock struct {
	returnError bool
	wasCalled   bool
}

func NewRegistrationGatewayMock(returnError bool) RegistrationGatewayMock {
	return RegistrationGatewayMock{
		returnError: returnError,
	}
}

func (gateway RegistrationGatewayMock) RegisterPublicDomain(userId int64, publicDomain register_public_domain.
	PublicDomainEntity) error {
	if gateway.returnError {
		return errors.New("error with Registration Gateway")
	} else {
		return nil
	}
}
