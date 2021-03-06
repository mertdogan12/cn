package api

import (
	"fmt"
	"net/http"

	"github.com/mertdogan12/cn/internal/database"
	"go.mongodb.org/mongo-driver/mongo"
)

func ApiRespond(status int, message string, w http.ResponseWriter) {
	w.WriteHeader(status)
	w.Write([]byte(message))
	fmt.Println("users/me:", status, "|", message)
}

func ApiRespondErr(err error, w http.ResponseWriter) {
	ApiRespond(http.StatusInternalServerError, "Error", w)
	fmt.Println(err)
}

func GetName(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query()["id"]
	if len(id) == 0 {
		ApiRespond(http.StatusNotImplemented, "Parameter id fehlt", w)
		return
	}

	name, err := database.GetName(id[0])
	if err == mongo.ErrNoDocuments {
		ApiRespond(http.StatusNoContent, "Konnte keinen Name unter der Id "+id[0]+" finden", w)
		return
	}
	if err != nil {
		ApiRespondErr(err, w)
	}

	ApiRespond(http.StatusOK, name, w)
}
