package routers

import (
	"common_service/global"
	"common_service/internal/middleware"
	v1 "common_service/internal/routers/api/v1"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"

	_ "common_service/docs"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.Recovery())
	r.Use(middleware.AccessLogger())
	//r.Use(middleware.HandleErrors())   // TODO:still thinking about this middleware
	r.Use(middleware.Translations())
	url := ginSwagger.URL("./swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.StaticFS(global.AppSetting.UploadServerUrl, http.Dir(global.AppSetting.UploadSavePath))

	pprof.Register(r)

	auth := v1.NewAuth()
	user := v1.NewUser()
	role := v1.NewRole()
	permPolicy := v1.NewPermPolicy()
	upload := v1.NewUpload()
	example := v1.NewExample()

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

		authApi.GET("/perm/policies", permPolicy.List)
		authApi.POST("/perm/policies", permPolicy.Create)
		authApi.DELETE("/perm/policies", permPolicy.Delete)
		authApi.GET("/perm/policies/reload", permPolicy.ReLoad)

		authApi.POST("/upload", upload.UploadFile)

		// example for UserRateLimiter middleware
		authApi.Use(middleware.UserRateLimiter("2-M")).GET("/example/rate-limit", example.UserRateLimit)

	}

	return r
}
