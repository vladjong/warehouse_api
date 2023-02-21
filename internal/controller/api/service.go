package api

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
	"github.com/adough/warehouse_api/internal/repository"
	"github.com/adough/warehouse_api/internal/service/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

type Service struct {
	db      repository.Repository
	handler handler.Service
	cfg     config.Config
}

func New(db repository.Repository, handler handler.Service, cfg config.Config) *Service {
	return &Service{
		db:      db,
		handler: handler,
		cfg:     cfg,
	}
}

func (s *Service) Start() {
	server := &http.Server{
		Addr:    fmt.Sprintf("localhost%v", s.cfg.Listen.Port),
		Handler: s.initHandler(),
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

	log.Println("start service")

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf(err.Error())
	}

	<-serverCtx.Done()

}

func (s *Service) initHandler() http.Handler {
	rpcServer := rpc.NewServer()

	rpcServer.RegisterCodec(json.NewCodec(), "application/json")
	rpcServer.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	rpcServer.RegisterService(s.handler, "warehouse")

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Handle("/warehouse", rpcServer)
	return r
}
