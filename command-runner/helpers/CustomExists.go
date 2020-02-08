package helpers

import (
	"github.com/LiveSocket/bot/command-runner/models"
	"github.com/LiveSocket/bot/service"
)

// CustomExists Checks if a custom command exists for the channel with the same name
func CustomExists(service *service.Service, channel string, name string) (bool, error) {
	custom, err := models.GetCustom(service, name, channel)
	if err != nil {
		return false, err
	}
	return custom != nil, nil
}
