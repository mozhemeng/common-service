package routers

import (
	"common_service/global"
	"common_service/internal/middleware"
	v1 "common_service/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"

	_ "common_service/docs"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.AccessLogger())
	//r.Use(middleware.HandleErrors())   // TODO:still thinking about this middleware
	r.Use(middleware.Translations())
	url := ginSwagger.URL("./swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.StaticFS(global.AppSetting.UploadServerUrl, http.Dir(global.AppSetting.UploadSavePath))

	auth := v1.NewAuth()
	user := v1.NewUser()
	role := v1.NewRole()
	casbinPolicy := v1.NewCasbinPolicy()
	upload := v1.NewUpload()

	apiV1 := r.Group("/api/v1")
	authApi := apiV1.Group("")
	authApi.Use(middleware.Authorization())
	authApi.Use(middleware.CasbinHandler())
	{

		apiV1.POST("/sign_in", auth.SignIn)
		apiV1.POST("/refresh_token", auth.RefreshAccessToken)

		authApi.GET("/users/:id", user.GetByID)
		authApi.GET("/users", user.List)
		authApi.POST("/users", user.Create)
		authApi.PUT("/users/:id", user.Update)
		authApi.DELETE("/users/:id", user.Delete)

		authApi.GET("/roles", role.List)
		authApi.POST("/roles", role.Create)
		authApi.PUT("/roles/:id", role.Update)
		authApi.DELETE("/roles/:id", role.Delete)

		authApi.GET("/casbin/policies", casbinPolicy.List)
		authApi.POST("/casbin/policies", casbinPolicy.Create)
		authApi.DELETE("/casbin/policies", casbinPolicy.Delete)
		authApi.GET("/casbin/policies/reload", casbinPolicy.ReLoad)

		authApi.POST("/upload", upload.UploadFile)
	}

	return r
}
