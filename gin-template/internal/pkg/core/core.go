package core

// Server is the main server structure
type Server struct {
	Router *Router
	Redis  *Redis
}

// Server methods
func NewHttpServer() *Server {
	return &Server{
		Router: NewRouter(),
		Redis:  NewCacheServer(),
	}
}

func (s *Server) Run(addr string) error {
	return s.Router.Run(addr)
}
