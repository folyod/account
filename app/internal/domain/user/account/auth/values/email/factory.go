package email

func Make(email string) (Email, error) {
	vo := Email{
		value: email,
	}

	return vo, vo.Validate()
}
