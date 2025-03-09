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

type RouterGroup interface {
	Group(path string, handlers ...HandlerFunc) RouterGroup
}

// routerGroup implements RouterGroup interface
type routerGroup struct {
	group *gin.RouterGroup
}

// Router is the main router structure that wraps gin.Engine
type Router struct {
	engine *gin.Engine
	*routerGroup
}

// Implementation of RouterGroup interface methods
func (rg *routerGroup) Group(path string, handlers ...HandlerFunc) RouterGroup {
	return &routerGroup{group: rg.group.Group(path, handlers...)}
}

func (rg *routerGroup) GET(path string, handlers ...HandlerFunc) {
	rg.group.GET(path, handlers...)
}

func (rg *routerGroup) POST(path string, handlers ...HandlerFunc) {
	rg.group.POST(path, handlers...)
}

func (rg *routerGroup) PUT(path string, handlers ...HandlerFunc) {
	rg.group.PUT(path, handlers...)
}

func (rg *routerGroup) DELETE(path string, handlers ...HandlerFunc) {
	rg.group.DELETE(path, handlers...)
}

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
		routerGroup: &routerGroup{group: engine.Group("")},
	}
}

// Group creates a new route group while maintaining the same engine
func (r *Router) Group(path string, handlers ...HandlerFunc) *Router {
	return &Router{
		engine:      r.engine,
		routerGroup: &routerGroup{group: r.group.Group(path, handlers...)},
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
