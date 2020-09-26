package log

import (
	"fmt"
	"strconv"
)

type EnablerFunc func(level Level, scope string) bool

type core struct {
	writers []*Writer

	print  func(l Level, s string, skip int, messages []interface{})
	printf func(l Level, s string, skip int, format string, args []interface{})
	printv func(l Level, s string, skip int, message string, keysValues []interface{})
}

func (c *core) Close() {
	for _, w := range c.writers {
		w.close()
	}
}

func (c *core) Config(w []*Writer) {
	switch len(w) {
	case 0:
		c.print = func(l Level, s string, skip int, messages []interface{}) {}
		c.printf = func(l Level, s string, skip int, format string, args []interface{}) {}
		c.printv = func(l Level, s string, skip int, message string, keysValues []interface{}) {}
	case 1:
		c.print = c.print1
		c.printf = c.printf1
		c.printv = c.printv1
	// case 2:
	// 	c.Print = c.print2
	// 	c.Printf = c.printf2
	// 	c.Printv = c.printv2
	default:
		c.print = c.printAll
		c.printf = c.printfAll
		c.printv = c.printvAll
	}

	// close old writer
	c.Close()

	c.writers = w
}

func (c *core) print1(l Level, s string, skip int, messages []interface{}) {
	w := c.writers[0]
	if w.isEnable(l, s) {
		caller := ""
		var stacks []string
		if w.caller {
			caller = c.getCaller(skip)
		}
		if w.isStack(l, s) {
			stacks = c.getStacks(skip)
		}
		w.Print(l, s, caller, stacks, fmt.Sprint(messages...))
	}
}

func (c *core) printf1(l Level, s string, skip int, format string, args []interface{}) {
	w := c.writers[0]
	if w.isEnable(l, s) {
		caller := ""
		var stacks []string
		if w.caller {
			caller = c.getCaller(skip)
		}
		if w.isStack(l, s) {
			stacks = c.getStacks(skip)
		}

		w.Print(l, s, caller, stacks, fmt.Sprintf(format, args...))
	}
}

func (c *core) printv1(l Level, s string, skip int, message string, keysValues []interface{}) {
	w := c.writers[0]
	if w.isEnable(l, s) {
		caller := ""
		var stacks []string
		if w.caller {
			caller = c.getCaller(skip)
		}
		if w.isStack(l, s) {
			stacks = c.getStacks(skip)
		}
		w.Printv(l, s, caller, stacks, message, keysValues)
	}
}

// // func (c *core) print2(l Level, s string, messages ...interface{}) {
// // 	t := fmt.Sprint(messages...)

// // 	caller := ""
// // 	if c.writers[0].caller || c.writers[1].caller {
// // 		caller = c.getCaller()
// // 	}

// // 	stack := ""
// // 	if c.writers[0].stack(l, s) || c.writers[1].stack(l, s) {
// // 		stack = c.getStack()
// // 	}

// // 	w := c.writers[0]
// // 	if w.isEnable(l, s) {
// // 		c1 := ""
// // 		s1 := ""

// // 		if w.caller {
// // 			c1 = caller
// // 		}

// // 		if w.caller {
// // 			s1 = stack
// // 		}

// // 		w.Print(l, s, c1, s1, t)
// // 	}

// // 	w = c.writers[1]
// // 	if w.isEnable(l, s) {
// // 		if w.caller {
// // 			w.Print(l, s, caller, stack, t)
// // 		} else {
// // 			w.Print(l, s, "", stack, t)
// // 		}
// // 	}
// // }

// // func (c *core) printf2(l Level, s string, format string, args []interface{}) {
// // 	message := fmt.Sprintf(format, args...)
// // 	caller := ""
// // 	if c.writers[0].caller || c.writers[1].caller {
// // 		caller = c.getCaller()
// // 	}

// // 	w := c.writers[0]
// // 	if w.isEnable(l, s) {
// // 		if w.caller {
// // 			w.Print(l, s, caller, message)
// // 		} else {
// // 			w.Print(l, s, "", message)
// // 		}
// // 	}

// // 	w = c.writers[1]
// // 	if w.isEnable(l, s) {
// // 		if w.caller {
// // 			w.Print(l, s, caller, message)
// // 		} else {
// // 			w.Print(l, s, "", message)
// // 		}
// // 	}
// // }

// func (c *core) printv2(l Level, s string, message string, keysValues []interface{}) {
// 	w := c.writers[0]
// 	caller := ""
// 	if c.writers[0].caller || c.writers[1].caller {
// 		caller = c.getCaller()
// 	}

// 	if w.isEnable(l, s) {
// 		if w.caller {
// 			w.Printv(l, s, caller, message, keysValues)
// 		} else {
// 			w.Printv(l, s, "", message, keysValues)
// 		}
// 	}

// 	w = c.writers[1]
// 	if w.isEnable(l, s) {
// 		if w.caller {
// 			w.Printv(l, s, caller, message, keysValues)
// 		} else {
// 			w.Printv(l, s, "", message, keysValues)
// 		}
// 	}
// }

func (c *core) printAll(l Level, s string, skip int, messages []interface{}) {
	caller := ""
	var stacks []string

	callerS := ""
	var stacksS []string

	for _, w := range c.writers {
		if w.isEnable(l, s) {
			if w.caller {
				if caller == "" {
					caller = c.getCaller(skip)
				}
				callerS = caller
			}

			if w.isStack(l, s) {
				if len(stacks) == 0 {
					stacks = c.getStacks(skip)
				}
				stacksS = stacks
			}

			w.Print(l, s, callerS, stacksS, fmt.Sprint(messages...))
		}
	}
}

func (c *core) printfAll(l Level, s string, skip int, format string, args []interface{}) {
	caller := ""
	var stacks []string

	callerS := ""
	var stacksS []string
	message := fmt.Sprintf(format, args...)
	for _, w := range c.writers {
		if w.isEnable(l, s) {
			if w.caller {
				if caller == "" {
					caller = c.getCaller(skip)
				}
				callerS = caller
			}

			if w.isStack(l, s) {
				if len(stacks) == 0 {
					stacks = c.getStacks(skip)
				}
				stacksS = stacks
			}
			w.Print(l, s, callerS, stacksS, message)
		}
	}
}

func (c *core) printvAll(l Level, s string, skip int, message string, keysValues []interface{}) {
	caller := ""
	var stacks []string

	callerS := ""
	var stacksS []string
	for _, w := range c.writers {
		if w.isEnable(l, s) {
			if w.caller {
				if caller == "" {
					caller = c.getCaller(skip)
				}
				callerS = caller
			}

			if w.isStack(l, s) {
				if len(stacks) == 0 {
					stacks = c.getStacks(skip)
				}
				stacksS = stacks
			}
			w.Printv(l, s, callerS, stacksS, message, keysValues)
		}
	}
}

func (c *core) getCaller(skip int) string {
	frame, defined := getCallerFrame(skip) // log.callerSkip + callerSkipOffset
	if !defined {
		return ""
	}

	return getFolderFile(frame.File) + ":" + strconv.Itoa(frame.Line)
}

func (c *core) getStacks(skip int) []string {
	return getStacks(skip)
}
