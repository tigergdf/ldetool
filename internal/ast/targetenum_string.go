// Code generated by "stringer -type TargetEnum src/logparsergen/ast/target-enums.go"; DO NOT EDIT

package ast

import "fmt"

const _TargetEnum_name = "StringChar"

var _TargetEnum_index = [...]uint8{0, 6, 10}

func (i TargetEnum) String() string {
	if i < 0 || i >= TargetEnum(len(_TargetEnum_index)-1) {
		return fmt.Sprintf("TargetEnum(%d)", i)
	}
	return _TargetEnum_name[_TargetEnum_index[i]:_TargetEnum_index[i+1]]
}
