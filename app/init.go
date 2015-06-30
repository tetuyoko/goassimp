package app

import (
	"github.com/revel/revel"
	"time"

	"goassimp/lib/mgndb"
	"goassimp/lib/mgnredis"
)

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
	}

	// register startup functions with OnAppStart
	revel.OnAppStart(func() {
		mugendb.InitDB(user, password, host, dbname)
	}) // DBやテーブルの作成

	revel.OnAppStart(func() {
		host := revel.Config.StringDefault("redis.host", ":6379")
		capacity := revel.Config.IntDefault("redis.capacity_pool", 20)
		max_capacity := revel.Config.IntDefault("redis.max_capacity_pool", 200)
		dstr := revel.Config.StringDefault("redis.idleTimeout", "1m")
		duration, err := time.ParseDuration(dstr)
		if err != nil {
			panic(err)
		}
		mgnredis.InitRedis(host, capacity, max_capacity, duration)
	})
}

// TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	// Add some common security headers
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}
