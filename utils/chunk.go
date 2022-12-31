package utils

func ChunkSlice[T any](alist []T, sublen int) [][]T {
	tmp := [][]T{}
	for i := 0; i < len(alist); i += sublen {
		tmp = append(tmp, alist[i:i+sublen])
	}
	return tmp
}

func ChunkMap[T comparable, V any](amap map[T]V, sublen int) [][]V {
	tmp := [][]V{}
	i := 0
	for _, v := range amap {
		if i%sublen == 0 {
			tmp = append(tmp, []V{})
		}
		tmp[len(tmp)-1] = append(tmp[len(tmp)-1], v)
		i++
	}
	return tmp
}
