package startup

import "echo-starter/internal/models"

func (s *Startup) loadApiResources() (err error) {
	s.apiResources = append(s.apiResources, models.APIResource{
		Name:   "invoices",
		Scopes: []string{"invoices", "invoices.read", "invoices.write"},
	})
	s.apiResources = append(s.apiResources, models.APIResource{
		Name:   "users",
		Scopes: []string{"users", "users.read", "users.write"},
	})
	return
}
