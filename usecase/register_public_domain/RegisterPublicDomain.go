package register_public_domain

import (
	"errors"
	"github.com/attestify/go-domain/gateway"
)

type RegisterPublicDomain struct {
	identityGateway gateway.IdentityGateway
	registrationGateway RegistrationGateway
}

func New(identityGateway gateway.IdentityGateway, registrationGateway RegistrationGateway) (RegisterPublicDomain, error) {
	if identityGateway == nil {
		_error := errors.New("the provided identity gateway is nil, please provide a valid instance of an identity gateway")
		return RegisterPublicDomain{}, _error
	}

	return RegisterPublicDomain{
		identityGateway: identityGateway,
		registrationGateway: registrationGateway,
	}, nil
}
