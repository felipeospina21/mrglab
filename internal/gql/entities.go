package gql

type Author struct {
	Name string
}

type Labels struct {
	Count int
	Edges []struct {
		Node struct {
			Color     string
			Title     string
			TextColor string
			ID        string
		}
	}
}
