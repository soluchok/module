package main

import (
	"fmt"
	"github.com/soluchok/module/submodule1"
	submodule1v2 "github.com/soluchok/module/submodule1/v2"
	submodule1v3 "github.com/soluchok/module/submodule1/v3"
	"github.com/soluchok/module/submodule2"
	submodule2v2 "github.com/soluchok/module/submodule2/v2"
	submodule2v3 "github.com/soluchok/module/submodule2/v3"
)

func main() {
	fmt.Println(submodule1.V)
	fmt.Println(submodule2.V)
	fmt.Println(submodule1v2.V)
	fmt.Println(submodule2v2.V)
	fmt.Println(submodule1v3.V)
	fmt.Println(submodule2v3.V)

	// Output:
	// submodule1-v1
	// submodule2-v1
	// submodule1-v2
	// submodule2-v2
	// submodule1-v3
	// submodule2-v3
}
