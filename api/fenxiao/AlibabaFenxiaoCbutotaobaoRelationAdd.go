package fenxiao

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/fenxiao"
)

// AlibabaFenxiaoCbutotaobaoRelationAdd 1688分销铺货到淘宝关系添加
// alibaba.fenxiao.cbutotaobao.relation.add
//
// 1688分销铺货到淘宝关系添加
func AlibabaFenxiaoCbutotaobaoRelationAdd(ctx context.Context, clt *core.SDKClient, req *fenxiao.AlibabaFenxiaoCbutotaobaoRelationAddAPIRequest, resp *fenxiao.AlibabaFenxiaoCbutotaobaoRelationAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
