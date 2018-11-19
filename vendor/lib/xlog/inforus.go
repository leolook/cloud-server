package xlog

//
//import (
//	"runtime"
//	"sync"
//
//	log "github.com/sirupsen/logrus"
//	"path"
//)
//
//const skipFrameCnt = 9
//
//// AddHookDefault is ...
//func AddHookDefault() {
//	log.AddHook(Hook{
//		mu:       &sync.Mutex{},
//		file:     true,
//		line:     true,
//		function: true,
//		levels:   log.AllLevels,
//	})
//}
//
//// AddHook is ...
//func AddHook(file, line, function bool, levels []log.Level) {
//	log.AddHook(Hook{
//		mu:       &sync.Mutex{},
//		file:     file,
//		line:     line,
//		function: function,
//		levels:   levels,
//	})
//}
//
//// Hook is ...
//type Hook struct {
//	mu       *sync.Mutex
//	file     bool
//	line     bool
//	function bool
//	levels   []log.Level
//}
//
//// Levels is ...
//func (h Hook) Levels() []log.Level {
//	return h.levels
//}
//
//// Fire is ...
//func (h Hook) Fire(entry *log.Entry) error {
//	//pc := make([]uintptr, 64)
//	//cnt := runtime.Callers(skipFrameCnt, pc)
//	_, file, line, ok := runtime.Caller(9);
//
//	if !ok {
//		return nil;
//	}
//
//	//fu := runtime.FuncForPC(pc)
//	//name := fu.Name()
//	if h.file {
//		h.mu.Lock()
//		entry.Data["file"] = path.Base(file)
//		//entry.Data["file"] = file;
//		h.mu.Unlock()
//	}
//
//	/*
//	if h.function {
//		h.mu.Lock()
//		entry.Data["func"] = path.Base(name)
//		h.mu.Unlock()
//
//	}
//	*/
//
//	if h.line {
//		h.mu.Lock()
//		entry.Data["line"] = line
//		h.mu.Unlock()
//	}
//	return nil
//}
