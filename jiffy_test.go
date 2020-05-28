package jiffy

import (
	"testing"
	"time"
)

type TestDataDec struct {
	desc string
	unit time.Duration
}

func TestDescribe(t *testing.T) {
	desc := DescribeDuration(12*time.Second+13*time.Minute+14*time.Hour, NewWant())
	expect := "14h 13m 12s"
	if desc != expect {
		t.Errorf("Expect \"%s\" but \"%s\"", expect, desc)
	}

	desc = DescribeDuration(12*time.Second+13*time.Minute+14*time.Hour, &Want{Hour: true, Minute: true, Second: true, Verbose: true, Separator: " "})
	expect = "14 hours 13 minutes 12 seconds"
	if desc != expect {
		t.Errorf("Expect \"%s\" but \"%s\"", expect, desc)
	}
}

func TestDescribe2(t *testing.T) {
	desc := DescribeDuration(14*time.Hour, &Want{Hour: true, Verbose: false})
	expect := "14h"
	if desc != expect {
		t.Errorf("Expect \"%s\" but \"%s\"", expect, desc)
	}
}

func TestDescribe3(t *testing.T) {
	desc := DescribeDuration(14*time.Hour, &Want{Minute: true, Verbose: false})
	expect := "840m"
	if desc != expect {
		t.Errorf("Expect \"%s\" but \"%s\"", expect, desc)
	}
}

func TestDescribe4(t *testing.T) {
	desc := DescribeDuration(14*time.Hour, &Want{Second: true, Verbose: false})
	expect := "50400s"
	if desc != expect {
		t.Errorf("Expect \"%s\" but \"%s\"", expect, desc)
	}
}

func TestDescribe5(t *testing.T) {
	dur := 3*YearDuraton + 4*MonthDuration + 13*DayDuration + 14*time.Hour + 100*time.Microsecond
	desc := DescribeDuration(dur, NewWant())
	expect := "3y 4mon 13d 14h"
	if desc != expect {
		t.Errorf("Expect \"%s\" but \"%s\"", expect, desc)
	}
}

func TestDurationOf(t *testing.T) {
	tests := []*TestDataDec{
		&TestDataDec{"12 microsecond", 12 * time.Microsecond},
		&TestDataDec{"12 millisecond", 12 * time.Millisecond},
		&TestDataDec{"12 second", 12 * time.Second},
		&TestDataDec{"-12 second", (-12) * time.Second},
		&TestDataDec{"12 minute", 12 * time.Minute},
		&TestDataDec{"12 hour", 12 * time.Hour},
		&TestDataDec{"12 day", 12 * DayDuration},
		&TestDataDec{"12 month", 12 * MonthDuration},
		&TestDataDec{"12 year", 12 * YearDuraton},
		&TestDataDec{"12 s", 12 * time.Second},
		&TestDataDec{"12 m", 12 * time.Minute},
		&TestDataDec{"12 h", 12 * time.Hour},
		&TestDataDec{"12 d", 12 * DayDuration},
		&TestDataDec{"12 mon", 12 * MonthDuration},
		&TestDataDec{"12y", 12 * YearDuraton},
		&TestDataDec{"12s", 12 * time.Second},
		&TestDataDec{"12m", 12 * time.Minute},
		&TestDataDec{"12h", 12 * time.Hour},
		&TestDataDec{"12d", 12 * DayDuration},
		&TestDataDec{"12mon", 12 * MonthDuration},
		&TestDataDec{"12y", 12 * YearDuraton},
		&TestDataDec{"12.3 microsecond", time.Duration(12.3 * float64(time.Microsecond))},
		&TestDataDec{"12.3 millisecond", time.Duration(12.3 * float64(time.Millisecond))},
		&TestDataDec{"12.3 second", time.Duration(12.3 * float64(time.Second))},
		&TestDataDec{"12.3 minute", time.Duration(12.3 * float64(time.Minute))},
		&TestDataDec{"12.3 hour", time.Duration(12.3 * float64(time.Hour))},
		&TestDataDec{"12.3 day", time.Duration(12.3 * float64(DayDuration))},
		&TestDataDec{"12.3 month", time.Duration(12.3 * float64(MonthDuration))},
		&TestDataDec{"12.3 year", time.Duration(12.3 * float64(YearDuraton))},

		&TestDataDec{"12 microseconds", 12 * time.Microsecond},
		&TestDataDec{"12 milliseconds", 12 * time.Millisecond},
		&TestDataDec{"12 seconds", 12 * time.Second},
		&TestDataDec{"12 minutes", 12 * time.Minute},
		&TestDataDec{"12 hours", 12 * time.Hour},
		&TestDataDec{"12 days", 12 * DayDuration},
		&TestDataDec{"12 months", 12 * MonthDuration},
		&TestDataDec{"12 years", 12 * YearDuraton},
		&TestDataDec{"12.3 microseconds", time.Duration(12.3 * float64(time.Microsecond))},
		&TestDataDec{"12.3 milliseconds", time.Duration(12.3 * float64(time.Millisecond))},
		&TestDataDec{"12.3 seconds", time.Duration(12.3 * float64(time.Second))},
		&TestDataDec{"12.3 minutes", time.Duration(12.3 * float64(time.Minute))},
		&TestDataDec{"12.3 hours", time.Duration(12.3 * float64(time.Hour))},
		&TestDataDec{"12.3 days", time.Duration(12.3 * float64(DayDuration))},
		&TestDataDec{"12.3 months", time.Duration(12.3 * float64(MonthDuration))},
		&TestDataDec{"12.3 years", time.Duration(12.3 * float64(YearDuraton))},

		&TestDataDec{"12 millisecond 4 microsecond", (12 * time.Millisecond) + (4 * time.Microsecond)},
		&TestDataDec{"12 second 4 millisecond", (12 * time.Second) + (4 * time.Millisecond)},
		&TestDataDec{"12 minute 4 second", (12 * time.Minute) + (4 * time.Second)},
		&TestDataDec{"12 hour 4 minute", (12 * time.Hour) + (4 * time.Minute)},
		&TestDataDec{"12 day 4 hour", (12 * DayDuration) + (4 * time.Hour)},
		&TestDataDec{"12 month 4 day", (12 * MonthDuration) + (4 * DayDuration)},
		&TestDataDec{"12 year 4 month", (12 * YearDuraton) + (4 * MonthDuration)},

		&TestDataDec{"12s and 4millisecond", (12 * time.Second) + (4 * time.Millisecond)},
		&TestDataDec{"12m 4s", (12 * time.Minute) + (4 * time.Second)},
		&TestDataDec{"12h4m", (12 * time.Hour) + (4 * time.Minute)},
		&TestDataDec{"12d 4h", (12 * DayDuration) + (4 * time.Hour)},
		&TestDataDec{"12mon, 4d", (12 * MonthDuration) + (4 * DayDuration)},
		&TestDataDec{"12y 4mon", (12 * YearDuraton) + (4 * MonthDuration)},
	}

	for _, td := range tests {
		dur, err := DurationOf(td.desc)
		if err != nil {
			t.Logf("failed parsing %s, got %v", td.desc, err)
			t.Fail()
		} else {
			if dur != td.unit {
				t.Logf("incorrect result betwen %d and %d for %s", td.unit, dur, td.desc)
				t.Fail()
			}
		}
	}
}
