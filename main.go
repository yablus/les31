/*
Цель практической работы
Научиться:

писать микросервис и proxy,
тестировать написанное приложение.


Что нужно сделать
В прошлом домашнем задании вы писали приложение, которое принимает HTTP-запросы, создаёт пользователей, добавляет друзей и так далее. Давайте теперь приблизим наше приложение к реальному продукту.

Отрефакторьте приложение так, чтобы вы могли поднять две реплики данного приложения.
Используйте любую базу данных, чтобы сохранять информацию о пользователях, или можете сохранять информацию в файл, предварительно сереализуя в JSON.
Напишите proxy или используйте, например, nginx.
Протестируйте приложение.


Что оценивается
Дизайн API, работа в соответствии с функциональными требованиями и масштабирование. Тестирование сервиса.
*/

package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/yablus/les31/internal/handlers"
	"github.com/yablus/les31/internal/models"
)

func main() {
	r := SetupServer()
	http.ListenAndServe(":8080", r)
}

func SetupServer() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) { // GET "/"
		w.Write([]byte("OK"))
	})
	r.Mount("/users", UserRoutes())
	return r
}

func UserRoutes() chi.Router {
	r := chi.NewRouter()
	//u := &handlers.UserHandler{&test.FakeStorage{}}
	u := &handlers.UserHandler{models.NewStorage()}

	r.Get("/", u.ListUsers)                // GET /users
	r.Post("/", u.CreateUser)              // POST /users
	r.Put("/{id}", u.UpdateUser)           // PUT /users/{id}
	r.Delete("/", u.DeleteUser)            // DELETE /users
	r.Post("/make_friends", u.MakeFriends) // POST /users/make_friends
	r.Get("/{id}/friends", u.GetFriends)   // GET /users/{id}/friends
	return r
}
