package register_public_domain

// PublicDomainRequest is a default implementation of the RequestForPublicDomain
type PublicDomainRequest struct {
	domain string
}

func NewPublicDomainRequest(domain string) (*PublicDomainRequest, error) {
	return &PublicDomainRequest {
		domain: domain,
	}, nil
}

func (request *PublicDomainRequest) Domain() string {
	return request.domain
}