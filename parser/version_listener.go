// Code generated from Version.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Version
import "github.com/antlr/antlr4/runtime/Go/antlr"

// VersionListener is a complete listener for a parse tree produced by VersionParser.
type VersionListener interface {
	antlr.ParseTreeListener

	// EnterVersionLock is called when entering the versionLock production.
	EnterVersionLock(c *VersionLockContext)

	// EnterCmp is called when entering the cmp production.
	EnterCmp(c *CmpContext)

	// EnterInterval is called when entering the interval production.
	EnterInterval(c *IntervalContext)

	// EnterVersion is called when entering the version production.
	EnterVersion(c *VersionContext)

	// ExitVersionLock is called when exiting the versionLock production.
	ExitVersionLock(c *VersionLockContext)

	// ExitCmp is called when exiting the cmp production.
	ExitCmp(c *CmpContext)

	// ExitInterval is called when exiting the interval production.
	ExitInterval(c *IntervalContext)

	// ExitVersion is called when exiting the version production.
	ExitVersion(c *VersionContext)
}
