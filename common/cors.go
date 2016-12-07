package common

import "github.com/lunny/tango"

func MyHandler() tango.HandlerFunc {
	return func(ctx *tango.Context) {
		ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", " Origin, X-Requested-With, Content-Type, Accept")
		ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if "OPTIONS" == ctx.Req().Method {
			ctx.Abort(200, "SUCCESS")
		} else {
			ctx.Next()

		}
	}
}
