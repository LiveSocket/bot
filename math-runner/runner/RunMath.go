package runner

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/gammazero/nexus/v3/wamp"
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/LiveSocket/bot/service"
)

// RunMath Performs a calculation on a math message
func RunMath(service *service.Service, message *twitch.PrivateMessage, left string, right string, operator string) {
	// Convert string digits to floats for calculating
	var result float64
	l, err := strconv.ParseFloat(left, 64)
	if err != nil {
		log.Print(err)
	}
	r, err := strconv.ParseFloat(right, 64)
	if err != nil {
		log.Print(err)
	}

	// Check and perform the operation
	switch operator {
	case "+":
		result = l + r
		break
	case "-":
		result = l - r
		break
	case "*":
		result = l * r
	case "x":
		result = l * r
		break
	case "/":
		result = l / r
		break
	case "^":
		result = math.Pow(l, r)
		break
	case "%":
		result = math.Mod(l, r)
		break
	}

	// Speak a result in chat
	service.SimpleCall("private.twitch.chat.say", wamp.List{message.Channel, fmt.Sprintf("%f", result)}, nil)
}
