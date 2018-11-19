package xlog

//var g_log_level int = int(logrus.DebugLevel)

//func init() {
//flag.IntVar(&g_log_level, "log_level", 5, "log level")
//module.Define("xlog", xlog_init, nil)
//}

/*
func xlog_init() error {
	AddHookDefault()

	//log.SetFlags(log.Ldate | log.Lmicroseconds |log.Lshortfile);
	format := &TextFormatter {
		TimestampFormat: "20060102 15:04:05.999",
	}
	logrus.SetFormatter(format);

	logf, err := rotatelogs.New(
		g_log_file + ".%Y%m%d",
		//rotatelogs.WithMaxAge(-1),
		//rotatelogs.WithRotationCount(7),
	)

	if err != nil {
		log.Fatal("failed create rotatelogs: %s", err)
		return err
	}

	logrus.SetOutput(io.MultiWriter(logf, os.Stderr))
	logrus.SetLevel(logrus.Level(g_log_level))

	return nil
}
*/
