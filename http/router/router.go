package router

import (
	"errors"
	"github.com/andrew-petrenko/golang-todo-api/http/controllers"
	br "github.com/andrew-petrenko/golang-todo-api/http/resources/base-response"
	"github.com/andrew-petrenko/golang-todo-api/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"net/http"
	"os"
	"strings"
)

func InitRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/api/test", controllers.Test)

	r.Post("/api/auth/register", controllers.Register)
	r.Post("/api/auth/login", controllers.Login)

	authGroup := r.Group(nil)
	authGroup.Use(authMiddleware)
	authGroup.Get("/api/projects", controllers.GetAllProjects)
	authGroup.Post("/api/projects", controllers.CreateProject)
	//authGroup.Patch("/api/projects/{id}", controllers.UpdateProject)
	authGroup.Get("/api/projects/{id}", controllers.GetOneProject)
	authGroup.Delete("/api/projects/{id}", controllers.DeleteProject)

	return r
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		strArr := strings.Split(authHeader, " ")
		if len(strArr) != 2 {
			utils.WriteResponse(w, br.NewResponse(map[string]string{"message": "Forbidden"}, false), http.StatusForbidden)
			return
		}

		_, err := jwt.Parse(strArr[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("Unexpected signing method")
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			utils.WriteResponse(w, br.NewResponse(map[string]string{"message": "Forbidden"}, false), http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
