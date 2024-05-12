package factory

import (
	"github.com/abmpio/x/str"
)

// ParseUnderlyTypeId, parse type string descriptor
// if t is nil,return ""
// if t is primitive types,return type string, for example,return "string","int","int32"
// if t is custome struct,will contain package and type
func ParseUnderlyTypeId(t interface{}) string {
	if t == nil {
		return ""
	}

	pkgName, typeName := GetPkgAndName(t)
	if pkgName != "" && typeName != "" {
		return pkgName + "." + str.ToLowerCamel(typeName)
	}
	if pkgName != "" {
		return pkgName
	}
	if typeName != "" {
		return typeName
	}
	return ""
}
