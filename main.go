package main

import (
	"fmt"
	"github.com/soluchok/module/submodule1"
	submodule1v2 "github.com/soluchok/module/submodule1/v2"
	"github.com/soluchok/module/submodule2"
	submodule2v2 "github.com/soluchok/module/submodule2/v2"
)

func main() {
	fmt.Println(submodule1.V)
	fmt.Println(submodule2.V)
	fmt.Println(submodule1v2.V)
	fmt.Println(submodule2v2.V)

	// Output:
	// submodule1-v1
	// submodule2-v1
	// submodule1-v2
	// submodule2-v2
}
