package clogrus

import (
	"fmt"
	"path"
	"runtime"
	"strings"

	"github.com/Sirupsen/logrus"
)

func getProgDetails() string {
	pc, file, line, _ := runtime.Caller(2)
	_, fileName := path.Split(file)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)
	packageName := ""
	funcName := parts[pl-1]

	if parts[pl-2][0] == '(' {
		funcName = parts[pl-2] + "." + funcName
		packageName = strings.Join(parts[0:pl-2], ".")
	} else {
		packageName = strings.Join(parts[0:pl-1], ".")
	}

	return fmt.Sprintf("\n pkg:%s,file:%s,func:%s,line:%d",
		packageName, fileName, funcName, line)
}

//InfoLogWithDetails is a logrus wrapper that will provide additional
//loggging data like pkg, file, func, line and logs as an info
func InfoLogWithDetails(args ...interface{}) {
	logrus.Infoln(args, getProgDetails())
}

//DebugLogWithDetails is a logrus wrapper that will provide additional
//loggging data like pkg, file, func, line and logs as a debugln
func DebugLogWithDetails(args ...interface{}) {
	logrus.Debugln(args, getProgDetails())
}
