package main

import (
	"github.com/sklaski/gosea/status"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	//log.Print("starting service")
	/*	logfileName := "message.log"
		logfile, logErr := os.Create(logfileName)
		if logErr != nil {
			log.Fatalf("error opening %v, %s", logfileName, logErr.Error())
		}

		defer logfile.Close()
	*/
	logger := log.New(os.Stdout, "gosea", log.LstdFlags)

	sigChan := make(chan os.Signal)
	defer close(sigChan)

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", status.Health)

	srv := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	defer srv.Close()

	go func() {
		err := srv.ListenAndServe()
		//if err != nil && !errors.Is(err, http.ErrServerClosed) {
		if err != nil {
			logger.Fatalf("error starting service: %s", err.Error())
		}
	}()

	logger.Print("starting service")

	<-sigChan

	logger.Print("stopping service")

}
