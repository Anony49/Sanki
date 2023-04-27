package config

import (
	"fmt"
	"os"
	"path"
	"runtime"

	log "github.com/sirupsen/logrus"
)

var (
	AllowedUpdates     []string
	ValidLangCodes     []string
	BotToken           string
	DatabaseURI        string
	MainDbName         string
	WebhookURL         string
	BotVersion         string
	ApiServer          string
	RedisUri           string
	RedisPassword      string
	WorkingMode        = "worker"
	Debug              = false
	DropPendingUpdates = true
	EnableWebhook      = false
	WebhookPort        int
	OwnerId            int64
	MessageDump        int64
	LogChannel         int64
	SecretToken        string
)

func init() {
	// set logger config
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
	log.SetFormatter(
		&log.JSONFormatter{
			DisableHTMLEscape: true,
			PrettyPrint:       true,
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				return f.Function, fmt.Sprintf("%s:%d", path.Base(f.File), f.Line)
			},
		},
	)

	// set necessary variables
	Debug = typeConvertor{str: os.Getenv("DEBUG")}.Bool()
	DropPendingUpdates = typeConvertor{str: os.Getenv("DEDROP_PENDING_UPDATESBUG")}.Bool()
	DatabaseURI = os.Getenv("DB_URI")
	MainDbName = os.Getenv("DB_NAME")
	OwnerId = typeConvertor{str: os.Getenv("OWNER_ID", "6190680150")}.Int64()
	MessageDump = typeConvertor{str: os.Getenv("MESSAGE_DUMP", "-1001862524511")}.Int64()
	LogChannel = typeConvertor{str: os.Getenv("LOG_CHANNEL", "-1001862524511")}.Int64()
	WebhookURL = os.Getenv("WEBHOOK_URL")
	BotToken = os.Getenv("BOT_TOKEN")
	BotVersion = os.Getenv("BOT_VERSION")
	ApiServer = os.Getenv("API_SERVER")
	SecretToken = os.Getenv("SECRET_TOKEN")
	EnableWebhook = typeConvertor{str: os.Getenv("USE_WEBHOOKS")}.Bool()
	WebhookPort = typeConvertor{str: os.Getenv("PORT")}.Int()
	AllowedUpdates = typeConvertor{str: os.Getenv("ALLOWED_UPDATES")}.StringArray()
	ValidLangCodes = typeConvertor{str: os.Getenv("ENABLED_LOCALES")}.StringArray()
	RedisUri = os.Getenv("REDIS_URI")
	RedisPassword = os.Getenv("REDIS_PASSWORD")

	// if allowed updates is not set, set it to receive all updates
	if (len(AllowedUpdates) == 1 && AllowedUpdates[0] == "") || (len(AllowedUpdates) == 0) {
		AllowedUpdates = []string{
			"message",
			"edited_message",
			"channel_post",
			"edited_channel_post",
			"inline_query",
			"chosen_inline_result",
			"callback_query",
			"shipping_query",
			"pre_checkout_query",
			"poll",
			"poll_answer",
			"my_chat_member",
			"chat_member",
			"chat_join_request",
		}
	}

	// if valid lang codes is not set, set it to 'en' only
	if (len(ValidLangCodes) == 1 && ValidLangCodes[0] == "") || (len(ValidLangCodes) == 0) {
		ValidLangCodes = []string{"en"}
	}

	// set as default api server if not set
	if ApiServer == "" {
		ApiServer = "https://api.telegram.org"
	}
}
