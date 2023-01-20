package app

type Rule interface {
	Validate() (bool, error)
}
