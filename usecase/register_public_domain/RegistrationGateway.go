package register_public_domain

// RegistrationGateway is the interface which defines the behaviors expected
// when invoking a gateway implementation for the public domain registration usecase.
// The implementation is responsible for persisting the publi domain data, and is
// additionally responsible for constraint checking that public domain name does not already exist.
// Returns two (2) types of errors:
//   - InternalError - Returned if there is an error invoking the implementation of this interface
//	 - AlreadyExists - Returned if the entity being registered already exists
type RegistrationGateway interface {
	RegisterPublicDomain(userId int64, publicDomain PublicDomainEntity) error
}
