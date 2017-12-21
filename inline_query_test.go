package tgbot

import (
	"fmt"
	"testing"
)

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if actual == expected {
		return
	}

	errorMessage := fmt.Sprintf("Error:\nExpected - %v\nActual - %v", expected, actual)
	t.Fatal(errorMessage)
}

func TestInlineQueryResult(t *testing.T) {
	assertEqual(t, "article", InlineQueryResultArticle{}.getType())
	assertEqual(t, "photo", InlineQueryResultPhoto{}.getType())
}

func TestDummyMethods(t *testing.T) {
	InputTextMessageContent{}._dummyInputMessageContent()
	InputLocationMessageContent{}._dummyInputMessageContent()
	InputVenueMessageContent{}._dummyInputMessageContent()
	InputContactMessageContent{}._dummyInputMessageContent()
}
