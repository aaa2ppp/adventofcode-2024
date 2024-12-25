package uitls

import (
	"bufio"
	"io"
	"strconv"
	"unsafe"
)

func ScanWord(sc *bufio.Scanner) (string, error) {
	if sc.Scan() {
		return sc.Text(), nil
	}
	return "", io.EOF
}

func ScanInt(sc *bufio.Scanner) (int, error)                  { return ScanIntX[int](sc) }
func ScanTwoInt(sc *bufio.Scanner) (_, _ int, _ error)        { return ScanTwoIntX[int](sc) }
func ScanThreeInt(sc *bufio.Scanner) (_, _, _ int, _ error)   { return ScanThreeIntX[int](sc) }
func ScanFourInt(sc *bufio.Scanner) (_, _, _, _ int, _ error) { return ScanFourIntX[int](sc) }

func ScanIntX[T Int](sc *bufio.Scanner) (res T, err error) {
	if sc.Scan() {
		v, err := strconv.ParseInt(UnsafeString(sc.Bytes()), 0, int(unsafe.Sizeof(res))<<3)
		return T(v), err
	}

	err = sc.Err()
	if err == nil {
		err = io.EOF
	}

	return 0, err
}

func ScanTwoIntX[T Int](sc *bufio.Scanner) (v1, v2 T, err error) {
	v1, err = ScanIntX[T](sc)
	if err == nil {
		v2, err = ScanIntX[T](sc)
	}
	return v1, v2, err
}

func ScanThreeIntX[T Int](sc *bufio.Scanner) (v1, v2, v3 T, err error) {
	v1, err = ScanIntX[T](sc)
	if err == nil {
		v2, err = ScanIntX[T](sc)
	}
	if err == nil {
		v3, err = ScanIntX[T](sc)
	}
	return v1, v2, v3, err
}

func ScanFourIntX[T Int](sc *bufio.Scanner) (v1, v2, v3, v4 T, err error) {
	v1, err = ScanIntX[T](sc)
	if err == nil {
		v2, err = ScanIntX[T](sc)
	}
	if err == nil {
		v3, err = ScanIntX[T](sc)
	}
	if err == nil {
		v4, err = ScanIntX[T](sc)
	}
	return v1, v2, v3, v4, err
}

func ScanInts[T Int](sc *bufio.Scanner, a []T) error {
	for i := range a {
		v, err := ScanIntX[T](sc)
		if err != nil {
			return err
		}
		a[i] = v
	}
	return nil
}

type WriteOpts struct {
	Sep   byte
	Begin byte
	End   byte
}

func DefaultWriteOpts() WriteOpts {
	return WriteOpts{Sep: ' ', End: '\n'}
}

func WriteInt[I Int](bw *bufio.Writer, v I, opts WriteOpts) error {
	var buf [32]byte

	var err error
	if opts.Begin != 0 {
		err = bw.WriteByte(opts.Begin)
	}

	if err == nil {
		_, err = bw.Write(strconv.AppendInt(buf[:0], int64(v), 10))
	}

	if err == nil && opts.End != 0 {
		err = bw.WriteByte(opts.End)
	}

	return err
}

func WriteInts[I Int | byte](bw *bufio.Writer, a []I, opts WriteOpts) error {
	var err error
	if opts.Begin != 0 {
		err = bw.WriteByte(opts.Begin)
	}

	if len(a) != 0 {
		var buf [32]byte

		if opts.Sep == 0 {
			opts.Sep = ' '
		}

		_, err = bw.Write(strconv.AppendInt(buf[:0], int64(a[0]), 10))

		for i := 1; err == nil && i < len(a); i++ {
			err = bw.WriteByte(opts.Sep)
			if err == nil {
				_, err = bw.Write(strconv.AppendInt(buf[:0], int64(a[i]), 10))
			}
		}
	}

	if err == nil && opts.End != 0 {
		err = bw.WriteByte(opts.End)
	}

	return err
}
