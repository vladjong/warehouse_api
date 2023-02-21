package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/adough/warehouse_api/internal/config"
	"github.com/adough/warehouse_api/internal/service/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

type App struct {
	handler handler.Service
	cfg     config.Config
}

func New(handler handler.Service, cfg config.Config) *App {
	return &App{
		handler: handler,
		cfg:     cfg,
	}
}

func (a *App) Start() {
	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0%s", a.cfg.Listen.Port),
		Handler: a.initHandler(a.handler),
	}

	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig
		shutdownCtx, cancel := context.WithTimeout(serverCtx, 30*time.Second)
		defer cancel()

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Println("context deadline")
			}
		}()

		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Fatalf(err.Error())
		}
		serverStopCtx()
	}()

	log.Printf("start service at %s", server.Addr)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf(err.Error())
	}

	<-serverCtx.Done()

}

func (a *App) initHandler(handler handler.Service) http.Handler {
	rpcServer := rpc.NewServer()

	rpcServer.RegisterCodec(json.NewCodec(), "application/json")
	rpcServer.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	rpcServer.RegisterService(handler, "warehouse")

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Handle("/warehouse", rpcServer)
	return r
}
