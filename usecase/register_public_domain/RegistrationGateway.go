package register_public_domain

// RegistrationGateway is the interface which defines the behaviors expected
// when invoking a gateway implementation for the public domain registration usecase.
type RegistrationGateway interface {
	RegisterPublicDomain(publicDomain PublicDomainEntity) error
}
