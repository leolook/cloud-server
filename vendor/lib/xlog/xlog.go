package xlog

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/metadata"
	"lib/goid"
	"sync"
)

// Debug logs to the DEBUG log.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Debug(args ...interface{}) {
	gid := goid.Get()
	if color, ok := GetColorByKey(gid); ok {
		msg := fmt.Sprint(args)
		g_logger.Debugf(msg, "method", color.Method(), "traceid", color.TraceID(), "uid", color.Uid(), "clientversion", color.ClientVerson(), "clienttype", color.ClentType())
	} else {
		g_logger.Debug(args...)
	}
}

// Debugln logs to the INFO log.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Debugln(args ...interface{}) {
	gid := goid.Get()
	if color, ok := GetColorByKey(gid); ok {
		msg := fmt.Sprint(args)
		g_logger.Debugw(msg, "method", color.Method(), "traceid", color.TraceID(), "uid", color.Uid(), "clientversion", color.ClientVerson(), "clienttype", color.ClentType())
	} else {
		g_logger.Debug(args...)
	}
}

// Debugf logs to the INFO log.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Debugf(format string, args ...interface{}) {
	gid := goid.Get()
	if color, ok := GetColorByKey(gid); ok {
		msg := fmt.Sprintf(format, args...)
		g_logger.Debugw(msg, "method", color.Method(), "traceid", color.TraceID(), "uid", color.Uid(), "clientversion", color.ClientVerson(), "clienttype", color.ClentType())
	} else {
		g_logger.Debugf(format, args...)
	}
}

// Info logs to the INFO log.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Info(args ...interface{}) {
	gid := goid.Get()
	if color, ok := GetColorByKey(gid); ok {
		msg := fmt.Sprint(args)
		g_logger.Infow(msg, "method", color.Method(), "traceid", color.TraceID(), "uid", color.Uid(), "clientversion", color.ClientVerson(), "clienttype", color.ClentType())
	} else {
		g_logger.Info(args...)
	}
}

// Infoln logs to the INFO log.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Infoln(args ...interface{}) {
	gid := goid.Get()
	if color, ok := GetColorByKey(gid); ok {
		msg := fmt.Sprint(args)
		g_logger.Infow(msg, "method", color.Method(), "traceid", color.TraceID(), "uid", color.Uid(), "clientversion", color.ClientVerson(), "clienttype", color.ClentType())
	} else {
		g_logger.Info(args...)
	}
}

// Infof logs to the INFO log.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Infof(format string, args ...interface{}) {
	gid := goid.Get()
	if color, ok := GetColorByKey(gid); ok {
		msg := fmt.Sprintf(format, args...)
		g_logger.Infow(msg, "method", color.Method(), "traceid", color.TraceID(), "uid", color.Uid(), "clientversion", color.ClientVerson(), "clienttype", color.ClentType())
	} else {
		g_logger.Infof(format, args...)
	}
}

func InfoW(format string, keysAndValues ...interface{}) {
	g_logger.Infow(format, keysAndValues...)
}

// Warning logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Warn(args ...interface{}) {
	gid := goid.Get()
	if color, ok := GetColorByKey(gid); ok {
		msg := fmt.Sprint(args)
		g_logger.Warnw(msg, "method", color.Method(), "traceid", color.TraceID(), "uid", color.Uid(), "clientversion", color.ClientVerson(), "clienttype", color.ClentType())
	} else {
		g_logger.Warn(args...)
	}
}

func Warnw(format string, keysAndValues ...interface{}) {
	g_logger.Warnw(format, keysAndValues...)
}

// Warningln logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Warnln(args ...interface{}) {
	gid := goid.Get()
	if color, ok := GetColorByKey(gid); ok {
		msg := fmt.Sprint(args)
		g_logger.Warnw(msg, "method", color.Method(), "traceid", color.TraceID(), "uid", color.Uid(), "clientversion", color.ClientVerson(), "clienttype", color.ClentType())
	} else {
		g_logger.Warn(args...)
	}
}

// Warningf logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Warnf(format string, args ...interface{}) {
	gid := goid.Get()
	if color, ok := GetColorByKey(gid); ok {
		msg := fmt.Sprintf(format, args...)
		g_logger.Warnw(msg, "method", color.Method(), "traceid", color.TraceID(), "uid", color.Uid(), "clientversion", color.ClientVerson(), "clienttype", color.ClentType())
	} else {
		g_logger.Warnf(format, args...)
	}
}

// Error logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Error(args ...interface{}) {
	gid := goid.Get()
	if color, ok := GetColorByKey(gid); ok {
		msg := fmt.Sprint(args)
		g_logger.Errorw(msg, "method", color.Method(), "traceid", color.TraceID(), "uid", color.Uid(), "clientversion", color.ClientVerson(), "clienttype", color.ClentType())
	} else {
		g_logger.Error(args...)
	}
}

