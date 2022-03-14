// Code generated by entc, DO NOT EDIT.

package runtime

import (
	"github.com/yonidavidson/ent-hooks-examples/ent/dog"
	"github.com/yonidavidson/ent-hooks-examples/ent/schema"
	"github.com/yonidavidson/ent-hooks-examples/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	dogHooks := schema.Dog{}.Hooks()
	dog.Hooks[0] = dogHooks[0]
	dogFields := schema.Dog{}.Fields()
	_ = dogFields
	// dogDescName is the schema descriptor for name field.
	dogDescName := dogFields[0].Descriptor()
	// dog.NameValidator is a validator for the "name" field. It is called by the builders before save.
	dog.NameValidator = dogDescName.Validators[0].(func(string) error)
	userHooks := schema.User{}.Hooks()
	user.Hooks[0] = userHooks[0]
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescConnectionString is the schema descriptor for connection_string field.
	userDescConnectionString := userFields[1].Descriptor()
	// user.ConnectionStringValidator is a validator for the "connection_string" field. It is called by the builders before save.
	user.ConnectionStringValidator = userDescConnectionString.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
}

const (
	Version = "v0.10.1"                                         // Version of ent codegen.
	Sum     = "h1:dM5h4Zk6yHGIgw4dCqVzGw3nWgpGYJiV4/kyHEF6PFo=" // Sum of ent codegen.
)
