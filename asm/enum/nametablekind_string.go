// Code generated by "string2enum -linecomment -type NameTableKind /home/u/Desktop/go/src/github.com/llir/llvm/ir/enum"; DO NOT EDIT.

package enum

import "fmt"
import "github.com/llir/llvm/ir/enum"

const _NameTableKind_name = "DefaultGNUNone"

var _NameTableKind_index = [...]uint8{0, 7, 10, 14}

func NameTableKindFromString(s string) enum.NameTableKind {
	if len(s) == 0 {
		return 0
	}
	for i := range _NameTableKind_index[:len(_NameTableKind_index)-1] {
		if s == _NameTableKind_name[_NameTableKind_index[i]:_NameTableKind_index[i+1]] {
			return enum.NameTableKind(i)
		}
	}
	panic(fmt.Errorf("unable to locate NameTableKind enum corresponding to %q", s))
}
