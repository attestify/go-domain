package gateway

// IdentityGateway provides the behavior for generating sequential unique identifiers.
// Returns one (1) type of error:
//   - - InternalError - Returned if there is an error invoking the implementation of this interface
type IdentityGateway interface {
	GenerateId() (int64, error)
}
