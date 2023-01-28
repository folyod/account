package password

func Make(password string) Password {
	return Password{
		value: password,
	}
}

func Hashed(password string) Password {
	return Password{
		value: password,
	}
}
