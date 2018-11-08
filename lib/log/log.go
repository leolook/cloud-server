package log

// Debug logs to the DEBUG log.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Debug(args ...interface{}) {
	g_logger.Sugar().Debug(args...)
}

// Debugln logs to the INFO log.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Debugln(args ...interface{}) {
	g_logger.Sugar().Debug(args...)
}

// Debugf logs to the INFO log.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Debugf(format string, args ...interface{}) {
	g_logger.Sugar().Debugf(format, args...)
}

// Info logs to the INFO log.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Info(args ...interface{}) {
	g_logger.Sugar().Info(args...)
}

// Infoln logs to the INFO log.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Infoln(args ...interface{}) {
	g_logger.Sugar().Info(args...)
}

// Infof logs to the INFO log.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Infof(format string, args ...interface{}) {
	g_logger.Sugar().Infof(format, args...)
}

func InfoW(format string, keysAndValues ...interface{}) {
	g_logger.Sugar().Infow(format, keysAndValues...)
}

// Warning logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Warn(args ...interface{}) {
	g_logger.Sugar().Warn(args...)
}

func Warnw(format string, keysAndValues ...interface{}) {
	g_logger.Sugar().Warnw(format, keysAndValues...)
}

// Warningln logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Warnln(args ...interface{}) {
	g_logger.Sugar().Warn(args...)
}

// Warningf logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Warnf(format string, args ...interface{}) {
	g_logger.Sugar().Warnf(format, args...)
}

// Error logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Error(args ...interface{}) {
	g_logger.Sugar().Error(args...)
}

// Errorln logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Errorln(args ...interface{}) {
	g_logger.Sugar().Error(args...)
}

// Errorf logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Errorf(format string, args ...interface{}) {
	g_logger.Sugar().Errorf(format, args...)
}

func Errorw(format string, keysAndValues ...interface{}) {
	g_logger.Sugar().Errorw(format, keysAndValues...)
}

// Fatal logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Fatal(args ...interface{}) {
	g_logger.Sugar().Fatal(args...)
}

// Fatalln logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Fatalln(args ...interface{}) {
	g_logger.Sugar().Fatal(args...)
}

// Fatalf logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Fatalf(format string, args ...interface{}) {
	g_logger.Sugar().Fatalf(format, args...)
}
