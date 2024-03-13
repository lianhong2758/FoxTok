package sq

import (
	"FoxTok/server/template"
)

func GetVido(id int64) *template.Video {
	return nil
}

func GetUser(id int64) *template.User {
	return nil
}

func GetNextVideo(lastTime int64) *template.Video {
	return &template.Video{
		Id:       1,
		PlayUrl:  GetFileUri() + `D:/CSGO/6ed6ebbe7edf9331b7624e870c9a5142.mp4`,
		CoverUrl: GetFileUri() + "",
		Title:    "CSGO",
		FavoriteCount: 11,
		CommentCount: 11,
		IsFavorite: false,
		Author:&template.User{
			Id: 1,
			Name: "test",
		},
	}
}
