package service

import (
	"time"

	"github.com/LiveSocket/bot/conv"
	"github.com/gempir/go-twitch-irc/v2"
)

// MapToPrivateMessage Parses a map into a twitch.PrivateMessage
func MapToPrivateMessage(message map[string]interface{}) *twitch.PrivateMessage {
	user := message["User"].(map[string]interface{})
	bits, _ := conv.ToInt(message["Bits"])
	action, _ := conv.ToBool(message["Action"])
	return &twitch.PrivateMessage{
		User: twitch.User{
			ID:          conv.ToString(user["ID"]),
			Name:        conv.ToString(user["Name"]),
			DisplayName: conv.ToString(user["DisplayName"]),
			Color:       conv.ToString(user["Color"]),
			Badges:      badges(user),
		},
		Raw:     conv.ToString(message["Raw"]),
		RawType: conv.ToString(message["RawType"]),
		Tags:    tags(message),
		Type:    _type(message),
		Time:    _time(message),
		Message: conv.ToString(message["Message"]),
		Channel: conv.ToString(message["Channel"]),
		RoomID:  conv.ToString(message["RoomID"]),
		ID:      conv.ToString(message["ID"]),
		Bits:    bits,
		Action:  action,
		Emotes:  emotes(message),
	}
}

func badges(user map[string]interface{}) map[string]int {
	result := map[string]int{}
	if user["Badges"] != nil {
		for key, value := range user["Badges"].(map[string]interface{}) {
			v, err := conv.ToInt(value)
			if err != nil {
				v = 0
			}
			result[key] = v
		}
	}
	return result
}

func _type(message map[string]interface{}) twitch.MessageType {
	if message["MessageType"] != nil {
		return message["MessageType"].(twitch.MessageType)
	}
	return 0
}

func tags(message map[string]interface{}) map[string]string {
	result := map[string]string{}
	if message["Tags"] != nil {
		for key, value := range message["Tags"].(map[string]interface{}) {
			result[key] = conv.ToString(value)
		}
	}
	return result
}

func _time(message map[string]interface{}) time.Time {
	result := time.Unix(0, 0)
	if message["Time"] != nil {
		t, err := time.Parse(time.RFC3339, message["Time"].(string))
		if err == nil {
			result = t
		}
	}
	return result
}

func emotes(message map[string]interface{}) []*twitch.Emote {
	result := []*twitch.Emote{}
	if message["Emotes"] != nil {
		for _, emote := range message["Emotes"].([]interface{}) {
			item := emote.(map[string]interface{})
			count, err := conv.ToInt(item["Count"])
			if err != nil {
				count = 0
			}
			result = append(result, &twitch.Emote{
				Name:  conv.ToString(item["Name"]),
				ID:    conv.ToString(item["ID"]),
				Count: count,
			})
		}
	}
	return result
}
