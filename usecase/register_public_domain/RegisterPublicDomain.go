package register_public_domain

import (
	"errors"
	"github.com/attestify/go-domain/gateway"
	"github.com/attestify/go-domain/usecase/register_public_domain/public_domain"
	"github.com/attestify/go-domain/usecase/register_public_domain/public_domain_request"
)

type RegisterPublicDomain struct {
	identityGateway     gateway.IdentityGateway
	registrationGateway RegistrationGateway
}

func New(identityGateway gateway.IdentityGateway, registrationGateway RegistrationGateway) (RegisterPublicDomain, error) {

	if identityGateway == nil {
		_error := errors.New("the provided identity gateway is nil, please provide a valid instance of an identity gateway")
		return RegisterPublicDomain{}, _error
	}

	if registrationGateway == nil {
		_error := errors.New("the provided RegistrationGateway is nil, please provide a valid instance of a RegistrationGateway")
		return RegisterPublicDomain{}, _error
	}

	return RegisterPublicDomain{
		identityGateway:     identityGateway,
		registrationGateway: registrationGateway,
	}, nil
}

func (usecsae RegisterPublicDomain) MakeRequest(request *public_domain_request.PublicDomainRequest) error {
	domainId := usecsae.identityGateway.GenerateId()
	request.UpdateDomainId(domainId)

	publicDomain, err := public_domain.New(request.DomainId(), request.Domain())
	if err != nil {
		return errors.New("error creating the PublicDomain entity: /n " + err.Error())
	}

	err = usecsae.registrationGateway.RegisterPublicDomain(publicDomain)
	if err != nil {
		return errors.New("error registering the PublicDomain entity using the Registration Gateway: /n " + err.
			Error())
	}

	return nil

}
