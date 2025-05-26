package alisports

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alisports"
)

// AlibabaAlisportsPassportAccountGetaccountinfo 获取会员信息
// alibaba.alisports.passport.account.getaccountinfo
//
// 获取阿里体育会员信息
func AlibabaAlisportsPassportAccountGetaccountinfo(ctx context.Context, clt *core.SDKClient, req *alisports.AlibabaAlisportsPassportAccountGetaccountinfoAPIRequest, resp *alisports.AlibabaAlisportsPassportAccountGetaccountinfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
