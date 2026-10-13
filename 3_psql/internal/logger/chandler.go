package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"sync"
	"time"
)

type CHandler struct {
	// json              bool // true => output JSON; false => output text
	opts  *slog.HandlerOptions
	mu    *sync.Mutex
	w     io.Writer
	attrs []slog.Attr
	// preformattedAttrs []byte
	// groupPrefix is for the text handler only.
	// It holds the prefix for groups that were already pre-formatted.
	// A group will appear here when a call to WithGroup is followed by
	// a call to WithAttrs.
	// groupPrefix string
	// groups      []string // all groups started from WithGroup
	// nOpenGroups int      // the number of groups opened in preformattedAttrs
}

func newCustomHandler(w io.Writer, opts *slog.HandlerOptions) *CHandler {
	if opts == nil {
		opts = &slog.HandlerOptions{}
	}
	return &CHandler{
		w:    w,
		opts: opts,
		mu:   &sync.Mutex{},
	}
}

func (h *CHandler) Enabled(_ context.Context, level slog.Level) bool {
	return true
	// return level >= h.opts.Level.Level()
}

func (h *CHandler) Handle(_ context.Context, rec slog.Record) error {
	//заменить строку на слайс  байт или выделить под строку сразу размер
	logLine := fmt.Sprintf("%s [%s]\n    %s",
		rec.Time.Format(time.DateTime),
		rec.Level,
		rec.Message,
	)

	rec.Attrs(func(a slog.Attr) bool {
		logLine += fmt.Sprintf(" | %s=%v", a.Key, a.Value)
		return true
	})
	// добавить вывод атрибутов из слайса по нормальному, сейчас черновой вариант
	for _, v := range h.attrs {
		logLine += fmt.Sprintf(" | %s=%v", v.Key, v.Value)
	}
	logLine += "\n"
	h.mu.Lock()
	defer h.mu.Unlock()
	h.w.Write([]byte(logLine))
	return nil
}

func (h *CHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	if len(attrs) == 0 {
		return h
	}

	nh := *h
	newAttrs := make([]slog.Attr, len(h.attrs), len(h.attrs)+len(attrs))
	copy(newAttrs, h.attrs)
	nh.attrs = append(newAttrs, attrs...)
	//или просто: nh.attrs = slices.Clone(h.attrs)
	//nh.attrs = append(nh.attrs, attrs...)

	return &nh
}

func (h *CHandler) WithGroup(name string) slog.Handler {
	if name == "" {
		return h
	}
	return h
	// return c.withGroupOrAttrs(groupOrAttrs{group: name})
}
