// Code generated by hertz generator.

package checkout

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/wuyuesong/gomall/app/frontend/middleware"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{middleware.Auth()}
}

func _checkoutMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _checkout0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _checkoutresultMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _checkoutwaitingMw() []app.HandlerFunc {
	// your code...
	return nil
}
