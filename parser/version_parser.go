// Code generated from Version.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Version
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 9, 31, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 5, 2, 19, 10, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 2, 2, 6, 2, 4, 6, 8, 2, 3, 3, 2, 3, 6,
	2, 28, 2, 18, 3, 2, 2, 2, 4, 20, 3, 2, 2, 2, 6, 22, 3, 2, 2, 2, 8, 24,
	3, 2, 2, 2, 10, 19, 5, 8, 5, 2, 11, 12, 5, 4, 3, 2, 12, 13, 5, 8, 5, 2,
	13, 19, 3, 2, 2, 2, 14, 15, 5, 8, 5, 2, 15, 16, 5, 6, 4, 2, 16, 17, 5,
	8, 5, 2, 17, 19, 3, 2, 2, 2, 18, 10, 3, 2, 2, 2, 18, 11, 3, 2, 2, 2, 18,
	14, 3, 2, 2, 2, 19, 3, 3, 2, 2, 2, 20, 21, 9, 2, 2, 2, 21, 5, 3, 2, 2,
	2, 22, 23, 7, 7, 2, 2, 23, 7, 3, 2, 2, 2, 24, 25, 7, 9, 2, 2, 25, 26, 7,
	8, 2, 2, 26, 27, 7, 9, 2, 2, 27, 28, 7, 8, 2, 2, 28, 29, 7, 9, 2, 2, 29,
	9, 3, 2, 2, 2, 3, 18,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'<'", "'>'", "'>='", "'<='", "'~'", "'.'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "Number",
}

var ruleNames = []string{
	"versionLock", "cmp", "interval", "version",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type VersionParser struct {
	*antlr.BaseParser
}

func NewVersionParser(input antlr.TokenStream) *VersionParser {
	this := new(VersionParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Version.g4"

	return this
}

// VersionParser tokens.
const (
	VersionParserEOF    = antlr.TokenEOF
	VersionParserT__0   = 1
	VersionParserT__1   = 2
	VersionParserT__2   = 3
	VersionParserT__3   = 4
	VersionParserT__4   = 5
	VersionParserT__5   = 6
	VersionParserNumber = 7
)

// VersionParser rules.
const (
	VersionParserRULE_versionLock = 0
	VersionParserRULE_cmp         = 1
	VersionParserRULE_interval    = 2
	VersionParserRULE_version     = 3
)

// IVersionLockContext is an interface to support dynamic dispatch.
type IVersionLockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVersionLockContext differentiates from other interfaces.
	IsVersionLockContext()
}

type VersionLockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVersionLockContext() *VersionLockContext {
	var p = new(VersionLockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VersionParserRULE_versionLock
	return p
}

func (*VersionLockContext) IsVersionLockContext() {}

func NewVersionLockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VersionLockContext {
	var p = new(VersionLockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VersionParserRULE_versionLock

	return p
}

func (s *VersionLockContext) GetParser() antlr.Parser { return s.parser }

func (s *VersionLockContext) AllVersion() []IVersionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVersionContext)(nil)).Elem())
	var tst = make([]IVersionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVersionContext)
		}
	}

	return tst
}

func (s *VersionLockContext) Version(i int) IVersionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVersionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVersionContext)
}

func (s *VersionLockContext) Cmp() ICmpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICmpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICmpContext)
}

func (s *VersionLockContext) Interval() IIntervalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntervalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntervalContext)
}

func (s *VersionLockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VersionLockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VersionLockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VersionListener); ok {
		listenerT.EnterVersionLock(s)
	}
}

func (s *VersionLockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VersionListener); ok {
		listenerT.ExitVersionLock(s)
	}
}

func (p *VersionParser) VersionLock() (localctx IVersionLockContext) {
	localctx = NewVersionLockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, VersionParserRULE_versionLock)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(16)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(8)
			p.Version()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(9)
			p.Cmp()
		}
		{
			p.SetState(10)
			p.Version()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(12)
			p.Version()
		}
		{
			p.SetState(13)
			p.Interval()
		}
		{
			p.SetState(14)
			p.Version()
		}

	}

	return localctx
}

// ICmpContext is an interface to support dynamic dispatch.
type ICmpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCmpContext differentiates from other interfaces.
	IsCmpContext()
}

type CmpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmpContext() *CmpContext {
	var p = new(CmpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VersionParserRULE_cmp
	return p
}

func (*CmpContext) IsCmpContext() {}

func NewCmpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmpContext {
	var p = new(CmpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VersionParserRULE_cmp

	return p
}

func (s *CmpContext) GetParser() antlr.Parser { return s.parser }
func (s *CmpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VersionListener); ok {
		listenerT.EnterCmp(s)
	}
}

func (s *CmpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VersionListener); ok {
		listenerT.ExitCmp(s)
	}
}

func (p *VersionParser) Cmp() (localctx ICmpContext) {
	localctx = NewCmpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, VersionParserRULE_cmp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(18)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<VersionParserT__0)|(1<<VersionParserT__1)|(1<<VersionParserT__2)|(1<<VersionParserT__3))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IIntervalContext is an interface to support dynamic dispatch.
type IIntervalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntervalContext differentiates from other interfaces.
	IsIntervalContext()
}

type IntervalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntervalContext() *IntervalContext {
	var p = new(IntervalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VersionParserRULE_interval
	return p
}

func (*IntervalContext) IsIntervalContext() {}

func NewIntervalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntervalContext {
	var p = new(IntervalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VersionParserRULE_interval

	return p
}

func (s *IntervalContext) GetParser() antlr.Parser { return s.parser }
func (s *IntervalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntervalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntervalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VersionListener); ok {
		listenerT.EnterInterval(s)
	}
}

func (s *IntervalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VersionListener); ok {
		listenerT.ExitInterval(s)
	}
}

func (p *VersionParser) Interval() (localctx IIntervalContext) {
	localctx = NewIntervalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, VersionParserRULE_interval)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(20)
		p.Match(VersionParserT__4)
	}

	return localctx
}

// IVersionContext is an interface to support dynamic dispatch.
type IVersionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVersionContext differentiates from other interfaces.
	IsVersionContext()
}

type VersionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVersionContext() *VersionContext {
	var p = new(VersionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VersionParserRULE_version
	return p
}

func (*VersionContext) IsVersionContext() {}

func NewVersionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VersionContext {
	var p = new(VersionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VersionParserRULE_version

	return p
}

func (s *VersionContext) GetParser() antlr.Parser { return s.parser }

func (s *VersionContext) AllNumber() []antlr.TerminalNode {
	return s.GetTokens(VersionParserNumber)
}

func (s *VersionContext) Number(i int) antlr.TerminalNode {
	return s.GetToken(VersionParserNumber, i)
}

func (s *VersionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VersionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VersionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VersionListener); ok {
		listenerT.EnterVersion(s)
	}
}

func (s *VersionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VersionListener); ok {
		listenerT.ExitVersion(s)
	}
}

func (p *VersionParser) Version() (localctx IVersionContext) {
	localctx = NewVersionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, VersionParserRULE_version)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(22)
		p.Match(VersionParserNumber)
	}
	{
		p.SetState(23)
		p.Match(VersionParserT__5)
	}
	{
		p.SetState(24)
		p.Match(VersionParserNumber)
	}
	{
		p.SetState(25)
		p.Match(VersionParserT__5)
	}
	{
		p.SetState(26)
		p.Match(VersionParserNumber)
	}

	return localctx
}
