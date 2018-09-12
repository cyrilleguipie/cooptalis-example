package app

import (
	"cooptalis-example/app/controllers"
	"cooptalis-example/app/models"
	"cooptalis-example/app/utils"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/revel/revel"
	"log"
	"net/http"
)

var (
	// AppVersion revel app version (ldflags)
	AppVersion string

	// BuildTime revel app build-time (ldflags)
	BuildTime string
)

func CheckAuthentification(c *revel.Controller) revel.Result {
	log.Println("Authenticate!")
	log.Println(c.Params)
	var response utils.JsonResponse
	token, success, message := getToken(c)
	response.Message = message
	response.Success = success
	if !success {
		log.Println(message)
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(response)
	} else {
		user, _ := models.FindByToken(token)
		if user == nil {
			log.Println(message)
			c.Response.Status = http.StatusUnauthorized
			return c.RenderJSON(response)
		} else {
			log.Println("auth token success")
			c.Session["user_role"] = user.Role
			return nil
		}

	}
}

func getToken(c *revel.Controller) (token string, success bool, message string) {
	authHeader := c.Request.Header.Get("Authorization")
	if len(authHeader) < 1 {
		//log.Println()
		return "", false, "Empty token"
	} else {

		return authHeader, true, "Token exists"
	}

	//tokenSlice := strings.Split(authHeader, " ")
	//if len(tokenSlice) != 2 {
	//	return "", errInvalidTokenFormat
	//}
	//tokenString = tokenSlice[1]
	//return tokenString, nil

}

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
		//AuthzFilter,
	}

	// Register startup functions with OnAppStart
	// revel.DevMode and revel.RunMode only work inside of OnAppStart. See Example Startup Script
	// ( order dependent )
	// revel.OnAppStart(ExampleStartupScript)
	revel.InterceptFunc(CheckAuthentification, revel.BEFORE, &controllers.Menus{})
	revel.FilterController(controllers.Menus{}).Insert(AuthzFilter, revel.BEFORE, revel.ActionInvoker)
	revel.OnAppStart(models.InitDb)
	revel.OnAppStart(models.InitMenu)
	revel.OnAppStart(models.InitUser)
	revel.OnAppStart(models.InitClient)
	revel.OnAppStart(models.InitCollaborater)
	// revel.OnAppStart(FillCache)
}

// HeaderFilter adds common security headers
// There is a full implementation of a CSRF filter in
// https://github.com/revel/modules/tree/master/csrf
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")
	c.Response.Out.Header().Add("Referrer-Policy", "strict-origin-when-cross-origin")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

//func ExampleStartupScript() {
//	// revel.DevMod and revel.RunMode work here
//	// Use this script to check for dev mode and set dev/prod startup scripts here!
//	if revel.DevMode == true {
//		// Dev mode
//	}
//}
