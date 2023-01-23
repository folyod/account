package email

type createAuthDto interface {
	Email() string
	Password() string
	Phone() string
}

func New() {

}
