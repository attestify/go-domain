package public_domain_entity

import (
	"errors"
	"github.com/attestify/go-kernel/identity/id"
	"github.com/attestify/go-kernel/uri/registered_name"
)

// PublicDomain is the entity
type PublicDomain struct {
	domainId id.Id
	registeredName registered_name.RegisteredName
}

func NewPublicDomain(domainId id.Id, domain string) (PublicDomain, error) {

	_registeredName, err := registered_name.NewFromString(domain)

	if err != nil {
		errorMessage := "Error creating a Public Domain: " + err.Error()
		return PublicDomain{}, errors.New(errorMessage)
	}

	return PublicDomain{
		domainId: domainId,
		registeredName: _registeredName,
	}, nil

}

func (pd *PublicDomain) RegisteredName() registered_name.RegisteredName {
	return pd.registeredName
}

func (pd *PublicDomain) DomainId() id.Id {
	return pd.domainId
}
