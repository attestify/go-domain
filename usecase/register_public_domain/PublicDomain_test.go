package register_public_domain

import (
	"testing"
)

/** Happy Paths **/

// Test_Instantiate_PublicDomainRequest should successfully instantiate a PublicDomainRequest without error given the domain entered is "attestify.io"
func Test_Instantiate_PublicDomain(t *testing.T) {
	setup(t)
	_, err := NewPublicDomain("attestify.io")

	// Fatal use to end test if an error object was not returned because the expressions after this evaluate the error object
	if err != nil {
		t.Fatalf("An error was returned when no error was expected.\n Error: %s ", err.Error())
	}

}

// Test_Instantiate_PublicDomainRequest should instantiate without error and return a value of "attestify.
// io" when an argument of "attestify.io" is provided.
func Test_Instantiate_PublicDomain_ValueCheck(t *testing.T) {
	setup(t)
	request, err := NewPublicDomain("attestify.io")

	// Fatal use to end test if an error object was not returned because the expressions after this evaluate the error object
	if err != nil {
		t.Fatalf("An error was returned when no error was expected.\n Error: %s ", err.Error())
	}

	actual := request.Value()
	expected := "attestify.io"
	if expected!= actual {
		t.Errorf("Did not return the expected value.\nActual: %s\nExpected: %s", actual, expected)
	}

}

/** Sad Paths **/

// Test_Instantiate_PublicDomainRequest_Error should error on instantiation when a PublicDomainRequest with an
// argument of 1attestify.io-com is provided
func Test_Instantiate_PublicDomain_Error(t *testing.T) {
	setup(t)
	_, err := NewPublicDomain("1attestify.io-com")

	// Fatal use to end test if an error object was not returned because the expressions after this evaluate the error object
	if err == nil {
		t.Fatalf("An error was expceted, but none retured.")
	}

}
