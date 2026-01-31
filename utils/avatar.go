package utils

import (
	"fmt"
)

type AvatarUrls struct {
	AvatarUrl       string
	AvatarUrlMedium string
	AvatarUrlFull   string
}

func AvatarUrlsFromHash(h string) AvatarUrls {
	return AvatarUrls{
		AvatarUrl:       fmt.Sprintf("https://avatars.steamstatic.com/%s.jpg", h),
		AvatarUrlMedium: fmt.Sprintf("https://avatars.steamstatic.com/%s_medium.jpg", h),
		AvatarUrlFull:   fmt.Sprintf("https://avatars.steamstatic.com/%s_full.jpg", h),
	}
}
