// Generated from Jiffy.g4 by ANTLR 4.7.

package parser // Jiffy

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 33, 123,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 7, 3, 34, 10, 3, 12,
	3, 14, 3, 37, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5,
	4, 47, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 55, 10, 5, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 63, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 5, 7, 71, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8,
	79, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 87, 10, 9, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 95, 10, 10, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 5, 11, 103, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 5, 12, 111, 10, 12, 3, 13, 5, 13, 114, 10, 13, 3, 13, 3,
	13, 3, 14, 5, 14, 119, 10, 14, 3, 14, 3, 14, 3, 14, 2, 2, 15, 2, 4, 6,
	8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 2, 10, 3, 2, 9, 10, 3, 2, 11, 13,
	4, 2, 6, 6, 14, 16, 4, 2, 5, 5, 17, 19, 4, 2, 4, 4, 20, 22, 4, 2, 3, 3,
	23, 24, 3, 2, 25, 27, 4, 2, 7, 7, 28, 29, 2, 127, 2, 28, 3, 2, 2, 2, 4,
	31, 3, 2, 2, 2, 6, 46, 3, 2, 2, 2, 8, 54, 3, 2, 2, 2, 10, 62, 3, 2, 2,
	2, 12, 70, 3, 2, 2, 2, 14, 78, 3, 2, 2, 2, 16, 86, 3, 2, 2, 2, 18, 94,
	3, 2, 2, 2, 20, 102, 3, 2, 2, 2, 22, 110, 3, 2, 2, 2, 24, 113, 3, 2, 2,
	2, 26, 118, 3, 2, 2, 2, 28, 29, 5, 4, 3, 2, 29, 30, 7, 2, 2, 3, 30, 3,
	3, 2, 2, 2, 31, 35, 5, 6, 4, 2, 32, 34, 5, 6, 4, 2, 33, 32, 3, 2, 2, 2,
	34, 37, 3, 2, 2, 2, 35, 33, 3, 2, 2, 2, 35, 36, 3, 2, 2, 2, 36, 5, 3, 2,
	2, 2, 37, 35, 3, 2, 2, 2, 38, 47, 5, 8, 5, 2, 39, 47, 5, 10, 6, 2, 40,
	47, 5, 12, 7, 2, 41, 47, 5, 14, 8, 2, 42, 47, 5, 16, 9, 2, 43, 47, 5, 18,
	10, 2, 44, 47, 5, 20, 11, 2, 45, 47, 5, 22, 12, 2, 46, 38, 3, 2, 2, 2,
	46, 39, 3, 2, 2, 2, 46, 40, 3, 2, 2, 2, 46, 41, 3, 2, 2, 2, 46, 42, 3,
	2, 2, 2, 46, 43, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 46, 45, 3, 2, 2, 2, 47,
	7, 3, 2, 2, 2, 48, 49, 5, 24, 13, 2, 49, 50, 9, 2, 2, 2, 50, 55, 3, 2,
	2, 2, 51, 52, 5, 26, 14, 2, 52, 53, 9, 2, 2, 2, 53, 55, 3, 2, 2, 2, 54,
	48, 3, 2, 2, 2, 54, 51, 3, 2, 2, 2, 55, 9, 3, 2, 2, 2, 56, 57, 5, 24, 13,
	2, 57, 58, 9, 3, 2, 2, 58, 63, 3, 2, 2, 2, 59, 60, 5, 26, 14, 2, 60, 61,
	9, 3, 2, 2, 61, 63, 3, 2, 2, 2, 62, 56, 3, 2, 2, 2, 62, 59, 3, 2, 2, 2,
	63, 11, 3, 2, 2, 2, 64, 65, 5, 24, 13, 2, 65, 66, 9, 4, 2, 2, 66, 71, 3,
	2, 2, 2, 67, 68, 5, 26, 14, 2, 68, 69, 9, 4, 2, 2, 69, 71, 3, 2, 2, 2,
	70, 64, 3, 2, 2, 2, 70, 67, 3, 2, 2, 2, 71, 13, 3, 2, 2, 2, 72, 73, 5,
	24, 13, 2, 73, 74, 9, 5, 2, 2, 74, 79, 3, 2, 2, 2, 75, 76, 5, 26, 14, 2,
	76, 77, 9, 5, 2, 2, 77, 79, 3, 2, 2, 2, 78, 72, 3, 2, 2, 2, 78, 75, 3,
	2, 2, 2, 79, 15, 3, 2, 2, 2, 80, 81, 5, 24, 13, 2, 81, 82, 9, 6, 2, 2,
	82, 87, 3, 2, 2, 2, 83, 84, 5, 26, 14, 2, 84, 85, 9, 6, 2, 2, 85, 87, 3,
	2, 2, 2, 86, 80, 3, 2, 2, 2, 86, 83, 3, 2, 2, 2, 87, 17, 3, 2, 2, 2, 88,
	89, 5, 24, 13, 2, 89, 90, 9, 7, 2, 2, 90, 95, 3, 2, 2, 2, 91, 92, 5, 26,
	14, 2, 92, 93, 9, 7, 2, 2, 93, 95, 3, 2, 2, 2, 94, 88, 3, 2, 2, 2, 94,
	91, 3, 2, 2, 2, 95, 19, 3, 2, 2, 2, 96, 97, 5, 24, 13, 2, 97, 98, 9, 8,
	2, 2, 98, 103, 3, 2, 2, 2, 99, 100, 5, 26, 14, 2, 100, 101, 9, 8, 2, 2,
	101, 103, 3, 2, 2, 2, 102, 96, 3, 2, 2, 2, 102, 99, 3, 2, 2, 2, 103, 21,
	3, 2, 2, 2, 104, 105, 5, 24, 13, 2, 105, 106, 9, 9, 2, 2, 106, 111, 3,
	2, 2, 2, 107, 108, 5, 26, 14, 2, 108, 109, 9, 9, 2, 2, 109, 111, 3, 2,
	2, 2, 110, 104, 3, 2, 2, 2, 110, 107, 3, 2, 2, 2, 111, 23, 3, 2, 2, 2,
	112, 114, 7, 8, 2, 2, 113, 112, 3, 2, 2, 2, 113, 114, 3, 2, 2, 2, 114,
	115, 3, 2, 2, 2, 115, 116, 7, 31, 2, 2, 116, 25, 3, 2, 2, 2, 117, 119,
	7, 8, 2, 2, 118, 117, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119, 120, 3, 2,
	2, 2, 120, 121, 7, 32, 2, 2, 121, 27, 3, 2, 2, 2, 14, 35, 46, 54, 62, 70,
	78, 86, 94, 102, 110, 113, 118,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "", "'-'",
}
var symbolicNames = []string{
	"", "D", "H", "M", "S", "Y", "MINUS", "MICROSECOND", "MICROSECONDS", "MS",
	"MILLISECOND", "MILLISECONDS", "SEC", "SECOND", "SECONDS", "MIN", "MINUTE",
	"MINUTES", "HRS", "HOUR", "HOURS", "DAY", "DAYS", "MON", "MONTH", "MONTHS",
	"YEAR", "YEARS", "AND", "DECIMAL_LITERAL", "REAL_LITERAL", "SPACE",
}

