package core

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

// HandlerFunc is an alias for gin.HandlerFunc for abstraction
type HandlerFunc = gin.HandlerFunc

// RouterConfig holds configuration for the router
type RouterConfig struct {
	// TrustedProxies is a list of IP addresses or CIDR ranges that are trusted
	TrustedProxies []string

	// TrustedPlatform specifies the trusted platform header (e.g., "X-Real-Ip")
	TrustedPlatform string
}

// DefaultConfig returns secure default configuration
func DefaultConfig() *RouterConfig {
	return &RouterConfig{
		TrustedProxies:  nil, // Trust no proxies by default (most secure)
		TrustedPlatform: "",
	}
}

type IRouterGroup interface {
	Group(path string, handlers ...HandlerFunc) *RouterGroup
	IMethod
}

var _ IMethod = (*RouterGroup)(nil)

type IMethod interface {
	GET(string, ...HandlerFunc)
	POST(string, ...HandlerFunc)
	PUT(string, ...HandlerFunc)
	DELETE(string, ...HandlerFunc)
}

// routerGroup implements RouterGroup interface
type RouterGroup struct {
	group *gin.RouterGroup
}

// Implementation of RouterGroup interface methods
func (rg *RouterGroup) Group(path string, handlers ...HandlerFunc) *RouterGroup {
	return &RouterGroup{group: rg.group.Group(path, handlers...)}
}

func (rg *RouterGroup) GET(path string, handlers ...HandlerFunc) {
	rg.group.GET(path, handlers...)
}

func (rg *RouterGroup) POST(path string, handlers ...HandlerFunc) {
	rg.group.POST(path, handlers...)
}

func (rg *RouterGroup) PUT(path string, handlers ...HandlerFunc) {
	rg.group.PUT(path, handlers...)
}

func (rg *RouterGroup) DELETE(path string, handlers ...HandlerFunc) {
	rg.group.DELETE(path, handlers...)
}

// Router is the main router structure that wraps gin.Engine
type Router struct {
	engine *gin.Engine
	*RouterGroup
}

type IRouter interface {
	Use(...HandlerFunc)
	Run(string) error
}

var _ IRouter = (*Router)(nil)

// Router methods
func NewRouter() *Router {
	return NewWithConfig(DefaultConfig())
}

func NewWithConfig(config *RouterConfig) *Router {
	engine := gin.New()
	gin.SetMode(os.Getenv("GIN_MODE"))

	if err := engine.SetTrustedProxies(config.TrustedProxies); err != nil {
		// TODO: Consider using a proper logger
		fmt.Printf("Error setting trusted proxies: %v\n", err)
	}

	if config.TrustedPlatform != "" {
		engine.TrustedPlatform = config.TrustedPlatform
	}

	return &Router{
		engine:      engine,
		RouterGroup: &RouterGroup{group: engine.Group("")},
	}
}

// Use adds middleware to the router
func (r *Router) Use(middleware ...HandlerFunc) {
	r.group.Use(middleware...)
}

// Run starts the HTTP server
func (r *Router) Run(addr string) error {
	return r.engine.Run(addr)
}
