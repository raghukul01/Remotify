package api

import (
	"io"
	"net/http"
	"os/exec"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func Shutdown(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	key := vars["key"]
	actualKey := viper.GetString("SHUTDOWN_KEY")
	logrus.Info("got key = ", key)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if key != actualKey {
		io.WriteString(w, `{
			"shutdown": "no",
			"comment": "Invalid Key",
			}`,
		)
		return
	} else {
		io.WriteString(w, `{
			"shutdown": "yes",
			"comment": "closing...",
			}`,
		)
	}
	cmd := exec.Command("shutdown", "now")
	cmd.Run()
}

func addShutdownAPI(router *mux.Router) {
	router.HandleFunc(
		"/Shutdown/{key}", Shutdown,
	).Methods(http.MethodGet)
}

func addGetApis(router *mux.Router) {
	addShutdownAPI(router)
}
