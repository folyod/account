package email

func Make(email string) (*Email, error) {
	vo := &Email{
		value: email,
	}
	err := vo.Validate()
	if err != nil {
		return nil, err
	}
	return vo, nil
}
