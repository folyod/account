package phone

func Make(phone string) (Phone, error) {
	phoneValue := Phone{
		value: phone,
	}

	return phoneValue, nil
}