var ruleNames = []string{
	"root", "timeexpr", "timeatom", "microsecond", "millisecond", "second",
	"minute", "hour", "day", "month", "year", "decimalLiteral", "realLiteral",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type JiffyParser struct {
	*antlr.BaseParser
}

func NewJiffyParser(input antlr.TokenStream) *JiffyParser {
	this := new(JiffyParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Jiffy.g4"

	return this
}

// JiffyParser tokens.
const (
	JiffyParserEOF             = antlr.TokenEOF
	JiffyParserD               = 1
	JiffyParserH               = 2
	JiffyParserM               = 3
	JiffyParserS               = 4
	JiffyParserY               = 5
	JiffyParserMINUS           = 6
	JiffyParserMICROSECOND     = 7
	JiffyParserMICROSECONDS    = 8
	JiffyParserMS              = 9
	JiffyParserMILLISECOND     = 10
	JiffyParserMILLISECONDS    = 11
	JiffyParserSEC             = 12
	JiffyParserSECOND          = 13
	JiffyParserSECONDS         = 14
	JiffyParserMIN             = 15
	JiffyParserMINUTE          = 16
	JiffyParserMINUTES         = 17
	JiffyParserHRS             = 18
	JiffyParserHOUR            = 19
	JiffyParserHOURS           = 20
	JiffyParserDAY             = 21
	JiffyParserDAYS            = 22
	JiffyParserMON             = 23
	JiffyParserMONTH           = 24
	JiffyParserMONTHS          = 25
	JiffyParserYEAR            = 26
	JiffyParserYEARS           = 27
	JiffyParserAND             = 28
	JiffyParserDECIMAL_LITERAL = 29
	JiffyParserREAL_LITERAL    = 30
	JiffyParserSPACE           = 31
)

// JiffyParser rules.
const (
	JiffyParserRULE_root           = 0
	JiffyParserRULE_timeexpr       = 1
	JiffyParserRULE_timeatom       = 2
	JiffyParserRULE_microsecond    = 3
	JiffyParserRULE_millisecond    = 4
	JiffyParserRULE_second         = 5
	JiffyParserRULE_minute         = 6
	JiffyParserRULE_hour           = 7
	JiffyParserRULE_day            = 8
	JiffyParserRULE_month          = 9
	JiffyParserRULE_year           = 10
	JiffyParserRULE_decimalLiteral = 11
	JiffyParserRULE_realLiteral    = 12
)

// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRootContext differentiates from other interfaces.
	IsRootContext()
}

type RootContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootContext() *RootContext {
	var p = new(RootContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JiffyParserRULE_root
	return p
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JiffyParserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }

func (s *RootContext) Timeexpr() ITimeexprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeexprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeexprContext)
}

