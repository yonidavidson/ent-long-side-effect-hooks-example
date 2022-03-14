// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/yonidavidson/ent-hooks-examples/ent/cache"
	"github.com/yonidavidson/ent-hooks-examples/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// ConnectionString holds the value of the "connection_string" field.
	ConnectionString string `json:"connection_string,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"-"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges      UserEdges `json:"edges"`
	user_cache *int
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Pets holds the value of the pets edge.
	Pets []*Dog `json:"pets,omitempty"`
	// Cache holds the value of the cache edge.
	Cache *Cache `json:"cache,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// PetsOrErr returns the Pets value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PetsOrErr() ([]*Dog, error) {
	if e.loadedTypes[0] {
		return e.Pets, nil
	}
	return nil, &NotLoadedError{edge: "pets"}
}

// CacheOrErr returns the Cache value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) CacheOrErr() (*Cache, error) {
	if e.loadedTypes[1] {
		if e.Cache == nil {
			// The edge cache was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: cache.Label}
		}
		return e.Cache, nil
	}
	return nil, &NotLoadedError{edge: "cache"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldName, user.FieldConnectionString, user.FieldPassword:
			values[i] = new(sql.NullString)
		case user.ForeignKeys[0]: // user_cache
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldConnectionString:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field connection_string", values[i])
			} else if value.Valid {
				u.ConnectionString = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_cache", value)
			} else if value.Valid {
				u.user_cache = new(int)
				*u.user_cache = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryPets queries the "pets" edge of the User entity.
func (u *User) QueryPets() *DogQuery {
	return (&UserClient{config: u.config}).QueryPets(u)
}

// QueryCache queries the "cache" edge of the User entity.
func (u *User) QueryCache() *CacheQuery {
	return (&UserClient{config: u.config}).QueryCache(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", name=")
	builder.WriteString(u.Name)
	builder.WriteString(", connection_string=")
	builder.WriteString(u.ConnectionString)
	builder.WriteString(", password=<sensitive>")
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
