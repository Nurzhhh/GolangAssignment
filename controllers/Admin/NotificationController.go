package Admin

import (
	"GolangProject/services"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
)

type NotificationController interface {
	Send(context *gin.Context)
}

type notificationController struct {
	userService services.UserService
	jwtService  services.JWTService
}

// NewAuthController creates a new instance of AuthController
func NewNotificationController(userService services.UserService, jwtService services.JWTService) NotificationController {
	return &notificationController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *notificationController) Send(context *gin.Context) {
	data := url.Values{
		"to":    {"ExponentPushToken[-8l9PRNaTJ1_gW58FeVqXd]"},
		"sound": {"default"},
		"title": {"qwertyu"},
		"body":  {"test"},
	}
	//log.Println(data)

	resp, err := http.PostForm("https://exp.host/--/api/v2/push/send/", data)

	if err != nil {
		log.Fatal(err)
	}

	//response := helpers.BuildResponse(true, "OK", resp)
	//context.JSON(http.StatusCreated, response)

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)
	context.JSON(http.StatusOK, res)
}
