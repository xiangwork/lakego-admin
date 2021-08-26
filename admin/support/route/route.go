package route

import (
    "github.com/gin-gonic/gin"
    "lakego-admin/lakego/config"
    "lakego-admin/lakego/http/route"
)

// 路由
func AddRoute(engine *gin.Engine, f func(rg *gin.RouterGroup)) {
    // 配置
    conf := config.New("admin")

    // 后台路由及设置中间件
    m := route.GetMiddlewares(conf.GetString("Route.Middleware"))

    // 路由
    admin := engine.Group(conf.GetString("Route.Group"))
    {
        admin.Use(m...)
        {
            f(admin)
        }
    }
}
