package main

import (
	"errors"
	"fmt"
)

func function1(arg int) (int, error) {
	if arg == 45 {
		return -1, errors.New("can't work with 45")
	}

	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func function2(arg int) (int, error) {
	if arg == 45 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	for _, value := range []int{7, 45} {
		result, err := function1(value)
		if err != nil {
			fmt.Println("function 1 failed: ", err)
		} else {
			fmt.Println("function 1 worked: ", result)
		}
	}
	for _, value := range []int{7, 45} {
		result, err := function2(value)
		if err != nil {
			fmt.Println("function 2 failed: ", err)
		} else {
			fmt.Println("function 2 worked: ", result)
		}
	}

	_, err := function2(45)
	argumentError, typeAssertionOkay := err.(*argError)
	if typeAssertionOkay {
		fmt.Println(argumentError.arg)
		fmt.Println(argumentError.prob)
	}

}
