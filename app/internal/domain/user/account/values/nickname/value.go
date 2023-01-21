package nickname

type Nickname struct {
	value string
}

func (n *Nickname) Value() string {
	return n.value
}

func (n *Nickname) Equal(nickname Nickname) bool {
	return n.Value() == nickname.Value()
}
