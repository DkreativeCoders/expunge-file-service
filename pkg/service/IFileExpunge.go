package service

type IFileExpunger interface {
	executeDeleteTask() ()
}

