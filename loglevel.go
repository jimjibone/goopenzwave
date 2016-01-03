package goopenzwave

// #include "loglevel.h"
// #include <stdlib.h>
import "C"

type LogLevel int32

const (
	LogLevelInvalid      = LogLevel(C.loglevel_invalid)
	LogLevelNone         = LogLevel(C.loglevel_none)
	LogLevelAlways       = LogLevel(C.loglevel_always)
	LogLevelFatal        = LogLevel(C.loglevel_fatal)
	LogLevelError        = LogLevel(C.loglevel_error)
	LogLevelWarning      = LogLevel(C.loglevel_warning)
	LogLevelAlert        = LogLevel(C.loglevel_alert)
	LogLevelInfo         = LogLevel(C.loglevel_info)
	LogLevelDetail       = LogLevel(C.loglevel_detail)
	LogLevelDebug        = LogLevel(C.loglevel_debug)
	LogLevelStreamdetail = LogLevel(C.loglevel_streamdetail)
	LogLevelInternal     = LogLevel(C.loglevel_internal)
)
