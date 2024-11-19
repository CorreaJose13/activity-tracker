package track

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	expectedWishlistContentParts  = 2
	invalidURLMessage             = `mandame un link válido bob@ hp, le estoy hablando claro 😡😡😡`
	invalidWishlistCommandMessage = `no entiendo tu wishlist, usa el formato correcto bob@ hp --> /track wishlist <item> <link>`
)

// SendTrackWishlist tracks the wishlist activity
func SendTrackWishlist(bot *shared.Bot, userName, content string, chatID int64) error {
	if content == "" {
		return getWishlist(bot, userName, chatID)
	}

	contentParts := strings.Split(content, " ")
	if len(contentParts) != expectedWishlistContentParts {
		return shared.SendMessage(bot, chatID, invalidWishlistCommandMessage)
	}

	wishlistItem := contentParts[0]
	wishlistLink := contentParts[1]

	if wishlistItem == "" || wishlistLink == "" {
		return shared.SendMessage(bot, chatID, invalidWishlistCommandMessage)
	}

	if !shared.IsValidURL(wishlistLink) {
		return shared.SendMessage(bot, chatID, invalidURLMessage)
	}

	now, err := shared.GetNow()
	if err != nil {
		return shared.SendMessage(bot, chatID, err.Error())
	}

	nowStr := now.Format(time.RFC3339)

	userActivity := shared.UserActivity{
		ID:        shared.GenerateActivityItemID(now, userName, shared.Wishlist),
		Name:      userName,
		Activity:  shared.Wishlist,
		CreatedAt: nowStr,
		Content:   content,
	}

	err = storage.Create(userActivity)
	if err != nil {
		return shared.SendMessage(bot, chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	msg := fmt.Sprintf("Tu %s se agregó a la wishlist correctamente 🤑", wishlistItem)

	return shared.SendMessage(bot, chatID, msg)
}

func getWishlist(bot *shared.Bot, userName string, chatID int64) error {
	wishlistItems, err := storage.GetActivityHistory(userName, shared.Wishlist)
	if errors.Is(err, storage.ErrNoActivitiesFound) {
		return shared.SendMessage(bot, chatID, "no tienes nada en tu wishlist, es momento de dejarse llevar por la fiebre del billete 🤑🥵\n\nusate /track wishlist <item> <link>")
	}

	if err != nil {
		return shared.SendMessage(bot, chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	wishlist := "Weno mi fafá, tu wishlist es la siguiente: \n\n"

	for _, item := range wishlistItems {
		wishlist += fmt.Sprintf("- %s 🤑\n", item.Content)
	}

	return shared.SendMessage(bot, chatID, wishlist)
}
