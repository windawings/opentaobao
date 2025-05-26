package jst

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/jst"
)

// TaobaoJstSmsSignnameDelete 淘宝短信签名删除
// taobao.jst.sms.signname.delete
//
// 淘宝短信签名删除
func TaobaoJstSmsSignnameDelete(ctx context.Context, clt *core.SDKClient, req *jst.TaobaoJstSmsSignnameDeleteAPIRequest, resp *jst.TaobaoJstSmsSignnameDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
