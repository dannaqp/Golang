package operaciones

func suma(a, b int) int {
	return a + b
}

func Sumaresta(c, d int) (int, int) {
	if c > d {
		return c + d, c - d
	}
	return c + d, 0
}

func sumatoria(nums ...int) int {
	total := 0
	for _, nro := range nums {
		total += nro
	}
	return total
}
