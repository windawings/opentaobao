package jst

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/jst"
)

// TaobaoRdsDbCreateaccount rds创建数据库账户
// taobao.rds.db.createaccount
//
// rds创建数据库账户
func TaobaoRdsDbCreateaccount(ctx context.Context, clt *core.SDKClient, req *jst.TaobaoRdsDbCreateaccountAPIRequest, resp *jst.TaobaoRdsDbCreateaccountAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
