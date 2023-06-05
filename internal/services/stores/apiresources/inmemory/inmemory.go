package inmemory

import (
	contracts_stores_apiresources "echo-starter/internal/contracts/stores/apiresources"
	models "echo-starter/internal/models"

	di "github.com/dozm/di"
	core_hashset "github.com/fluffy-bunny/fluffycore/gods/sets/hashset"
)

type (
	service struct {
		apiResources    []models.APIResource
		apiResourcesMap map[string]models.APIResource
		scopeMap        map[string]string
		scopes          *core_hashset.StringSet
	}
)

var stemService *service = new(service)

func init() {
	var _ contracts_stores_apiresources.IAPIResources = (*service)(nil)
}

func (s *service) Ctor(apiResources []models.APIResource) (*service, error) {
	obj := &service{
		apiResources:    []models.APIResource{},
		apiResourcesMap: make(map[string]models.APIResource),
		scopeMap:        make(map[string]string),
		scopes:          core_hashset.NewStringSet(),
	}

	for _, item := range apiResources {
		obj.apiResourcesMap[item.Name] = item
		for _, scope := range item.Scopes {
			obj.scopeMap[scope] = item.Name
			obj.scopes.Add(scope)
		}
	}
	return obj, nil
}

// AddSingletonIAPIResources registers the *service as a singleton.
func AddSingletonIAPIResources(builder di.ContainerBuilder, apiResources []models.APIResource) {
	obj, err := stemService.Ctor(apiResources)
	if err != nil {
		panic(err)
	}

	di.AddInstance[contracts_stores_apiresources.IAPIResources](builder,
		obj)

}
func (s *service) GetApiResourceByScope(scope string) (*models.APIResource, bool, error) {
	name, found := s.scopeMap[scope]
	if !found {
		return nil, false, nil
	}
	item, found := s.apiResourcesMap[name]
	if !found {
		return nil, false, nil
	}
	copyOfItem := item
	return &copyOfItem, true, nil
}

func (s *service) GetAPIResources() ([]models.APIResource, error) {
	return s.apiResources, nil
}
func (s *service) GetAPIResource(name string) (*models.APIResource, bool, error) {
	item, found := s.apiResourcesMap[name]
	if !found {
		return nil, false, nil
	}
	copyOfItem := item
	return &copyOfItem, true, nil
}
func (s *service) GetApiResourceScopes() (*core_hashset.StringSet, error) {
	return s.scopes, nil
}
