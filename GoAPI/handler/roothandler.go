package handler

import "net/http"

/*
*Method name here starts with caps letter 'R'
*which makes function 'RootHandler' public which is
*visible across all packages
*@param writer which is a response writer
*@param reader which is a pointer to http.Request
 */
func RootHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Asset not found"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Go Lang Rest service \n"))
}
