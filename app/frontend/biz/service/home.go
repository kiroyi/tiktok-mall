// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package service

import (
	"context"

	common "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/infra/rpc"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (map[string]any, error) {
	products, err := rpc.ProductClient.ListProducts(h.Context, &product.ListProductsReq{})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"title": "Hot sale",
		"items": products.Products,
	}, nil

	// resp := make(map[string]any)
	// items := []map[string]any{
	// 	{"Name": "T-shirt-1", "Price": 100, "Picture": "/static/image/t-shirt-1.jpeg"},
	// 	{"Name": "T-shirt-2", "Price": 110, "Picture": "/static/image/t-shirt-1.jpeg"},
	// 	{"Name": "T-shirt-3", "Price": 120, "Picture": "/static/image/t-shirt-2.jpeg"},
	// 	{"Name": "T-shirt-4", "Price": 130, "Picture": "/static/image/notebook.jpeg"},
	// 	{"Name": "T-shirt-5", "Price": 140, "Picture": "/static/image/t-shirt-1.jpeg"},
	// 	{"Name": "T-shirt-6", "Price": 150, "Picture": "/static/image/t-shirt.jpeg"},
	// }
	// resp["Title"] = "Hot Sales"
	// resp["Items"] = items
	// return resp, nil
}
