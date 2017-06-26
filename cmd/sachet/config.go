package main

import (
	"io/ioutil"

	"github.com/messagebird/sachet/provider/cm"
	"github.com/messagebird/sachet/provider/exotel"
	"github.com/messagebird/sachet/provider/infobip"
	"github.com/messagebird/sachet/provider/messagebird"
	"github.com/messagebird/sachet/provider/nexmo"
	"github.com/messagebird/sachet/provider/telegram"
	"github.com/messagebird/sachet/provider/twilio"

	"gopkg.in/yaml.v2"
)

type ReceiverConf struct {
	Name     string
	Provider string
	To       []string
	From     string
}

var config struct {
	Providers struct {
		MessageBird messagebird.MessageBirdConfig
		Nexmo       nexmo.NexmoConfig
		Twilio      twilio.TwilioConfig
		Infobip     infobip.InfobipConfig
		Exotel      exotel.ExotelConfig
		CM          cm.CMConfig
		Telegram    telegram.TelegramConfig
	}

	Receivers []ReceiverConf
}

// LoadConfig loads the specified YAML configuration file.
func LoadConfig(filename string) error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return err
	}

	return nil
}
