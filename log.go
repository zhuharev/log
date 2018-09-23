// Copyright 2018 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package log

import (
	"log"
	"os"
	"strconv"
)

var (
	logVerbose, _ = strconv.ParseBool(os.Getenv("VERBOSE"))
	defaultLogger = New(Verbose(logVerbose))
)

// Printf log to stdout if verbose enabled
func Printf(format string, args ...interface{}) {
	defaultLogger.Printf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}

type Logger struct {
	debug, err *log.Logger
}

type Opt func(*Logger)

func Verbose(verboses ...bool) Opt {
	return func(l *Logger) {
		verbose := true
		if len(verboses) > 0 {
			verbose = verboses[0]
		}
		if verbose {
			l.debug = log.New(os.Stdout, "D", log.LstdFlags|log.Lshortfile)
		} else {
			l.debug = nil
		}
	}
}

func New(opt ...Opt) *Logger {
	l := &Logger{
		err: log.New(os.Stderr, "E", log.LstdFlags|log.Lshortfile),
	}
	for _, fn := range opt {
		fn(l)
	}
	return l
}

func (l *Logger) Printf(format string, args ...interface{}) {
	if l == nil {
		return
	}
	if l.debug != nil {
		l.debug.Printf(format, args...)
	}
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	if l == nil {
		return
	}
	if l.err == nil {
		return
	}
	l.err.Printf(format, args...)
}
