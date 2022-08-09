package public_domain

import (
	"errors"
	"github.com/attestify/go-kernel/identity/id"
	"github.com/attestify/go-kernel/uri/registered_name"
)

// PublicDomain is the default implementation for the PublicDomainEntity interface
type PublicDomain struct {
	domainId       id.Id
	registeredName registered_name.RegisteredName
}

func New(domainId int64, domain string) (PublicDomain, error) {

	_registeredName, err := registered_name.NewFromString(domain)

	if err != nil {
		errorMessage := "Error creating a Public Domain: " + err.Error()
		return PublicDomain{}, errors.New(errorMessage)
	}

	return PublicDomain{
		domainId:       id.New(domainId),
		registeredName: _registeredName,
	}, nil

}

func (pd *PublicDomain) Domain() registered_name.RegisteredName {
	return pd.registeredName
}

func (pd *PublicDomain) Id() id.Id {
	return pd.domainId
}
