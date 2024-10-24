package data

import "github.com/felipeospina21/mrglab/internal/gql"

var GQLDiscussionsMock = gql.MergeRequestNotesConnection{
	Count: 1,
	Edges: []gql.MergeRequestNotesEdges{
		{
			Node: gql.MergeRequestNotesNode{
				Id: "1",
				Discussion: gql.MergeRequestDiscussion{
					Notes: gql.MergeRequestNoteConnection{
						Count: 1,
						Edges: []gql.MergeRequestNoteEdges{
							{
								Node: gql.MergeRequestNote{
									Author: gql.Author{Name: "Mock User"},
									Body:   "some body",
								},
							},
						},
					},
				},
			},
		},
	},
}
