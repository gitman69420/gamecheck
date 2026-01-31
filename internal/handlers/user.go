package handlers

import (
	"gamecheck-backend/utils"
	"html"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	SteamId         string `json:"steamid"`
	PersonaName     string `json:"personaname"`
	AvatarUrl       string `json:"avatar_url"`
	AvatarUrlMedium string `json:"avatar_url_medium"`
	AvatarUrlFull   string `json:"avatar_url_full"`
}

func (h *Handler) GetMyInfoHandler(ctx *gin.Context) {

	userId := ctx.GetInt("user_id")

	if userId == 0 {
		utils.GinUnauthorizedResponse(ctx)
	}

	// TODO: Should we cache this DB call?
	// But this query would ideally get called 'once per user per session' on the frontend
	// So we can leave it be for now

	myInfo, err := h.q.GetUserInfo(ctx, int32(userId))

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	avatarUrls := utils.AvatarUrlsFromHash(myInfo.Avatarhash)

	response := UserResponse{
		SteamId:         myInfo.Steamid,
		PersonaName:     html.EscapeString(myInfo.Personaname),
		AvatarUrl:       avatarUrls.AvatarUrl,
		AvatarUrlMedium: avatarUrls.AvatarUrlMedium,
		AvatarUrlFull:   avatarUrls.AvatarUrlFull,
	}

	utils.GinSuccessResponse(ctx, response)
}
