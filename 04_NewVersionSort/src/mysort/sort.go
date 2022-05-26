package mysort

import (
	"04_Sort2/src/config"
	"04_Sort2/src/utils"
	"fmt"
	"os"
)

type Sort struct {
	data     StringerElements
	KeyFlags *config.Keys
}

//TODO delete after debug
func (s *Sort) DebugSetKeys(keys ...byte) {
	s.KeyFlags.DebugSetKeys(keys...)
}

//TODO delete after debug
func (s *Sort) DebugSetData(pathFile string) {
	f, ok := os.Open(pathFile)
	if ok != nil {
		panic(ok)
	}
	rawData := utils.ReadToString(f)
	s.data = SliceToStringers(rawData, " ")
}

//TODO delete after debug
func (s *Sort) DebugSetKeyFlags(keys *config.Keys) {
	s.KeyFlags = keys
}

func NewSort() (sort *Sort) {
	sort = new(Sort)
	conf := config.NewConfig()
	sort.KeyFlags = conf.KeysFlags
	rawData := utils.ReadToString(conf.Input)
	sort.data = SliceToStringers(rawData, " ")
	return sort
}

func (s *Sort) Strategy(method Operation) {
	method.Do(s.data)
	fmt.Print(s.data.ToString())
}
