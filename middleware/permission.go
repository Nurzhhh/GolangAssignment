package middleware

import (
	"GolangProject/controllers"
	"GolangProject/services"
	"github.com/gin-gonic/gin"
)

func CheckPermission(check services.PermissionService, u *controllers.UserController) gin.HandlerFunc {
	return func(context *gin.Context) {
		//authHeader := context.GetHeader("Authorization")
		//userID := u.getUserIDByToken(authHeader)
		//convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		//authHeader := context.GetHeader("Authorization")
		//token, err := context.jwtService.ValidateToken(authHeader)
		//if err != nil {
		//	panic(err.Error())
		//}
		//claims := token.Claims.(jwt.MapClaims)
		//id := fmt.Sprintf("%v", claims["user_id"])
		//user := c.userService.Profile(id)
	}
}
