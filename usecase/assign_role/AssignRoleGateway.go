package assign_role

import "github.com/attestify/go-kernel/identity/id"

// AssignRoleGateway provides the behavior for persisting a role record
// Expected Alternative Behaviors
//	- If the combination already exists, AssignRoleGateway gateway will respond as if it was a successful record attempt.
// Returns one (1) type of error:
//   - - InternalError - Returned if there is an error invoking the implementation of this interface
type AssignRoleGateway interface {
	RecordAssignment(userId id.Id, entityId id.Id, entity string) error
}
