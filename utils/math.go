package uitls

func Gcd[I Int](a, b I) I {
	if a > b {
		a, b = b, a
	}
	for a > 0 {
		a, b = b%a, a
	}
	return b
}

func GcdX(a, b int, x, y *int) int {
	if a == 0 {
		*x = 0
		*y = 1
		return b
	}
	var x1, y1 int
	d := GcdX(b%a, a, &x1, &y1)
	*x = y1 - (b/a)*x1
	*y = x1
	return d
}

func Abs[N Number](a N) N {
	if a < 0 {
		return -a
	}
	return a
}

func Sign[N Number](a N) N {
	if a < 0 {
		return -1
	} else if a > 0 {
		return 1
	}
	return 0
}

type Ordered interface {
	Number | ~string
}

func Max[T Ordered](a, b T) T {
	if a < b {
		return b
	}
	return a
}

func Min[T Ordered](a, b T) T {
	if a > b {
		return b
	}
	return a
}
