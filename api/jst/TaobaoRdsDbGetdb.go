package jst

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/jst"
)

// TaobaoRdsDbGetdb rds获取RDS的DB
// taobao.rds.db.getdb
//
// rds获取RDS的DB
func TaobaoRdsDbGetdb(ctx context.Context, clt *core.SDKClient, req *jst.TaobaoRdsDbGetdbAPIRequest, resp *jst.TaobaoRdsDbGetdbAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
