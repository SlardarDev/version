// Code generated from Version.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Version
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseVersionListener is a complete listener for a parse tree produced by VersionParser.
type BaseVersionListener struct{}

var _ VersionListener = &BaseVersionListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVersionListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVersionListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVersionListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVersionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterVersionLock is called when production versionLock is entered.
func (s *BaseVersionListener) EnterVersionLock(ctx *VersionLockContext) {}

// ExitVersionLock is called when production versionLock is exited.
func (s *BaseVersionListener) ExitVersionLock(ctx *VersionLockContext) {}

// EnterCmp is called when production cmp is entered.
func (s *BaseVersionListener) EnterCmp(ctx *CmpContext) {}

// ExitCmp is called when production cmp is exited.
func (s *BaseVersionListener) ExitCmp(ctx *CmpContext) {}

// EnterInterval is called when production interval is entered.
func (s *BaseVersionListener) EnterInterval(ctx *IntervalContext) {}

// ExitInterval is called when production interval is exited.
func (s *BaseVersionListener) ExitInterval(ctx *IntervalContext) {}

// EnterVersion is called when production version is entered.
func (s *BaseVersionListener) EnterVersion(ctx *VersionContext) {}

// ExitVersion is called when production version is exited.
func (s *BaseVersionListener) ExitVersion(ctx *VersionContext) {}
