package log

import "go.uber.org/zap"

var Logger *zap.Logger

func InitLogger() {
	var err error
	Logger, err = NewLogger()
	if err != nil {
		panic(err)
	}
}

func NewLogger() (*zap.Logger, error) {
	pro := zap.NewProductionConfig()
	pro.OutputPaths = append(pro.OutputPaths, "D:\\goProject\\src\\gologin\\log\\login.log")
	return pro.Build()
}
