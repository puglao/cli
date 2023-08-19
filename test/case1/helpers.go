//go:build integration
// +build integration

package case1

import (
	"fmt"
)

func skipper() {
	fmt.Println("skipping test")
}
