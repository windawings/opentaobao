package mos

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/mos"
)

// AlibabaMjOcWritesaleslip 开票占库
// alibaba.mj.oc.writesaleslip
//
// 开票占库
func AlibabaMjOcWritesaleslip(ctx context.Context, clt *core.SDKClient, req *mos.AlibabaMjOcWritesaleslipAPIRequest, resp *mos.AlibabaMjOcWritesaleslipAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
