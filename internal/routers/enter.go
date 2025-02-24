package routers

import (
	"go-ecommerce-be/internal/routers/manage"
	"go-ecommerce-be/internal/routers/user"
)

type RouterGroup struct {
	User user.UserRouterGroup
	Admin manage.AdminRouterGroup
}

var RouterGroupApp = new(RouterGroup)
