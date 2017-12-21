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
	assertEqual(t, "gif", InlineQueryResultGif{}.getType())
	assertEqual(t, "mpeg4_gif", InlineQueryResultMpeg4Gif{}.getType())
	assertEqual(t, "video", InlineQueryResultVideo{}.getType())
	assertEqual(t, "audio", InlineQueryResultAudio{}.getType())
	assertEqual(t, "voice", InlineQueryResultVoice{}.getType())
	assertEqual(t, "document", InlineQueryResultDocument{}.getType())
	assertEqual(t, "location", InlineQueryResultLocation{}.getType())
	assertEqual(t, "venue", InlineQueryResultVenue{}.getType())
	assertEqual(t, "contact", InlineQueryResultContact{}.getType())
}

func TestDummyInputMessageContent(t *testing.T) {
	InputTextMessageContent{}._dummyInputMessageContent()
	InputLocationMessageContent{}._dummyInputMessageContent()
	InputVenueMessageContent{}._dummyInputMessageContent()
	InputContactMessageContent{}._dummyInputMessageContent()
}
