package register_public_domain

// RegistrationGateway
/*
Interface Description
The interface which defines the behaviors expected when invoking a gateway implementation for the public domain
registration usecase.

Constraints
The implementation is responsible for persisting the public domain data,
and is additionally responsible for constraint checking that public domain name does not already exist.

Returns two (2) types of errors:
- InternalError - Returned if there is an error invoking the implementation of this interface
- AlreadyExists - Returned if the entity being registered already exists
 */
type RegistrationGateway interface {
	RegisterPublicDomain(userId int64, publicDomain PublicDomainEntity) error
}
