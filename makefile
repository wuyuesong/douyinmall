.PHONY:	gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/home.proto  --service frontend -module github.com/wuyuesong/gomall/app/frontend -I ../../idl

.PHONY:	gen-auth-page
gen-auth-page:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/auth_page.proto  --service frontend -module github.com/wuyuesong/gomall/app/frontend -I ../../idl

.PHONY:	gen-rpc-client
gen-rpc-client:
	@cd rpc_gen && cwgo client --type RPC --idl ../idl/user.proto --service user -module github.com/wuyuesong/gomall/rpc_gen -I ../idl 

.PHONY:	gen-rpc-server
gen-rpc-server:
	@cd app/user && cwgo server --type RPC --idl ../../idl/user.proto --service user -module github.com/wuyuesong/gomall/app/user -I ../../idl --pass "-use github.com/wuyuesong/gomall/rpc_gen/kitex_gen"