package register_public_domain

// RequestForPublicDomain is the interface which defines the behaviors expected when incoking the use case to
// register a public domain
type RequestForPublicDomain interface {
	Domain() string
}