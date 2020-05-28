package jiffy

import (
	"bytes"
	"fmt"
	"strconv"
	"time"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/hyperjumptech/jiffy/jiffyparser"
)

const (
	// DayDuration is a duration in one day
	DayDuration = time.Hour * 24
	// MonthDuration is a duration in one month or 31 days
	MonthDuration = DayDuration * 31
	// YearDuraton is a duration in one year
	YearDuraton = MonthDuration * 12
)

// NewWant creates a default Want instance
func NewWant() *Want {
	return &Want{
		Year:      true,
		Month:     true,
		Day:       true,
		Hour:      true,
		Minute:    true,
		Second:    true,
		Verbose:   false,
		Separator: " ",
	}
}

// Want describe the time unit to produce when describing a duration.
type Want struct {
	Year        bool
	Month       bool
	Day         bool
	Hour        bool
	Minute      bool
	Second      bool
	Millisecond bool
	Microsecond bool
	Separator   string
	Verbose     bool
}

// DurationOf try to caculate time duration by the string descripton
func DurationOf(description string) (time.Duration, error) {
	is := antlr.NewInputStream(description)
	lexer := parser.NewJiffyLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	var parseError error

	listener := &AntlrAstListener{
		errorCallBack: func(err error) {
			parseError = err
		},
	}
	psr := parser.NewJiffyParser(stream)
	psr.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(listener, psr.Root())
	if parseError != nil {
		return 0, parseError
	}
	return listener.duration, nil
}

// DescribeDuration will produce a time description out of a duration
func DescribeDuration(duration time.Duration, want *Want) string {
	var yearval, monthval, dayval, hourval, minuteval, secondval, millival, microval int64
	var buff bytes.Buffer
	// buff.WriteString(">")
	left := duration
	if want.Year {
		yearval = int64(left) / int64(YearDuraton)
		left = left % YearDuraton
	}
	if want.Month {
		monthval = int64(left) / int64(MonthDuration)
		left = left % MonthDuration
	}
	if want.Day {
		dayval = int64(left) / int64(DayDuration)
		left = left % DayDuration
	}
	if want.Hour {
		hourval = int64(left) / int64(time.Hour)
		left = left % time.Hour
	}
	if want.Minute {
		minuteval = int64(left) / int64(time.Minute)
		left = left % time.Minute
	}
	if want.Second {
		secondval = int64(left) / int64(time.Second)
		left = left % time.Second
	}
	if want.Millisecond {
		millival = int64(left) / int64(time.Millisecond)
		left = left % time.Millisecond
	}
	if want.Microsecond {
		microval = int64(left) / int64(time.Microsecond)
		left = left % time.Microsecond
	}

	// fmt.Printf("%d %d %d %d %d %d %d %d\n", yearval, monthval, dayval, hourval, minuteval, secondval, millival, microval)

	if yearval > 0 {
		if want.Verbose {
			if yearval == 1 {
				buff.WriteString(fmt.Sprintf("%d year", yearval))
			} else {
				buff.WriteString(fmt.Sprintf("%d years", yearval))
			}
		} else {
			buff.WriteString(fmt.Sprintf("%dy", yearval))
		}
	}
	if monthval > 0 {
		if buff.Len() > 0 {
			buff.WriteString(want.Separator)
		}
		if want.Verbose {
			if monthval == 1 {
				buff.WriteString(fmt.Sprintf("%d month", monthval))
			} else {
				buff.WriteString(fmt.Sprintf("%d months", monthval))
			}
		} else {
			buff.WriteString(fmt.Sprintf("%dmon", monthval))
		}
	}
	if dayval > 0 {
		if buff.Len() > 0 {
			buff.WriteString(want.Separator)
		}
		if want.Verbose {
			if dayval == 1 {
				buff.WriteString(fmt.Sprintf("%d day", dayval))
			} else {
				buff.WriteString(fmt.Sprintf("%d days", dayval))
			}
		} else {
			buff.WriteString(fmt.Sprintf("%dd", dayval))
		}
	}
	if hourval > 0 {
		if buff.Len() > 0 {
			buff.WriteString(want.Separator)
		}
		if want.Verbose {
			if hourval == 1 {
				buff.WriteString(fmt.Sprintf("%d hour", hourval))
			} else {
				buff.WriteString(fmt.Sprintf("%d hours", hourval))
			}
		} else {
			buff.WriteString(fmt.Sprintf("%dh", hourval))
		}
	}
	if minuteval > 0 {
		if buff.Len() > 0 {
			buff.WriteString(want.Separator)
		}
		if want.Verbose {
			if minuteval == 1 {
				buff.WriteString(fmt.Sprintf("%d minute", minuteval))
			} else {
				buff.WriteString(fmt.Sprintf("%d minutes", minuteval))
			}
		} else {
			buff.WriteString(fmt.Sprintf("%dm", minuteval))
		}
	}
	if secondval > 0 {
		if buff.Len() > 0 {
			buff.WriteString(want.Separator)
		}
		if want.Verbose {
			if secondval == 1 {
				buff.WriteString(fmt.Sprintf("%d second", secondval))
			} else {
				buff.WriteString(fmt.Sprintf("%d seconds", secondval))
			}
		} else {
			buff.WriteString(fmt.Sprintf("%ds", secondval))
		}
	}
	if millival > 0 {
		if buff.Len() > 0 {
			buff.WriteString(want.Separator)
		}
		if want.Verbose {
			if millival == 1 {
				buff.WriteString(fmt.Sprintf("%d millisecond", millival))
			} else {
				buff.WriteString(fmt.Sprintf("%d milliseconds", millival))
			}
		} else {
			buff.WriteString(fmt.Sprintf("%dms", millival))
		}
	}
	if microval > 0 {
		if buff.Len() > 0 {
			buff.WriteString(want.Separator)
		}
		if want.Verbose {
			if microval == 1 {
				buff.WriteString(fmt.Sprintf("%d microsecond", microval))
			} else {
				buff.WriteString(fmt.Sprintf("%d microseconds", microval))
			}
		} else {
			buff.WriteString(fmt.Sprintf("%dµs", microval))
		}
	}
	return buff.String()
}

