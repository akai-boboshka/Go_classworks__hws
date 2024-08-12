package main

import (
	"strings"
)

// #1

// type StringProcessor interface {
// 	Length(string) int
// 	// WordCount()
// }

// type LengthOfString struct{}
// type WordCount struct{}

// func (l LengthOfString) Length(str string) int {
// 	return len(str)
// }

// func main() {
// 	str := "Hello world"
// 	var s StringProcessor
// 	s = LengthOfString{}
// 	fmt.Println(s.Length(str))
// }

// #2
type Formatter interface {
	ToUpper() string
	ToLower() string
}

type MyFormatter struct {
	s string
}

func (mf MyFormatter) ToUpper() string {
	return strings.ToUpper(mf.s)
}

func (mf MyFormatter) ToLower() string {
	return strings.ToLower(mf.s)
}

type Formater struct {
	Formatter
}

func NewFormatter(fm Formatter) Formater {
	return Formater{
		fm,
	}
}

// #3
type PointerOperations interface {
	Increment()
	Decrement()
}

type IntPointer struct {
	n *int
}

func (ip IntPointer) Incerement() {
	*ip.n++
}

func (ip IntPointer) Decrement() {
	*ip.n--
}

type Pointer struct {
	PointerOperations
}

func NewPointer(po PointerOperations) Pointer {
	return Pointer{
		po,
	}
}

// #4
type StringCleaner interface {
	TrimSpaces() string
	// RemoveSpaces() string
}

type TextCleaner struct {
	s string
}

func (tc TextCleaner) TrimSpaces() string {
	return strings.TrimSpace(tc.s)
}

// func (tc TextCleaner) RemoveSpaces() {

// }

type NewText struct {
	StringCleaner
}

func NewText2(sc StringCleaner) NewText {
	return NewText{
		sc,
	}
}
