// PKGPATH: gno.land/r/std_test
package std_test

import (
	"std"
)

// Like zrealm_std1.go but interface var.
var aset std.AddressSet

func init() {
	caller := std.GetOrigCaller()
	aset = std.NewAddressList()
	aset.AddAddress(caller)
}

func main() {
	println(*(aset.(*std.AddressList)))
	caller := std.GetOrigCaller()
	err := aset.AddAddress(caller)
	println("error:", err)
	has := aset.HasAddress(caller)
	println("has:", has)
	has = aset.HasAddress(std.Address(""))
	println("has:", has)
}

// Output:
// slice[ref(1ed29bd278d735e20e296bd4afe927501941392f:4)]
// error: address already exists
// has: true
// has: false
