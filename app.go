package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/fabian-lapotre/document-api/server"
	"github.com/oklog/oklog/pkg/group"
)

func main() {

	var g group.Group

	{
		myServer := server.Create()
		srv := &http.Server{
			Addr:    ":8080",
			Handler: myServer,
		}
		g.Add(func() error {
			return srv.ListenAndServe()
		}, func(error) {
			srv.Close()
		})
	}
	{
		// This function just sits and waits for ctrl-C.
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}

	if err := g.Run(); err != nil {
		log.Panicln(err)
	}

}
