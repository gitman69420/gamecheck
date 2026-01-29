package handlers

import (
	"gamecheck-backend/auth"
	"gamecheck-backend/db"
	"gamecheck-backend/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateAuthSessionResponse struct {
	AccessToken string    `json:"access_token"`
	ExpiresIn   time.Time `json:"expires_in"`
}

const TOKEN_EXPIRY_TIME time.Duration = 45 * time.Minute

// CreateAuthSessionHandler creates a session with the provided Steam details
//
// IMPORTANT: We only call this once the OpenID authentication is externally verified. If not, we shouldn't call this
func (h *Handler) CreateAuthSessionHandler(ctx *gin.Context) {
	// accept the user info from the steam API (steamid, displayname, avatarhash)
	var requestBody auth.SteamUserDetails

	if err := ctx.ShouldBindBodyWithJSON(&requestBody); err != nil {
		utils.GinBadRequest(ctx, "Invalid body structure")
		return
	}

	// create a user record if it doesn't exist
	userId, err := h.q.PotentiallyCreateUserAndGetUserID(ctx, db.PotentiallyCreateUserAndGetUserIDParams{
		Steamid:     requestBody.SteamId,
		Avatarhash:  requestBody.AvatarHash,
		Personaname: requestBody.PersonaName,
	})

	if err != nil {
		h.logger.Printf("Unable to create or retrieve user record: %s", err)
		utils.GinInternalServerError(ctx)
		return
	}

	// create the claims with the user details
	userDetails := auth.UserDetails{
		UserId:      int(userId),
		SteamId:     requestBody.SteamId,
		PersonaName: requestBody.PersonaName,
		AvatarHash:  requestBody.AvatarHash,
	}

	// create a JWT using the claims and the secret
	accessToken, accessTokenExpiry, err := auth.GenerateJWT(userDetails, TOKEN_EXPIRY_TIME)

	if err != nil {
		h.logger.Println("Unable to create access token")
		utils.GinInternalServerError(ctx)
		return
	}

	// return the JWT value
	utils.GinSuccessResponse(ctx, CreateAuthSessionResponse{
		AccessToken: accessToken,
		ExpiresIn:   accessTokenExpiry,
	})
}
