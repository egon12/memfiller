package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// s size of this struct is 8 bytes
type s struct {
	v int64
}

const (
	kb = 1024
	mb = 1024 * kb
	gb = 1024 * mb
	tb = 1024 * gb

	// 100 MB for shortcut
	mb100 = 100 * mb
)

var sizeType = regexp.MustCompile(`^(\d+)([kKmMgGtT])$`)

func main() {
	max := getInputSize("Enter size of memory to allocate (e.g. 1G): ")

	var ss []s = make([]s, max/8, max/8)

	println(len(ss) * 8)

	incr := getInputSize("Enter increment size per second (e.g. 1M): ")
	incrs := make([]s, incr/8, incr/8)
	c := time.Tick(1 * time.Second)
	for {
		<-c
		ss = append(ss, incrs...)
		println(len(ss) * 8)
	}
}

func getInputSize(message string) int {
	println(message)

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		println("error reading input")
		return -1
	}

	return getSize(input)
}

func getSize(input string) int {
	match := sizeType.FindStringSubmatch(input)
	if len(match) != 3 {
		println("error parsing input")
		return -1
	}
	num, _ := strconv.Atoi(match[1])
	switch match[2] {
	case "k", "K":
		return num * kb
	case "m", "M":
		return num * mb
	case "g", "G":
		return num * gb
	case "t", "T":
		return num * tb
	default:
		return num
	}

	return 0
}
