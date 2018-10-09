package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/nlopes/slack"
)

// SlackListener is
type SlackListener struct {
	client    *slack.Client
	botID     string
	channelID string
}

// ListenAndResponse listens slack events and response
// particular messages. It replies by slack message button.
func (s *SlackListener) ListenAndResponse() {
	rtm := s.client.NewRTM()

	// Start listening slack events
	go rtm.ManageConnection()

	// Handle slack events
	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			if err := s.handleMessageEvent(ev); err != nil {
				log.Printf("[ERROR] Failed to handle message: %s", err)
			}
		}
	}
}

// handleMesageEvent handles message events.
func (s *SlackListener) handleMessageEvent(ev *slack.MessageEvent) error {
	// Only response in specific channel. Ignore else.
	if ev.Channel != s.channelID {
		log.Printf("%s %s", ev.Channel, ev.Msg.Text)
		return nil
	}

	// ignore converted text
	if strings.Contains(ev.Msg.Text, "|sup_id=") {
		return nil
	}

	supIDs := getSupID(ev.Msg.Text)

	if supIDs == nil {
		return nil
	}

	text := makeSupIDStr(supIDs)

	if _, _, err := s.client.PostMessage(
		ev.Channel,
		text,
		slack.PostMessageParameters{},
	); err != nil {
		return fmt.Errorf("failed to post message: %s", err)
	}

	return nil
}
func makeSupIDStr(supIDs [][]string) (res string) {

	for _, supArr := range supIDs {
		res = res + "<https://opendb.middle.nec.co.jp/intora/oracle/#/sup/" + supArr[1] + "|sup_id=" + supArr[1] + ">\n"
	}
	return res
}

func getSupID(text string) [][]string {

	r, _ := regexp.Compile("sup_id=([0-9]+)")

	res := r.FindAllStringSubmatch(text, -1)
	fmt.Println(res)

	if res == nil {
		return nil
	}
	return res
}
