package miyanbor

import (
	"regexp"

	telegramAPI "gopkg.in/telegram-bot-api.v4"
)

// CallbackFunction will be called to push new update
type CallbackFunction func(*UserSession, interface{})

// FallbackCallbackFunction will be called to push new update
type FallbackCallbackFunction func(*UserSession, *telegramAPI.Update)

// callback contains a CallbackFunction and it's related Regex pattern
type callback struct {
	Pattern  *regexp.Regexp
	Function CallbackFunction
}
