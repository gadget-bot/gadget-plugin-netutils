package netutils

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gadget-bot/gadget/router"
	"github.com/likexian/whois"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

func queryWhois() *router.MentionRoute {
	var pluginRoute router.MentionRoute
	pluginRoute.Permissions = append(pluginRoute.Permissions, "*")
	pluginRoute.Name = "network_utils.queryWhois"
	pluginRoute.Pattern = `(?i)^whois <?([^>]+)>?$`
	pluginRoute.Description = "Looks up WHOIS info for a given domain, IP, or ASN"
	pluginRoute.Help = "whois <DOMAIN|IP|ASN>"
	pluginRoute.Plugin = func(router router.Router, route router.Route, api slack.Client, ev slackevents.AppMentionEvent, message string) {
		// Here's how we can react to the message
		msgRef := slack.NewRefToMessage(ev.Channel, ev.TimeStamp)
		api.AddReaction("male-detective", msgRef)

		re := regexp.MustCompile(pluginRoute.Pattern)
		results := re.FindStringSubmatch(message)
		input := results[1]

		names := strings.Split(input, "|")

		var replyTS string

		if len(names) > 1 {
			input = names[1]
		} else {
			input = names[0]
		}

		result, err := whois.Whois(input)
		if err != nil {
			result = fmt.Sprintf("Something went wrong looking up WHOIS info for '%s': %s", input, err)
		}

		if ev.ThreadTimeStamp != "" {
			replyTS = ev.ThreadTimeStamp
		} else {
			replyTS = ev.TimeStamp
		}

		api.PostMessage(
			ev.Channel,
			slack.MsgOptionText(
				fmt.Sprintf("```%s```\n", result),
				false,
			),
			slack.MsgOptionTS(replyTS),
		)
	}

	// We've got to return the MentionRoute
	return &pluginRoute
}
