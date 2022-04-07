// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package probe

import (
	"reflect"
	"strings"

	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

// ReflectTypeIProbe used when your service claims to implement IProbe
var ReflectTypeIProbe = di.GetInterfaceReflectType((*IProbe)(nil))

// AddSingletonIProbe adds a type that implements IProbe
func AddSingletonIProbe(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIProbe)
	_logAddIProbe("SINGLETON", implType, _getImplementedIProbeNames(implementedTypes...),
		_logIProbeExtra{
			Name:  "DI-BY",
			Value: "type",
		})
	di.AddSingleton(builder, implType, implementedTypes...)
}

// AddSingletonIProbeWithMetadata adds a type that implements IProbe
func AddSingletonIProbeWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIProbe)
	_logAddIProbe("SINGLETON", implType, _getImplementedIProbeNames(implementedTypes...),
		_logIProbeExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logIProbeExtra{
			Name:  "DI-M",
			Value: metaData,
		})
	di.AddSingletonWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddSingletonIProbeByObj adds a prebuilt obj
func AddSingletonIProbeByObj(builder *di.Builder, obj interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIProbe)
	_logAddIProbe("SINGLETON", reflect.TypeOf(obj), _getImplementedIProbeNames(implementedTypes...),
		_logIProbeExtra{
			Name:  "DI-BY",
			Value: "obj",
		})
	di.AddSingletonWithImplementedTypesByObj(builder, obj, implementedTypes...)
}

// AddSingletonIProbeByObjWithMetadata adds a prebuilt obj
func AddSingletonIProbeByObjWithMetadata(builder *di.Builder, obj interface{}, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIProbe)
	_logAddIProbe("SINGLETON", reflect.TypeOf(obj), _getImplementedIProbeNames(implementedTypes...),
		_logIProbeExtra{
			Name:  "DI-BY",
			Value: "obj",
		},
		_logIProbeExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddSingletonWithImplementedTypesByObjWithMetadata(builder, obj, metaData, implementedTypes...)
}

// AddSingletonIProbeByFunc adds a type by a custom func
func AddSingletonIProbeByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIProbe)
	_logAddIProbe("SINGLETON", implType, _getImplementedIProbeNames(implementedTypes...),
		_logIProbeExtra{
			Name:  "DI-BY",
			Value: "func",
		})
	di.AddSingletonWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddSingletonIProbeByFuncWithMetadata adds a type by a custom func
func AddSingletonIProbeByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIProbe)
	_logAddIProbe("SINGLETON", implType, _getImplementedIProbeNames(implementedTypes...),
		_logIProbeExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logIProbeExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddSingletonWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddTransientIProbe adds a type that implements IProbe
func AddTransientIProbe(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIProbe)
	_logAddIProbe("TRANSIENT", implType, _getImplementedIProbeNames(implementedTypes...),
		_logIProbeExtra{
			Name:  "DI-BY",
			Value: "type",
		})

	di.AddTransientWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddTransientIProbeWithMetadata adds a type that implements IProbe
func AddTransientIProbeWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIProbe)
	_logAddIProbe("TRANSIENT", implType, _getImplementedIProbeNames(implementedTypes...),
		_logIProbeExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logIProbeExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddTransientWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddTransientIProbeByFunc adds a type by a custom func
func AddTransientIProbeByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIProbe)
	_logAddIProbe("TRANSIENT", implType, _getImplementedIProbeNames(implementedTypes...),
		_logIProbeExtra{
			Name:  "DI-BY",
			Value: "func",
		})

	di.AddTransientWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddTransientIProbeByFuncWithMetadata adds a type by a custom func
func AddTransientIProbeByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIProbe)
	_logAddIProbe("TRANSIENT", implType, _getImplementedIProbeNames(implementedTypes...),
		_logIProbeExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logIProbeExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddTransientWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddScopedIProbe adds a type that implements IProbe
func AddScopedIProbe(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIProbe)
	_logAddIProbe("SCOPED", implType, _getImplementedIProbeNames(implementedTypes...),
		_logIProbeExtra{
			Name:  "DI-BY",
			Value: "type",
		})
	di.AddScopedWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddScopedIProbeWithMetadata adds a type that implements IProbe
func AddScopedIProbeWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIProbe)
	_logAddIProbe("SCOPED", implType, _getImplementedIProbeNames(implementedTypes...),
		_logIProbeExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logIProbeExtra{
			Name:  "DI-M",
			Value: metaData,
		})
	di.AddScopedWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddScopedIProbeByFunc adds a type by a custom func
func AddScopedIProbeByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIProbe)
	_logAddIProbe("SCOPED", implType, _getImplementedIProbeNames(implementedTypes...),
		_logIProbeExtra{
			Name:  "DI-BY",
			Value: "func",
		})
	di.AddScopedWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddScopedIProbeByFuncWithMetadata adds a type by a custom func
func AddScopedIProbeByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIProbe)
	_logAddIProbe("SCOPED", implType, _getImplementedIProbeNames(implementedTypes...),
		_logIProbeExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logIProbeExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddScopedWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// RemoveAllIProbe removes all IProbe from the DI
func RemoveAllIProbe(builder *di.Builder) {
	builder.RemoveAllByType(ReflectTypeIProbe)
}

// GetIProbeFromContainer alternative to SafeGetIProbeFromContainer but panics of object is not present
func GetIProbeFromContainer(ctn di.Container) IProbe {
	return ctn.GetByType(ReflectTypeIProbe).(IProbe)
}

// GetManyIProbeFromContainer alternative to SafeGetManyIProbeFromContainer but panics of object is not present
func GetManyIProbeFromContainer(ctn di.Container) []IProbe {
	objs := ctn.GetManyByType(ReflectTypeIProbe)
	var results []IProbe
	for _, obj := range objs {
		results = append(results, obj.(IProbe))
	}
	return results
}

// SafeGetIProbeFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetIProbeFromContainer(ctn di.Container) (IProbe, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeIProbe)
	if err != nil {
		return nil, err
	}
	return obj.(IProbe), nil
}

// GetIProbeDefinition returns that last definition registered that this container can provide
func GetIProbeDefinition(ctn di.Container) *di.Def {
	def := ctn.GetDefinitionByType(ReflectTypeIProbe)
	return def
}

// GetIProbeDefinitions returns all definitions that this container can provide
func GetIProbeDefinitions(ctn di.Container) []*di.Def {
	defs := ctn.GetDefinitionsByType(ReflectTypeIProbe)
	return defs
}

// SafeGetManyIProbeFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetManyIProbeFromContainer(ctn di.Container) ([]IProbe, error) {
	objs, err := ctn.SafeGetManyByType(ReflectTypeIProbe)
	if err != nil {
		return nil, err
	}
	var results []IProbe
	for _, obj := range objs {
		results = append(results, obj.(IProbe))
	}
	return results, nil
}

type _logIProbeExtra struct {
	Name  string
	Value interface{}
}

func _logAddIProbe(scopeType string, implType reflect.Type, interfaces string, extra ..._logIProbeExtra) {
	infoEvent := log.Info().
		Str("DI", scopeType).
		Str("DI-I", interfaces).
		Str("DI-B", implType.Elem().String())

	for _, extra := range extra {
		infoEvent = infoEvent.Interface(extra.Name, extra.Value)
	}

	infoEvent.Send()

}
func _getImplementedIProbeNames(implementedTypes ...reflect.Type) string {
	builder := strings.Builder{}
	for idx, implementedType := range implementedTypes {
		builder.WriteString(implementedType.Name())
		if idx < len(implementedTypes)-1 {
			builder.WriteString(", ")
		}
	}
	return builder.String()
}
