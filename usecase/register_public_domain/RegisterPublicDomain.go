package register_public_domain

import "github.com/attestify/go-domain/gateway"

type RegisterPublicDomain struct {
	identityGateway gateway.IdentityGateway
	registrationGateway RegistrationGateway
}

func New(identityGateway gateway.IdentityGateway, registrationGateway RegistrationGateway) (RegisterPublicDomain, error) {
	return RegisterPublicDomain{
		identityGateway: identityGateway,
		registrationGateway: registrationGateway,
	}, nil
}
