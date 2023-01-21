package nickname

func Make(nickname string) Nickname {
	return Nickname{
		value: nickname,
	}
}
