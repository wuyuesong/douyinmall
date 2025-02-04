# .PHONY:	gen-frontend
# gen-frontend:
# 	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/home.proto  --service frontend -module github.com/wuyuesong/gomall/app/frontend -I ../../idl

# .PHONY:	gen-auth-page
# gen-auth-page:
# 	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/auth_page.proto  --service frontend -module github.com/wuyuesong/gomall/app/frontend -I ../../idl

# .PHONY:	gen-rpc-client
# gen-rpc-client:
# 	@cd rpc_gen && cwgo client --type RPC --idl ../idl/user.proto --service user -module github.com/wuyuesong/gomall/rpc_gen -I ../idl 

# .PHONY:	gen-rpc-server
# gen-rpc-server:
# 	@cd app/user && cwgo server --type RPC --idl ../../idl/user.proto --service user -module github.com/wuyuesong/gomall/app/user -I ../../idl --pass "-use github.com/wuyuesong/gomall/rpc_gen/kitex_gen"

.PHONY:	all
all:	gen-frontend gen-auth-page gen-rpc-client gen-rpc-server

.PHONY: gen-frontend
gen-frontend: gen-frontend-home gen-product-page gen-category-page gen-auth-page gen-cart-page

.PHONY:	gen-frontend-home
gen-frontend-home:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/home.proto  --service frontend -module github.com/wuyuesong/gomall/app/frontend -I ../../idl

.PHONY:	gen-auth-page
gen-auth-page:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/auth_page.proto  --service frontend -module github.com/wuyuesong/gomall/app/frontend -I ../../idl


.PHONY: gen-user
gen-user: gen-user-rpc-client gen-user-rpc-server

.PHONY:	gen-user-rpc-client
gen-user-rpc-client:
	@cd rpc_gen && cwgo client --type RPC --idl ../idl/user.proto --service user -module github.com/wuyuesong/gomall -I ../idl 

.PHONY:	gen-user-rpc-server
gen-user-rpc-server:
	@cd app/user && cwgo server --type RPC --idl ../../idl/user.proto --service user -module github.com/wuyuesong/gomall -I ../../idl --pass "-use github.com/wuyuesong/gomall/rpc_gen/kitex_gen"

.PHONY: gen-product
gen-product:
	@cd rpc_gen && cwgo client --type RPC --idl ../idl/product.proto --service product -module github.com/wuyuesong/gomall -I ../idl 
	@cd app/product && cwgo server --type RPC --idl ../../idl/product.proto --service product -module github.com/wuyuesong/gomall -I ../../idl --pass "-use github.com/wuyuesong/gomall/rpc_gen/kitex_gen"


.PHONY:	gen-product-page
gen-product-page:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/product_page.proto  --service frontend -module github.com/wuyuesong/gomall/app/frontend -I ../../idl

.PHONY:	gen-category-page
gen-category-page:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/category_page.proto  --service frontend -module github.com/wuyuesong/gomall/app/frontend -I ../../idl

.PHONY:	gen-cart-page
gen-cart-page:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/cart_page.proto  --service frontend -module github.com/wuyuesong/gomall/app/frontend -I ../../idl

.PHONY: gen-cart
gen-cart:
	@cd rpc_gen && cwgo client --type RPC --idl ../idl/cart.proto --service cart -module github.com/wuyuesong/gomall -I ../idl 
	@cd app/cart && cwgo server --type RPC --idl ../../idl/cart.proto --service cart -module github.com/wuyuesong/gomall -I ../../idl --pass "-use github.com/wuyuesong/gomall/rpc_gen/kitex_gen"

