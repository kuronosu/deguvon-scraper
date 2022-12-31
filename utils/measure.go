package utils

import (
	"fmt"
	"time"
)

func MeasureTime(name string, f func()) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Printf("%s took %s seconds to complete\n", name, elapsed)
}
