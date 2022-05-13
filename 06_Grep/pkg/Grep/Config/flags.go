package Config

import (
	"flag"
	"io"
	"log"
	"os"
	"strings"
)

const (
	constAlnum  = "[[:alnum:]]"
	constAlpha  = "[[:alpha:]]"
	constBlank  = "[[:blank:]]"
	constCntrl  = "[[:cntrl:]]"
	constDigit  = "[[:digit:]]"
	constGraph  = "[[:graph:]]"
	constPunct  = "[[:punct:]]"
	constLower  = "[[:lower:]]"
	constPrint  = "[[:print:]]"
	constSpace  = "[[:space:]]"
	constUpper  = "[[:upper:]]"
	constXdigit = "[[:xdigit:]]"

	constAlnumEX  = "[0-9A-Za-z]"
	constAlphaEX  = "[A-Za-z]"
	constBlankEX  = "\t\n\v\r\b\f\r"
	constCntrlEX  = "[\000\037\177]"
	constDigitEX  = "[0-9]"
	constGraphEX  = "[:alnum:]|[:punct:]"
	constPunctEx  = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~."
	constLowerEX  = "[a-z]"
	constPrintEX  = "([:alnum:]|[:punct:]| )"
	constSpaceEX  = "\n\v\b\r "
	constUpperEX  = "[A-Z]"
	constXdigitEX = "[0-9a-fA-F]"
)

func PrepareRequest(request string) string {
	request = strings.ReplaceAll(request, constAlnum, constAlnumEX)
	request = strings.ReplaceAll(request, constAlpha, constAlphaEX)
	request = strings.ReplaceAll(request, constBlank, constBlankEX)
	request = strings.ReplaceAll(request, constCntrl, constCntrlEX)
	request = strings.ReplaceAll(request, constDigit, constDigitEX)
	request = strings.ReplaceAll(request, constPunct, constPunctEx)
	request = strings.ReplaceAll(request, constGraph, constGraphEX)
	request = strings.ReplaceAll(request, constLower, constLowerEX)
	request = strings.ReplaceAll(request, constPrint, constPrintEX)
	request = strings.ReplaceAll(request, constSpace, constSpaceEX)
	request = strings.ReplaceAll(request, constUpper, constUpperEX)
	request = strings.ReplaceAll(request, constXdigit, constXdigitEX)
	return request
}

type Conf struct {
	KeyA    int  // "after" печатать +N строк после совпадения
	KeyB    int  // "before" печатать +N строк до совпадения
	KeyC    int  // "context" (A+B) печатать ±N строк вокруг совпадения
	Keyc    bool // "count" (количество строк)
	Keyi    bool // "ignore-case" (игнорировать регистр)
	Keyv    bool // "invert" (вместо совпадения, исключать)
	KeyF    bool // "fixed", точное совпадение со строкой, не паттерн
	Keyn    bool // "line num", напечатать номер строки
	Input   io.Reader
	Request string
}

func prepareFlagsABC(conf *Conf) {
	if conf.KeyC != -1 {
		if conf.KeyA == -1 {
			conf.KeyA = conf.KeyC
		}
		if conf.KeyB == -1 {
			conf.KeyB = conf.KeyC
		}
	} else {
		if conf.KeyA == -1 {
			conf.KeyA = 0
		}
		if conf.KeyB == -1 {
			conf.KeyB = 0
		}
	}
}

func GetConfig() Conf {
	conf := Conf{}
	flag.IntVar(&conf.KeyA, "A", -1, "\"after\" печатать +N строк после совпадения")
	flag.IntVar(&conf.KeyB, "B", -1, "\"before\" печатать +N строк до совпадения")
	flag.IntVar(&conf.KeyC, "C", -1, "\"context\" (A+B) печатать ±N строк вокруг совпадения")
	flag.BoolVar(&conf.Keyc, "c", false, "\"count\" (количество строк)")
	flag.BoolVar(&conf.Keyi, "i", false, "\"ignore-case\" (игнорировать регистр)")
	flag.BoolVar(&conf.Keyv, "v", false, "\"invert\" (вместо совпадения, исключать)")
	flag.BoolVar(&conf.KeyF, "F", false, "\"fixed\" точное совпадение со строкой, не паттерн")
	flag.BoolVar(&conf.Keyn, "n", false, "\"line num\", напечатать номер строки")
	flag.Parse()

	prepareFlagsABC(&conf)
	var filename string
	conf.Request, filename = findFile()
	conf.Input = GetFile(filename)
	return conf
}

func findFile() (request, filename string) {
	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-A", "-B", "-C":
			i++
		case "-c", "-i", "-v", "-F", "-n":
		default:
			if request == "" {
				request = os.Args[i]
			} else {
				filename = os.Args[i]
			}
		}
	}
	return request, filename
}

func GetFile(filename string) io.Reader {
	var ok error
	var file *os.File
	if filename != "" {
		if file, ok = os.Open(filename); ok == nil {
			return file
		}
		log.Fatal(ok)
	}
	return os.Stdin
}
