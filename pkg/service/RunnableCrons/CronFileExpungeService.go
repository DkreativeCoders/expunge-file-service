package RunnableCrons

import (
	"github.com/DkreativeCoders/expunge-file-service/pkg/service/Expunge"
	"log"
	"os"
)

type CronFileExpungeService struct {
	fileExpunge Expunge.IFileExpunge
}

func (c CronFileExpungeService) Execute() {
	enableCronInterval := os.Getenv("ENABLE.CRONJOB")
	if enableCronInterval == "true"{
		log.Println("About Running CronFileExpungeService")
		c.fileExpunge.ExecuteDeleteTask()
		log.Println("Done Running CronFileExpungeService")
	}
}

func (c CronFileExpungeService) GetCronTime() string {
	cronInterval := os.Getenv("CRON.INTERVAL")
	return cronInterval
	//return "@every 0h0m30s"
}
