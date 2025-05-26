package kclub

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/kclub"
)

// AlibabaKclubKcQaSearchPage 知识云-知识检索(分页)
// alibaba.kclub.kc.qa.search.page
//
// 知识云-知识搜索服务
func AlibabaKclubKcQaSearchPage(ctx context.Context, clt *core.SDKClient, req *kclub.AlibabaKclubKcQaSearchPageAPIRequest, resp *kclub.AlibabaKclubKcQaSearchPageAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
