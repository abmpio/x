package factory

import (
	"github.com/abmpio/x/str"
)

// ParseUnderlyTypeId, parse type string descriptor
// if t is nil,return ""
func ParseUnderlyTypeId(t interface{}) string {
	if t == nil {
		return ""
	}

	name := ""
	pkgName, typeName := GetPkgAndName(t)
	if pkgName != "" {
		name = pkgName + "." + str.ToLowerCamel(typeName)
	}
	return name
}
