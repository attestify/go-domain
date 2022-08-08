package register_public_domain

import "github.com/attestify/go-domain/usecase/register_public_domain/public_domain_entity"

// GatewayForRegistration is the interface which defines the behaviors expected when invoking a gateway
// implementation for the public domain registration usecase.
type GatewayForRegistration interface {
	 RegisterDomain(publicDomain public_domain_entity.PublicDomain) error
}
