package common

import (
	"github.com/go-xorm/xorm"
)

var mapXormEngine map[string]*xorm.Engine = make(map[string]*xorm.Engine)

func SetEngine(key string, e *xorm.Engine) {
	mapXormEngine[key] = e
}

func GetEngine(key string) *xorm.Engine {
	e, ok := mapXormEngine[key]
	if !ok {
		panic("engine not exist")
	}
	return e
}
