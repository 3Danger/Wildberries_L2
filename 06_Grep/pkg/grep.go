package pkg

import (
	"fmt"
	"grep/pkg/Config"
	"grep/pkg/io"
	"log"
	"regexp"
	"strings"
)

/*
Наконец, определённые именованные классы символов предопределены внутри выражений в квадратных скобках как показано ниже.
	Их интерпретация зависит от LC_CTYPE локали; например, «[[:alnum:]]» означает класс символов из чисел и букв в текущей локали.
[:alnum:]
Алфавитные символы: «[:alpha:]» и «[:digit:]»; в локали «C» и кодировке символов ASCII, это то же самое что и «[0-9A-Za-z]».
[:alpha:]
Алфавитные символы: «[:lower:]» и «[:upper:]»; в локали «C» и кодировке символов ASCII, это то же самое что и «[A-Za-z]».
[:blank:]
Пустые символы: пробел и табуляция.
[:cntrl:]
Управляющие символы. В ASCII эти символы имеют восьмеричные коды от 000 до 037 и 177 (DEL).
	В других наборах символов это эквивалентные символы, если они есть.
[:digit:]
Цифры: 0 1 2 3 4 5 6 7 8 9.
[:graph:]
Графические символы: «[:alnum:]» и «[:punct:]».
[:lower:]
Буквы в нижнем регистре, в локали «C» и кодировке символов ASCII это a b c d e f g h i j k l m n o p q r s t u v w x y z.
[:print:]
Печатные символы: «[:alnum:]», «[:punct:]», и пробел.
[:punct:]
Пунктуационные символы; в локали «C» и кодировке символов ASCII,
	это ! " # $ % & ' ( ) * + , - . / : ; < = > ? @ [ \ ] ^ _ ` { | } ~.
[:space:]
Пробельные символы: в локали «C», это табуляция, новая строка, вертикальная табуляция, разрыв страницы, возврат каретки и пробел.
	Смотрите раздел Использование grep для дополнительной информации о совпадении новой строки.
[:upper:]
Буквы в верхнем регистре: в локали «C» и кодировке символов ASCII, это A B C D E F G H I J K L M N O P Q R S T U V W X Y Z.
[:xdigit:]
	Шестнадцатеричные цифры: 0 1 2 3 4 5 6 7 8 9 A B C D E F a b c d e f.
*/

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

type Found struct {
	*Grep
	index int
}

func (f Found) GetData() []string {
	var data []string
	var start = f.index - f.cnf.KeyB
	var finish = f.index + f.cnf.KeyA + 1

	if start < 0 {
		start = 0
	}
	if finish > len(f.rawData) {
		finish = len(f.rawData)
	}
	if f.cnf.Keyn {
		data = make([]string, finish-start)
		for i := range data {
			data = append(data, fmt.Sprintf("%d:%s", i+1, f.rawData[start]))
		}
		return data
	} else {
		return append(data, f.rawData[start:finish]...)
	}
}

type Grep struct {
	cnf     Config.Conf
	rawData []string
}

func prepareRequest(request string) string {
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

func NewGrep() *Grep {
	cnf := Config.GetConfig()
	var rawData = io.GetData(cnf)
	cnf.Request = prepareRequest(cnf.Request)
	return &Grep{cnf, rawData}
}

func (g *Grep) Run() string {
	var found = g.findAllMatches()
	sb := strings.Builder{}
	for _, v := range found {
		sb.WriteString(strings.Join(v.GetData(), "\n") + "\n")
	}
	return sb.String()
}

func (g *Grep) findAllMatches() (found []Found) {
	var reg *regexp.Regexp
	var err error
	var pref string
	var post string

	if g.cnf.Keyi {
		pref = "(?i)"
	}
	if g.cnf.KeyF {
		pref = "^" + pref
		post = "$"
	}
	reg, err = regexp.Compile(pref + g.cnf.Request + post)
	if err != nil {
		log.Fatal(err)
	}
	for i, v := range g.rawData {
		index := reg.FindIndex([]byte(v))
		if index != nil {
			found = append(found, Found{g, i})
		}
	}
	return found
}
