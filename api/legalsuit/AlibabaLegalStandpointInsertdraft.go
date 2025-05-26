package legalsuit

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/legalsuit"
)

// AlibabaLegalStandpointInsertdraft 插入草稿
// alibaba.legal.standpoint.insertdraft
//
// 插入草稿
func AlibabaLegalStandpointInsertdraft(ctx context.Context, clt *core.SDKClient, req *legalsuit.AlibabaLegalStandpointInsertdraftAPIRequest, resp *legalsuit.AlibabaLegalStandpointInsertdraftAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
