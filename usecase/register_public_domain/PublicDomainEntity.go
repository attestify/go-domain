package register_public_domain

// PublicDomainEntity defines the interface required when passing an
// instantiation of the public domain entity across software boundaries.
type PublicDomainEntity interface {
	Domain() string
	DomainId() int64
	Id() int64
}
