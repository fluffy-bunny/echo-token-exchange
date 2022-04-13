package inmemory

import (
	"reflect"

	contracts_stores_apiresources "echo-starter/internal/contracts/stores/apiresources"
	"echo-starter/internal/models"

	core_hashset "github.com/fluffy-bunny/grpcdotnetgo/pkg/gods/sets/hashset"

	contracts_logger "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/logger"
	di "github.com/fluffy-bunny/sarulabsdi"
)

type (
	service struct {
		Logger          contracts_logger.ILogger `inject:""`
		apiResources    []models.APIResource
		apiResourcesMap map[string]models.APIResource
		scopeMap        map[string]string
		scopes          *core_hashset.StringSet
	}
)

func assertImplementation() {
	var _ contracts_stores_apiresources.IAPIResources = (*service)(nil)
}

var reflectType = reflect.TypeOf((*service)(nil))

func (s *service) Ctor() {
	s.scopes = core_hashset.NewStringSet()
	s.scopeMap = make(map[string]string)
	s.apiResourcesMap = make(map[string]models.APIResource)
}

// AddSingletonIAPIResources registers the *service as a singleton.
func AddSingletonIAPIResources(builder *di.Builder, apiResources []models.APIResource) {
	contracts_stores_apiresources.AddSingletonIAPIResourcesByFunc(builder, reflectType, func(ctn di.Container) (interface{}, error) {
		obj := &service{}
		obj.Ctor()

		scopeMap := make(map[string]string)
		apiResourcesMap := make(map[string]models.APIResource)
		for _, item := range apiResources {
			apiResourcesMap[item.Name] = item
			for _, scope := range item.Scopes {
				scopeMap[scope] = item.Name
				obj.scopes.Add(scope)
			}
		}
		obj.apiResources = apiResources
		obj.apiResourcesMap = apiResourcesMap
		obj.scopeMap = scopeMap

		return obj, nil
	})
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
	var copyOfItem models.APIResource
	copyOfItem = item
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
	var copyOfItem models.APIResource
	copyOfItem = item
	return &copyOfItem, true, nil
}
func (s *service) GetApiResourceScopes() (*core_hashset.StringSet, error) {
	return s.scopes, nil
}
