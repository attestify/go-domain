package register_public_domain

import "github.com/attestify/go-domain/usecase/register_public_domain/public_domain"

// RegistrationGateway is the interface which defines the behaviors expected when invoking a gateway implementation for the public domain registration usecase.
type RegistrationGateway interface {
	 RegisterPublicDomain(publicDomain public_domain.PublicDomain) error
}
