package account

type CreateAccountCommand struct {
	name     string
	login    string
	email    string
	password string
}

func (command CreateAccountCommand) Name() string {
	return command.name
}

func (command CreateAccountCommand) Login() string {
	return command.login
}

func (command CreateAccountCommand) Email() string {
	return command.email
}

func (command CreateAccountCommand) Password() string {
	return command.password
}
