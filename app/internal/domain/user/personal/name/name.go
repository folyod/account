package name

type Name struct {
	value string
}

func Make(name string) Name {
	return Name{
		value: name,
	}
}

func (n *Name) Value() string {
	return n.value
}

func (n *Name) Equal(name Name) bool {
	return n.Value() == name.Value()
}
