package wishlist

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	expectedWishlistContentParts = 2
	errInvalidURL                = errors.New(`mandame un link válido bob@ hp, le estoy hablando claro 😡😡😡`)
	errInvalidWishlistCommand    = errors.New(`no entiendo tu wishlist, usa el formato correcto bob@ hp --> /wishlist <item> <link>`)

	emptyWishlistMessage            = `no tienes nada en tu wishlist, es momento de dejarse llevar por la fiebre del billete 🤑🥵\n\nusate /wishlist <item> <link>`
	successWishlistItemAddedMessage = `Se agregó '%s' a la wishlist correctamente 🤑`
	wishlistFinalMessage            = `Weno mi fafá, tu wishlist es la siguiente: \n\n`
)

// HandleWishlist handles the user wishlist command
func HandleWishlist(client *shared.Client, chatID int64, userName, content string) error {
	contentParts := strings.Split(content, " ")
	if len(contentParts) != expectedWishlistContentParts {
		return errInvalidWishlistCommand
	}

	wishlistItem := contentParts[0]
	wishlistLink := contentParts[1]

	if wishlistItem == "" || wishlistLink == "" {
		return errInvalidWishlistCommand
	}

	if !shared.IsValidURL(wishlistLink) {
		return errInvalidURL
	}

	now, err := shared.GetNow()
	if err != nil {
		return err
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
		return client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	msg := fmt.Sprintf(successWishlistItemAddedMessage, wishlistItem)

	return client.SendMessage(chatID, msg)
}

// GetWishlist returns the user wishlist
func GetWishlist(client *shared.Client, userName string, chatID int64) error {
	wishlistItems, err := storage.GetActivityHistory(userName, shared.Wishlist)
	if errors.Is(err, storage.ErrNoActivitiesFound) {
		return client.SendMessage(chatID, emptyWishlistMessage)
	}

	if err != nil {
		return client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	wishlist := wishlistFinalMessage

	for _, item := range wishlistItems {
		wishlist += fmt.Sprintf("- %s 🤑\n", item.Content)
	}

	return client.SendMessage(chatID, wishlist)
}
