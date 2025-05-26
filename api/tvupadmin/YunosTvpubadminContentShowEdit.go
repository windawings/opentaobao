package tvupadmin

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentShowEdit 媒资节目信息修改
// yunos.tvpubadmin.content.show.edit
//
// 供迎客松修改媒资节目信息
func YunosTvpubadminContentShowEdit(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentShowEditAPIRequest, resp *tvupadmin.YunosTvpubadminContentShowEditAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
