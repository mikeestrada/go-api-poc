package model

type HealthCheck struct {
	Alive bool `json:"alive"`
}

func (healthCheck *HealthCheck) status() *HealthCheck {
	// do validation that all the necessary dependencies are up, for example database, service-b, etc
	healthCheck.Alive = true
	return healthCheck
}