func (s *RootContext) EOF() antlr.TerminalNode {
	return s.GetToken(JiffyParserEOF, 0)
}

func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JiffyListener); ok {
		listenerT.EnterRoot(s)
	}
}

func (s *RootContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JiffyListener); ok {
		listenerT.ExitRoot(s)
	}
}

func (p *JiffyParser) Root() (localctx IRootContext) {
	localctx = NewRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JiffyParserRULE_root)

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
		p.SetState(26)
		p.Timeexpr()
	}
	{
		p.SetState(27)
		p.Match(JiffyParserEOF)
	}

	return localctx
}

// ITimeexprContext is an interface to support dynamic dispatch.
type ITimeexprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeexprContext differentiates from other interfaces.
	IsTimeexprContext()
}

type TimeexprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeexprContext() *TimeexprContext {
	var p = new(TimeexprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JiffyParserRULE_timeexpr
	return p
}

func (*TimeexprContext) IsTimeexprContext() {}

func NewTimeexprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeexprContext {
	var p = new(TimeexprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JiffyParserRULE_timeexpr

	return p
}

func (s *TimeexprContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeexprContext) AllTimeatom() []ITimeatomContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITimeatomContext)(nil)).Elem())
	var tst = make([]ITimeatomContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITimeatomContext)
		}
	}

	return tst
}

func (s *TimeexprContext) Timeatom(i int) ITimeatomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeatomContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITimeatomContext)
}

func (s *TimeexprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeexprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeexprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JiffyListener); ok {
		listenerT.EnterTimeexpr(s)
	}
}

func (s *TimeexprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JiffyListener); ok {
		listenerT.ExitTimeexpr(s)
	}
}

func (p *JiffyParser) Timeexpr() (localctx ITimeexprContext) {
	localctx = NewTimeexprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JiffyParserRULE_timeexpr)
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
		p.SetState(29)
		p.Timeatom()
	}
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JiffyParserMINUS)|(1<<JiffyParserDECIMAL_LITERAL)|(1<<JiffyParserREAL_LITERAL))) != 0 {
		{
			p.SetState(30)
			p.Timeatom()
		}

		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITimeatomContext is an interface to support dynamic dispatch.
type ITimeatomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeatomContext differentiates from other interfaces.
	IsTimeatomContext()
}

type TimeatomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeatomContext() *TimeatomContext {
	var p = new(TimeatomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JiffyParserRULE_timeatom
	return p
}

func (*TimeatomContext) IsTimeatomContext() {}

func NewTimeatomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeatomContext {
	var p = new(TimeatomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JiffyParserRULE_timeatom

	return p
}

func (s *TimeatomContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeatomContext) Microsecond() IMicrosecondContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMicrosecondContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMicrosecondContext)
}

func (s *TimeatomContext) Millisecond() IMillisecondContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMillisecondContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMillisecondContext)
}

func (s *TimeatomContext) Second() ISecondContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISecondContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISecondContext)
}

func (s *TimeatomContext) Minute() IMinuteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMinuteContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMinuteContext)
}

func (s *TimeatomContext) Hour() IHourContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHourContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHourContext)
}

func (s *TimeatomContext) Day() IDayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDayContext)
}

func (s *TimeatomContext) Month() IMonthContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMonthContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMonthContext)
}

func (s *TimeatomContext) Year() IYearContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IYearContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IYearContext)
}

func (s *TimeatomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeatomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeatomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JiffyListener); ok {
		listenerT.EnterTimeatom(s)
	}
}

func (s *TimeatomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JiffyListener); ok {
		listenerT.ExitTimeatom(s)
	}
}

func (p *JiffyParser) Timeatom() (localctx ITimeatomContext) {
	localctx = NewTimeatomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, JiffyParserRULE_timeatom)

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

	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(36)
			p.Microsecond()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(37)
			p.Millisecond()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(38)
			p.Second()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(39)
			p.Minute()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(40)
			p.Hour()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(41)
			p.Day()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(42)
			p.Month()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(43)
			p.Year()
		}

	}

	return localctx
}

// IMicrosecondContext is an interface to support dynamic dispatch.
type IMicrosecondContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMicrosecondContext differentiates from other interfaces.
	IsMicrosecondContext()
}

type MicrosecondContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMicrosecondContext() *MicrosecondContext {
	var p = new(MicrosecondContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JiffyParserRULE_microsecond
	return p
}

func (*MicrosecondContext) IsMicrosecondContext() {}

func NewMicrosecondContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MicrosecondContext {
	var p = new(MicrosecondContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JiffyParserRULE_microsecond

	return p
}

func (s *MicrosecondContext) GetParser() antlr.Parser { return s.parser }

func (s *MicrosecondContext) DecimalLiteral() IDecimalLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecimalLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDecimalLiteralContext)
}

func (s *MicrosecondContext) MICROSECOND() antlr.TerminalNode {
	return s.GetToken(JiffyParserMICROSECOND, 0)
}

func (s *MicrosecondContext) MICROSECONDS() antlr.TerminalNode {
	return s.GetToken(JiffyParserMICROSECONDS, 0)
}

func (s *MicrosecondContext) RealLiteral() IRealLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealLiteralContext)
}

func (s *MicrosecondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MicrosecondContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MicrosecondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JiffyListener); ok {
		listenerT.EnterMicrosecond(s)
	}
}

func (s *MicrosecondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JiffyListener); ok {
		listenerT.ExitMicrosecond(s)
	}
}

func (p *JiffyParser) Microsecond() (localctx IMicrosecondContext) {
	localctx = NewMicrosecondContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, JiffyParserRULE_microsecond)
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

	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(46)
			p.DecimalLiteral()
		}
		p.SetState(47)
		_la = p.GetTokenStream().LA(1)

		if !(_la == JiffyParserMICROSECOND || _la == JiffyParserMICROSECONDS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(49)
			p.RealLiteral()
		}
		p.SetState(50)
		_la = p.GetTokenStream().LA(1)

		if !(_la == JiffyParserMICROSECOND || _la == JiffyParserMICROSECONDS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	}

	return localctx
}

// IMillisecondContext is an interface to support dynamic dispatch.
type IMillisecondContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMillisecondContext differentiates from other interfaces.
	IsMillisecondContext()
}

type MillisecondContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMillisecondContext() *MillisecondContext {
	var p = new(MillisecondContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JiffyParserRULE_millisecond
	return p
}

func (*MillisecondContext) IsMillisecondContext() {}

func NewMillisecondContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MillisecondContext {
	var p = new(MillisecondContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JiffyParserRULE_millisecond

	return p
}

func (s *MillisecondContext) GetParser() antlr.Parser { return s.parser }

func (s *MillisecondContext) DecimalLiteral() IDecimalLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecimalLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDecimalLiteralContext)
}

func (s *MillisecondContext) MS() antlr.TerminalNode {
	return s.GetToken(JiffyParserMS, 0)
}

func (s *MillisecondContext) MILLISECOND() antlr.TerminalNode {
	return s.GetToken(JiffyParserMILLISECOND, 0)
}

func (s *MillisecondContext) MILLISECONDS() antlr.TerminalNode {
	return s.GetToken(JiffyParserMILLISECONDS, 0)
}

func (s *MillisecondContext) RealLiteral() IRealLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealLiteralContext)
}

func (s *MillisecondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MillisecondContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MillisecondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JiffyListener); ok {
		listenerT.EnterMillisecond(s)
	}
}

func (s *MillisecondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JiffyListener); ok {
		listenerT.ExitMillisecond(s)
	}
}

func (p *JiffyParser) Millisecond() (localctx IMillisecondContext) {
	localctx = NewMillisecondContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, JiffyParserRULE_millisecond)
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

	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(54)
			p.DecimalLiteral()
		}
		p.SetState(55)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JiffyParserMS)|(1<<JiffyParserMILLISECOND)|(1<<JiffyParserMILLISECONDS))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(57)
			p.RealLiteral()
		}
		p.SetState(58)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JiffyParserMS)|(1<<JiffyParserMILLISECOND)|(1<<JiffyParserMILLISECONDS))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	}

	return localctx
}

// ISecondContext is an interface to support dynamic dispatch.
type ISecondContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSecondContext differentiates from other interfaces.
	IsSecondContext()
}

type SecondContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySecondContext() *SecondContext {
	var p = new(SecondContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JiffyParserRULE_second
	return p
}

func (*SecondContext) IsSecondContext() {}

func NewSecondContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SecondContext {
	var p = new(SecondContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JiffyParserRULE_second

	return p
}

func (s *SecondContext) GetParser() antlr.Parser { return s.parser }

func (s *SecondContext) DecimalLiteral() IDecimalLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecimalLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDecimalLiteralContext)
}

func (s *SecondContext) S() antlr.TerminalNode {
	return s.GetToken(JiffyParserS, 0)
}

func (s *SecondContext) SEC() antlr.TerminalNode {
	return s.GetToken(JiffyParserSEC, 0)
}

func (s *SecondContext) SECOND() antlr.TerminalNode {
	return s.GetToken(JiffyParserSECOND, 0)
}

func (s *SecondContext) SECONDS() antlr.TerminalNode {
	return s.GetToken(JiffyParserSECONDS, 0)
}

func (s *SecondContext) RealLiteral() IRealLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealLiteralContext)
}

func (s *SecondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SecondContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SecondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JiffyListener); ok {
		listenerT.EnterSecond(s)
	}
}

func (s *SecondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JiffyListener); ok {
		listenerT.ExitSecond(s)
	}
}

func (p *JiffyParser) Second() (localctx ISecondContext) {
	localctx = NewSecondContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, JiffyParserRULE_second)
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

	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(62)
			p.DecimalLiteral()
		}
		p.SetState(63)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JiffyParserS)|(1<<JiffyParserSEC)|(1<<JiffyParserSECOND)|(1<<JiffyParserSECONDS))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(65)
			p.RealLiteral()
		}
		p.SetState(66)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JiffyParserS)|(1<<JiffyParserSEC)|(1<<JiffyParserSECOND)|(1<<JiffyParserSECONDS))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	}

	return localctx
}

// IMinuteContext is an interface to support dynamic dispatch.
type IMinuteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMinuteContext differentiates from other interfaces.
	IsMinuteContext()
}

type MinuteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMinuteContext() *MinuteContext {
	var p = new(MinuteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JiffyParserRULE_minute
	return p
}

func (*MinuteContext) IsMinuteContext() {}

func NewMinuteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MinuteContext {
	var p = new(MinuteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JiffyParserRULE_minute

	return p
}

func (s *MinuteContext) GetParser() antlr.Parser { return s.parser }

func (s *MinuteContext) DecimalLiteral() IDecimalLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecimalLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDecimalLiteralContext)
}

func (s *MinuteContext) M() antlr.TerminalNode {
	return s.GetToken(JiffyParserM, 0)
}

func (s *MinuteContext) MIN() antlr.TerminalNode {
	return s.GetToken(JiffyParserMIN, 0)
}

func (s *MinuteContext) MINUTE() antlr.TerminalNode {
	return s.GetToken(JiffyParserMINUTE, 0)
}

func (s *MinuteContext) MINUTES() antlr.TerminalNode {
	return s.GetToken(JiffyParserMINUTES, 0)
}

func (s *MinuteContext) RealLiteral() IRealLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealLiteralContext)
}

func (s *MinuteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MinuteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MinuteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JiffyListener); ok {
		listenerT.EnterMinute(s)
	}
}

func (s *MinuteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JiffyListener); ok {
		listenerT.ExitMinute(s)
	}
}

func (p *JiffyParser) Minute() (localctx IMinuteContext) {
	localctx = NewMinuteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, JiffyParserRULE_minute)
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

	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(70)
			p.DecimalLiteral()
		}
		p.SetState(71)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JiffyParserM)|(1<<JiffyParserMIN)|(1<<JiffyParserMINUTE)|(1<<JiffyParserMINUTES))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(73)
			p.RealLiteral()
		}
		p.SetState(74)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JiffyParserM)|(1<<JiffyParserMIN)|(1<<JiffyParserMINUTE)|(1<<JiffyParserMINUTES))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	}

	return localctx
}

// IHourContext is an interface to support dynamic dispatch.
type IHourContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHourContext differentiates from other interfaces.
	IsHourContext()
}

type HourContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHourContext() *HourContext {
	var p = new(HourContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JiffyParserRULE_hour
	return p
}

func (*HourContext) IsHourContext() {}

func NewHourContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HourContext {
	var p = new(HourContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JiffyParserRULE_hour

	return p
}

func (s *HourContext) GetParser() antlr.Parser { return s.parser }

func (s *HourContext) DecimalLiteral() IDecimalLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecimalLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDecimalLiteralContext)
}

func (s *HourContext) H() antlr.TerminalNode {
	return s.GetToken(JiffyParserH, 0)
}

func (s *HourContext) HRS() antlr.TerminalNode {
	return s.GetToken(JiffyParserHRS, 0)
}

func (s *HourContext) HOUR() antlr.TerminalNode {
	return s.GetToken(JiffyParserHOUR, 0)
}

func (s *HourContext) HOURS() antlr.TerminalNode {
	return s.GetToken(JiffyParserHOURS, 0)
}

func (s *HourContext) RealLiteral() IRealLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealLiteralContext)
}

func (s *HourContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HourContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HourContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JiffyListener); ok {
		listenerT.EnterHour(s)
	}
}

func (s *HourContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JiffyListener); ok {
		listenerT.ExitHour(s)
	}
}

func (p *JiffyParser) Hour() (localctx IHourContext) {
	localctx = NewHourContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, JiffyParserRULE_hour)
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

	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(78)
			p.DecimalLiteral()
		}
		p.SetState(79)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JiffyParserH)|(1<<JiffyParserHRS)|(1<<JiffyParserHOUR)|(1<<JiffyParserHOURS))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(81)
			p.RealLiteral()
		}
		p.SetState(82)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JiffyParserH)|(1<<JiffyParserHRS)|(1<<JiffyParserHOUR)|(1<<JiffyParserHOURS))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	}

	return localctx
}

// IDayContext is an interface to support dynamic dispatch.
type IDayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDayContext differentiates from other interfaces.
	IsDayContext()
}

type DayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDayContext() *DayContext {
	var p = new(DayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JiffyParserRULE_day
	return p
}

func (*DayContext) IsDayContext() {}

func NewDayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DayContext {
	var p = new(DayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JiffyParserRULE_day

	return p
}

func (s *DayContext) GetParser() antlr.Parser { return s.parser }

func (s *DayContext) DecimalLiteral() IDecimalLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecimalLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDecimalLiteralContext)
}

func (s *DayContext) D() antlr.TerminalNode {
	return s.GetToken(JiffyParserD, 0)
}

func (s *DayContext) DAY() antlr.TerminalNode {
	return s.GetToken(JiffyParserDAY, 0)
}

func (s *DayContext) DAYS() antlr.TerminalNode {
	return s.GetToken(JiffyParserDAYS, 0)
}

func (s *DayContext) RealLiteral() IRealLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealLiteralContext)
}

func (s *DayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JiffyListener); ok {
		listenerT.EnterDay(s)
	}
}

func (s *DayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JiffyListener); ok {
		listenerT.ExitDay(s)
	}
}

func (p *JiffyParser) Day() (localctx IDayContext) {
	localctx = NewDayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, JiffyParserRULE_day)
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

	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(86)
			p.DecimalLiteral()
		}
		p.SetState(87)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JiffyParserD)|(1<<JiffyParserDAY)|(1<<JiffyParserDAYS))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)
			p.RealLiteral()
		}
		p.SetState(90)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JiffyParserD)|(1<<JiffyParserDAY)|(1<<JiffyParserDAYS))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	}

	return localctx
}

// IMonthContext is an interface to support dynamic dispatch.
type IMonthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMonthContext differentiates from other interfaces.
	IsMonthContext()
}

type MonthContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMonthContext() *MonthContext {
	var p = new(MonthContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JiffyParserRULE_month
	return p
}

func (*MonthContext) IsMonthContext() {}

func NewMonthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MonthContext {
	var p = new(MonthContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JiffyParserRULE_month

	return p
}

func (s *MonthContext) GetParser() antlr.Parser { return s.parser }

func (s *MonthContext) DecimalLiteral() IDecimalLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecimalLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDecimalLiteralContext)
}

func (s *MonthContext) MON() antlr.TerminalNode {
	return s.GetToken(JiffyParserMON, 0)
}

func (s *MonthContext) MONTH() antlr.TerminalNode {
	return s.GetToken(JiffyParserMONTH, 0)
}

func (s *MonthContext) MONTHS() antlr.TerminalNode {
	return s.GetToken(JiffyParserMONTHS, 0)
}

func (s *MonthContext) RealLiteral() IRealLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealLiteralContext)
}

func (s *MonthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MonthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MonthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JiffyListener); ok {
		listenerT.EnterMonth(s)
	}
}

func (s *MonthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JiffyListener); ok {
		listenerT.ExitMonth(s)
	}
}

func (p *JiffyParser) Month() (localctx IMonthContext) {
	localctx = NewMonthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, JiffyParserRULE_month)
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

	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(94)
			p.DecimalLiteral()
		}
		p.SetState(95)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JiffyParserMON)|(1<<JiffyParserMONTH)|(1<<JiffyParserMONTHS))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(97)
			p.RealLiteral()
		}
		p.SetState(98)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JiffyParserMON)|(1<<JiffyParserMONTH)|(1<<JiffyParserMONTHS))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	}

	return localctx
}

// IYearContext is an interface to support dynamic dispatch.
type IYearContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsYearContext differentiates from other interfaces.
	IsYearContext()
}

type YearContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyYearContext() *YearContext {
	var p = new(YearContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JiffyParserRULE_year
	return p
}

func (*YearContext) IsYearContext() {}

func NewYearContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *YearContext {
	var p = new(YearContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JiffyParserRULE_year

	return p
}

func (s *YearContext) GetParser() antlr.Parser { return s.parser }

func (s *YearContext) DecimalLiteral() IDecimalLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecimalLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDecimalLiteralContext)
}

