package callback

import (
	"bc-labs-lunch-bot/di"
	"bc-labs-lunch-bot/internal/callback/service"
	"bc-labs-lunch-bot/internal/config"
	"encoding/json"
	"github.com/slack-go/slack"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	callbackService service.CallbackService
)

func init() {
	callbackService = di.InitializeCallbackServiceImpl()
}

func CallBack(c *gin.Context) {
	var i slack.InteractionCallback

	err := json.Unmarshal([]byte(c.PostForm("payload")), &i)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	switch i.View.PrivateMetadata {
	case "/slash":
		log.Println("check slash")
		err = callbackService.SaveLunchPayment(i, config.GetAccessBotToken())
	default:
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.String(http.StatusAccepted, "")
}
