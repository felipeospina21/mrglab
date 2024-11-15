package api

import (
	"context"

	"github.com/felipeospina21/mrglab/internal/config"
	"github.com/felipeospina21/mrglab/internal/gql"
)

// Respond to a discussion
func CreateNote(
	input gql.CreateNoteInput,
) (gql.CreateNoteResponse, error) {
	cfg := &config.GlobalConfig

	if cfg.DevMode {
		return gql.CreateNoteResponse{}, nil
	}

	var mutation gql.CreateNote

	variables := gql.CreateNoteVariables(input)

	client := newClient()

	err := client.Mutate(context.Background(), &mutation, variables)
	if err != nil {
		return gql.CreateNoteResponse{}, err
	}

	return mutation.CreateNote, nil
}
