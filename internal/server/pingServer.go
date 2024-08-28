package server

import (
	"fmt"
	"github.com/tpl-x/conec/internal/config"
	"github.com/tpl-x/conec/internal/service"
	"github.com/tpl-x/conec/pkg/ping/v1/pingv1connect"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net/http"
	"strconv"
)

type PingServer struct {
	serverConfig *config.ServerConfig
	pingSvc      *service.PingService
	logger       *zap.Logger
}

func NewPingServer(serverConfig *config.ServerConfig, logger *zap.Logger, pingSvc *service.PingService) *PingServer {
	return &PingServer{
		serverConfig: serverConfig,
		logger:       logger,
		pingSvc:      pingSvc,
	}
}

func (p *PingServer) ListenAndServe() error {
	mux := http.NewServeMux()
	path, handler := pingv1connect.NewPingServiceHandler(p.pingSvc)
	mux.Handle(path, handler)
	p.logger.Info("try to start serve on", zap.String("port", strconv.FormatInt(p.serverConfig.BindPort, 10)))
	return http.ListenAndServe(
		fmt.Sprintf(":%d", p.serverConfig.BindPort),
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
