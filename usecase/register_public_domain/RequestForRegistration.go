package register_public_domain

// RequestForRegistration is the interface which defines the behaviors expected when invoking the use case to
//register a public domain
type RequestForRegistration interface {
	Domain() string
}