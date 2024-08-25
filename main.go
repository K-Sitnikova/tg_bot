package main
import(
	"log"
	"os" 
	"github.com/go-telegram-bot-api/telegram-bot-api"
)


func main() {
	botToken := os.Getenv("TG_TOKEN")
	if botToken == "" {
		log.Fatal("TG_TOKEN environment variable is required")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
        log.Panic(err)
    }

	for update := range updates {
        if update.Message == nil {
            continue
        }

        log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

   
        if update.Message.IsCommand() {
            switch update.Message.Command() {
            case "admin":
                msg := tgbotapi.NewMessage(update.Message.Chat.ID, "@blackberry_fox spam is here")
                bot.Send(msg)
			case "iakov_job":
                msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Иаков, ищи работу")
                bot.Send(msg)
			}
        }
    }
}