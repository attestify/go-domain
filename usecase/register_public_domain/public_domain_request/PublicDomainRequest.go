package public_domain_request

import "github.com/attestify/go-kernel/identity/id"

// PublicDomainRequest is a default implementation of the RequestForPublicDomain
type PublicDomainRequest struct {
	userId id.Id
	domainId id.Id
	domain string
}

func NewPublicDomainRequest(userId int64, domain string) (*PublicDomainRequest, error) {
	return &PublicDomainRequest{
		userId: id.New(userId),
		domainId: id.New(0),
		domain: domain,
	}, nil
}

func (request *PublicDomainRequest) Domain() string {
	return request.domain
}