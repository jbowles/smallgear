package smallgear

import "net/http"

func KnnWildClassifyHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("knn wild text classifier"))
}