func (s *YearContext) Y() antlr.TerminalNode {
	return s.GetToken(JiffyParserY, 0)
}

func (s *YearContext) YEAR() antlr.TerminalNode {
	return s.GetToken(JiffyParserYEAR, 0)
}

func (s *YearContext) YEARS() antlr.TerminalNode {
	return s.GetToken(JiffyParserYEARS, 0)
}

func (s *YearContext) RealLiteral() IRealLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealLiteralContext)
}

func (s *YearContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *YearContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *YearContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JiffyListener); ok {
		listenerT.EnterYear(s)
	}
}

func (s *YearContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JiffyListener); ok {
		listenerT.ExitYear(s)
	}
}

func (p *JiffyParser) Year() (localctx IYearContext) {
	localctx = NewYearContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, JiffyParserRULE_year)
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

	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(102)
			p.DecimalLiteral()
		}
		p.SetState(103)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JiffyParserY)|(1<<JiffyParserYEAR)|(1<<JiffyParserYEARS))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(105)
			p.RealLiteral()
		}
		p.SetState(106)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JiffyParserY)|(1<<JiffyParserYEAR)|(1<<JiffyParserYEARS))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	}

	return localctx
}

// IDecimalLiteralContext is an interface to support dynamic dispatch.
type IDecimalLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDecimalLiteralContext differentiates from other interfaces.
	IsDecimalLiteralContext()
}

type DecimalLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecimalLiteralContext() *DecimalLiteralContext {
	var p = new(DecimalLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JiffyParserRULE_decimalLiteral
	return p
}

func (*DecimalLiteralContext) IsDecimalLiteralContext() {}

func NewDecimalLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecimalLiteralContext {
	var p = new(DecimalLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JiffyParserRULE_decimalLiteral

	return p
}

func (s *DecimalLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *DecimalLiteralContext) DECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(JiffyParserDECIMAL_LITERAL, 0)
}

func (s *DecimalLiteralContext) MINUS() antlr.TerminalNode {
	return s.GetToken(JiffyParserMINUS, 0)
}

func (s *DecimalLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecimalLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecimalLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JiffyListener); ok {
		listenerT.EnterDecimalLiteral(s)
	}
}

func (s *DecimalLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JiffyListener); ok {
		listenerT.ExitDecimalLiteral(s)
	}
}

func (p *JiffyParser) DecimalLiteral() (localctx IDecimalLiteralContext) {
	localctx = NewDecimalLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, JiffyParserRULE_decimalLiteral)
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
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JiffyParserMINUS {
		{
			p.SetState(110)
			p.Match(JiffyParserMINUS)
		}

	}
	{
		p.SetState(113)
		p.Match(JiffyParserDECIMAL_LITERAL)
	}

	return localctx
}

// IRealLiteralContext is an interface to support dynamic dispatch.
type IRealLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRealLiteralContext differentiates from other interfaces.
	IsRealLiteralContext()
}

type RealLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRealLiteralContext() *RealLiteralContext {
	var p = new(RealLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JiffyParserRULE_realLiteral
	return p
}

func (*RealLiteralContext) IsRealLiteralContext() {}

func NewRealLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RealLiteralContext {
	var p = new(RealLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JiffyParserRULE_realLiteral

	return p
}

func (s *RealLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *RealLiteralContext) REAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(JiffyParserREAL_LITERAL, 0)
}

func (s *RealLiteralContext) MINUS() antlr.TerminalNode {
	return s.GetToken(JiffyParserMINUS, 0)
}

func (s *RealLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RealLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RealLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JiffyListener); ok {
		listenerT.EnterRealLiteral(s)
	}
}

func (s *RealLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JiffyListener); ok {
		listenerT.ExitRealLiteral(s)
	}
}

func (p *JiffyParser) RealLiteral() (localctx IRealLiteralContext) {
	localctx = NewRealLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, JiffyParserRULE_realLiteral)
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
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JiffyParserMINUS {
		{
			p.SetState(115)
			p.Match(JiffyParserMINUS)
		}

	}
	{
		p.SetState(118)
		p.Match(JiffyParserREAL_LITERAL)
	}

	return localctx
}
