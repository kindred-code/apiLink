package routing

import (
	"github.com/gin-gonic/gin"
	end "mpolitakis.LinkApi/Endpoints"
)

func Routing(router *gin.Engine) {

	router.GET("/details/:profileId", func(ctx *gin.Context) {
		end.GetDetails(ctx)
	})
	router.POST("/details/:profileId", func(ctx *gin.Context) {
		end.PostDetails(ctx)
	})
	router.GET("/profile/:profileId", func(ctx *gin.Context) {
		end.GetProfileById(ctx)
	})
	router.POST("/profile", func(ctx *gin.Context) {
		end.PostProfile(ctx)
	})
	router.GET("/profile", func(ctx *gin.Context) {
		end.GetAllProfiles(ctx)
	})
	router.POST("/photoPost/", func(ctx *gin.Context) {
		end.PostPhoto(ctx)
	})
	router.GET("/photo/:profileId", func(ctx *gin.Context) {
		end.GetPhoto(ctx)
	})

}
