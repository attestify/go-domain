package gateway

type IdentityGateway interface {
	GenerateId() int64
}