// AntlrAstListener an implementation of the ANTLR4 AST walker
type AntlrAstListener struct {
	parser.BaseJiffyListener
	duration      time.Duration
	errorCallBack func(err error)
	lastNode      string
}

// VisitTerminal is called when a terminal node is visited.
func (s *AntlrAstListener) VisitTerminal(node antlr.TerminalNode) {
	s.lastNode = node.GetText()
}

// VisitErrorNode is called when an error node is visited.
func (s *AntlrAstListener) VisitErrorNode(node antlr.ErrorNode) {
	s.errorCallBack(fmt.Errorf("the time description is not correct ... %s", s.lastNode))
}

// EnterMicrosecond is called when production microsecond is entered.
func (s *AntlrAstListener) EnterMicrosecond(ctx *parser.MicrosecondContext) {
	if ctx.DecimalLiteral() != nil {
		num, _ := strconv.Atoi(ctx.DecimalLiteral().GetText())
		s.duration += time.Duration(num) * time.Microsecond
	} else {
		f, _ := strconv.ParseFloat(ctx.RealLiteral().GetText(), 64)
		s.duration += time.Duration(f * float64(time.Microsecond))
	}
}

// EnterMillisecond is called when production millisecond is entered.
func (s *AntlrAstListener) EnterMillisecond(ctx *parser.MillisecondContext) {
	if ctx.DecimalLiteral() != nil {
		num, _ := strconv.Atoi(ctx.DecimalLiteral().GetText())
		s.duration += time.Duration(num) * time.Millisecond
	} else {
		f, _ := strconv.ParseFloat(ctx.RealLiteral().GetText(), 64)
		s.duration += time.Duration(f * float64(time.Millisecond))
	}
}

// EnterSecond is called when production second is entered.
func (s *AntlrAstListener) EnterSecond(ctx *parser.SecondContext) {
	if ctx.DecimalLiteral() != nil {
		num, _ := strconv.Atoi(ctx.DecimalLiteral().GetText())
		s.duration += time.Duration(num) * time.Second
	} else {
		f, _ := strconv.ParseFloat(ctx.RealLiteral().GetText(), 64)
		s.duration += time.Duration(f * float64(time.Second))
	}
}

// EnterMinute is called when production minute is entered.
func (s *AntlrAstListener) EnterMinute(ctx *parser.MinuteContext) {
	if ctx.DecimalLiteral() != nil {
		num, _ := strconv.Atoi(ctx.DecimalLiteral().GetText())
		s.duration += time.Duration(num) * time.Minute
	} else {
		f, _ := strconv.ParseFloat(ctx.RealLiteral().GetText(), 64)
		s.duration += time.Duration(f * float64(time.Minute))
	}
}

// EnterHour is called when production hour is entered.
func (s *AntlrAstListener) EnterHour(ctx *parser.HourContext) {
	if ctx.DecimalLiteral() != nil {
		num, _ := strconv.Atoi(ctx.DecimalLiteral().GetText())
		s.duration += time.Duration(num) * time.Hour
	} else {
		f, _ := strconv.ParseFloat(ctx.RealLiteral().GetText(), 64)
		s.duration += time.Duration(f * float64(time.Hour))
	}
}

// EnterDay is called when production day is entered.
func (s *AntlrAstListener) EnterDay(ctx *parser.DayContext) {
	if ctx.DecimalLiteral() != nil {
		num, _ := strconv.Atoi(ctx.DecimalLiteral().GetText())
		s.duration += time.Duration(num) * DayDuration
	} else {
		f, _ := strconv.ParseFloat(ctx.RealLiteral().GetText(), 64)
		s.duration += time.Duration(f * float64(DayDuration))
	}
}

// EnterMonth is called when production month is entered.
func (s *AntlrAstListener) EnterMonth(ctx *parser.MonthContext) {
	if ctx.DecimalLiteral() != nil {
		num, _ := strconv.Atoi(ctx.DecimalLiteral().GetText())
		s.duration += time.Duration(num) * MonthDuration
	} else {
		f, _ := strconv.ParseFloat(ctx.RealLiteral().GetText(), 64)
		s.duration += time.Duration(f * float64(MonthDuration))
	}
}

// EnterYear is called when production year is entered.
func (s *AntlrAstListener) EnterYear(ctx *parser.YearContext) {
	if ctx.DecimalLiteral() != nil {
		num, _ := strconv.Atoi(ctx.DecimalLiteral().GetText())
		s.duration += time.Duration(num) * YearDuraton
	} else {
		f, _ := strconv.ParseFloat(ctx.RealLiteral().GetText(), 64)
		s.duration += time.Duration(f * float64(YearDuraton))
	}
}
