package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"gvb_server/api"
	"gvb_server/middleware"
)

var store = cookie.NewStore([]byte("HyvCD89g3VDJ9646BFGEh37GFJ"))

func (router RouterGroup) UserRouter() {
	app := api.ApiGroupApp.UserApi
	router.Use(sessions.Sessions("sessionid", store))
	router.POST("email_login", app.EmailLoginView)
	router.GET("user_list", middleware.JwtAuth(), app.UserListView)
	router.PUT("user_role", middleware.JwtAuth(), app.UserUpdateRoleView)
	router.PUT("user_password", middleware.JwtAuth(), app.UserUpdatePassword)
	router.POST("logout", middleware.JwtAuth(), app.LogoutView)
	router.DELETE("user_remove", middleware.JwtAuth(), app.UserRemoveView)
	router.POST("user_bind_email", middleware.JwtAuth(), app.UserBindEmailView)
	router.POST("qq_login_view", app.QQLoginView)
	router.POST("user_create", middleware.JwtAuth(), app.UserCreateView)
}
