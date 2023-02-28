package server

// APIServer は API サーバーを表す。
type APIServer struct {
	*HealthCheckController
	*OrganizationsController
}

// Controller はコントローラーのインターフェース。
type Controller interface {
	apply(s *APIServer)
}

// Controllers はコントローラーへ実体を apply する。
func (s *APIServer) Controllers(cs ...Controller) {
	for _, c := range cs {
		c.apply(s)
	}
}
