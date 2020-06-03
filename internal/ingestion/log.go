package ingestion

import (
	"fmt"
	"go.uber.org/zap"
)

var logger *zap.Logger

func InitLog(){
	var err error
	log, err := zap.NewDevelopment()
	if err != nil {
		fmt.Errorf("can't initialize zap logger: %v", err)
	}
	defer log.Sync()
	logger = log
	log.Info("Zap Logger Started")
}