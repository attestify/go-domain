package public_domain

import (
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
		return PublicDomain{}, err
	}

	return PublicDomain{
		domainId:       id.New(domainId),
		registeredName: _registeredName,
	}, nil

}

func NewFromDomainName(domain string) (PublicDomain, error) {

	_registeredName, err := registered_name.NewFromString(domain)
	if err != nil {
		return PublicDomain{}, err
	}

	return PublicDomain{
		domainId:       id.New(0),
		registeredName: _registeredName,
	}, nil

}

func (pd *PublicDomain) UpdateId(domainId int64) {
	pd.domainId = id.New(domainId)
}

func (pd PublicDomain) Domain() string {
	return pd.registeredName.Value()
}

func (pd PublicDomain) Id() int64 {
	return pd.domainId.AsInteger()
}
