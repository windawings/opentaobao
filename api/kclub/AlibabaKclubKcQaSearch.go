package kclub

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/kclub"
)

// AlibabaKclubKcQaSearch 知识云-知识检索
// alibaba.kclub.kc.qa.search
//
// 知识云-知识搜索服务
func AlibabaKclubKcQaSearch(ctx context.Context, clt *core.SDKClient, req *kclub.AlibabaKclubKcQaSearchAPIRequest, resp *kclub.AlibabaKclubKcQaSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
