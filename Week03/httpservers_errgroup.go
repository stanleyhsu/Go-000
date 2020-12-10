package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func helloHttpMux() *http.ServeMux {
	helloMux := http.NewServeMux()
	helloMux.HandleFunc("/hello", hello)
	return helloMux
}

func ping(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Pong!")
}

func pingHttpMux() *http.ServeMux {
	pingMux := http.NewServeMux()
	pingMux.HandleFunc("/ping", ping)
	return pingMux
}

func userTimeout(name string ) <-chan time.Time {
	if name == "hello" {
		return time.After(10 * time.Second)
	} else {
		return time.After(20 * time.Second)
	}
}

func serveHttpServer(ctx context.Context, name string, url string, handle *http.ServeMux, stop chan struct{}) error {
	log.Printf("Start httpserver %s on %s\n", name, url)
	s := http.Server{Addr: url, Handler: handle}

	go func() {
		select {
		case <- userTimeout(name): //fixme: just for goroutine self quit, to test 一个退出，全部注销退出
			log.Println("user-defined timeout for stop httpserver", name)

		case <-stop:
			log.Println("user fire stop httpserver", name)

		case <-ctx.Done():
			log.Println("context trigger stop httpserver", name)

		}
		s.Shutdown(context.Background())
	}()
	return s.ListenAndServe()
}

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	stopCh := make(chan struct{})
	done := make(chan error, 2)

	servers := []struct {
		name   string
		url    string
		mux    *http.ServeMux
		stopCh chan struct{}
	}{
		{"hello", ":8080", helloHttpMux(), stopCh},
		{"ping", ":8081", pingHttpMux(), stopCh},
	}

	for _, server := range servers {
		s := server
		g.Go(func() error {
			return serveHttpServer(ctx, s.name, s.url, s.mux, s.stopCh)
		})
	}

	go func() {
		done <- g.Wait()
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	select {
	case s := <-c:
		log.Printf("get a signal %s\n", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Println("fire for stop all servers")
			close(stopCh)
		case syscall.SIGHUP:
		default:
		}
	case err := <-done:
		if err != nil {
			log.Println("http servers group error", err)
		}
	}

	time.Sleep(2 * time.Second)
}
