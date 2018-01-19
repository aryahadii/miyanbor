package miyanbor

import (
	"regexp"

	"github.com/pkg/errors"
)

func (b *Bot) SetFallbackCallbackHandler(function FallbackCallbackFunction) {
	fallbackCallbackFunction = function
}

func (b *Bot) AddCallbackHandler(callbackQueryPattern string, function CallbackFunction) error {
	// Compile pattern
	regexPattern, err := regexp.Compile(callbackQueryPattern)
	if err != nil {
		return errors.Wrap(err, "can't compile pattern")
	}

	// Add to callbacks list
	callbacks = append(callbacks, callback{
		Pattern:  regexPattern,
		Function: function,
	})
	return nil
}
