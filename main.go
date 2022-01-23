package main

import (
	"akabane_nyanko/backend/src/controllers"
	"time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		// アクセスを許可したいHTTPメソッド
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))

	// TODO: ルーティング周りの設定は全て別ファイルに切り出す
	apiEngine := engine.Group("/api")
	{
		adminEngine := apiEngine.Group("/admin")
		{
			catEngine := adminEngine.Group("/cat")
			{
				catEngine.GET("/list", controllers.CatList)
			}

			areaEngine := adminEngine.Group("/area")
			{
				areaEngine.GET("/list", controllers.AreaList)
			}
		}
	}
	engine.Run()
}
