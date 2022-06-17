package snapshots

type Order struct {
	Name     string
	Asc      bool
	Selected bool
}

var Orders = []Order{
	{Name: "Name", Asc: true},
	{Name: "Release Date", Asc: true, Selected: true},
}
