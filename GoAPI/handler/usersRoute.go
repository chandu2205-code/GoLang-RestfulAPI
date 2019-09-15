package handler

import (
	"net/http"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

func UsersRouters(w http.ResponseWriter, r *http.Request) {

	trimmedURL := strings.TrimSuffix(r.URL.Path, "/")

	if trimmedURL == "/users" {

		switch r.Method {

		case http.MethodGet:
			usersGetAll(w, r)
			return
		case http.MethodPost:
			usersPostOne(w, r)
			return
		default:
			postError(w, http.StatusMethodNotAllowed)
		}
	}

	path := strings.TrimPrefix(r.URL.Path, "/users/")

	if !bson.IsObjectIdHex(path) {
		postError(w, http.StatusNotFound)
		return
	}

	id := bson.ObjectIdHex(path)

	switch r.Method {
	case http.MethodGet:
		usersGetOne(w, r, id)
		return
	case http.MethodPost:
		return
	case http.MethodPatch:
		usersPatchOne(w, r, id)
		return
	case http.MethodPut:
		usersPutOne(w, r, id)
		return
	case http.MethodDelete:
		usersDeleteOne(w, r, id)
		return
	default:
		postError(w, http.StatusMethodNotAllowed)
		return
	}

}
