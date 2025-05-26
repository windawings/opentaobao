package aliexpress

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/aliexpress"
)

// AliexpressSocialInsDirectresultUpdate ISV更新INS私信发送的结果
// aliexpress.social.ins.directresult.update
//
// ISV更新INS私信发送的结果
func AliexpressSocialInsDirectresultUpdate(ctx context.Context, clt *core.SDKClient, req *aliexpress.AliexpressSocialInsDirectresultUpdateAPIRequest, resp *aliexpress.AliexpressSocialInsDirectresultUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
