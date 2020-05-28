// Generated from Jiffy.g4 by ANTLR 4.7.

package parser // Jiffy

import "github.com/antlr/antlr4/runtime/Go/antlr"

// JiffyListener is a complete listener for a parse tree produced by JiffyParser.
type JiffyListener interface {
	antlr.ParseTreeListener

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterTimeexpr is called when entering the timeexpr production.
	EnterTimeexpr(c *TimeexprContext)

	// EnterTimeatom is called when entering the timeatom production.
	EnterTimeatom(c *TimeatomContext)

	// EnterMicrosecond is called when entering the microsecond production.
	EnterMicrosecond(c *MicrosecondContext)

	// EnterMillisecond is called when entering the millisecond production.
	EnterMillisecond(c *MillisecondContext)

	// EnterSecond is called when entering the second production.
	EnterSecond(c *SecondContext)

	// EnterMinute is called when entering the minute production.
	EnterMinute(c *MinuteContext)

	// EnterHour is called when entering the hour production.
	EnterHour(c *HourContext)

	// EnterDay is called when entering the day production.
	EnterDay(c *DayContext)

	// EnterMonth is called when entering the month production.
	EnterMonth(c *MonthContext)

	// EnterYear is called when entering the year production.
	EnterYear(c *YearContext)

	// EnterDecimalLiteral is called when entering the decimalLiteral production.
	EnterDecimalLiteral(c *DecimalLiteralContext)

	// EnterRealLiteral is called when entering the realLiteral production.
	EnterRealLiteral(c *RealLiteralContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitTimeexpr is called when exiting the timeexpr production.
	ExitTimeexpr(c *TimeexprContext)

	// ExitTimeatom is called when exiting the timeatom production.
	ExitTimeatom(c *TimeatomContext)

	// ExitMicrosecond is called when exiting the microsecond production.
	ExitMicrosecond(c *MicrosecondContext)

	// ExitMillisecond is called when exiting the millisecond production.
	ExitMillisecond(c *MillisecondContext)

	// ExitSecond is called when exiting the second production.
	ExitSecond(c *SecondContext)

	// ExitMinute is called when exiting the minute production.
	ExitMinute(c *MinuteContext)

	// ExitHour is called when exiting the hour production.
	ExitHour(c *HourContext)

	// ExitDay is called when exiting the day production.
	ExitDay(c *DayContext)

	// ExitMonth is called when exiting the month production.
	ExitMonth(c *MonthContext)

	// ExitYear is called when exiting the year production.
	ExitYear(c *YearContext)

	// ExitDecimalLiteral is called when exiting the decimalLiteral production.
	ExitDecimalLiteral(c *DecimalLiteralContext)

	// ExitRealLiteral is called when exiting the realLiteral production.
	ExitRealLiteral(c *RealLiteralContext)
}
