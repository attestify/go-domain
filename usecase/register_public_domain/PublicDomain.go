package register_public_domain

import "github.com/attestify/go-kernel/uri/registered_name"

// PublicDomain is the entity
type PublicDomain struct {
	registeredName registered_name.RegisteredName
}
