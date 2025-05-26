package b2bcert

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/b2bcert"
)

// AlibabaAuthCertGet 获取证书数据
// alibaba.auth.cert.get
//
// 获取证书数据
func AlibabaAuthCertGet(ctx context.Context, clt *core.SDKClient, req *b2bcert.AlibabaAuthCertGetAPIRequest, resp *b2bcert.AlibabaAuthCertGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
