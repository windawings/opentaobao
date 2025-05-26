package servicecenter

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/servicecenter"
)

// TaobaoVasServiceValidate 增值服务订购服务验证
// taobao.vas.service.validate
//
// 增值服务订购服务验证
func TaobaoVasServiceValidate(ctx context.Context, clt *core.SDKClient, req *servicecenter.TaobaoVasServiceValidateAPIRequest, resp *servicecenter.TaobaoVasServiceValidateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
