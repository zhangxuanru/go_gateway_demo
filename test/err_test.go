/*
@Time : 2020/9/28 19:19 
@Author : zxr
@File : err_test
@Software: GoLand
*/
package test

import (
	"fmt"
	"io"
	"path"
	"runtime"
	"strings"
	"testing"
)

type stack []uintptr
type Frame uintptr

type AA struct {
	msg string
	*stack
}

func callers() *stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	var st stack = pcs[0:n]
	return &st
}

func (a *AA) Error() string {
	return a.msg
}

func (a *AA) Format(s fmt.State, verb rune) {
	fmt.Println("aa:Format",s,"----",verb)
	switch verb {
	case 'v':
		if s.Flag('+') {
			io.WriteString(s, a.msg)
			a.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, a.msg)
	case 'q':
		fmt.Fprintf(s, "%q", a.msg)
	}
}

func (s *stack) Format(st fmt.State, verb rune) {
	switch verb {
	case 'v':
		switch {
		case st.Flag('+'):
			for _, pc := range *s {
				f := Frame(pc)
				fmt.Fprintf(st, "\n%+v", f)
			}
		}
	}
}

func (f Frame) Format(s fmt.State, verb rune) {
	switch verb {
	case 's':
		switch {
		case s.Flag('+'):
			pc := f.pc()
			fn := runtime.FuncForPC(pc)
			if fn == nil {
				io.WriteString(s, "unknown")
			} else {
				file, _ := fn.FileLine(pc)
				fmt.Fprintf(s, "%s\n\t%s", fn.Name(), file)
			}
		default:
			io.WriteString(s, path.Base(f.file()))
		}
	case 'd':
		fmt.Fprintf(s, "%d", f.line())
	case 'n':
		name := runtime.FuncForPC(f.pc()).Name()
		io.WriteString(s, funcname(name))
	case 'v':
		f.Format(s, 's')
		io.WriteString(s, ":")
		f.Format(s, 'd')
	}
}

func (f Frame) pc() uintptr { return uintptr(f) - 1 }
func (f Frame) file() string {
	fn := runtime.FuncForPC(f.pc())
	if fn == nil {
		return "unknown"
	}
	file, _ := fn.FileLine(f.pc())
	return file
}
// function for this Frame's pc.
func (f Frame) line() int {
	fn := runtime.FuncForPC(f.pc())
	if fn == nil {
		return 0
	}
	_, line := fn.FileLine(f.pc())
	return line
}

// funcname removes the path prefix component of a function's name reported by func.Name().
func funcname(name string) string {
	i := strings.LastIndex(name, "/")
	name = name[i+1:]
	i = strings.Index(name, ".")
	return name[i+1:]
}

func New(msg string) error  {
	return &AA{
		msg:msg,
		stack:callers(),
	}
}


type A string
func (c A) Format(s fmt.State, verb rune) {
   fmt.Println("A---a -- Format")
}
func TestErr(t *testing.T)  {
    var aa A
	aa = "hello test"
	fmt.Printf("%+v\n",aa)

	//err:=New("hello world")
	//stack := fmt.Sprintf("%+v",err)
	//fmt.Printf(stack)
}
