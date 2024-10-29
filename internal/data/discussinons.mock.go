package data

import (
	"time"

	"github.com/felipeospina21/mrglab/internal/gql"
)

var GQLDiscussionsMock = gql.MergeRequestDiscussionsConnection{
	Nodes: []gql.DiscussionNode{
		{
			Resolvable: true,
			Resolved:   false,
			Notes: gql.NoteConnection{
				Nodes: []gql.Note{
					{
						Resolvable: true,
						Author:     gql.Author{Name: "Mock User"},
						Body:       "Question about these functions that were a `useCallback`: \n\ni.e `handleToggleOpen` is a prop for an `Accordion` component, don't we need to keep the `useCallback` and do something like\n\n```\n  const handleToggleOpen = useCallback(\n    (panelId: string, nextValue: boolean): void => {\n      togglePanel(panelId, nextValue, openPanels, setOpenPanels);\n    },\n    [openPanels, setOpenPanels]\n  );\n```",
						CreatedAt:  time.Now(),
					},
					{
						Resolvable: true,
						Author:     gql.Author{Name: "Mock User"},
						Body:       "not really. The useCallback only makes sense if the function is defined within a React component. Since we extracted the function it won't be re-declared when the component re-renders, since it will always be pointing to the function reference (in another file). So basically when you import a function and used it in the component it has the same behavior as if it was declared inside the function within a useCallback (regarding the preservation of the reference).",
						CreatedAt:  time.Now(),
					},
					{
						Resolvable: true,
						Author:     gql.Author{Name: "Mock User"},
						Body:       "to be more specific, the getHandleToggleOpen function returns a function, so handleToggleOpen is just a reference/pointer.",
						CreatedAt:  time.Now(),
					},
				},
			},
		},
		{
			Resolvable: true,
			Resolved:   true,
			Notes: gql.NoteConnection{
				Nodes: []gql.Note{
					{
						Resolvable: true,
						Author:     gql.Author{Name: "Mock User"},
						Body:       "question",
						CreatedAt:  time.Now(),
					},
					{
						Resolvable: true,
						Author:     gql.Author{Name: "Mock User"},
						Body:       "response 1",
						CreatedAt:  time.Now(),
					},
					{
						Resolvable: false,
						Author:     gql.Author{Name: "Mock User"},
						Body:       "add commit 1 [in versin 4] ()",
						CreatedAt:  time.Now(),
					},
					{
						Resolvable: true,
						Author:     gql.Author{Name: "Mock User"},
						Body:       "response 2",
						CreatedAt:  time.Now(),
					},
				},
			},
		},
	},
}
