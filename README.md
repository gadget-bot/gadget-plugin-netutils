# Gadget Network Utilities Plugin

This is a [Gadget](https://github.com/gadget-bot/gadget) plugin that allows Gadget to perform various network utility functions.

Note that Gadget -- and therefore any plugin for Gadget -- is still very much a **work in progress**, so please don't use it in production yet (or if you do, don't complain).

## How do I use this plugin?

In your `main.go` for using Gadget, your `main()` function should instruct you bot, which is created using `gadget.Setup()` needs to be instructed to `Router.AddMentionRoutes(netutils.GetMentionRoutes())`. This _usually_ looks like this:

```golang
package main

import (
	gadget "github.com/gadget-bot/gadget/core"
	netutils "github.com/gadget-bot/gadget-plugin-netutils"
)

func main() {
	// This is the Gadget bot
	myBot := gadget.Setup()

	// Add your custom plugins here
	myBot.Router.AddMentionRoutes(netutils.GetMentionRoutes())

	// This launches your bot
	myBot.Run()
}
```

## What can this plugin do?

In your chat, you can perform the following actions:

* `whois <DOMAIN|IP|ASN>` - Queries public [WHOIS](https://en.wikipedia.org/wiki/WHOIS) for the given object.
  * Parameters:
	  - `<DOMAIN|IP|ASN>` either a domain name, an IP address, or an [ASN](https://www.arin.net/resources/guide/asn/).
* `hping [get|post|head] URL [COUNT] [INTERVAL(s|ms)]` - Performs an HTTP request to the given `URL` with some optional parameters.
	* Parameters:
	  - `[get|post|head]` - _optional_ Allows specificy the HTTP Request Type (default: `get`)
		- `URL` - The HTTP or HTTPS URL
		- `[COUNT]` - _optional_ The number of times to attempt the "ping" (default: `3`)
		- `[INTERVAL(s|ms)]` - _optional_ The time to wait, in either seconds (`s`) or milliseconds (`ms`), between requests (default: `2s`)
	* Notes:
	  - A big shout out to [https://github.com/mehrdadrad/mylg](https://github.com/mehrdadrad/mylg), from which most of the `hping` feature's work is derived
