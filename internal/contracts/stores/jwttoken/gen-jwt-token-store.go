// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package jwttoken

import (
	"reflect"
	"strings"

	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

// ReflectTypeIJwtTokenStore used when your service claims to implement IJwtTokenStore
var ReflectTypeIJwtTokenStore = di.GetInterfaceReflectType((*IJwtTokenStore)(nil))

// AddSingletonIJwtTokenStore adds a type that implements IJwtTokenStore
func AddSingletonIJwtTokenStore(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIJwtTokenStore)
	_logAddIJwtTokenStore("SINGLETON", implType, _getImplementedIJwtTokenStoreNames(implementedTypes...),
		_logIJwtTokenStoreExtra{
			Name:  "DI-BY",
			Value: "type",
		})
	di.AddSingleton(builder, implType, implementedTypes...)
}

// AddSingletonIJwtTokenStoreWithMetadata adds a type that implements IJwtTokenStore
func AddSingletonIJwtTokenStoreWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIJwtTokenStore)
	_logAddIJwtTokenStore("SINGLETON", implType, _getImplementedIJwtTokenStoreNames(implementedTypes...),
		_logIJwtTokenStoreExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logIJwtTokenStoreExtra{
			Name:  "DI-M",
			Value: metaData,
		})
	di.AddSingletonWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddSingletonIJwtTokenStoreByObj adds a prebuilt obj
func AddSingletonIJwtTokenStoreByObj(builder *di.Builder, obj interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIJwtTokenStore)
	_logAddIJwtTokenStore("SINGLETON", reflect.TypeOf(obj), _getImplementedIJwtTokenStoreNames(implementedTypes...),
		_logIJwtTokenStoreExtra{
			Name:  "DI-BY",
			Value: "obj",
		})
	di.AddSingletonWithImplementedTypesByObj(builder, obj, implementedTypes...)
}

// AddSingletonIJwtTokenStoreByObjWithMetadata adds a prebuilt obj
func AddSingletonIJwtTokenStoreByObjWithMetadata(builder *di.Builder, obj interface{}, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIJwtTokenStore)
	_logAddIJwtTokenStore("SINGLETON", reflect.TypeOf(obj), _getImplementedIJwtTokenStoreNames(implementedTypes...),
		_logIJwtTokenStoreExtra{
			Name:  "DI-BY",
			Value: "obj",
		},
		_logIJwtTokenStoreExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddSingletonWithImplementedTypesByObjWithMetadata(builder, obj, metaData, implementedTypes...)
}

// AddSingletonIJwtTokenStoreByFunc adds a type by a custom func
func AddSingletonIJwtTokenStoreByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIJwtTokenStore)
	_logAddIJwtTokenStore("SINGLETON", implType, _getImplementedIJwtTokenStoreNames(implementedTypes...),
		_logIJwtTokenStoreExtra{
			Name:  "DI-BY",
			Value: "func",
		})
	di.AddSingletonWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddSingletonIJwtTokenStoreByFuncWithMetadata adds a type by a custom func
func AddSingletonIJwtTokenStoreByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIJwtTokenStore)
	_logAddIJwtTokenStore("SINGLETON", implType, _getImplementedIJwtTokenStoreNames(implementedTypes...),
		_logIJwtTokenStoreExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logIJwtTokenStoreExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddSingletonWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddTransientIJwtTokenStore adds a type that implements IJwtTokenStore
func AddTransientIJwtTokenStore(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIJwtTokenStore)
	_logAddIJwtTokenStore("TRANSIENT", implType, _getImplementedIJwtTokenStoreNames(implementedTypes...),
		_logIJwtTokenStoreExtra{
			Name:  "DI-BY",
			Value: "type",
		})

	di.AddTransientWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddTransientIJwtTokenStoreWithMetadata adds a type that implements IJwtTokenStore
func AddTransientIJwtTokenStoreWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIJwtTokenStore)
	_logAddIJwtTokenStore("TRANSIENT", implType, _getImplementedIJwtTokenStoreNames(implementedTypes...),
		_logIJwtTokenStoreExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logIJwtTokenStoreExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddTransientWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddTransientIJwtTokenStoreByFunc adds a type by a custom func
func AddTransientIJwtTokenStoreByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIJwtTokenStore)
	_logAddIJwtTokenStore("TRANSIENT", implType, _getImplementedIJwtTokenStoreNames(implementedTypes...),
		_logIJwtTokenStoreExtra{
			Name:  "DI-BY",
			Value: "func",
		})

	di.AddTransientWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddTransientIJwtTokenStoreByFuncWithMetadata adds a type by a custom func
func AddTransientIJwtTokenStoreByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIJwtTokenStore)
	_logAddIJwtTokenStore("TRANSIENT", implType, _getImplementedIJwtTokenStoreNames(implementedTypes...),
		_logIJwtTokenStoreExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logIJwtTokenStoreExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddTransientWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddScopedIJwtTokenStore adds a type that implements IJwtTokenStore
func AddScopedIJwtTokenStore(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIJwtTokenStore)
	_logAddIJwtTokenStore("SCOPED", implType, _getImplementedIJwtTokenStoreNames(implementedTypes...),
		_logIJwtTokenStoreExtra{
			Name:  "DI-BY",
			Value: "type",
		})
	di.AddScopedWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddScopedIJwtTokenStoreWithMetadata adds a type that implements IJwtTokenStore
func AddScopedIJwtTokenStoreWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIJwtTokenStore)
	_logAddIJwtTokenStore("SCOPED", implType, _getImplementedIJwtTokenStoreNames(implementedTypes...),
		_logIJwtTokenStoreExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logIJwtTokenStoreExtra{
			Name:  "DI-M",
			Value: metaData,
		})
	di.AddScopedWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddScopedIJwtTokenStoreByFunc adds a type by a custom func
func AddScopedIJwtTokenStoreByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIJwtTokenStore)
	_logAddIJwtTokenStore("SCOPED", implType, _getImplementedIJwtTokenStoreNames(implementedTypes...),
		_logIJwtTokenStoreExtra{
			Name:  "DI-BY",
			Value: "func",
		})
	di.AddScopedWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddScopedIJwtTokenStoreByFuncWithMetadata adds a type by a custom func
func AddScopedIJwtTokenStoreByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIJwtTokenStore)
	_logAddIJwtTokenStore("SCOPED", implType, _getImplementedIJwtTokenStoreNames(implementedTypes...),
		_logIJwtTokenStoreExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logIJwtTokenStoreExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddScopedWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// RemoveAllIJwtTokenStore removes all IJwtTokenStore from the DI
func RemoveAllIJwtTokenStore(builder *di.Builder) {
	builder.RemoveAllByType(ReflectTypeIJwtTokenStore)
}

// GetIJwtTokenStoreFromContainer alternative to SafeGetIJwtTokenStoreFromContainer but panics of object is not present
func GetIJwtTokenStoreFromContainer(ctn di.Container) IJwtTokenStore {
	return ctn.GetByType(ReflectTypeIJwtTokenStore).(IJwtTokenStore)
}

// GetManyIJwtTokenStoreFromContainer alternative to SafeGetManyIJwtTokenStoreFromContainer but panics of object is not present
func GetManyIJwtTokenStoreFromContainer(ctn di.Container) []IJwtTokenStore {
	objs := ctn.GetManyByType(ReflectTypeIJwtTokenStore)
	var results []IJwtTokenStore
	for _, obj := range objs {
		results = append(results, obj.(IJwtTokenStore))
	}
	return results
}

// SafeGetIJwtTokenStoreFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetIJwtTokenStoreFromContainer(ctn di.Container) (IJwtTokenStore, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeIJwtTokenStore)
	if err != nil {
		return nil, err
	}
	return obj.(IJwtTokenStore), nil
}

// GetIJwtTokenStoreDefinition returns that last definition registered that this container can provide
func GetIJwtTokenStoreDefinition(ctn di.Container) *di.Def {
	def := ctn.GetDefinitionByType(ReflectTypeIJwtTokenStore)
	return def
}

// GetIJwtTokenStoreDefinitions returns all definitions that this container can provide
func GetIJwtTokenStoreDefinitions(ctn di.Container) []*di.Def {
	defs := ctn.GetDefinitionsByType(ReflectTypeIJwtTokenStore)
	return defs
}

// SafeGetManyIJwtTokenStoreFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetManyIJwtTokenStoreFromContainer(ctn di.Container) ([]IJwtTokenStore, error) {
	objs, err := ctn.SafeGetManyByType(ReflectTypeIJwtTokenStore)
	if err != nil {
		return nil, err
	}
	var results []IJwtTokenStore
	for _, obj := range objs {
		results = append(results, obj.(IJwtTokenStore))
	}
	return results, nil
}

type _logIJwtTokenStoreExtra struct {
	Name  string
	Value interface{}
}

func _logAddIJwtTokenStore(scopeType string, implType reflect.Type, interfaces string, extra ..._logIJwtTokenStoreExtra) {
	infoEvent := log.Info().
		Str("DI", scopeType).
		Str("DI-I", interfaces).
		Str("DI-B", implType.Elem().String())

	for _, extra := range extra {
		infoEvent = infoEvent.Interface(extra.Name, extra.Value)
	}

	infoEvent.Send()

}
func _getImplementedIJwtTokenStoreNames(implementedTypes ...reflect.Type) string {
	builder := strings.Builder{}
	for idx, implementedType := range implementedTypes {
		builder.WriteString(implementedType.Name())
		if idx < len(implementedTypes)-1 {
			builder.WriteString(", ")
		}
	}
	return builder.String()
}
