package ping

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

type PingLogic struct{}

func NewPingLogic() *PingLogic {
	return &PingLogic{}
}

func (l *PingLogic) Ping(ctx context.Context, c *app.RequestContext) {
	c.JSON(200, "pong")
}
