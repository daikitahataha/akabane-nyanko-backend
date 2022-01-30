package server

import (
	"akabane_nyanko/backend/src/config"
	"akabane_nyanko/backend/src/controllers/admin"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	engine := gin.Default()

	cors := config.GetCorsConfig()
	engine.Use(cors)

	apiEngine := engine.Group("/api")
	{
		adminEngine := apiEngine.Group("/admin")
		{
			catEngine := adminEngine.Group("/cat")
			{
				catEngine.GET("/list", admin.CatList)
			}

			areaEngine := adminEngine.Group("/area")
			{
				areaEngine.GET("/list", admin.AreaList)
			}

			tagEngine := adminEngine.Group("/tag")
			{
				var controller admin.CatTagController
				tagEngine.GET("/list", controller.TagList)
			}
		}
	}

	return engine
}
