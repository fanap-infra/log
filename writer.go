package log

type encore interface {
	Print(l Level, s string, caller string, stacks []string, message string)
	Printv(l Level, s string, caller string, stacks []string, message string, keysValues []interface{})
	close()
}

type Writer struct {
	encore
	enabler EnablerFunc // ToDo: immutable: use map for cache reponse
	stack   EnablerFunc // ToDo: immutable: use map for cache reponse
	caller  bool
}
