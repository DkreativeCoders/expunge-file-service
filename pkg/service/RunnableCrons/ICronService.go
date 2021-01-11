package RunnableCrons

type ICronService interface {
	Execute()
	GetCronTime() string
}

