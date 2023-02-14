package external_id

type ExternalId struct {
	value string
}

func (e *ExternalId) Value() string {
	return e.value
}

func (e *ExternalId) String() string {
	return e.Value()
}

func (e *ExternalId) Equal(id ExternalId) bool {
	return e.Value() == id.Value()
}
