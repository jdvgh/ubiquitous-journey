package main

import (
	"fmt"
	"time"
)

func main() {

	p := fmt.Println

	now := time.Now()
	p(now)
	then := time.Date(
		2009, 11, 17, 20, 34, 59, 651387237, time.UTC)
	p(then)

	p(then.Year())

	p(then.Local())
	p(then.Location())
	p(then.Weekday())

	p(then.Before(now))
	p(now.Sub(then).Hours())

	p(now.Add(now.Sub(then)))

}
