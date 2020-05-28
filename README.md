# Jiffy

[![Build Status](https://travis-ci.org/hyperjumptech/jiffy.svg?branch=master)](https://travis-ci.org/hyperjumptech/jiffy)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

Jiffy is a golang library to work with `time.Duration`.
It provides a simpler way on defining and describing `time.Duration`.

*Jiffy* extends the `func (d Duration) String() string` where Jiffy can describe a duration up to day, month even year. Verbose or non-verbose<br>
*Jiffy* also extends the `func ParseDuration(s string) (Duration, error)` where Jiffy can parse day, month and year element in a more flexible formating.

Usecases :

* You have measured a processing time duration and want to display it in a format easily understand by human.

```go
startTime := time.Now()
/*
    you do a heavy time consuming process here
*/
endTime := time.Now()
duration := startTime.Sub(endTime)

// print using time.Duration.String() function
fmt.Println(duration.String()) 
// this print like "30086h0m0s", not so easy to understand

// print using jiffy
fmt.Println(jiffy.DescribeDuration(duration, jiffy.NewWant()))
// this print "3y 4mon 13d 14h"
```

* You want to provide a simple way for your user to define some time duration.

```go
workTime := getUserInput() // user key in his working time as string. eg. "3h 15minutes" or "3 hours 15 minutes" or "3h and 15m"
workDuration, err := jiffy.DurationOf(workTime)
if err != nil {
    panic(err)
}
if workDuration > 2 * time.Hour {
    // do something.
}
```

Too add *Jiffy* into your project.

```
$ go get github.com/hyperjumptech/jiffy
```

To import *Jiffy* in your coode

```go
import "gitlab.com/hyperjujptech/jiffy"
```

## Creating time.Duration from a Description

```go
dur, err := jiffy.DurationOf("2 hours 30 minutes")
if err != nil {
    panic(err)
}
```

Examples of duration descriptions

* `5d 3h 12m`
* `2y 3mon 15d`
* `12 days 16hours`
* `1minute 15seconds`
* `4seconds and 15milliseconds`
* `25s123ms59microsecond`

### Time units

| Unit        | Notation | Example(s) |
| ----------- | -------- | ---------- |
| Year        | `y` or `year` or `years` | `25y` or `25 y` or `25 years` |
| Month       | `mon` or `month` or `months` | `2.5mon` or `3 month` or `6 months` |
| Day         | `d` or `day` or `days` | `0.5d` or `1.5day` or `2 days` |
| Hour        | `h` or `hrs` or `hour` or `hours` | `1h` or `2hour` or `3 hours` |
| Minute      | `m` or `min` or `minute` or `minutes` | `4m` or `5minute` or `6minutes`  |
| Second      | `s` or `sec` or `second` or `seconds` | `7s` or `8 second` or `9seconds` |
| Millisecond | `ms` or `millisecond` or `milliseconds` | `120ms` or `450 millisecond` or `500 milliseconds` |
| Microsecond | `us` or `microsecond` or `microseconds` | `300us` or `1200.5microsecond` or `0.1 microseconds` |

_Note_ :

* You can add the word `and` between time unit. For example : `12h and 30m`
* Your duration description is *case insensitive*
* You could always use *real* literals. For example : `12.5 days`
* You also can use a negative sign. For example : `-5d`
* Whitespaces (eg. tabs or space) is ignored, so you can always add as many spaces or tabs as you want.
* `CR` or `LF` is an illegal character.
* When you specify each unit, it can be in any order. For example `12 days 5 hour 3sec` is the same as `3 second 5hours 12day`
* You can repeat same time units. For example : `5d 5d 1d` equals to `11d`

## Describing a time.Duration

To have a `time.Duration` get described, you can use 

```go
func jiffy.DescribeDuration(dur time.Duration, want *jiffy.Want) string
```

`jiffy.Want` is a _struct_ to specify what time unit should be displayed to describe the `time.Duration`.
Because, a duration of `3 hours` can also be described in `180 minutes` or `10800 seconds`. 

Following is the `jiffy.Want` _struct_

```go
type Want struct {
	Year        bool   // set this true if you want year to be calcuated and displayed.
	Month       bool   // or this one for month
	Day         bool   // or this one for day
	Hour        bool   // or this one for hour
	Minute      bool   // or this one for minute
	Second      bool   // or this one for second
	Millisecond bool   // or this one for millisecond
	Microsecond bool   // or this one for microsecond
	Separator   string // specify this string as separator between time units
	Verbose     bool   // verbose/nonverbose mode
}
```

`Want` also have a default _constructor_ function 

```go
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
```

### Verbosity

If you want to describe a `time.Duration` verbosely the `DescribeDuration` function will return for example

```
1 hour 2 minutes 4 seconds
```

And if you set the `Want.Verbose = false` it will produce

```
1h 2m 4s
```

### Describe example

```go
aDuration := 1 * time.Second + 1 * time.Minute + 1 * time.Hour

fmt.Println( jiffy.DescribeDuration(aDuration, &jiffy.Want{
    Hour: true,
    Minute: true,
    Second: true,
    Verbose: false,
    Separator: " ",
}) ) // this will print "1h 1m 1s"

fmt.Println( jiffy.DescribeDuration(aDuration, &jiffy.Want{
    Hour: true,
    Minute: true,
    Second: true,
    Verbose: true,
    Separator: " ",
}) ) // this will print "1 hour 1 minute 1 second"

fmt.Println( jiffy.DescribeDuration(aDuration, &jiffy.Want{
    Minute: true,
    Second: true,
    Verbose: true,
    Separator: ", ",
}) ) // this will print "61 minutes, 1 seconds"

fmt.Println( jiffy.DescribeDuration(aDuration, &jiffy.Want{
    Second: true,
    Verbose: true,
    Separator: ", ",
}) ) // this will print "3661 seconds"
```

## FAQ

*Q* : Why *jiffy* does not support *Nanosecond* ?<br>
*A* : Nanosecond is too much. No human can comprehend that. Or if you want it so badly,  please submit a PR.

*Q* : Why *jiffy* ? Not the golang's `func (d Duration) String() string` and `func ParseDuration(s string) (Duration, error)` ?<br>
*A* : Because *jiffy* parsing and producing up to  year, month, date, and it parses string in more flexible way

# Tasks and Help Wanted.

Yes. We need contributor to make **Jiffy** even better and useful to Open Source Community.

If you really want to help us, simply `Fork` the project and apply for Pull Request.
Please read our [Contribution Manual](CONTRIBUTING.md) and [Code of Conduct](CODE_OF_CONDUCTS.md)