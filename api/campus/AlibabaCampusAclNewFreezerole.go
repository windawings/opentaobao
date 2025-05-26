package campus

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/campus"
)

// AlibabaCampusAclNewFreezerole 冻结角色
// alibaba.campus.acl.new.freezerole
//
// 冻结角色
func AlibabaCampusAclNewFreezerole(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusAclNewFreezeroleAPIRequest, resp *campus.AlibabaCampusAclNewFreezeroleAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
