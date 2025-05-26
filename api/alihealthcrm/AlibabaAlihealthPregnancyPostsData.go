package alihealthcrm

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alihealthcrm"
)

// AlibabaAlihealthPregnancyPostsData 发回帖子信息同步
// alibaba.alihealth.pregnancy.posts.data
//
// 发回帖子信息同步
func AlibabaAlihealthPregnancyPostsData(ctx context.Context, clt *core.SDKClient, req *alihealthcrm.AlibabaAlihealthPregnancyPostsDataAPIRequest, resp *alihealthcrm.AlibabaAlihealthPregnancyPostsDataAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
