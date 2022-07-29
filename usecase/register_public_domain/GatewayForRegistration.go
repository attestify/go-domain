package register_public_domain

// GatewayForRegistration is the interface which defines the behaviors expected when invoking a gateway
// implementation for the public domain registration usecase.
type GatewayForRegistration interface {
	 RegisterDomain(publicDomain PublicDomain) error
}
