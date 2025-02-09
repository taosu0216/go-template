package v1

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"go-template/service/logic/ping"
)

func GeneratePingLogic() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		l := ping.NewPingLogic()
		l.Ping(ctx, c)
	}
}
