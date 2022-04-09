// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package apiresources

import (
	"reflect"
	"strings"

	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

// ReflectTypeIAPIResources used when your service claims to implement IAPIResources
var ReflectTypeIAPIResources = di.GetInterfaceReflectType((*IAPIResources)(nil))

// AddSingletonIAPIResources adds a type that implements IAPIResources
func AddSingletonIAPIResources(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIAPIResources)
	_logAddIAPIResources("SINGLETON", implType, _getImplementedIAPIResourcesNames(implementedTypes...),
		_logIAPIResourcesExtra{
			Name:  "DI-BY",
			Value: "type",
		})
	di.AddSingleton(builder, implType, implementedTypes...)
}

// AddSingletonIAPIResourcesWithMetadata adds a type that implements IAPIResources
func AddSingletonIAPIResourcesWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIAPIResources)
	_logAddIAPIResources("SINGLETON", implType, _getImplementedIAPIResourcesNames(implementedTypes...),
		_logIAPIResourcesExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logIAPIResourcesExtra{
			Name:  "DI-M",
			Value: metaData,
		})
	di.AddSingletonWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddSingletonIAPIResourcesByObj adds a prebuilt obj
func AddSingletonIAPIResourcesByObj(builder *di.Builder, obj interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIAPIResources)
	_logAddIAPIResources("SINGLETON", reflect.TypeOf(obj), _getImplementedIAPIResourcesNames(implementedTypes...),
		_logIAPIResourcesExtra{
			Name:  "DI-BY",
			Value: "obj",
		})
	di.AddSingletonWithImplementedTypesByObj(builder, obj, implementedTypes...)
}

// AddSingletonIAPIResourcesByObjWithMetadata adds a prebuilt obj
func AddSingletonIAPIResourcesByObjWithMetadata(builder *di.Builder, obj interface{}, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIAPIResources)
	_logAddIAPIResources("SINGLETON", reflect.TypeOf(obj), _getImplementedIAPIResourcesNames(implementedTypes...),
		_logIAPIResourcesExtra{
			Name:  "DI-BY",
			Value: "obj",
		},
		_logIAPIResourcesExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddSingletonWithImplementedTypesByObjWithMetadata(builder, obj, metaData, implementedTypes...)
}

// AddSingletonIAPIResourcesByFunc adds a type by a custom func
func AddSingletonIAPIResourcesByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIAPIResources)
	_logAddIAPIResources("SINGLETON", implType, _getImplementedIAPIResourcesNames(implementedTypes...),
		_logIAPIResourcesExtra{
			Name:  "DI-BY",
			Value: "func",
		})
	di.AddSingletonWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddSingletonIAPIResourcesByFuncWithMetadata adds a type by a custom func
func AddSingletonIAPIResourcesByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIAPIResources)
	_logAddIAPIResources("SINGLETON", implType, _getImplementedIAPIResourcesNames(implementedTypes...),
		_logIAPIResourcesExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logIAPIResourcesExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddSingletonWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddTransientIAPIResources adds a type that implements IAPIResources
func AddTransientIAPIResources(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIAPIResources)
	_logAddIAPIResources("TRANSIENT", implType, _getImplementedIAPIResourcesNames(implementedTypes...),
		_logIAPIResourcesExtra{
			Name:  "DI-BY",
			Value: "type",
		})

	di.AddTransientWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddTransientIAPIResourcesWithMetadata adds a type that implements IAPIResources
func AddTransientIAPIResourcesWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIAPIResources)
	_logAddIAPIResources("TRANSIENT", implType, _getImplementedIAPIResourcesNames(implementedTypes...),
		_logIAPIResourcesExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logIAPIResourcesExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddTransientWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddTransientIAPIResourcesByFunc adds a type by a custom func
func AddTransientIAPIResourcesByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIAPIResources)
	_logAddIAPIResources("TRANSIENT", implType, _getImplementedIAPIResourcesNames(implementedTypes...),
		_logIAPIResourcesExtra{
			Name:  "DI-BY",
			Value: "func",
		})

	di.AddTransientWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddTransientIAPIResourcesByFuncWithMetadata adds a type by a custom func
func AddTransientIAPIResourcesByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIAPIResources)
	_logAddIAPIResources("TRANSIENT", implType, _getImplementedIAPIResourcesNames(implementedTypes...),
		_logIAPIResourcesExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logIAPIResourcesExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddTransientWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddScopedIAPIResources adds a type that implements IAPIResources
func AddScopedIAPIResources(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIAPIResources)
	_logAddIAPIResources("SCOPED", implType, _getImplementedIAPIResourcesNames(implementedTypes...),
		_logIAPIResourcesExtra{
			Name:  "DI-BY",
			Value: "type",
		})
	di.AddScopedWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddScopedIAPIResourcesWithMetadata adds a type that implements IAPIResources
func AddScopedIAPIResourcesWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIAPIResources)
	_logAddIAPIResources("SCOPED", implType, _getImplementedIAPIResourcesNames(implementedTypes...),
		_logIAPIResourcesExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logIAPIResourcesExtra{
			Name:  "DI-M",
			Value: metaData,
		})
	di.AddScopedWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddScopedIAPIResourcesByFunc adds a type by a custom func
func AddScopedIAPIResourcesByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIAPIResources)
	_logAddIAPIResources("SCOPED", implType, _getImplementedIAPIResourcesNames(implementedTypes...),
		_logIAPIResourcesExtra{
			Name:  "DI-BY",
			Value: "func",
		})
	di.AddScopedWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddScopedIAPIResourcesByFuncWithMetadata adds a type by a custom func
func AddScopedIAPIResourcesByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIAPIResources)
	_logAddIAPIResources("SCOPED", implType, _getImplementedIAPIResourcesNames(implementedTypes...),
		_logIAPIResourcesExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logIAPIResourcesExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddScopedWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// RemoveAllIAPIResources removes all IAPIResources from the DI
func RemoveAllIAPIResources(builder *di.Builder) {
	builder.RemoveAllByType(ReflectTypeIAPIResources)
}

// GetIAPIResourcesFromContainer alternative to SafeGetIAPIResourcesFromContainer but panics of object is not present
func GetIAPIResourcesFromContainer(ctn di.Container) IAPIResources {
	return ctn.GetByType(ReflectTypeIAPIResources).(IAPIResources)
}

// GetManyIAPIResourcesFromContainer alternative to SafeGetManyIAPIResourcesFromContainer but panics of object is not present
func GetManyIAPIResourcesFromContainer(ctn di.Container) []IAPIResources {
	objs := ctn.GetManyByType(ReflectTypeIAPIResources)
	var results []IAPIResources
	for _, obj := range objs {
		results = append(results, obj.(IAPIResources))
	}
	return results
}

// SafeGetIAPIResourcesFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetIAPIResourcesFromContainer(ctn di.Container) (IAPIResources, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeIAPIResources)
	if err != nil {
		return nil, err
	}
	return obj.(IAPIResources), nil
}

// GetIAPIResourcesDefinition returns that last definition registered that this container can provide
func GetIAPIResourcesDefinition(ctn di.Container) *di.Def {
	def := ctn.GetDefinitionByType(ReflectTypeIAPIResources)
	return def
}

// GetIAPIResourcesDefinitions returns all definitions that this container can provide
func GetIAPIResourcesDefinitions(ctn di.Container) []*di.Def {
	defs := ctn.GetDefinitionsByType(ReflectTypeIAPIResources)
	return defs
}

// SafeGetManyIAPIResourcesFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetManyIAPIResourcesFromContainer(ctn di.Container) ([]IAPIResources, error) {
	objs, err := ctn.SafeGetManyByType(ReflectTypeIAPIResources)
	if err != nil {
		return nil, err
	}
	var results []IAPIResources
	for _, obj := range objs {
		results = append(results, obj.(IAPIResources))
	}
	return results, nil
}

type _logIAPIResourcesExtra struct {
	Name  string
	Value interface{}
}

func _logAddIAPIResources(scopeType string, implType reflect.Type, interfaces string, extra ..._logIAPIResourcesExtra) {
	infoEvent := log.Info().
		Str("DI", scopeType).
		Str("DI-I", interfaces).
		Str("DI-B", implType.Elem().String())

	for _, extra := range extra {
		infoEvent = infoEvent.Interface(extra.Name, extra.Value)
	}

	infoEvent.Send()

}
func _getImplementedIAPIResourcesNames(implementedTypes ...reflect.Type) string {
	builder := strings.Builder{}
	for idx, implementedType := range implementedTypes {
		builder.WriteString(implementedType.Name())
		if idx < len(implementedTypes)-1 {
			builder.WriteString(", ")
		}
	}
	return builder.String()
}