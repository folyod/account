package email

func Make(email string) (Email, error) {
	object := Email{
		value: email,
	}

	return object, object.Validate()
}
