package startup

import "echo-starter/internal/models"

func (s *Startup) loadApiResources() (err error) {
	s.apiResources = append(s.apiResources, models.APIResource{
		Name:   "invoices",
		Scopes: []string{"read", "write"},
	})
	s.apiResources = append(s.apiResources, models.APIResource{
		Name:   "users",
		Scopes: []string{"read", "write"},
	})
	return
}
