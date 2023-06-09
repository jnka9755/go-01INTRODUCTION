package main

import (
	"fmt"
	"time"
)

func main() {

	p := fmt.Println

	now := time.Now()
	p(now)
	p()

	then := time.Date(2023, 04, 27, 12, 30, 25, 12124, time.UTC)
	p(then)
	p()

	then = then.Add(-time.Hour * 1)
	p(then)
	p()

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday())

	p("Then es antes que Now: ", then.Before(now))
	p("Then es despues que Now: ", then.After(now))
	p("Then es igual que Now: ", then.Equal(now))

	p()

	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Milliseconds())
}
