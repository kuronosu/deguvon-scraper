package utils

import (
	"fmt"
	"io"
	"strings"
	"sync"
)

type ProgressbarFunc[T any] func([]T, string, int, io.Writer, func(T))

func show(i, size, count int, prefix string, out io.Writer) {
	if out == nil {
		return
	}
	x := i * size / count
	fmt.Fprintf(out, "\r%s[%s%s] %d/%d", prefix, strings.Repeat("=", x), strings.Repeat(" ", size-x), i, count)
}

func Progressbar[T any](data []T, prefix string, size int, out io.Writer, callback func(T)) {
	count := len(data)
	show(0, size, count, prefix, out)
	for i, item := range data {
		callback(item)
		show(i+1, size, count, prefix, out)
	}
	fmt.Fprintln(out)
}

func ProgressbarAsync[T any](data []T, prefix string, size int, out io.Writer, callback func(T)) {
	count := len(data)
	show(0, size, count, prefix, out)
	var wg sync.WaitGroup
	counter := NewSafeCounter()
	for i, item := range data {
		wg.Add(1)
		go func(item T, i int) {
			defer wg.Done()
			callback(item)
			show(counter.IncAndGet(), size, count, prefix, out)
		}(item, i)
	}
	wg.Wait()
	if out != nil {
		fmt.Fprintln(out)
	}
}
