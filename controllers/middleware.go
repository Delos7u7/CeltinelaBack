package controllers

import (
	"CeltinelaBack/models"
	"fmt"
	"net/http"
)

type MiddleWareController struct {
	MiddleWareService models.UserService
}

func (mc MiddleWareController) AuthenticationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		fmt.Println(cookie)
		if err != nil {
			if err == http.ErrNoCookie {
				fmt.Println("Cookie no encontradaa")
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		token := cookie.Value
		id, err := mc.MiddleWareService.ConsultaID(token)
		if err != nil || id == 0 {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	})
}
