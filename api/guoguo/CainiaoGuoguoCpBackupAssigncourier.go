package guoguo

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/guoguo"
)

// CainiaoGuoguoCpBackupAssigncourier CP兜底后指定接单的小件员
// cainiao.guoguo.cp.backup.assigncourier
//
// CP兜底后指定接单的小件员；CP改派小件员
func CainiaoGuoguoCpBackupAssigncourier(ctx context.Context, clt *core.SDKClient, req *guoguo.CainiaoGuoguoCpBackupAssigncourierAPIRequest, resp *guoguo.CainiaoGuoguoCpBackupAssigncourierAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
