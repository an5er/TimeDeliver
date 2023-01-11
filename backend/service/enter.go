package service

import "timedeliver/service/deliver"

type ServiceGroup struct {
	DeliverServiceGroup deliver.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
