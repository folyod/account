package name

type Name struct {
	value string
}

func (n *Name) Value() string {
	return n.value
}

func (n *Name) Equal(name Name) bool {
	return n.Value() == name.Value()
}
