package http

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"nhooyr.io/websocket"
	"qubes/internal/config"
	"qubes/internal/ws"
)

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Server struct {
	ctx      context.Context
	logger   *zap.SugaredLogger
	router   *mux.Router
	cfg      *config.AppConfig
	wsServer *ws.Server
}

func New(ctx context.Context, logger *zap.SugaredLogger, cfg *config.AppConfig, serv *ws.Server) *Server {
	return &Server{
		ctx:      ctx,
		logger:   logger,
		cfg:      cfg,
		wsServer: serv,
	}
}

func (s *Server) Start() error {
	serv := http.Server{
		Addr:    fmt.Sprintf("%s:%s", s.cfg.HTTP.Host, s.cfg.HTTP.Port),
		Handler: s.initRouter(),
	}
	s.logger.Infof("Listening at %s...", serv.Addr)
	return serv.ListenAndServe()
}
func (s *Server) initRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", hello)
	r.HandleFunc("/ws", s.handleWS)
	r.Handle("/metrics", promhttp.Handler())
	return r
}

func hello(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("hello!"))
}

func (s *Server) handleWS(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})
	if err != nil {
		s.logger.Info(err)
		return
	}

	go s.wsServer.HandleConn(c)

}
