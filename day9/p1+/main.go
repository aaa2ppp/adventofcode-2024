package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"

	. "adventofcode-2024/utils"
)

func solve(zipMap []byte) int {
	if n := len(zipMap) / 2; n == 0 {
		// усекаем пустое место в конце
		zipMap = zipMap[:n-1]
	}

	var unzipMap []int
	count := 0

	for i, l, r := 0, 0, len(zipMap)-1; l <= r; {
		if l%2 == 0 {

			// это файл l/2 - считаем его

			if zipMap[l] == '0' {
				l++
				continue
			}

			zipMap[l]--
			count += i * (l / 2)
			i++

			if debugEnable {
				unzipMap = append(unzipMap, l/2)
			}

		} else {

			// это пустое место - перекладываем на него файл r/2 и считаем

			if zipMap[l] == '0' {
				l++
				continue
			}
			if zipMap[r] == '0' {
				r -= 2
				continue
			}

			zipMap[l]--
			zipMap[r]--
			count += i * (r / 2)
			i++

			if debugEnable {
				unzipMap = append(unzipMap, r/2)
			}
		}
	}

	if debugEnable {
		log.Println("unzipMap:", unzipMap)
	}

	return count
}

func run(in io.Reader, out io.Writer) {

	buf, err := io.ReadAll(in)
	if err != nil {
		panic(err)
	}
	bw := bufio.NewWriter(out)
	defer bw.Flush()

	zipMap := bytes.TrimSpace(buf)

	ans := solve(zipMap)
	WriteInt(bw, ans, DefaultWriteOpts())
}

// ----------------------------------------------------------------------------

var _, debugEnable = os.LookupEnv("DEBUG")

func main() {
	_ = debugEnable
	run(os.Stdin, os.Stdout)
}
