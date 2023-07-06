package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	t := time.Now()

	p(t.Format(time.RFC3339))

	t1, _ := time.Parse(
		time.RFC3339,
		"2023-06-27T09:39:20+00:00")
	p(t1)

	/*
	 Layouts must use the reference time Mon Jan 2 15:04:05 MST 2006 to show
	 the pattern with which to format/parse a given time/string. The example
	 time must be exactly as shown: the year 2006, 15 for the hour,
	 Monday for the day of the week, etc.
	 refer to:
	 https://gobyexample.com/time-formatting-parsing

	*/
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))

	form := "Mon 3:04 PM"
	t2, _ := time.Parse(form, "Tue 8:39 PM")
	p(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e := time.Parse(ansic, "8:41PM")
	p(e)
}
