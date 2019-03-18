package version

import (
	"github.com/lixi520/version/parser"
	"os"
)

type VersionListener struct {
	*parser.BaseVersionListener
	IsOk bool
}

func (s *VersionListener) EnterVersionLock(ctx *parser.VersionLockContext) {

	if ctx.Cmp() != nil {

		requiredVersion := ctx.Version(0).GetText()
		switch ctx.Cmp().GetText() {
		case "<":
			if Smaller(Get(), requiredVersion) {
				s.IsOk = true
			}
			return
		case ">":
			if Greater(Get(), requiredVersion) {
				s.IsOk = true
			}
			return
		case "<=":
			if SmallerOrEqual(Get(), requiredVersion) {
				s.IsOk = true
			}
			return
		case ">=":
			if GreaterOrEqual(Get(), requiredVersion) {
				s.IsOk = true
			}
			return
		default:
			os.Stderr.WriteString("version syntax is not right\n")
			os.Exit(-1)
		}
	} else if ctx.Interval() != nil {

		lower := ctx.Version(0).GetText()
		upper := ctx.Version(1).GetText()
		current := Get()
		if GreaterOrEqual(current, lower) && SmallerOrEqual(current, upper) {
			s.IsOk = true
		}
		return
	} else {
		requiredVersion := ctx.Version(0).GetText()
		if Equal(requiredVersion, Get()) {
			s.IsOk = true
		}
		return
	}

}
