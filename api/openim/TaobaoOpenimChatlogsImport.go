package openim

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/openim"
)

// TaobaoOpenimChatlogsImport openim单聊消息导入
// taobao.openim.chatlogs.import
//
// 提供openim账号的聊天消息导入功能
func TaobaoOpenimChatlogsImport(ctx context.Context, clt *core.SDKClient, req *openim.TaobaoOpenimChatlogsImportAPIRequest, resp *openim.TaobaoOpenimChatlogsImportAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
