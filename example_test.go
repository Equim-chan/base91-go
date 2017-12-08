package base91_test

import (
	"fmt"

	"ekyu.moe/base91"
)

func Example() {
	fmt.Println(base91.EncodeToString([]byte("Hello, 世界")))
	fmt.Println(string(base91.DecodeString(">OwJh>}AFU~PUh%Y")))
	// Output:
	// >OwJh>}AFU~PUh%Y
	// Hello, 世界
}
