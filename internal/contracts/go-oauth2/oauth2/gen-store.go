// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package oauth2

import (
	"reflect"
	"strings"

	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

// ReflectTypeITokenStore used when your service claims to implement ITokenStore
var ReflectTypeITokenStore = di.GetInterfaceReflectType((*ITokenStore)(nil))

// AddSingletonITokenStore adds a type that implements ITokenStore
func AddSingletonITokenStore(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITokenStore)
	_logAddITokenStore("SINGLETON", implType, _getImplementedITokenStoreNames(implementedTypes...),
		_logITokenStoreExtra{
			Name:  "DI-BY",
			Value: "type",
		})
	di.AddSingleton(builder, implType, implementedTypes...)
}

// AddSingletonITokenStoreWithMetadata adds a type that implements ITokenStore
func AddSingletonITokenStoreWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITokenStore)
	_logAddITokenStore("SINGLETON", implType, _getImplementedITokenStoreNames(implementedTypes...),
		_logITokenStoreExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logITokenStoreExtra{
			Name:  "DI-M",
			Value: metaData,
		})
	di.AddSingletonWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddSingletonITokenStoreByObj adds a prebuilt obj
func AddSingletonITokenStoreByObj(builder *di.Builder, obj interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITokenStore)
	_logAddITokenStore("SINGLETON", reflect.TypeOf(obj), _getImplementedITokenStoreNames(implementedTypes...),
		_logITokenStoreExtra{
			Name:  "DI-BY",
			Value: "obj",
		})
	di.AddSingletonWithImplementedTypesByObj(builder, obj, implementedTypes...)
}

// AddSingletonITokenStoreByObjWithMetadata adds a prebuilt obj
func AddSingletonITokenStoreByObjWithMetadata(builder *di.Builder, obj interface{}, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITokenStore)
	_logAddITokenStore("SINGLETON", reflect.TypeOf(obj), _getImplementedITokenStoreNames(implementedTypes...),
		_logITokenStoreExtra{
			Name:  "DI-BY",
			Value: "obj",
		},
		_logITokenStoreExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddSingletonWithImplementedTypesByObjWithMetadata(builder, obj, metaData, implementedTypes...)
}

// AddSingletonITokenStoreByFunc adds a type by a custom func
func AddSingletonITokenStoreByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITokenStore)
	_logAddITokenStore("SINGLETON", implType, _getImplementedITokenStoreNames(implementedTypes...),
		_logITokenStoreExtra{
			Name:  "DI-BY",
			Value: "func",
		})
	di.AddSingletonWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddSingletonITokenStoreByFuncWithMetadata adds a type by a custom func
func AddSingletonITokenStoreByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITokenStore)
	_logAddITokenStore("SINGLETON", implType, _getImplementedITokenStoreNames(implementedTypes...),
		_logITokenStoreExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logITokenStoreExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddSingletonWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddTransientITokenStore adds a type that implements ITokenStore
func AddTransientITokenStore(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITokenStore)
	_logAddITokenStore("TRANSIENT", implType, _getImplementedITokenStoreNames(implementedTypes...),
		_logITokenStoreExtra{
			Name:  "DI-BY",
			Value: "type",
		})

	di.AddTransientWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddTransientITokenStoreWithMetadata adds a type that implements ITokenStore
func AddTransientITokenStoreWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITokenStore)
	_logAddITokenStore("TRANSIENT", implType, _getImplementedITokenStoreNames(implementedTypes...),
		_logITokenStoreExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logITokenStoreExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddTransientWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddTransientITokenStoreByFunc adds a type by a custom func
func AddTransientITokenStoreByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITokenStore)
	_logAddITokenStore("TRANSIENT", implType, _getImplementedITokenStoreNames(implementedTypes...),
		_logITokenStoreExtra{
			Name:  "DI-BY",
			Value: "func",
		})

	di.AddTransientWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddTransientITokenStoreByFuncWithMetadata adds a type by a custom func
func AddTransientITokenStoreByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITokenStore)
	_logAddITokenStore("TRANSIENT", implType, _getImplementedITokenStoreNames(implementedTypes...),
		_logITokenStoreExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logITokenStoreExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddTransientWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddScopedITokenStore adds a type that implements ITokenStore
func AddScopedITokenStore(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITokenStore)
	_logAddITokenStore("SCOPED", implType, _getImplementedITokenStoreNames(implementedTypes...),
		_logITokenStoreExtra{
			Name:  "DI-BY",
			Value: "type",
		})
	di.AddScopedWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddScopedITokenStoreWithMetadata adds a type that implements ITokenStore
func AddScopedITokenStoreWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITokenStore)
	_logAddITokenStore("SCOPED", implType, _getImplementedITokenStoreNames(implementedTypes...),
		_logITokenStoreExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logITokenStoreExtra{
			Name:  "DI-M",
			Value: metaData,
		})
	di.AddScopedWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddScopedITokenStoreByFunc adds a type by a custom func
func AddScopedITokenStoreByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITokenStore)
	_logAddITokenStore("SCOPED", implType, _getImplementedITokenStoreNames(implementedTypes...),
		_logITokenStoreExtra{
			Name:  "DI-BY",
			Value: "func",
		})
	di.AddScopedWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddScopedITokenStoreByFuncWithMetadata adds a type by a custom func
func AddScopedITokenStoreByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITokenStore)
	_logAddITokenStore("SCOPED", implType, _getImplementedITokenStoreNames(implementedTypes...),
		_logITokenStoreExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logITokenStoreExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddScopedWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// RemoveAllITokenStore removes all ITokenStore from the DI
func RemoveAllITokenStore(builder *di.Builder) {
	builder.RemoveAllByType(ReflectTypeITokenStore)
}

// GetITokenStoreFromContainer alternative to SafeGetITokenStoreFromContainer but panics of object is not present
func GetITokenStoreFromContainer(ctn di.Container) ITokenStore {
	return ctn.GetByType(ReflectTypeITokenStore).(ITokenStore)
}

// GetManyITokenStoreFromContainer alternative to SafeGetManyITokenStoreFromContainer but panics of object is not present
func GetManyITokenStoreFromContainer(ctn di.Container) []ITokenStore {
	objs := ctn.GetManyByType(ReflectTypeITokenStore)
	var results []ITokenStore
	for _, obj := range objs {
		results = append(results, obj.(ITokenStore))
	}
	return results
}

// SafeGetITokenStoreFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetITokenStoreFromContainer(ctn di.Container) (ITokenStore, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeITokenStore)
	if err != nil {
		return nil, err
	}
	return obj.(ITokenStore), nil
}

// GetITokenStoreDefinition returns that last definition registered that this container can provide
func GetITokenStoreDefinition(ctn di.Container) *di.Def {
	def := ctn.GetDefinitionByType(ReflectTypeITokenStore)
	return def
}

// GetITokenStoreDefinitions returns all definitions that this container can provide
func GetITokenStoreDefinitions(ctn di.Container) []*di.Def {
	defs := ctn.GetDefinitionsByType(ReflectTypeITokenStore)
	return defs
}

// SafeGetManyITokenStoreFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetManyITokenStoreFromContainer(ctn di.Container) ([]ITokenStore, error) {
	objs, err := ctn.SafeGetManyByType(ReflectTypeITokenStore)
	if err != nil {
		return nil, err
	}
	var results []ITokenStore
	for _, obj := range objs {
		results = append(results, obj.(ITokenStore))
	}
	return results, nil
}

type _logITokenStoreExtra struct {
	Name  string
	Value interface{}
}

func _logAddITokenStore(scopeType string, implType reflect.Type, interfaces string, extra ..._logITokenStoreExtra) {
	infoEvent := log.Info().
		Str("DI", scopeType).
		Str("DI-I", interfaces).
		Str("DI-B", implType.Elem().String())

	for _, extra := range extra {
		infoEvent = infoEvent.Interface(extra.Name, extra.Value)
	}

	infoEvent.Send()

}
func _getImplementedITokenStoreNames(implementedTypes ...reflect.Type) string {
	builder := strings.Builder{}
	for idx, implementedType := range implementedTypes {
		builder.WriteString(implementedType.Name())
		if idx < len(implementedTypes)-1 {
			builder.WriteString(", ")
		}
	}
	return builder.String()
}

