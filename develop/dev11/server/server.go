package server

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"net/http"
)

type Server struct {
	ctx    context.Context
	db     *pgxpool.Pool
	server *http.Server
}

func NewServer(
	ctx context.Context,
	server *http.Server,
	db *pgxpool.Pool,
) *Server {
	return &Server{
		ctx:    ctx,
		db:     db,
		server: server,
	}
}

func (s *Server) Init() {
	mux := http.NewServeMux()

	mux.HandleFunc("/create_event", s.createEvent)
	mux.HandleFunc("/update_event", s.updateEvent)
	mux.HandleFunc("/delete_event", s.deleteEvent)
	mux.HandleFunc("/events_for_day", s.eventsForDay)
	mux.HandleFunc("/events_for_week", s.eventsForWeek)
	mux.HandleFunc("/events_for_month", s.eventsForMonth)

	wrappedMux := NewLogger(mux)
	s.server.Handler = wrappedMux
}

func (s *Server) Run(ctx context.Context) {
	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			fmt.Print(err)
		}
	}()
}

func (s *Server) Shutdown() {
	s.server.Shutdown(s.ctx)
}
