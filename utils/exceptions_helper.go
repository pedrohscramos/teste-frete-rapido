package utils

import "net/http"

func Error(err error, w http.ResponseWriter) {
	if err != nil {
		if w != nil {
			if err.Error() == "404" {
				http.Error(w, "Not Found", http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
		} else {
			panic(err.Error())
		}
	}
}
