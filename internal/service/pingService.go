package service

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	v1 "github.com/tpl-x/conec/pkg/ping/v1"
	"github.com/tpl-x/conec/pkg/ping/v1/pingv1connect"
	"go.uber.org/zap"
	"time"
)

var _ pingv1connect.PingServiceHandler = (*PingService)(nil)

type PingService struct {
	logger *zap.Logger
}

func NewPingService(logger *zap.Logger) *PingService {
	return &PingService{logger: logger}
}

func (p PingService) Ping(ctx context.Context, req *connect.Request[v1.PingRequest]) (*connect.Response[v1.PongResponse], error) {
	p.logger.Info("Request ", zap.Any("headers", req.Header()))
	res := connect.NewResponse(&v1.PongResponse{
		Text:      fmt.Sprintf("Hello, %s!", req.Msg.Text),
		Timestamp: time.Now().UnixNano(),
	})
	res.Header().Set("Ping-Version", "v1")
	return res, nil
}
