package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yash-sojitra-20/address-book-backend/internal/repositories"
	"github.com/yash-sojitra-20/address-book-backend/internal/utils"
)

func EnsureUserExistsMiddleware(userRepo *repositories.UserRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.GetUint("user_id") // get from JWT middleware
		if userID == 0 {
			utils.Error(ctx, http.StatusUnauthorized, "user not found in context")
			ctx.Abort()
			return
		}

		exists, err := userRepo.ExistsByID(userID)
		if err != nil || !exists {
			utils.Error(ctx, http.StatusUnauthorized, "user does not exist")
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
