package log

type encore interface {
	Print(l Level, s string, caller string, stacks []string, message string)
	Printv(l Level, s string, caller string, stacks []string, message string, keysValues []interface{})
	close()
}

type Writer struct {
	encore
	enabler EnablerFunc
	stack   EnablerFunc
	caller  bool
}
