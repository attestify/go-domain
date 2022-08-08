package public_domain_request

import (
	"testing"
)

func setup(t *testing.T) {
	t.Parallel()
}
/** Happy Path Tests **/

// Test_Instantiate_PublicDomainRequest should successfully instantiate a PublicDomainRequest without error given the domain entered is "attestify.io"
// todo - updated description
func Test_Instantiate_PublicDomainRequest(t *testing.T) {
	setup(t)
	// Arrange
	var userId int64 = 1541815603606036480
	registeredName := "attestify.io"

	// Act
	_, err := New(userId, registeredName)

	// Assert
	if err != nil {
		t.Errorf("An error was returned when no error was expected.\n Error: %s ", err.Error())
	}

}

// todo - updated description
func Test_Instantiate_PublicDomainRequest_Get_Domain_Successfully(t *testing.T) {
	setup(t)
	// Arrange
	var userId int64 = 1541815603606036480
	registeredName := "attestify.io"

	// Act
	request, err := New(userId, registeredName)

	// Assert
	if err != nil {
		t.Fatalf("An error was returned when no error was expected.\n Error: %s ", err.Error())
	}

	actual := request.Domain()
	expected := "attestify.io"
	if  actual != expected {
	t.Errorf("The exptected domain was not returned. \n Actual: %s \n Expected: %s", actual, expected)

}
// todo - validate the get-userId function
// todo - mutate and a public-domain-id without error, then get the public-domain-id

//actual := request.Domain()
//expected := "attestify.io"
//if  actual != expected {
//t.Errorf("The exptected domain was not returned. \n Actual: %s \n Expected: %s", actual, expected)
//
}

// todo - updated description
func Test_Instantiate_PublicDomainRequest_Get_UserId_Successfully(t *testing.T) {
	setup(t)
	// Arrange
	var userId int64 = 1541815603606036480
	registeredName := "attestify.io"

	// Act
	request, err := New(userId, registeredName)

	// Assert
	if err != nil {
		t.Fatalf("An error was returned when no error was expected.\n Error: %s ", err.Error())
	}

	actual := request.UserId()
	var expected int64 = 1541815603606036480
	if  actual != expected {
		t.Errorf("The exptected domain was not returned. \n Actual: %d \n Expected: %d", actual, expected)

	}
	// todo - validate the get-userId function
	// todo - mutate and a public-domain-id without error, then get the public-domain-id

	//actual := request.Domain()
	//expected := "attestify.io"
	//if  actual != expected {
	//t.Errorf("The exptected domain was not returned. \n Actual: %s \n Expected: %s", actual, expected)
	//
}