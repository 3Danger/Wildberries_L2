package Config

import (
	"errors"
	"flag"
	"io"
	"log"
	"os"
)

type Conf struct {
	A     int  // "after" печатать +N строк после совпадения
	B     int  // "before" печатать +N строк до совпадения
	C     int  // "context" (A+B) печатать ±N строк вокруг совпадения
	c     bool // "count" (количество строк)
	i     bool // "ignore-case" (игнорировать регистр)
	v     bool // "invert" (вместо совпадения, исключать)
	F     bool // "fixed", точное совпадение со строкой, не паттерн
	n     bool // "line num", напечатать номер строки
	input io.Reader
}

func find(s string, haystack *[]string) (int, error) {
	for i, v := range *haystack {
		if v == s {
			return i, nil
		}
	}
	return -1, errors.New("not found")
}

func GetConfig() Conf {
	conf := Conf{}
	flag.IntVar(&conf.A, "A", 0, "\"after\" печатать +N строк после совпадения")
	flag.IntVar(&conf.B, "B", 0, "\"before\" печатать +N строк до совпадения")
	flag.IntVar(&conf.C, "C", 0, "\"context\" (A+B) печатать ±N строк вокруг совпадения")
	flag.BoolVar(&conf.c, "c", false, "\"count\" (количество строк)")
	flag.BoolVar(&conf.i, "i", false, "\"ignore-case\" (игнорировать регистр)")
	flag.BoolVar(&conf.v, "v", false, "\"invert\" (вместо совпадения, исключать)")
	flag.BoolVar(&conf.F, "F", false, "\"fixed\" точное совпадение со строкой, не паттерн")
	flag.BoolVar(&conf.n, "n", false, "\"line num\", напечатать номер строки")
	flag.Parse()
	conf.input = GetFile()
	return conf
}

func findFile() *string {
	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-A", "-B", "-C":
			i++
		case "-c", "-i", "-v", "-F", "-n":
		default:
			return &os.Args[i]
		}
	}
	return nil
}

func GetFile() io.Reader {
	var ok error
	var file *os.File
	flname := findFile()
	if flname != nil {
		if file, ok = os.Open(*flname); ok == nil {
			return file
		}
		log.Fatal(ok)
	}
	return os.Stdin
}
