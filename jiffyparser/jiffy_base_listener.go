// Generated from Jiffy.g4 by ANTLR 4.7.

package parser // Jiffy

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseJiffyListener is a complete listener for a parse tree produced by JiffyParser.
type BaseJiffyListener struct{}

var _ JiffyListener = &BaseJiffyListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJiffyListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJiffyListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJiffyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJiffyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseJiffyListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseJiffyListener) ExitRoot(ctx *RootContext) {}

// EnterTimeexpr is called when production timeexpr is entered.
func (s *BaseJiffyListener) EnterTimeexpr(ctx *TimeexprContext) {}

// ExitTimeexpr is called when production timeexpr is exited.
func (s *BaseJiffyListener) ExitTimeexpr(ctx *TimeexprContext) {}

// EnterTimeatom is called when production timeatom is entered.
func (s *BaseJiffyListener) EnterTimeatom(ctx *TimeatomContext) {}

// ExitTimeatom is called when production timeatom is exited.
func (s *BaseJiffyListener) ExitTimeatom(ctx *TimeatomContext) {}

// EnterMicrosecond is called when production microsecond is entered.
func (s *BaseJiffyListener) EnterMicrosecond(ctx *MicrosecondContext) {}

// ExitMicrosecond is called when production microsecond is exited.
func (s *BaseJiffyListener) ExitMicrosecond(ctx *MicrosecondContext) {}

// EnterMillisecond is called when production millisecond is entered.
func (s *BaseJiffyListener) EnterMillisecond(ctx *MillisecondContext) {}

// ExitMillisecond is called when production millisecond is exited.
func (s *BaseJiffyListener) ExitMillisecond(ctx *MillisecondContext) {}

// EnterSecond is called when production second is entered.
func (s *BaseJiffyListener) EnterSecond(ctx *SecondContext) {}

// ExitSecond is called when production second is exited.
func (s *BaseJiffyListener) ExitSecond(ctx *SecondContext) {}

// EnterMinute is called when production minute is entered.
func (s *BaseJiffyListener) EnterMinute(ctx *MinuteContext) {}

// ExitMinute is called when production minute is exited.
func (s *BaseJiffyListener) ExitMinute(ctx *MinuteContext) {}

// EnterHour is called when production hour is entered.
func (s *BaseJiffyListener) EnterHour(ctx *HourContext) {}

// ExitHour is called when production hour is exited.
func (s *BaseJiffyListener) ExitHour(ctx *HourContext) {}

// EnterDay is called when production day is entered.
func (s *BaseJiffyListener) EnterDay(ctx *DayContext) {}

// ExitDay is called when production day is exited.
func (s *BaseJiffyListener) ExitDay(ctx *DayContext) {}

// EnterMonth is called when production month is entered.
func (s *BaseJiffyListener) EnterMonth(ctx *MonthContext) {}

// ExitMonth is called when production month is exited.
func (s *BaseJiffyListener) ExitMonth(ctx *MonthContext) {}

// EnterYear is called when production year is entered.
func (s *BaseJiffyListener) EnterYear(ctx *YearContext) {}

// ExitYear is called when production year is exited.
func (s *BaseJiffyListener) ExitYear(ctx *YearContext) {}

// EnterDecimalLiteral is called when production decimalLiteral is entered.
func (s *BaseJiffyListener) EnterDecimalLiteral(ctx *DecimalLiteralContext) {}

// ExitDecimalLiteral is called when production decimalLiteral is exited.
func (s *BaseJiffyListener) ExitDecimalLiteral(ctx *DecimalLiteralContext) {}

// EnterRealLiteral is called when production realLiteral is entered.
func (s *BaseJiffyListener) EnterRealLiteral(ctx *RealLiteralContext) {}

// ExitRealLiteral is called when production realLiteral is exited.
func (s *BaseJiffyListener) ExitRealLiteral(ctx *RealLiteralContext) {}
