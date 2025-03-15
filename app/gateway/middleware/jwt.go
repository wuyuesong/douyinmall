/*
 * Copyright 2022 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package middleware

import (
	"context"
	"net/http"
	"time"

	// "github.com/cloudwego/hertz-examples/bizdemo/hertz_jwt/biz/model"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"
	"github.com/wuyuesong/douyinmall/app/gateway/biz/constants"
	"github.com/wuyuesong/douyinmall/app/gateway/biz/service"
	gatewayUtils "github.com/wuyuesong/douyinmall/app/gateway/biz/utils"
	"github.com/wuyuesong/douyinmall/app/gateway/hertz_gen/gateway/auth"
	"github.com/wuyuesong/douyinmall/rpc_gen/kitex_gen/user"
)

var JwtMiddleware *jwt.HertzJWTMiddleware

func InitJwt(secretKey string) {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:         "test zone",
		Key:           []byte(secretKey),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			if cartNum, exists := c.Get("cartNum"); exists {
				c.JSON(http.StatusOK, utils.H{
					"code":    code,
					"token":   token,
					"expire":  expire.Format(time.RFC3339),
					"cartNum": cartNum, // 添加购物车数量字段
					"message": "success",
				})
			} else {
				// 容错处理（根据业务需求选择是否强制返回错误）
				c.JSON(http.StatusOK, utils.H{
					"code":    code,
					"token":   token,
					"expire":  expire.Format(time.RFC3339),
					"cartNum": 0, // 默认值或错误标识
					"message": "success",
				})
			}
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var err error
			var req auth.LoginReq
			err = c.BindAndValidate(&req)
			if err != nil {
				gatewayUtils.SendErrResponse(ctx, c, consts.StatusOK, err)
				return nil, err
			}

			LoginResp, err := service.NewLoginService(ctx, c).Run(&req)
			if err != nil {
				gatewayUtils.SendErrResponse(ctx, c, consts.StatusUnauthorized, err)
				return nil, err
			}
			c.Set("cartNum", LoginResp.CartNum)
			return LoginResp, nil
		},
		IdentityKey: constants.IdentityKey,
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			resp, _ := c.Get(constants.IdentityKey)
			if resp.(user.LoginResp).Role == "admin" || resp.(user.LoginResp).Role == "user" {
				return true
			}
			return false
		},
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return user.LoginResp{
				UserId: int32(claims[constants.IdentityKey].(float64)),
				Role:   claims["role"].(string),
			}
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*user.LoginResp); ok {
				return jwt.MapClaims{
					constants.IdentityKey: v.UserId,
					"role":                v.Role,
				}
			}
			return jwt.MapClaims{}
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "jwt biz err = %+v", e.Error())
			return e.Error()
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusUnauthorized, utils.H{
				"code":    code,
				"message": message,
			})
		},
	})
	if err != nil {
		panic(err)
	}
}
