package tanx

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tanx"
)

// TaobaoTanxAuditCreativeModify 创意修改接口
// taobao.tanx.audit.creative.modify
//
// 创意修改接口
func TaobaoTanxAuditCreativeModify(ctx context.Context, clt *core.SDKClient, req *tanx.TaobaoTanxAuditCreativeModifyAPIRequest, resp *tanx.TaobaoTanxAuditCreativeModifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
