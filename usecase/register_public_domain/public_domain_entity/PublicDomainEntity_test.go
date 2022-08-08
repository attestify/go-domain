package public_domain_entity

import (
	"github.com/attestify/go-kernel/identity/id"
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
	domainId := id.New(1541815603606036480)
	registeredName := "attestify.io"
	_, err := NewPublicDomain(domainId, registeredName)

	// Fatal use to end test if an error object was not returned because the expressions after this evaluate the error object
	if err != nil {
		t.Fatalf("An error was returned when no error was expected.\n Error: %s ", err.Error())
	}

}

// Test_Instantiate_PublicDomain_RegisteredName_Check return a value of "attestify.io" without error
// given the registered domain name provided is "attestify.io" and,
// given the domain id provided is "1541815603606036480"
func Test_Instantiate_PublicDomain_ValueCheck(t *testing.T) {
	setup(t)
	domainId := id.New(1541815603606036480)
	registeredName := "attestify.io"
	entity, err := NewPublicDomain(domainId, registeredName)

	// Fatal use to end test if an error object was not returned because the expressions after this evaluate the error object
	if err != nil {
		t.Fatalf("An error was returned when no error was expected.\n Error: %s ", err.Error())
	}

	name := entity.RegisteredName()
	actual := name.Value()
	expected := "attestify.io"
	if expected != actual {
		t.Errorf("Did not return the expected value.\nActual: %s\nExpected: %s", actual, expected)
	}

}

/** Sad Paths **/

// Test_Instantiate_PublicDomainRequest_Error should error on instantiation when a PublicDomainRequest with an
// argument of 1attestify.io-com is provided
func Test_Instantiate_PublicDomain_Error(t *testing.T) {
	setup(t)
	domainId := id.New(1541815603606036480)
	badRegisteredName := "1attestify.io-com"
	_, err := NewPublicDomain(domainId, badRegisteredName)

	// Fatal use to end test if an error object was not returned because the expressions after this evaluate the error object
	if err == nil {
		t.Fatalf("An error was expceted, but none retured.")
	}

}
