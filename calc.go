package main

import (
	"errors"
	// "fmt"
	"strconv"
	"strings"
)

var ErrInvalidInput = errors.New("invalid input")

func Calc(expression string) (float64, error) {
	if strings.Replace(expression, " ", "", -1) == "" {
		return 0, ErrInvalidInput
	}
	if isFloat(expression) {
		return strconv.ParseFloat(expression, 64)
	}
	expression = strings.Replace(expression, " ", "", -1)
	expression = strings.Replace(expression, "+", " + ", -1)
	expression = strings.Replace(expression, "-", " - ", -1)
	expression = strings.Replace(expression, "*", " * ", -1)
	expression = strings.Replace(expression, "/", " / ", -1)
	expression = strings.Replace(expression, "(", "( ", -1)
	expression = strings.Replace(expression, ")", " )", -1)
	expr := strings.Split(expression, " ")

	// fmt.Printf("%q\n", expr)
	st := make(stack, 0)
	out := make(stack, 0)
	for _, chr := range expr {
		switch getType(chr) {
		case 'd':
			out.Push(chr)
		case 'l':
			for !st.Empty() &&
				(getType(st.Top()) == 'h' || getType(st.Top()) == 'l') {
				out.Push(st.Pop())
			}
			st.Push(chr)
		case 'h':
			for !st.Empty() && getType(st.Top()) == 'h' {
				out.Push(st.Pop())
			}
			st.Push(chr)
		case '(':
			st.Push(chr)
		case ')':
			if st.Empty() {
				return 0, ErrInvalidInput
			}
			for st.Top() != "(" {
				out.Push(st.Pop())
				if st.Empty() {
					return 0, ErrInvalidInput
				}
			}
			st.Pop()
		}
	}
	for !st.Empty() {
		out.Push(st.Pop())
	}
	// fmt.Printf("%q | %q\n\n", st, out)

	stack := make([]float64, 0)
	for _, i := range out {
		if isFloat(i) {
			r, _ := strconv.ParseFloat(i, 64)
			stack = append(stack, r)
			continue
		}
		if len(stack) < 2 {
			return 0, ErrInvalidInput
		}
		q := stack[len(stack)-1]
		p := stack[len(stack)-2]
		stack = stack[:len(stack)-2]
		switch i {
		case "+":
			stack = append(stack, p+q)
		case "-":
			stack = append(stack, p-q)
		case "*":
			stack = append(stack, p*q)
		case "/":
			stack = append(stack, p/q)
		}
	}
	if len(stack) != 1 {
		return 0, ErrInvalidInput
	}
	return stack[0], nil
}

func getType(c string) byte {
	switch {
	case isFloat(c):
		return 'd'
	case c == "+", c == "-":
		return 'l'
	case c == "*", c == "/":
		return 'h'
	case c == "(":
		return '('
	case c == ")":
		return ')'
	}
	return '0'
}

func isFloat(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

type stack []string

func (s *stack) Push(x string) {
	*s = append(*s, x)
}

func (s *stack) Pop() string {
	r := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return r
}

func (s *stack) Top() string {
	return (*s)[len(*s)-1]
}

func (s *stack) Empty() bool {
	return len(*s) == 0
}
