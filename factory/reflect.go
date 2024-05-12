package factory

import (
	"reflect"

	"github.com/abmpio/x/str"
)

// GetObjectType get the function output data type
func GetObjectType(object interface{}) (typ reflect.Type, ok bool) {
	if object == nil {
		return nil, false
	}
	var reflectType reflect.Type
	var reflectValue reflect.Value
	switch v := object.(type) {
	case reflect.Value:
		reflectType = v.Type()
	case reflect.Type:
		reflectType = v
	default:
		reflectValue = reflect.ValueOf(object)
		reflectType = reflectValue.Type()
	}
	typName := reflectType.Name()
	typKind := reflectType.Kind()
	if typKind == reflect.Func {
		numOut := reflectType.NumOut()
		if numOut > 0 {
			typ = IndirectType(reflectType.Out(0))
			ok = true
		}
	} else if typKind == reflect.Struct {
		if typName == "Method" {
			method := object.(reflect.Method)
			methodTyp := method.Func.Type()
			numOut := methodTyp.NumOut()
			if numOut > 0 {
				typ = IndirectType(methodTyp.Out(0))
				ok = true
			}
		} else {
			typ = reflect.TypeOf(object)
			ok = true
		}
	} else {
		// TODO: check if it effects others
		typ = IndirectType(reflectType)
		ok = true
	}

	return
}

// GetLowerCamelFullName get the object name with package name, e.g. pkg.objectName
func GetLowerCamelFullName(object interface{}) (name string) {

	pn, n := GetPkgAndName(object)
	name = pn + "." + str.ToLowerCamel(n)

	return
}

// GetPkgAndName get the package name and the object name with, e.g. pkg, Object
func GetPkgAndName(object interface{}) (pkgName, name string) {
	typ, ok := GetObjectType(object)
	if ok {
		pkgName = typ.PkgPath()
		name = typ.Name()
	}
	return
}

// IndirectType get indirect type
func IndirectType(reflectType reflect.Type) reflect.Type {
	for reflectType.Kind() == reflect.Ptr || reflectType.Kind() == reflect.Slice {
		reflectType = reflectType.Elem()
	}
	return reflectType
}

func normalPkgPath(pkgPath string) string {
	// path :=  io.DirName(pkgPath)
	return pkgPath
}

// GetLowerCamelFullNameByType get the object name with package name by type, e.g. pkg.Object
func GetLowerCamelFullNameByType(objType reflect.Type) (name string) {
	indTyp := IndirectType(objType)
	depPkgName := normalPkgPath(indTyp.PkgPath())
	name = depPkgName + "." + str.ToLowerCamel(indTyp.Name())
	return
}
