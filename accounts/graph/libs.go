package graph

import (
	"learn-apollo-federation-gqlgen/accounts/graph/model"
	"learn-apollo-federation-gqlgen/accounts/proto/generated"
	"strconv"
)

func convertUser(user *generated.User) *model.User {
	if user == nil {
		return nil
	}

	return &model.User{
		ID:       strconv.FormatUint(user.Id, 10),
		Username: user.Username,
	}
}
