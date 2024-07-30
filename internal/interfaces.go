package internal

import "github.com/AHTI6IOTIK/anti-bruteforce/internal/models"

type UseCases interface {
	AllowRequest(request models.Request) bool
	ClearBucket(request models.Request)
	Add(subnet string, list string)
	Remove(subnet string, list string)
	CheckSubnetInList(subnet string) error
}