// Errorln logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Errorln(args ...interface{}) {
	gid := goid.Get()
	if color, ok := GetColorByKey(gid); ok {
		msg := fmt.Sprint(args)
		g_logger.Errorw(msg, "method", color.Method(), "traceid", color.TraceID(), "uid", color.Uid(), "clientversion", color.ClientVerson(), "clienttype", color.ClentType())
	} else {
		g_logger.Error(args...)
	}
}

// Errorf logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Errorf(format string, args ...interface{}) {
	gid := goid.Get()
	if color, ok := GetColorByKey(gid); ok {
		msg := fmt.Sprintf(format, args...)
		g_logger.Errorw(msg, "method", color.Method(), "traceid", color.TraceID(), "uid", color.Uid(), "clientversion", color.ClientVerson(), "clienttype", color.ClentType())
	} else {
		g_logger.Errorf(format, args...)
	}
}

func Errorw(format string, keysAndValues ...interface{}) {
	g_logger.Errorw(format, keysAndValues...)
}

// Fatal logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Fatal(args ...interface{}) {
	gid := goid.Get()
	if color, ok := GetColorByKey(gid); ok {
		msg := fmt.Sprint(args)
		g_logger.Errorw(msg, "method", color.Method(), "traceid", color.TraceID(), "uid", color.Uid(), "clientversion", color.ClientVerson(), "clienttype", color.ClentType())
	} else {
		g_logger.Fatal(args...)
	}
}

// Fatalln logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Fatalln(args ...interface{}) {
	gid := goid.Get()
	if color, ok := GetColorByKey(gid); ok {
		msg := fmt.Sprint(args)
		g_logger.Fatalw(msg, "method", color.Method(), "traceid", color.TraceID(), "uid", color.Uid(), "clientversion", color.ClientVerson(), "clienttype", color.ClentType())
	} else {
		g_logger.Fatal(args...)
	}
}

// Fatalf logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Fatalf(format string, args ...interface{}) {
	gid := goid.Get()
	if color, ok := GetColorByKey(gid); ok {
		msg := fmt.Sprintf(format, args...)
		g_logger.Fatalw(msg, "method", color.Method(), "traceid", color.TraceID(), "uid", color.Uid(), "clientversion", color.ClientVerson(), "clienttype", color.ClentType())
	} else {
		g_logger.Fatalf(format, args...)
	}
}
func With(format string, fields []zap.Field) {
	g_logger.Desugar().With(fields...).Sugar().Debugf(format)
}

func getHeader(md metadata.MD, h string) string {
	v, ok := md[h]
	if !ok {
		return ""
	}

	if len(v) > 0 {
		return v[0]
	} else {
		return ""
	}
}

func getclientType(clienttype string) string {
	if clienttype == "1" {
		return "ios"
	} else if clienttype == "2" {
		return "android "
	} else {
		return "unknown client type"
	}
}

func TagsToFields(md metadata.MD) []zapcore.Field {
	fields := []zapcore.Field{}

	uid := getHeader(md, "uid")
	seq := getHeader(md, "seq")
	rid := getHeader(md, "rid")
	method := getHeader(md, "method")
	clienttype := getHeader(md, "client_type")
	clinetversion := getHeader(md, "client_version")
	trcaeid := getHeader(md, "traceid")

	fields = append(fields, zap.Any("uid", uid))
	fields = append(fields, zap.Any("seq", seq))
	fields = append(fields, zap.Any("rid", rid))
	fields = append(fields, zap.Any("method", method))
	fields = append(fields, zap.Any("client_type", getclientType(clienttype)))
	fields = append(fields, zap.Any("client_versiond", clinetversion))
	fields = append(fields, zap.Any("trcaeid", trcaeid))

	return fields
}

func Extract(md metadata.MD) *zap.SugaredLogger {
	fields := TagsToFields(md)
	baseLog := g_logger.Desugar()
	return baseLog.WithOptions(zap.AddCallerSkip(-1)).With(fields...).Sugar()
}

func GetZapLog() *zap.SugaredLogger {
	return g_logger
}

type colorI interface {
	String() string
	Uid() string
	ClientVerson() string
	Method() string
	TraceID() string
	ClentType() string
}

var g_color sync.Map

func SetColor(val colorI) bool {
	gid := goid.Get()
	_, loaded := g_color.LoadOrStore(gid, val)
	return !loaded
}

func GetColorByKey(key int64) (colorI, bool) {
	val, ok := g_color.Load(key)
	if !ok {
		return nil, ok
	} else {
		return val.(colorI), ok
	}
}

func GetColor() (colorI, bool) {
	gid := goid.Get()
	val, ok := g_color.Load(gid)
	return val.(colorI), ok
}

func DelColor() {
	gid := goid.Get()
	g_color.Delete(gid)
}
