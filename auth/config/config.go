package config

type Config struct {
	ZapLogger ZapLogger `yaml:"zap_logger"`
}

type ZapLogger struct {
	Test             bool     `yaml:"test"`
	Level            string   `yaml:"level"`
	OutputPaths      []string `yaml:"output_paths"`
	ErrorOutputPaths []string `yaml:"error_output_paths"`
}

func NewConfig() {
}
