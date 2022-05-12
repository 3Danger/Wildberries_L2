package Parse

import "strings"

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
