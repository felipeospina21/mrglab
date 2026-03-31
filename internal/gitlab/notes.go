package gitlab

import (
	"context"
	"time"
)

// CreateNote posts a new comment on a discussion thread.
func (c *Client) CreateNote(input CreateNoteInput) (CreateNoteResponse, error) {
	if c.devMode {
		c.sleep(500 * time.Millisecond)
		return CreateNoteResponse{}, nil
	}

	var mutation createNoteMutation
	variables := createNoteVariables(input)

	err := c.gql.Mutate(context.Background(), &mutation, variables)
	if err != nil {
		return CreateNoteResponse{}, err
	}

	return mutation.CreateNote, nil
}
