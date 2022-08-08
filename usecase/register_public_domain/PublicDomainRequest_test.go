package register_public_domain

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
	var userId int64 = 1541815603606036480
	registeredName := "attestify.io"
	_, err := NewPublicDomainRequest(userId, registeredName)

	// Fatal use to end test if an error object was not returned because the expressions after this evaluate the error object
	if err != nil {
		t.Fatalf("An error was returned when no error was expected.\n Error: %s ", err.Error())
	}

}

// todo - validate the get-public-domain function
// todo - validate the get-userId function
// todo - mutate and a public-domain-id without error, then get the public-domain-id

//actual := request.Domain()
//expected := "attestify.io"
//if  actual != expected {
//t.Errorf("The exptected domain was not returned. \n Actual: %s \n Expected: %s", actual, expected)
//}