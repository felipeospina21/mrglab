package gitlab

import "context"

// CreateNote posts a new comment on a discussion thread.
func (c *Client) CreateNote(input CreateNoteInput) (CreateNoteResponse, error) {
	if c.devMode {
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
