package interfaces

func (s *server) Routes() {
	s.Router.GET("/health", s.ph.Health())
}
