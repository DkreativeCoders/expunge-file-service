package RunnableCrons

import (
	"github.com/DkreativeCoders/expunge-file-service/pkg/service/Expunge"
	"github.com/kpango/glg"
	"os"
)

type CronFileExpungeService struct {
	fileExpunge Expunge.IFileExpunge
}

func (c CronFileExpungeService) Execute() {
	enableCronInterval := os.Getenv("ENABLE.CRONJOB")
	if enableCronInterval == "true"{
		glg.Log("About Running CronFileExpungeService")
		c.fileExpunge.ExecuteDeleteTask()
		glg.Log("Done Running CronFileExpungeService")
	}
}

func (c CronFileExpungeService) GetCronTime() string {
	cronInterval := os.Getenv("CRON.INTERVAL")
	return cronInterval
	//return "@every 0h0m30s"
}
