package interfaces

type IHealthCheckService interface {
	HealthCheckService() (string, error)
}

type IHealthCheckRepository interface {
	HealthCheckRepository() (string, error)
}
