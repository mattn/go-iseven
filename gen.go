package iseven

import (
	"fmt"
)

//go:generate go run gen/generate.go

type Number interface{}

//go:noinline
func is(lhs, rhs interface{}) bool {
	//switch t := lhs.(type) {
	//case string:
	//return t == fmt.Sprint(rhs)
	//case int, int32, int64, uint, uint32, uint64:
	//return t == rhs.
	//}
	return fmt.Sprint(lhs) == fmt.Sprint(rhs)
}
