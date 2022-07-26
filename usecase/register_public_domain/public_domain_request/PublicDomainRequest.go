package public_domain_request

// PublicDomainRequest describes the data required to register a public domain for the RegisterPublicDomain usecase.
type PublicDomainRequest struct {
	userId   int64
	domainId int64
	domain   string
}

func New(userId int64, domain string) PublicDomainRequest {
	return PublicDomainRequest{
		userId:   userId,
		domainId: 0,
		domain:   domain,
	}
}

func (request PublicDomainRequest) Domain() string {
	return request.domain
}

func (request PublicDomainRequest) UserId() int64 {
	return request.userId
}

func (request PublicDomainRequest) DomainId() int64 {
	return request.domainId
}

func (request *PublicDomainRequest) UpdateDomainId(id int64) {
	request.domainId = id
}
