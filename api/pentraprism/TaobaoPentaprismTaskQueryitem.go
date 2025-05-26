package pentraprism

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/pentraprism"
)

// TaobaoPentaprismTaskQueryitem 查询任务当前进度
// taobao.pentaprism.task.queryitem
//
// 外网用户查询五棱镜任务系统当前进度
func TaobaoPentaprismTaskQueryitem(ctx context.Context, clt *core.SDKClient, req *pentraprism.TaobaoPentaprismTaskQueryitemAPIRequest, resp *pentraprism.TaobaoPentaprismTaskQueryitemAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
