package graphql

import (
	"fmt"
	"strconv"
)

// ID represents GraphQL's "ID" scalar type. A custom type may be used instead.
type ID string

func (ID) ImplementsGraphQLType(name string) bool {
	return name == "ID"
}

func (id *ID) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		*id = ID(input)
	case int32:
		*id = ID(strconv.Itoa(int(input)))
	default:
		err = fmt.Errorf("wrong type for ID: %T", input)
	}
	return err
}

func (id ID) MarshalJSON() ([]byte, error) {
	return strconv.AppendQuote(nil, string(id)), nil
}
