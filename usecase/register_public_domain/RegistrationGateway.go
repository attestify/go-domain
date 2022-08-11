package register_public_domain

// RegistrationGateway is the interface which defines the behaviors expected
// when invoking a gateway implementation for the public domain registration usecase.
// Returns an InternalError when there are issues invoking the interface
type RegistrationGateway interface {
	RegisterPublicDomain(userId int64, publicDomain PublicDomainEntity) error
}
