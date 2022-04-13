// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package referencetoken

import (
	"reflect"
	"strings"

	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

// ReflectTypeIReferenceTokenStore used when your service claims to implement IReferenceTokenStore
var ReflectTypeIReferenceTokenStore = di.GetInterfaceReflectType((*IReferenceTokenStore)(nil))

// AddSingletonIReferenceTokenStore adds a type that implements IReferenceTokenStore
func AddSingletonIReferenceTokenStore(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIReferenceTokenStore)
	_logAddIReferenceTokenStore("SINGLETON", implType, _getImplementedIReferenceTokenStoreNames(implementedTypes...),
		_logIReferenceTokenStoreExtra{
			Name:  "DI-BY",
			Value: "type",
		})
	di.AddSingleton(builder, implType, implementedTypes...)
}

// AddSingletonIReferenceTokenStoreWithMetadata adds a type that implements IReferenceTokenStore
func AddSingletonIReferenceTokenStoreWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIReferenceTokenStore)
	_logAddIReferenceTokenStore("SINGLETON", implType, _getImplementedIReferenceTokenStoreNames(implementedTypes...),
		_logIReferenceTokenStoreExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logIReferenceTokenStoreExtra{
			Name:  "DI-M",
			Value: metaData,
		})
	di.AddSingletonWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddSingletonIReferenceTokenStoreByObj adds a prebuilt obj
func AddSingletonIReferenceTokenStoreByObj(builder *di.Builder, obj interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIReferenceTokenStore)
	_logAddIReferenceTokenStore("SINGLETON", reflect.TypeOf(obj), _getImplementedIReferenceTokenStoreNames(implementedTypes...),
		_logIReferenceTokenStoreExtra{
			Name:  "DI-BY",
			Value: "obj",
		})
	di.AddSingletonWithImplementedTypesByObj(builder, obj, implementedTypes...)
}

// AddSingletonIReferenceTokenStoreByObjWithMetadata adds a prebuilt obj
func AddSingletonIReferenceTokenStoreByObjWithMetadata(builder *di.Builder, obj interface{}, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIReferenceTokenStore)
	_logAddIReferenceTokenStore("SINGLETON", reflect.TypeOf(obj), _getImplementedIReferenceTokenStoreNames(implementedTypes...),
		_logIReferenceTokenStoreExtra{
			Name:  "DI-BY",
			Value: "obj",
		},
		_logIReferenceTokenStoreExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddSingletonWithImplementedTypesByObjWithMetadata(builder, obj, metaData, implementedTypes...)
}

// AddSingletonIReferenceTokenStoreByFunc adds a type by a custom func
func AddSingletonIReferenceTokenStoreByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIReferenceTokenStore)
	_logAddIReferenceTokenStore("SINGLETON", implType, _getImplementedIReferenceTokenStoreNames(implementedTypes...),
		_logIReferenceTokenStoreExtra{
			Name:  "DI-BY",
			Value: "func",
		})
	di.AddSingletonWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddSingletonIReferenceTokenStoreByFuncWithMetadata adds a type by a custom func
func AddSingletonIReferenceTokenStoreByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIReferenceTokenStore)
	_logAddIReferenceTokenStore("SINGLETON", implType, _getImplementedIReferenceTokenStoreNames(implementedTypes...),
		_logIReferenceTokenStoreExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logIReferenceTokenStoreExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddSingletonWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddTransientIReferenceTokenStore adds a type that implements IReferenceTokenStore
func AddTransientIReferenceTokenStore(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIReferenceTokenStore)
	_logAddIReferenceTokenStore("TRANSIENT", implType, _getImplementedIReferenceTokenStoreNames(implementedTypes...),
		_logIReferenceTokenStoreExtra{
			Name:  "DI-BY",
			Value: "type",
		})

	di.AddTransientWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddTransientIReferenceTokenStoreWithMetadata adds a type that implements IReferenceTokenStore
func AddTransientIReferenceTokenStoreWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIReferenceTokenStore)
	_logAddIReferenceTokenStore("TRANSIENT", implType, _getImplementedIReferenceTokenStoreNames(implementedTypes...),
		_logIReferenceTokenStoreExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logIReferenceTokenStoreExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddTransientWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddTransientIReferenceTokenStoreByFunc adds a type by a custom func
func AddTransientIReferenceTokenStoreByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIReferenceTokenStore)
	_logAddIReferenceTokenStore("TRANSIENT", implType, _getImplementedIReferenceTokenStoreNames(implementedTypes...),
		_logIReferenceTokenStoreExtra{
			Name:  "DI-BY",
			Value: "func",
		})

	di.AddTransientWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddTransientIReferenceTokenStoreByFuncWithMetadata adds a type by a custom func
func AddTransientIReferenceTokenStoreByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIReferenceTokenStore)
	_logAddIReferenceTokenStore("TRANSIENT", implType, _getImplementedIReferenceTokenStoreNames(implementedTypes...),
		_logIReferenceTokenStoreExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logIReferenceTokenStoreExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddTransientWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddScopedIReferenceTokenStore adds a type that implements IReferenceTokenStore
func AddScopedIReferenceTokenStore(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIReferenceTokenStore)
	_logAddIReferenceTokenStore("SCOPED", implType, _getImplementedIReferenceTokenStoreNames(implementedTypes...),
		_logIReferenceTokenStoreExtra{
			Name:  "DI-BY",
			Value: "type",
		})
	di.AddScopedWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddScopedIReferenceTokenStoreWithMetadata adds a type that implements IReferenceTokenStore
func AddScopedIReferenceTokenStoreWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIReferenceTokenStore)
	_logAddIReferenceTokenStore("SCOPED", implType, _getImplementedIReferenceTokenStoreNames(implementedTypes...),
		_logIReferenceTokenStoreExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logIReferenceTokenStoreExtra{
			Name:  "DI-M",
			Value: metaData,
		})
	di.AddScopedWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddScopedIReferenceTokenStoreByFunc adds a type by a custom func
func AddScopedIReferenceTokenStoreByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIReferenceTokenStore)
	_logAddIReferenceTokenStore("SCOPED", implType, _getImplementedIReferenceTokenStoreNames(implementedTypes...),
		_logIReferenceTokenStoreExtra{
			Name:  "DI-BY",
			Value: "func",
		})
	di.AddScopedWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddScopedIReferenceTokenStoreByFuncWithMetadata adds a type by a custom func
func AddScopedIReferenceTokenStoreByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIReferenceTokenStore)
	_logAddIReferenceTokenStore("SCOPED", implType, _getImplementedIReferenceTokenStoreNames(implementedTypes...),
		_logIReferenceTokenStoreExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logIReferenceTokenStoreExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddScopedWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// RemoveAllIReferenceTokenStore removes all IReferenceTokenStore from the DI
func RemoveAllIReferenceTokenStore(builder *di.Builder) {
	builder.RemoveAllByType(ReflectTypeIReferenceTokenStore)
}

// GetIReferenceTokenStoreFromContainer alternative to SafeGetIReferenceTokenStoreFromContainer but panics of object is not present
func GetIReferenceTokenStoreFromContainer(ctn di.Container) IReferenceTokenStore {
	return ctn.GetByType(ReflectTypeIReferenceTokenStore).(IReferenceTokenStore)
}

// GetManyIReferenceTokenStoreFromContainer alternative to SafeGetManyIReferenceTokenStoreFromContainer but panics of object is not present
func GetManyIReferenceTokenStoreFromContainer(ctn di.Container) []IReferenceTokenStore {
	objs := ctn.GetManyByType(ReflectTypeIReferenceTokenStore)
	var results []IReferenceTokenStore
	for _, obj := range objs {
		results = append(results, obj.(IReferenceTokenStore))
	}
	return results
}

// SafeGetIReferenceTokenStoreFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetIReferenceTokenStoreFromContainer(ctn di.Container) (IReferenceTokenStore, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeIReferenceTokenStore)
	if err != nil {
		return nil, err
	}
	return obj.(IReferenceTokenStore), nil
}

// GetIReferenceTokenStoreDefinition returns that last definition registered that this container can provide
func GetIReferenceTokenStoreDefinition(ctn di.Container) *di.Def {
	def := ctn.GetDefinitionByType(ReflectTypeIReferenceTokenStore)
	return def
}

// GetIReferenceTokenStoreDefinitions returns all definitions that this container can provide
func GetIReferenceTokenStoreDefinitions(ctn di.Container) []*di.Def {
	defs := ctn.GetDefinitionsByType(ReflectTypeIReferenceTokenStore)
	return defs
}

// SafeGetManyIReferenceTokenStoreFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetManyIReferenceTokenStoreFromContainer(ctn di.Container) ([]IReferenceTokenStore, error) {
	objs, err := ctn.SafeGetManyByType(ReflectTypeIReferenceTokenStore)
	if err != nil {
		return nil, err
	}
	var results []IReferenceTokenStore
	for _, obj := range objs {
		results = append(results, obj.(IReferenceTokenStore))
	}
	return results, nil
}

type _logIReferenceTokenStoreExtra struct {
	Name  string
	Value interface{}
}

func _logAddIReferenceTokenStore(scopeType string, implType reflect.Type, interfaces string, extra ..._logIReferenceTokenStoreExtra) {
	infoEvent := log.Info().
		Str("DI", scopeType).
		Str("DI-I", interfaces).
		Str("DI-B", implType.Elem().String())

	for _, extra := range extra {
		infoEvent = infoEvent.Interface(extra.Name, extra.Value)
	}

	infoEvent.Send()

}
func _getImplementedIReferenceTokenStoreNames(implementedTypes ...reflect.Type) string {
	builder := strings.Builder{}
	for idx, implementedType := range implementedTypes {
		builder.WriteString(implementedType.Name())
		if idx < len(implementedTypes)-1 {
			builder.WriteString(", ")
		}
	}
	return builder.String()
}