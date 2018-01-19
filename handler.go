package miyanbor

import (
	cache "github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	telegramAPI "gopkg.in/telegram-bot-api.v4"
)

var (
	callbacks                []callback
	fallbackCallbackFunction FallbackCallbackFunction

	usersSessionCache *cache.Cache
)

func (b *Bot) handleNewUpdate(update *telegramAPI.Update) {
	// Get UserInfo
	_, _, userSession := getUserInfo(update)

	// Find and call callback function
	if update.CallbackQuery != nil {
		for _, callback := range callbacks {
			if matches := callback.Pattern.FindStringSubmatch(update.CallbackQuery.Data); matches != nil {
				callback.Function(userSession, matches)
				return
			}
		}
	} else if update.Message != nil {
		userSession.messageCallback(userSession, update.Message.Text)
		return
	} else {
		logrus.Errorf("Unknown update")
	}

	// Call fallback callback function
	fallbackCallbackFunction(userSession, update)
}

func getUserInfo(update *telegramAPI.Update) (userID int, chatID int64, userSession *UserSession) {
	// Get userID
	if update.CallbackQuery != nil {
		chatID = update.CallbackQuery.Message.Chat.ID
		userID = update.CallbackQuery.From.ID
	} else if update.Message != nil {
		chatID = update.Message.Chat.ID
		userID = update.Message.From.ID
	} else {
		logrus.Errorf("can't get userID/chatID")
		return
	}

	// Find userSession
	userSessionInterface, found := usersSessionCache.Get(getUserSessionKey(userID, chatID))
	if !found {
		userSession = &UserSession{
			chatID: chatID,
			userID: userID,
		}
		usersSessionCache.Add(getUserSessionKey(userID, chatID), userSession, cache.DefaultExpiration)
	} else {
		userSession = userSessionInterface.(*UserSession)
	}
	return
}
