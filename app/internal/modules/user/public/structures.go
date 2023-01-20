package account

import "time"

type Account struct {
	uuid      string
	login     string
	createdAt time.Time
	personal  Personal
	auth      Auth
}

type Personal struct {
	name string
}

type Auth struct {
	email    string
	password string
}
