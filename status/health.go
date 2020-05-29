package status

import (
	"fmt"
	"net/http"
)

func Health(w http.ResponseWriter, _ *http.Request) {

	w.Header().Set("content-type", "text/plain")

	status := "Ok"
	/*
		output := strings.NewReader(status)
		io.Copy(w, output)
	*/

	//w.Write([]byte(status))

	_, _ = fmt.Fprintf(w, "Health: %v", status)

}
