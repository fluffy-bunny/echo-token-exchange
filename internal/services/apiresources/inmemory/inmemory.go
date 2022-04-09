package inmemory

import (
	"reflect"

	contracts_apiresources "echo-starter/internal/contracts/apiresources"
	"echo-starter/internal/models"

	contracts_logger "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/logger"
	di "github.com/fluffy-bunny/sarulabsdi"
)

type (
	service struct {
		Logger          contracts_logger.ILogger `inject:""`
		apiResources    []models.APIResource
		apiResourcesMap map[string]models.APIResource
	}
)

func assertImplementation() {
	var _ contracts_apiresources.IAPIResources = (*service)(nil)
}

var reflectType = reflect.TypeOf((*service)(nil))

// AddSingletonIAPIResources registers the *service as a singleton.
func AddSingletonIAPIResources(builder *di.Builder, apiResources []models.APIResource) {
	contracts_apiresources.AddSingletonIAPIResourcesByFunc(builder, reflectType, func(ctn di.Container) (interface{}, error) {

		cMap := make(map[string]models.APIResource)
		for _, item := range apiResources {
			cMap[item.Name] = item
		}

		obj := &service{
			apiResources:    apiResources,
			apiResourcesMap: cMap,
		}
		return obj, nil
	})
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
