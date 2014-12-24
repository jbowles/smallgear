package smallgear

import "net/http"

func BayesWildClassifyHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("bayes wild text classifier"))
}
