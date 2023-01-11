package v1

import "timedeliver/api/v1/deliver"

type ApiGroup struct {
	DeliverApiGroup deliver.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
