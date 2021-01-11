package RunnableCrons

import (
	"github.com/DkreativeCoders/expunge-file-service/pkg/service/Expunge"
)

type CronFactory struct {
	CronService                                []ICronService
}

func NewCronFactory( fileExpunge Expunge.IFileExpunge) *CronFactory {
	cronService:= buildProcessFileFactory(fileExpunge)
	return &CronFactory{CronService: cronService }
}



func buildProcessFileFactory(fileExpunge Expunge.IFileExpunge) []ICronService {
	//fileExpungeServiceObject :=object[0]
	//fileExpungeService, ok:= fileExpungeServiceObject.(Expunge.FileExpunge)
	//if ok {
	//	log.Fatal("Unable to TypeCast FileExpungeService")
	//}
	//fileExpunge.ExecuteDeleteTask()
	cronService := make([]ICronService, 0) // same as []int{0, 0}
	cronService = append(cronService, CronFileExpungeService{fileExpunge: fileExpunge})
	return cronService

}

