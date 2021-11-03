package utils

import (
	"encoding/json"
	"strconv"
	"sync"
)

func Once(f func()) {
	var once sync.Once
	once.Do(f)
}

func Atoint32(istr string) int32 {
	i, _ := strconv.Atoi(istr)
	return int32(i)
}

func IntToA(i int) string {
	s := strconv.Itoa(i)
	return s
}

func ToJson(obj interface{}) string {
	//jj, ee := json.Marshal(infos)
	jj, ee := json.MarshalIndent(obj, "", "     ")
	if ee == nil {
		return string(jj)
	}
	return ""

}
