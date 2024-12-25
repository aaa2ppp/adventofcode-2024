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

	const (
		POS  = 0
		SIZE = 1
	)

	files := make([][2]int, 0, len(zipMap)/2+1)
	spaces := make([][2]int, 0, len(zipMap)/2)

	// разделим файлы и пустые пространства
	pos := 0
	for i, dig := range zipMap {
		size := int(dig - '0')
		if i%2 == 0 {
			files = append(files, [2]int{pos, size})
		} else {
			spaces = append(spaces, [2]int{pos, size})
		}
		pos += size
	}

	// преместим файлы в пустые простанства
	for i := len(files) - 1; i >= 0; i-- {
		fileSize := files[i][SIZE]

		// (*) если использовать самобалансирующее дерево, можно за logN искать.
		// для наших данных, на один раз и так сойдет
		for j := 0; j < i; j++ {
			if fileSize <= spaces[j][SIZE] {
				files[i][POS] = spaces[j][POS]
				spaces[j][POS] += fileSize
				spaces[j][SIZE] -= fileSize
				break
			}
		}
	}

	if debugEnable {
		size := 0
		for _, dig := range zipMap {
			size += int(dig - '0')
		}

		unzipMap := make([]int, size)
		for i := range unzipMap {
			unzipMap[i] = -1
		}

		for i := range files {
			pos := files[i][POS]
			end := pos + files[i][SIZE]
			for ; pos < end; pos++ {
				unzipMap[pos] = i
			}
		}

		log.Println("unzipMap:", unzipMap)
	}

	// сосчитаем суммы для каждого блока файла
	count := 0
	for i := range files {
		pos := files[i][POS]
		end := pos + files[i][SIZE]
		for ; pos < end; pos++ {
			count += pos * i
		}
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
