package http

import (
	"net/http"

	"github.com/golang/glog"
	"github.com/justinsb/scaler/cmd/scaler/options"
	"github.com/justinsb/scaler/pkg/graph"
	"github.com/justinsb/scaler/pkg/simulate"
)

type APIServer struct {
	server *http.Server
}

type HasState interface {
	Query() interface{}
}

func NewAPIServer(options *options.AutoScalerConfig, state HasState) (*APIServer, error) {
	mux := http.NewServeMux()

	//if *profiling {
	//	mux.HandleFunc("/debug/pprof/", pprof.Index)
	//	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	//	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	//}

	mux.Handle("/api/statz", &Targets{state: state})

	ui := &UI{
		simulatable: state.(simulate.Simulatable),
		graphable:   state.(graph.Graphable),
	}
	ui.AddHandlers(mux)

	server := &http.Server{
		Addr:    options.ListenAPI,
		Handler: mux,
	}
	a := &APIServer{
		server: server,
	}
	return a, nil
}

func (s *APIServer) Start(stopCh <-chan struct{}) error {
	go func() {
		<-stopCh
		s.server.Close()
	}()

	glog.Infof("API listening on %s", s.server.Addr)
	err := s.server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
