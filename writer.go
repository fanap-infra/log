package log

import "sync"

type encoder interface {
	Print(l Level, s string, caller string, stacks []string, message string)
	Printv(l Level, s string, caller string, stacks []string, message string, keysValues []interface{})
	close()
}
type enablerFuncParam struct {
	l Level
	s string
}

type Writer struct {
	encoder
	enabler EnablerFunc
	stack   EnablerFunc
	caller  bool

	enablerCache sync.Map
	stackCache   sync.Map
}

func newWriter(enabler EnablerFunc, stack EnablerFunc, caller bool, encoder encoder) *Writer {
	return &Writer{
		encoder: encoder,
		enabler: enabler,
		caller:  caller,
		stack:   stack,

		// enablerCache: make(map[enablerFuncParam]bool),
		// stackCache:   make(map[enablerFuncParam]bool),
	}
}

func (w *Writer) isEnable(level Level, scope string) bool {
	p := enablerFuncParam{level, scope}

	if r, ok := w.enablerCache.Load(p); ok {
		return r.(bool)
	}

	r := w.enabler(level, scope)
	w.enablerCache.Store(p, r)

	return r
}

func (w *Writer) isStack(level Level, scope string) bool {
	p := enablerFuncParam{level, scope}

	if r, ok := w.stackCache.Load(p); ok {
		return r.(bool)
	}

	r := w.stack(level, scope)
	w.stackCache.Store(p, r)

	return r
}
