package external_id

func New(id string) *ExternalId {
	return &ExternalId{
		value: id,
	}
}
