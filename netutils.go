package netutils

import "github.com/gadget-bot/gadget/router"

// GetMentionRoutes is used to retrieve all Mention Routes from this plugin
func GetMentionRoutes() []router.MentionRoute {
	return []router.MentionRoute{
		*runHTTPPing(),
		*queryWhois(),
	}
}
