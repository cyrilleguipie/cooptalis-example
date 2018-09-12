package app

import (
	"cooptalis-example/app/utils"
	"fmt"
	"github.com/casbin/casbin"
	"github.com/revel/revel"
	"net/http"
)

var AuthzFilter = func(c *revel.Controller, fc []revel.Filter) {
	fmt.Println("AuthzFilter")
	e := casbin.NewEnforcer("conf/authz_model.conf", "conf/authz_policy.csv")
	fmt.Println("AuthzFilter Session USER ROLE " + c.Session["user_role"])
	if !CheckPermission(c.Session["user_role"], e, c.Request) {
		var response utils.JsonResponse
		c.Response.Status = http.StatusUnauthorized
		//c.Result = c.Forbidden("Access denied by the Authz plugin.")
		response.Success = false
		response.Message = "Access denied to user with role :" + c.Session["user_role"] + " for resource :" + c.Request.URL.Path
		c.RenderJSON(response)
		return
	} else {
		fc[0](c, fc[1:])
	}
}

func CheckPermission(role string, e *casbin.Enforcer, r *revel.Request) bool {
	method := r.Method
	path := r.URL.Path
	return e.Enforce(role, path, method)
}
