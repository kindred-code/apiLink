package routing

import (
	"github.com/gin-gonic/gin"
	end "mpolitakis.LinkApi/Endpoints"
)

func Routing(router gin.Context) {

	router.GET("/details/{profileId}", end.GetDetails)
	router.POST("/details/{profileId}", end.PostDetails)
	router.GET("/profile/{profileId}", end.GetProfileById)
	router.GET("/profile", end.GetAllProfiles)
	router.POST("/photo/", end.PostPhoto)
	router.GET("/photo/{profileId}", end.GetPhoto)

}
