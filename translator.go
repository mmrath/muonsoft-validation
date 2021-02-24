package validation

import (
	"fmt"

	"github.com/muonsoft/validation/message/translations"
	"github.com/muonsoft/validation/message/translations/english"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type Translator struct {
	fallbackLanguage language.Tag
	defaultLanguage  language.Tag
	messages         *catalog.Builder
	printers         map[language.Tag]*message.Printer
}

func (translator *Translator) init() error {
	if translator.defaultLanguage == language.Und {
		translator.defaultLanguage = language.English
	}

	err := translator.loadMessages(english.Messages)
	if err != nil {
		return err
	}

	return nil
}

func (translator *Translator) loadMessages(messages map[language.Tag]map[string]catalog.Message) error {
	if translator.messages == nil {
		translator.messages = catalog.NewBuilder()
	}
	if translator.printers == nil {
		translator.printers = make(map[language.Tag]*message.Printer, len(messages))
	}

	err := translations.LoadMessages(translator.messages, messages)
	if err != nil {
		return fmt.Errorf("failed to load translations: %w", err)
	}

	for tag := range messages {
		translator.printers[tag] = message.NewPrinter(tag, message.Catalog(translator.messages))
	}

	return nil
}

func (translator *Translator) translate(tag language.Tag, msg string, pluralCount int) string {
	if tag == language.Und {
		tag = translator.defaultLanguage
	}
	printer := translator.printers[tag]
	if printer == nil {
		printer = translator.printers[translator.fallbackLanguage]
	}

	return printer.Sprintf(msg, pluralCount)
}