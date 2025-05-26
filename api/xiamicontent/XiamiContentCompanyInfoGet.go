package xiamicontent

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/xiamicontent"
)

// XiamiContentCompanyInfoGet 获取厂牌信息
// xiami.content.company.info.get
//
// 获取厂牌信息
func XiamiContentCompanyInfoGet(ctx context.Context, clt *core.SDKClient, req *xiamicontent.XiamiContentCompanyInfoGetAPIRequest, resp *xiamicontent.XiamiContentCompanyInfoGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
