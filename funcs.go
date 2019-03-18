package version

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/lixi520/version/parser"
	"strconv"
	"strings"
)

var (
	version string
)

func Set(v string) {
	version = v
}

func CheckVersionOk(newVersion string) bool {
	if version == "" {
		return false
	}
	charstream := antlr.NewInputStream(newVersion)
	lexer := parser.NewVersionLexer(charstream)
	lexer.RemoveErrorListeners()
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	vParser := parser.NewVersionParser(tokenStream)
	vParser.BuildParseTrees = true
	service := vParser.VersionLock()
	listener := &VersionListener{}
	antlr.ParseTreeWalkerDefault.Walk(listener, service)
	return listener.IsOk
}

func Get() string {
	if version == "" {
		panic("version not support")
	}
	return version
}

func Equal(new, old string) bool {
	new = strings.Trim(new, "\n\r")
	newArr := strings.Split(new, ".")
	old = strings.Trim(old, "\n\r")
	oldArr := strings.Split(old, ".")

	for i := range [3]int{} {
		n, _ := strconv.ParseInt(newArr[i], 10, 64)
		o, _ := strconv.ParseInt(oldArr[i], 10, 64)
		if n == o {
			continue
		} else {
			return false
		}
	}
	return true
}

func Smaller(new, old string) bool {
	new = strings.Trim(new, "\n\r")
	newArr := strings.Split(new, ".")
	old = strings.Trim(old, "\n\r")
	oldArr := strings.Split(old, ".")

	for i := range [3]int{} {
		n, _ := strconv.ParseInt(newArr[i], 10, 64)
		o, _ := strconv.ParseInt(oldArr[i], 10, 64)
		if n < o {
			return true
		} else if n == o {
			continue
		} else {
			return false
		}
	}
	return false
}

func Greater(new, old string) bool {
	new = strings.Trim(new, "\n\r")
	newArr := strings.Split(new, ".")
	old = strings.Trim(old, "\n\r")
	oldArr := strings.Split(old, ".")

	for i := range [3]int{} {
		n, _ := strconv.ParseInt(newArr[i], 10, 64)
		o, _ := strconv.ParseInt(oldArr[i], 10, 64)
		if n > o {
			return true
		} else if n == o {
			continue
		} else {
			return false
		}
	}
	return false
}

func GreaterOrEqual(new, old string) bool {
	new = strings.Trim(new, "\n\r")
	newArr := strings.Split(new, ".")
	old = strings.Trim(old, "\n\r")
	oldArr := strings.Split(old, ".")

	for i := range [3]int{} {
		n, _ := strconv.ParseInt(newArr[i], 10, 64)
		o, _ := strconv.ParseInt(oldArr[i], 10, 64)
		if n > o {
			return true
		} else if n == o {
			continue
		} else {
			return false
		}
	}
	return true
}

func SmallerOrEqual(new, old string) bool {
	new = strings.Trim(new, "\n\r")
	newArr := strings.Split(new, ".")
	old = strings.Trim(old, "\n\r")
	oldArr := strings.Split(old, ".")

	for i := range [3]int{} {
		n, _ := strconv.ParseInt(newArr[i], 10, 64)
		o, _ := strconv.ParseInt(oldArr[i], 10, 64)
		if n < o {
			return true
		} else if n == o {
			continue
		} else {
			return false
		}
	}
	return true
}
