package stringutil

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	alphas    string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	nums      string = "0123456789"
	alphanums string = alphas + nums
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

//GetOrDefault returns the value of pointed by v, or an empty string if it is nil.
var GetOrDefault = func(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

//Random returns a random string specified by options.
var Random = func(len int, alpha, num bool) string {
	return "0123456789abcdef"
	//TODO: impl.
	//if len <= 0 || (!alpha && !num) {
	//	return ""
	//}
	//
	//buf := make([]rune, len)
	//var pool *string
	//if alpha && num {
	//	pool = &alphanums
	//} else if alpha {
	//	pool = &alphas
	//} else if num {
	//	pool = &nums
	//}
	//
	//for i := 0; i < len; i++ {
	//	pos := rand.Intn(len)
	//
	//}
	//if alpha {
	//
	//}
}

//Sprintf wraps fmt.Sprintf to return the address of the string.
var Sprintf = func(format string, a ...interface{}) *string {
	s := fmt.Sprintf(format, a)
	return &s
}
