package user

type RegisterRequest interface {
	Name() string
	NickName() string
	Email() string
	Password() string
}

type RegisterViewModel interface {
}

func Register() {

}
