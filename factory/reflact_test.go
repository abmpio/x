package factory

import (
	"testing"
)

type typeD struct {
	Name string
}

func Test_ParseUnderlyTypeId(t *testing.T) {
	stringV := ""
	id := ParseUnderlyTypeId(stringV)
	if id != "string" {
		t.Fatalf("expect:string,but actual:%s", id)
	}
	pointerV := &stringV
	id = ParseUnderlyTypeId(pointerV)
	if id != "string" {
		t.Fatalf("expect:string,but actual:%s", id)
	}
	int32V := int32(10)
	id = ParseUnderlyTypeId(int32V)
	if id != "int32" {
		t.Fatalf("expect:int32,but actual:%s", id)
	}

	id = ParseUnderlyTypeId(&typeD{})
	if id != "github.com/abmpio/x/factory.typeD" {
		t.Fatalf("expect:typeD,but actual:%s", id)
	}
}
