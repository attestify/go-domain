package public_domain_test

import (
	"errors"
	"github.com/attestify/go-domain/usecase/register_public_domain/public_domain"
	"github.com/attestify/go-kernel/error/validation_error"
	"testing"
)

func setup(t *testing.T) {
	t.Parallel()
}

/** Happy Paths **/

// Test_Instantiate_PublicDomain should successfully instantiate a PublicDomain entity without error
// given the registered domain name provided is "attestify.io" and,
// given the domain id provided is "1541815603606036480"
// then a PublicDomain entity should be created without an error.
func Test_Instantiate_PublicDomain(t *testing.T) {
	setup(t)
	var domainId int64 = 1541815603606036480
	registeredName := "attestify.io"
	_, err := public_domain.New(domainId, registeredName)

	// Fatal use to end test if an error object was not returned because the expressions after this evaluate the error object
	if err != nil {
		t.Fatalf("An error was returned when no error was expected.\n Error: %s ", err.Error())
	}

}

// Test_Instantiate_PublicDomain should successfully instantiate a PublicDomain entity without error
// given the registered domain name provided is "attestify.io"
// then a PublicDomain entity should be created without an error
// and the DomainId should be "0"
func Test_Instantiate_PublicDomain_From_DomainName_Only(t *testing.T) {
	setup(t)
	registeredName := "attestify.io"
	entity, err := public_domain.NewFromDomainName(registeredName)

	// Fatal use to end test if an error object was not returned because the expressions after this evaluate the error object
	if err != nil {
		t.Fatalf("An error was returned when no error was expected.\n Error: %s ", err.Error())
	}

	actual := entity.Id()
	var expected int64 = 0
	if expected != actual {
		t.Errorf("Did not return the expected value.\nActual: %d\nExpected: %d", actual, expected)
	}

}

// Test_Instantiate_PublicDomain_RegisteredName_Check return a value of "attestify.io" without error
// given the registered domain name provided is "attestify.io" and,
// given the domain id provided is "1541815603606036480"
func Test_Instantiate_PublicDomain_Check_Domain(t *testing.T) {
	setup(t)
	var domainId int64 = 1541815603606036480
	registeredName := "attestify.io"
	entity, err := public_domain.New(domainId, registeredName)

	// Fatal use to end test if an error object was not returned because the expressions after this evaluate the error object
	if err != nil {
		t.Fatalf("An error was returned when no error was expected.\n Error: %s ", err.Error())
	}

	actual := entity.Domain()
	expected := "attestify.io"
	if expected != actual {
		t.Errorf("Did not return the expected value.\nActual: %s\nExpected: %s", actual, expected)
	}

}

// Test_Instantiate_PublicDomain_RegisteredName_Check return a value of "1541815603606036480" without error
// given the registered domain name provided is "attestify.io" and,
// given the domain id provided is "1541815603606036480"
func Test_Instantiate_PublicDomain_Check_DomainId(t *testing.T) {
	setup(t)
	var domainId int64 = 1541815603606036480
	registeredName := "attestify.io"
	entity, err := public_domain.New(domainId, registeredName)

	// Fatal use to end test if an error object was not returned because the expressions after this evaluate the error object
	if err != nil {
		t.Fatalf("An error was returned when no error was expected.\n Error: %s ", err.Error())
	}

	actual := entity.Id()
	var expected int64 = 1541815603606036480
	if expected != actual {
		t.Errorf("Did not return the expected value.\nActual: %d\nExpected: %d", actual, expected)
	}

}

// Given only the registered domain name provided is "attestify.io",.
// and given a PublicDomain entity is created is created successfully
// When a PublicDomain entity is created
//
func Test_Updated_DomainId(t *testing.T) {
	setup(t)
	registeredName := "attestify.io"
	entity, err := public_domain.NewFromDomainName(registeredName)

	// Fatal use to end test if an error object was not returned because the expressions after this evaluate the error object
	if err != nil {
		t.Fatalf("An error was returned when no error was expected.\n Error: %s ", err.Error())
	}

	// Act
	entity.UpdateId(1541815603606036480)

	// Assert
	actual := entity.Id()
	var expected int64 = 1541815603606036480
	if expected != actual {
		t.Errorf("Did not return the expected value.\nActual: %d\nExpected: %d", actual, expected)
	}


}


/** Sad Paths **/

// Test_Instantiate_PublicDomainRequest_Error should error on instantiation when a PublicDomainRequest with an
// argument of 1attestify.io-com is provided
func Test_Instantiate_PublicDomain_Error(t *testing.T) {
	setup(t)
	var domainId int64 = 1541815603606036480
	badRegisteredName := "1attestify.io-com"
	_, err := public_domain.New(domainId, badRegisteredName)

	// Fatal use to end test if an error object was not returned because the expressions after this evaluate the error object
	if err == nil {
		t.Fatalf("An error was expceted, but none retured.")
	}

	if !errors.As(err, &validation_error.ValidationError{}) {
		t.Errorf("did not get the epected error of type ValidationError")
	}

}
