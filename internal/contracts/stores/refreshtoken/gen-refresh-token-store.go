// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package refreshtoken

import (
	"reflect"
	"strings"

	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

// ReflectTypeIRefreshTokenStore used when your service claims to implement IRefreshTokenStore
var ReflectTypeIRefreshTokenStore = di.GetInterfaceReflectType((*IRefreshTokenStore)(nil))

// AddSingletonIRefreshTokenStore adds a type that implements IRefreshTokenStore
func AddSingletonIRefreshTokenStore(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIRefreshTokenStore)
	_logAddIRefreshTokenStore("SINGLETON", implType, _getImplementedIRefreshTokenStoreNames(implementedTypes...),
		_logIRefreshTokenStoreExtra{
			Name:  "DI-BY",
			Value: "type",
		})
	di.AddSingleton(builder, implType, implementedTypes...)
}

// AddSingletonIRefreshTokenStoreWithMetadata adds a type that implements IRefreshTokenStore
func AddSingletonIRefreshTokenStoreWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIRefreshTokenStore)
	_logAddIRefreshTokenStore("SINGLETON", implType, _getImplementedIRefreshTokenStoreNames(implementedTypes...),
		_logIRefreshTokenStoreExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logIRefreshTokenStoreExtra{
			Name:  "DI-M",
			Value: metaData,
		})
	di.AddSingletonWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddSingletonIRefreshTokenStoreByObj adds a prebuilt obj
func AddSingletonIRefreshTokenStoreByObj(builder *di.Builder, obj interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIRefreshTokenStore)
	_logAddIRefreshTokenStore("SINGLETON", reflect.TypeOf(obj), _getImplementedIRefreshTokenStoreNames(implementedTypes...),
		_logIRefreshTokenStoreExtra{
			Name:  "DI-BY",
			Value: "obj",
		})
	di.AddSingletonWithImplementedTypesByObj(builder, obj, implementedTypes...)
}

// AddSingletonIRefreshTokenStoreByObjWithMetadata adds a prebuilt obj
func AddSingletonIRefreshTokenStoreByObjWithMetadata(builder *di.Builder, obj interface{}, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIRefreshTokenStore)
	_logAddIRefreshTokenStore("SINGLETON", reflect.TypeOf(obj), _getImplementedIRefreshTokenStoreNames(implementedTypes...),
		_logIRefreshTokenStoreExtra{
			Name:  "DI-BY",
			Value: "obj",
		},
		_logIRefreshTokenStoreExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddSingletonWithImplementedTypesByObjWithMetadata(builder, obj, metaData, implementedTypes...)
}

// AddSingletonIRefreshTokenStoreByFunc adds a type by a custom func
func AddSingletonIRefreshTokenStoreByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIRefreshTokenStore)
	_logAddIRefreshTokenStore("SINGLETON", implType, _getImplementedIRefreshTokenStoreNames(implementedTypes...),
		_logIRefreshTokenStoreExtra{
			Name:  "DI-BY",
			Value: "func",
		})
	di.AddSingletonWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddSingletonIRefreshTokenStoreByFuncWithMetadata adds a type by a custom func
func AddSingletonIRefreshTokenStoreByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIRefreshTokenStore)
	_logAddIRefreshTokenStore("SINGLETON", implType, _getImplementedIRefreshTokenStoreNames(implementedTypes...),
		_logIRefreshTokenStoreExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logIRefreshTokenStoreExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddSingletonWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddTransientIRefreshTokenStore adds a type that implements IRefreshTokenStore
func AddTransientIRefreshTokenStore(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIRefreshTokenStore)
	_logAddIRefreshTokenStore("TRANSIENT", implType, _getImplementedIRefreshTokenStoreNames(implementedTypes...),
		_logIRefreshTokenStoreExtra{
			Name:  "DI-BY",
			Value: "type",
		})

	di.AddTransientWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddTransientIRefreshTokenStoreWithMetadata adds a type that implements IRefreshTokenStore
func AddTransientIRefreshTokenStoreWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIRefreshTokenStore)
	_logAddIRefreshTokenStore("TRANSIENT", implType, _getImplementedIRefreshTokenStoreNames(implementedTypes...),
		_logIRefreshTokenStoreExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logIRefreshTokenStoreExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddTransientWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddTransientIRefreshTokenStoreByFunc adds a type by a custom func
func AddTransientIRefreshTokenStoreByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIRefreshTokenStore)
	_logAddIRefreshTokenStore("TRANSIENT", implType, _getImplementedIRefreshTokenStoreNames(implementedTypes...),
		_logIRefreshTokenStoreExtra{
			Name:  "DI-BY",
			Value: "func",
		})

	di.AddTransientWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddTransientIRefreshTokenStoreByFuncWithMetadata adds a type by a custom func
func AddTransientIRefreshTokenStoreByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIRefreshTokenStore)
	_logAddIRefreshTokenStore("TRANSIENT", implType, _getImplementedIRefreshTokenStoreNames(implementedTypes...),
		_logIRefreshTokenStoreExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logIRefreshTokenStoreExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddTransientWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddScopedIRefreshTokenStore adds a type that implements IRefreshTokenStore
func AddScopedIRefreshTokenStore(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIRefreshTokenStore)
	_logAddIRefreshTokenStore("SCOPED", implType, _getImplementedIRefreshTokenStoreNames(implementedTypes...),
		_logIRefreshTokenStoreExtra{
			Name:  "DI-BY",
			Value: "type",
		})
	di.AddScopedWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddScopedIRefreshTokenStoreWithMetadata adds a type that implements IRefreshTokenStore
func AddScopedIRefreshTokenStoreWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIRefreshTokenStore)
	_logAddIRefreshTokenStore("SCOPED", implType, _getImplementedIRefreshTokenStoreNames(implementedTypes...),
		_logIRefreshTokenStoreExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logIRefreshTokenStoreExtra{
			Name:  "DI-M",
			Value: metaData,
		})
	di.AddScopedWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddScopedIRefreshTokenStoreByFunc adds a type by a custom func
func AddScopedIRefreshTokenStoreByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIRefreshTokenStore)
	_logAddIRefreshTokenStore("SCOPED", implType, _getImplementedIRefreshTokenStoreNames(implementedTypes...),
		_logIRefreshTokenStoreExtra{
			Name:  "DI-BY",
			Value: "func",
		})
	di.AddScopedWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddScopedIRefreshTokenStoreByFuncWithMetadata adds a type by a custom func
func AddScopedIRefreshTokenStoreByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIRefreshTokenStore)
	_logAddIRefreshTokenStore("SCOPED", implType, _getImplementedIRefreshTokenStoreNames(implementedTypes...),
		_logIRefreshTokenStoreExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logIRefreshTokenStoreExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddScopedWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// RemoveAllIRefreshTokenStore removes all IRefreshTokenStore from the DI
func RemoveAllIRefreshTokenStore(builder *di.Builder) {
	builder.RemoveAllByType(ReflectTypeIRefreshTokenStore)
}

// GetIRefreshTokenStoreFromContainer alternative to SafeGetIRefreshTokenStoreFromContainer but panics of object is not present
func GetIRefreshTokenStoreFromContainer(ctn di.Container) IRefreshTokenStore {
	return ctn.GetByType(ReflectTypeIRefreshTokenStore).(IRefreshTokenStore)
}

// GetManyIRefreshTokenStoreFromContainer alternative to SafeGetManyIRefreshTokenStoreFromContainer but panics of object is not present
func GetManyIRefreshTokenStoreFromContainer(ctn di.Container) []IRefreshTokenStore {
	objs := ctn.GetManyByType(ReflectTypeIRefreshTokenStore)
	var results []IRefreshTokenStore
	for _, obj := range objs {
		results = append(results, obj.(IRefreshTokenStore))
	}
	return results
}

// SafeGetIRefreshTokenStoreFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetIRefreshTokenStoreFromContainer(ctn di.Container) (IRefreshTokenStore, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeIRefreshTokenStore)
	if err != nil {
		return nil, err
	}
	return obj.(IRefreshTokenStore), nil
}

// GetIRefreshTokenStoreDefinition returns that last definition registered that this container can provide
func GetIRefreshTokenStoreDefinition(ctn di.Container) *di.Def {
	def := ctn.GetDefinitionByType(ReflectTypeIRefreshTokenStore)
	return def
}

// GetIRefreshTokenStoreDefinitions returns all definitions that this container can provide
func GetIRefreshTokenStoreDefinitions(ctn di.Container) []*di.Def {
	defs := ctn.GetDefinitionsByType(ReflectTypeIRefreshTokenStore)
	return defs
}

// SafeGetManyIRefreshTokenStoreFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetManyIRefreshTokenStoreFromContainer(ctn di.Container) ([]IRefreshTokenStore, error) {
	objs, err := ctn.SafeGetManyByType(ReflectTypeIRefreshTokenStore)
	if err != nil {
		return nil, err
	}
	var results []IRefreshTokenStore
	for _, obj := range objs {
		results = append(results, obj.(IRefreshTokenStore))
	}
	return results, nil
}

type _logIRefreshTokenStoreExtra struct {
	Name  string
	Value interface{}
}

func _logAddIRefreshTokenStore(scopeType string, implType reflect.Type, interfaces string, extra ..._logIRefreshTokenStoreExtra) {
	infoEvent := log.Info().
		Str("DI", scopeType).
		Str("DI-I", interfaces).
		Str("DI-B", implType.Elem().String())

	for _, extra := range extra {
		infoEvent = infoEvent.Interface(extra.Name, extra.Value)
	}

	infoEvent.Send()

}
func _getImplementedIRefreshTokenStoreNames(implementedTypes ...reflect.Type) string {
	builder := strings.Builder{}
	for idx, implementedType := range implementedTypes {
		builder.WriteString(implementedType.Name())
		if idx < len(implementedTypes)-1 {
			builder.WriteString(", ")
		}
	}
	return builder.String()
}
