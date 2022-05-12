package Grep

import (
	"fmt"
	"grep/pkg/Config"
	"grep/pkg/Grep/Found"
	"grep/pkg/Grep/Parse"
	"grep/pkg/io"
	"log"
	"regexp"
	"strings"
)

type Grep struct {
	cnf     Config.Conf
	rawData []string
}

func (g Grep) GetConf() Config.Conf { return g.cnf }
func (g Grep) GetData() []string    { return g.rawData }

func NewGrep() *Grep {
	cnf := Config.GetConfig()
	var rawData = io.GetData(cnf)
	cnf.Request = Parse.PrepareRequest(cnf.Request)
	return &Grep{cnf, rawData}
}

func (g *Grep) Run() string {
	var reg *regexp.Regexp
	var pref, post string
	var err error

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
	if g.cnf.Keyc {
		count := 0
		for _, v := range g.rawData {
			seg := reg.FindIndex([]byte(v))
			if seg != nil {
				count++
			}
		}
		return fmt.Sprint(count)
	}
	found := CreateFoundGroup(g, reg)
	sb := strings.Builder{}
	for _, v := range found {
		sb.WriteString(strings.Join(v.GetData(), ""))
	}
	return sb.String()
}

func CreateFoundGroup(g *Grep, reg *regexp.Regexp) []*Found.Found {
	var pointIndex = make([]*Found.PointIndex, 0, 10)
	for i, v := range g.rawData {
		seg := reg.FindIndex([]byte(v))
		if seg != nil {
			pointIndex = append(pointIndex, Found.NewPointIndex(i, g.cnf.KeyB, g.cnf.KeyA))
		}
	}
	pointIndex = Found.MixPoints(pointIndex...)
	var found = make([]*Found.Found, 0, len(pointIndex))
	for i := range pointIndex {
		found = append(found, Found.NewFound(g.GetConf(), g.rawData, pointIndex[i]))
	}
	return found
}
