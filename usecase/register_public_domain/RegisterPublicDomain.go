package register_public_domain

import (
	"github.com/attestify/go-domain/gateway"
	"github.com/attestify/go-domain/usecase/register_public_domain/public_domain"
	"github.com/attestify/go-domain/usecase/register_public_domain/public_domain_request"
	"github.com/attestify/go-kernel/error/internal_error"
)

type RegisterPublicDomain struct {
	identityGateway     gateway.IdentityGateway
	registrationGateway RegistrationGateway
}

// New constructs a RegisterPublicDomain instance
// Returns an InternalError if eiter of the arguments passed are nil
func New(identityGateway gateway.IdentityGateway, registrationGateway RegistrationGateway) (RegisterPublicDomain, error) {

	if identityGateway == nil {
		_error := internal_error.New("the provided identity gateway is nil, " +
			"please provide a valid instance of an identity gateway")
		return RegisterPublicDomain{}, _error
	}

	if registrationGateway == nil {
		_error := internal_error.New("the provided RegistrationGateway is nil, please provide a valid instance of a RegistrationGateway")
		return RegisterPublicDomain{}, _error
	}

	return RegisterPublicDomain{
		identityGateway:     identityGateway,
		registrationGateway: registrationGateway,
	}, nil
}

// Register invokes the registration logic for the PublicDomainRequest
// Returns two (2) types of errors:
//   - ValidationError - Returned if the request data does not validate per business rules
//   - InternalError - Returned if there is an error invoking actions across software boundaries
func (usecase RegisterPublicDomain) Register(request *public_domain_request.PublicDomainRequest) error {

	// First, Apply business validation logic to the request before invoking other systems.
	publicDomain, err := public_domain.New(request.DomainId(), request.Domain())
	if err != nil {
		return err
	}

	// Second, retrieve a unique-id for the entity
	domainId, err := usecase.identityGateway.GenerateId()
	if err != nil {
		return internal_error.New(err.Error())
	}

	// Third, update the domain id for the entity
	publicDomain.UpdateId(domainId)

	// Fourth, attempt to register the public domain with the gateway
	err = usecase.registrationGateway.RegisterPublicDomain(request.UserId(), publicDomain)
	if err != nil {
		return err
	}

	// Finally, update the request with the new domain Id, and return nil.
	// This signifies the use case transaction was successful
	request.UpdateDomainId(domainId)
	return nil

}
