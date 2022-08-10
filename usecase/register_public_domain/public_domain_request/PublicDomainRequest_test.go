package public_domain_request

import (
	"testing"
)

func setup(t *testing.T) {
	t.Parallel()
}

/** Happy Path Tests **/

// Given a user id with the value of "1541815603606036480" is provided, and
// given a registered name with the value of "attestify.io" is provided
// when the PublicDomainRequest object is constructed successfully
// then no errors should be returned
// and the domain id should default to "0"
func Test_Instantiate_PublicDomainRequest_Successfully(t *testing.T) {
	setup(t)
	// Arrange
	var userId int64 = 1541815603606036480
	registeredName := "attestify.io"

	// Act
	request, err := New(userId, registeredName)

	// Assert
	if err != nil {
		t.Errorf("An error was returned when no error was expected.\n Error: %s ", err.Error())
	}

	actualDefaultDomainId := request.DomainId()
	var expectedDefaultDomainId int64 = 0
	if actualDefaultDomainId != expectedDefaultDomainId {
		t.Errorf("The exptected domain was not returned. \n Actual: %d \n Expected: %d", actualDefaultDomainId, expectedDefaultDomainId)

	}
}

// Given a user id with the value of "1541815603606036480" is provided, and
// given a registered name with the value of "attestify.io" is provided
// when the PublicDomainRequest object is constructed successfully
// then no errors should be returned
// and it should return the domain name of "attestify.io"
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
	if actual != expected {
		t.Errorf("The exptected domain was not returned. \n Actual: %s \n Expected: %s", actual, expected)
	}
}

// Given a user id with the value of "1541815603606036480" is provided, and
// given a registered name with the value of "attestify.io" is provided
// when the PublicDomainRequest object is constructed successfully
// then no errors should be returned
// and it should return the user id of "1541815603606036480"
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
	if actual != expected {
		t.Errorf("The exptected domain was not returned. \n Actual: %d \n Expected: %d", actual, expected)
	}

}

// Given a user id with the value of "1541815603606036480" is provided, and
// given a registered name with the value of "attestify.io" is provided
// when the PublicDomainRequest object is constructed successfully
// then no errors should be returned
// and when the domain id is updated to "1541815603606036481"
// then the id returned from DomainId() should be "1541815603606036480"
func Test_Update_Default_DomainId_Successfully(t *testing.T) {
	setup(t)
	// Arrange
	var userId int64 = 1541815603606036480
	registeredName := "attestify.io"

	// Act
	request, err := New(userId, registeredName)
	if err != nil {
		t.Errorf("An error was returned when no error was expected.\n Error: %s ", err.Error())
	}
	request.UpdateDomainId(1541815603606036481)

	// Assert
	actual := request.DomainId()
	var expected int64 = 1541815603606036481
	if actual != expected {
		t.Errorf("The exptected domain was not returned. \n Actual: %d \n Expected: %d", actual, expected)

	}
}
