package config

type LogConfig struct {
	EventLogFileName string `yaml:"event_log_file_name"`
	TraceLogFileName string `yaml:"trace_log_file_name"`
}

var Log LogConfig

func initLog() {
	_ = initConf(BasePath+"log_conf.yml", &Log)
}
