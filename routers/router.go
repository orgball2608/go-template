package routers

import (
	"elearning/config"
	"elearning/middlewares"
	v1Routers "elearning/routers/v1"
	"elearning/usecase/auth"
	"elearning/validations"
	"fmt"
	"github.com/gin-contrib/cors"
	"net/http"

	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

const (
	xForwardedProtoHeader = "x-forwarded-proto"
)

func InitRouter(
	config *config.Environment,
	middleware middlewares.Middleware,
	authUseCase auth.AuthUseCase,
) *gin.Engine {
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins: strings.Split(config.CorsAllowOrigins, ","),
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders: []string{
			"Origin",
			"Content-Length",
			"Content-Type",
			"Access-Control-Allow-Headers",
			"Authorization",
			"X-XSRF-TOKEN",
			"screenId",
			"apiOrder",
		},
		ExposeHeaders: []string{
			"Content-Disposition",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.Use(gin.Recovery())
	// Validations
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("customEmail", validations.CustomEmail)
		v.RegisterValidation("customPassword", validations.CustomPassword)
	}
	apiRouter := router.Group("/api")
	apiRouter.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1Routers.InitV1Router(
		apiRouter.Group("/v1", middleware.RestLogger),
		middleware,
		authUseCase,
	)
	router.NoRoute(func(c *gin.Context) {
		reverseProxy(c, config)
	})
	return router
}

func reverseProxy(c *gin.Context, config *config.Environment) {
	if c.GetHeader(xForwardedProtoHeader) != "https" {
		sslUrl := "https://" + c.Request.Host + c.Request.RequestURI
		c.Redirect(http.StatusFound, sslUrl)
		return
	}
	remote, _ := url.Parse(fmt.Sprintf("http://localhost:%d", config.FrontendPort))
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL = c.Request.URL
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}
