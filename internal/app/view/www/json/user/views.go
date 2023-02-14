package user

import (
	controller "account/internal/app/controler/user"
	"net/http"
)

func Register() func(r *http.Request, w *http.Response) {
	return func(r *http.Request, w *http.Response) {
		body : r.MultipartForm().Value
		request := controller.RegistrationRequest{
			Nickname: r.,
			Name:     "",
			Email:    "",
			Password: "",
		}

		controller.Register()
	}
}