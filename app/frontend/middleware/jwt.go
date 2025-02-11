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
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"
	"github.com/wuyuesong/douyinmall/app/frontend/biz/service"
	frontendUtils "github.com/wuyuesong/douyinmall/app/frontend/biz/utils"
	"github.com/wuyuesong/douyinmall/app/frontend/hertz_gen/frontend/auth"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	IdentityKey   = "identity"
)

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
			var req auth.LoginReq
			err = c.BindAndValidate(&req)
			redirect := "/"
			if req.Next != "" {
				redirect = req.Next
			}

			c.SetCookie("jwt", token, int(expire.Sub(time.Now()).Seconds()), "/", "", protocol.CookieSameSiteLaxMode, false, true)

			c.Redirect(consts.StatusOK, []byte(redirect))
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var err error
			var req auth.LoginReq
			err = c.BindAndValidate(&req)
			if err != nil {
				frontendUtils.SendErrResponse(ctx, c, consts.StatusOK, err)
				return nil, err
			}

			userId, err := service.NewLoginService(ctx, c).Run(&req)
			if err != nil {
				frontendUtils.SendErrResponse(ctx, c, consts.StatusOK, err)
				return nil, err
			}
			return userId, nil
		},
		IdentityKey: IdentityKey,
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return claims[IdentityKey]
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int32); ok {
				return jwt.MapClaims{
					IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "jwt biz err = %+v", e.Error())
			return e.Error()
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				"code":    code,
				"message": message,
			})
		},
	})
	if err != nil {
		panic(err)
	}
}