// ReflectTypeIClientStore used when your service claims to implement IClientStore
var ReflectTypeIClientStore = di.GetInterfaceReflectType((*IClientStore)(nil))

// AddSingletonIClientStore adds a type that implements IClientStore
func AddSingletonIClientStore(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIClientStore)
	_logAddIClientStore("SINGLETON", implType, _getImplementedIClientStoreNames(implementedTypes...),
		_logIClientStoreExtra{
			Name:  "DI-BY",
			Value: "type",
		})
	di.AddSingleton(builder, implType, implementedTypes...)
}

// AddSingletonIClientStoreWithMetadata adds a type that implements IClientStore
func AddSingletonIClientStoreWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIClientStore)
	_logAddIClientStore("SINGLETON", implType, _getImplementedIClientStoreNames(implementedTypes...),
		_logIClientStoreExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logIClientStoreExtra{
			Name:  "DI-M",
			Value: metaData,
		})
	di.AddSingletonWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddSingletonIClientStoreByObj adds a prebuilt obj
func AddSingletonIClientStoreByObj(builder *di.Builder, obj interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIClientStore)
	_logAddIClientStore("SINGLETON", reflect.TypeOf(obj), _getImplementedIClientStoreNames(implementedTypes...),
		_logIClientStoreExtra{
			Name:  "DI-BY",
			Value: "obj",
		})
	di.AddSingletonWithImplementedTypesByObj(builder, obj, implementedTypes...)
}

// AddSingletonIClientStoreByObjWithMetadata adds a prebuilt obj
func AddSingletonIClientStoreByObjWithMetadata(builder *di.Builder, obj interface{}, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIClientStore)
	_logAddIClientStore("SINGLETON", reflect.TypeOf(obj), _getImplementedIClientStoreNames(implementedTypes...),
		_logIClientStoreExtra{
			Name:  "DI-BY",
			Value: "obj",
		},
		_logIClientStoreExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddSingletonWithImplementedTypesByObjWithMetadata(builder, obj, metaData, implementedTypes...)
}

// AddSingletonIClientStoreByFunc adds a type by a custom func
func AddSingletonIClientStoreByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIClientStore)
	_logAddIClientStore("SINGLETON", implType, _getImplementedIClientStoreNames(implementedTypes...),
		_logIClientStoreExtra{
			Name:  "DI-BY",
			Value: "func",
		})
	di.AddSingletonWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddSingletonIClientStoreByFuncWithMetadata adds a type by a custom func
func AddSingletonIClientStoreByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIClientStore)
	_logAddIClientStore("SINGLETON", implType, _getImplementedIClientStoreNames(implementedTypes...),
		_logIClientStoreExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logIClientStoreExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddSingletonWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddTransientIClientStore adds a type that implements IClientStore
func AddTransientIClientStore(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIClientStore)
	_logAddIClientStore("TRANSIENT", implType, _getImplementedIClientStoreNames(implementedTypes...),
		_logIClientStoreExtra{
			Name:  "DI-BY",
			Value: "type",
		})

	di.AddTransientWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddTransientIClientStoreWithMetadata adds a type that implements IClientStore
func AddTransientIClientStoreWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIClientStore)
	_logAddIClientStore("TRANSIENT", implType, _getImplementedIClientStoreNames(implementedTypes...),
		_logIClientStoreExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logIClientStoreExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddTransientWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddTransientIClientStoreByFunc adds a type by a custom func
func AddTransientIClientStoreByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIClientStore)
	_logAddIClientStore("TRANSIENT", implType, _getImplementedIClientStoreNames(implementedTypes...),
		_logIClientStoreExtra{
			Name:  "DI-BY",
			Value: "func",
		})

	di.AddTransientWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddTransientIClientStoreByFuncWithMetadata adds a type by a custom func
func AddTransientIClientStoreByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIClientStore)
	_logAddIClientStore("TRANSIENT", implType, _getImplementedIClientStoreNames(implementedTypes...),
		_logIClientStoreExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logIClientStoreExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddTransientWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddScopedIClientStore adds a type that implements IClientStore
func AddScopedIClientStore(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIClientStore)
	_logAddIClientStore("SCOPED", implType, _getImplementedIClientStoreNames(implementedTypes...),
		_logIClientStoreExtra{
			Name:  "DI-BY",
			Value: "type",
		})
	di.AddScopedWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddScopedIClientStoreWithMetadata adds a type that implements IClientStore
func AddScopedIClientStoreWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIClientStore)
	_logAddIClientStore("SCOPED", implType, _getImplementedIClientStoreNames(implementedTypes...),
		_logIClientStoreExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logIClientStoreExtra{
			Name:  "DI-M",
			Value: metaData,
		})
	di.AddScopedWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddScopedIClientStoreByFunc adds a type by a custom func
func AddScopedIClientStoreByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIClientStore)
	_logAddIClientStore("SCOPED", implType, _getImplementedIClientStoreNames(implementedTypes...),
		_logIClientStoreExtra{
			Name:  "DI-BY",
			Value: "func",
		})
	di.AddScopedWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddScopedIClientStoreByFuncWithMetadata adds a type by a custom func
func AddScopedIClientStoreByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIClientStore)
	_logAddIClientStore("SCOPED", implType, _getImplementedIClientStoreNames(implementedTypes...),
		_logIClientStoreExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logIClientStoreExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddScopedWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// RemoveAllIClientStore removes all IClientStore from the DI
func RemoveAllIClientStore(builder *di.Builder) {
	builder.RemoveAllByType(ReflectTypeIClientStore)
}

// GetIClientStoreFromContainer alternative to SafeGetIClientStoreFromContainer but panics of object is not present
func GetIClientStoreFromContainer(ctn di.Container) IClientStore {
	return ctn.GetByType(ReflectTypeIClientStore).(IClientStore)
}

// GetManyIClientStoreFromContainer alternative to SafeGetManyIClientStoreFromContainer but panics of object is not present
func GetManyIClientStoreFromContainer(ctn di.Container) []IClientStore {
	objs := ctn.GetManyByType(ReflectTypeIClientStore)
	var results []IClientStore
	for _, obj := range objs {
		results = append(results, obj.(IClientStore))
	}
	return results
}

// SafeGetIClientStoreFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetIClientStoreFromContainer(ctn di.Container) (IClientStore, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeIClientStore)
	if err != nil {
		return nil, err
	}
	return obj.(IClientStore), nil
}

// GetIClientStoreDefinition returns that last definition registered that this container can provide
func GetIClientStoreDefinition(ctn di.Container) *di.Def {
	def := ctn.GetDefinitionByType(ReflectTypeIClientStore)
	return def
}

// GetIClientStoreDefinitions returns all definitions that this container can provide
func GetIClientStoreDefinitions(ctn di.Container) []*di.Def {
	defs := ctn.GetDefinitionsByType(ReflectTypeIClientStore)
	return defs
}

// SafeGetManyIClientStoreFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetManyIClientStoreFromContainer(ctn di.Container) ([]IClientStore, error) {
	objs, err := ctn.SafeGetManyByType(ReflectTypeIClientStore)
	if err != nil {
		return nil, err
	}
	var results []IClientStore
	for _, obj := range objs {
		results = append(results, obj.(IClientStore))
	}
	return results, nil
}

type _logIClientStoreExtra struct {
	Name  string
	Value interface{}
}

func _logAddIClientStore(scopeType string, implType reflect.Type, interfaces string, extra ..._logIClientStoreExtra) {
	infoEvent := log.Info().
		Str("DI", scopeType).
		Str("DI-I", interfaces).
		Str("DI-B", implType.Elem().String())

	for _, extra := range extra {
		infoEvent = infoEvent.Interface(extra.Name, extra.Value)
	}

	infoEvent.Send()

}
func _getImplementedIClientStoreNames(implementedTypes ...reflect.Type) string {
	builder := strings.Builder{}
	for idx, implementedType := range implementedTypes {
		builder.WriteString(implementedType.Name())
		if idx < len(implementedTypes)-1 {
			builder.WriteString(", ")
		}
	}
	return builder.String()
}
