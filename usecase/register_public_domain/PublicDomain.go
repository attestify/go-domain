package register_public_domain

import (
	"errors"
	"github.com/attestify/go-kernel/uri/registered_name"
)

// PublicDomain is the entity
type PublicDomain struct {
	registeredName registered_name.RegisteredName
}

func NewPublicDomain(domain string) (*PublicDomain, error) {

	_registeredName, err := registered_name.NewFromString(domain)

	if err != nil {
		errorMessage := "Error creating a Public Domain: " + err.Error()
		return &PublicDomain{}, errors.New(errorMessage)
	}

	return &PublicDomain{
		registeredName: *_registeredName,
	}, nil

}

func (pd *PublicDomain) Value() string {
	return pd.registeredName.Value()
}
