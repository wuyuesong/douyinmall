# .PHONY:	gen-frontend
# gen-frontend:
# 	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/home.proto  --service frontend -module ${ROOT_MOD}/app/frontend -I ../../idl

# .PHONY:	gen-auth-page
# gen-auth-page:
# 	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/auth_page.proto  --service frontend -module ${ROOT_MOD}/app/frontend -I ../../idl

# .PHONY:	gen-rpc-client
# gen-rpc-client:
# 	@cd rpc_gen && cwgo client --type RPC --idl ../idl/user.proto --service user -module ${ROOT_MOD}/rpc_gen -I ../idl 

# .PHONY:	gen-rpc-server
# gen-rpc-server:
# 	@cd app/user && cwgo server --type RPC --idl ../../idl/user.proto --service user -module ${ROOT_MOD}/app/user -I ../../idl --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"
export ROOT_MOD=github.com/wuyuesong/douyinmall


.PHONY:	all
all: gen-gateway gen-rpc

.PHONY: gen-gateway
gen-gateway: gen-gateway-home gen-product-page gen-category-page gen-auth-page gen-cart-page gen-checkout-page gen-order-page gen-admin-page gen-image-api

.PHONY:	gen-gateway-home
gen-gateway-home:
	@cd app/gateway && cwgo server --type HTTP --idl ../../idl/gateway/home.proto  --service gateway --module ${ROOT_MOD}/app/gateway -I ../../idl

.PHONY:	gen-auth-page
gen-auth-page:
	@cd app/gateway && cwgo server --type HTTP --idl ../../idl/gateway/auth_page.proto  --service gateway --module ${ROOT_MOD}/app/gateway -I ../../idl

.PHONY:	gen-product-page
gen-product-page:
	@cd app/gateway && cwgo server --type HTTP --idl ../../idl/gateway/product_page.proto  --service gateway --module ${ROOT_MOD}/app/gateway -I ../../idl

.PHONY:	gen-category-page
gen-category-page:
	@cd app/gateway && cwgo server --type HTTP --idl ../../idl/gateway/category_page.proto  --service gateway --module ${ROOT_MOD}/app/gateway -I ../../idl

.PHONY:	gen-cart-page
gen-cart-page:
	@cd app/gateway && cwgo server --type HTTP --idl ../../idl/gateway/cart_page.proto  --service gateway --module ${ROOT_MOD}/app/gateway -I ../../idl

.PHONY:	gen-checkout-page
gen-checkout-page:
	@cd app/gateway && cwgo server --type HTTP --idl ../../idl/gateway/checkout_page.proto  --service gateway --module ${ROOT_MOD}/app/gateway -I ../../idl

.PHONY:	gen-order-page
gen-order-page:
	@cd app/gateway && cwgo server --type HTTP --idl ../../idl/gateway/order_page.proto  --service gateway --module ${ROOT_MOD}/app/gateway -I ../../idl

.PHONY:	gen-admin-page
gen-admin-page:
	@cd app/gateway && cwgo server --type HTTP --idl ../../idl/gateway/admin_page.proto  --service gateway --module ${ROOT_MOD}/app/gateway -I ../../idl

.PHONY:	gen-image-api
gen-image-api:
	@cd app/gateway && cwgo server --type HTTP --idl ../../idl/gateway/image_api.proto  --service gateway --module ${ROOT_MOD}/app/gateway -I ../../idl

.PHONY:	gen-rpc
gen-rpc: gen-user gen-product gen-cart gen-payment gen-checkout gen-order gen-email


.PHONY:	gen-user
gen-user:
	@cd rpc_gen && cwgo client --type RPC --idl ../idl/user.proto --service user -module ${ROOT_MOD}/rpc_gen -I ../idl 
	@cd app/user && cwgo server --type RPC --idl ../../idl/user.proto --service user -module ${ROOT_MOD}/app/user -I ../../idl --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"

.PHONY: gen-product
gen-product:
	@cd rpc_gen && cwgo client --type RPC --idl ../idl/product.proto --service product -module ${ROOT_MOD}/rpc_gen -I ../idl 
	@cd app/product && cwgo server --type RPC --idl ../../idl/product.proto --service product -module ${ROOT_MOD}/app/product -I ../../idl --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"


.PHONY: gen-cart
gen-cart:
	@cd rpc_gen && cwgo client --type RPC --idl ../idl/cart.proto --service cart -module ${ROOT_MOD}/rpc_gen -I ../idl 
	@cd app/cart && cwgo server --type RPC --idl ../../idl/cart.proto --service cart -module ${ROOT_MOD}/app/cart -I ../../idl --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"

.PHONY: gen-payment
gen-payment:
	@cd rpc_gen && cwgo client --type RPC --idl ../idl/payment.proto --service payment -module ${ROOT_MOD}/rpc_gen -I ../idl 
	@cd app/payment && cwgo server --type RPC --idl ../../idl/payment.proto --service payment -module ${ROOT_MOD}/app/payment -I ../../idl --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"

.PHONY: gen-checkout
gen-checkout:
	@cd rpc_gen && cwgo client --type RPC --idl ../idl/checkout.proto --service checkout -module ${ROOT_MOD}/rpc_gen -I ../idl 
	@cd app/checkout && cwgo server --type RPC --idl ../../idl/checkout.proto --service checkout -module ${ROOT_MOD}/app/checkout -I ../../idl --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"

.PHONY: gen-order
gen-order:
	@cd rpc_gen && cwgo client --type RPC --idl ../idl/order.proto --service order -module ${ROOT_MOD}/rpc_gen -I ../idl 
	@cd app/order && cwgo server --type RPC --idl ../../idl/order.proto --service order -module ${ROOT_MOD}/app/order -I ../../idl --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"

.PHONY: gen-email
gen-email:
	@cd rpc_gen && cwgo client --type RPC --idl ../idl/email.proto --service email -module ${ROOT_MOD}/rpc_gen -I ../idl 
	@cd app/email && cwgo server --type RPC --idl ../../idl/email.proto --service email -module ${ROOT_MOD}/app/email -I ../../idl --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"





