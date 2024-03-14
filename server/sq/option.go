package sq

import (
	"FoxTok/server/template"

	"google.golang.org/protobuf/proto"
)

func GetVido(id int64) *template.Video {
	return nil
}

func GetUser(id int64) *template.User {
	return nil
}

func GetNextVideo(lastTime int64, UserID int64) (*template.Video, int64) {
	var v Video
	db.Where("createdat > ?", lastTime).Limit(1).Find(&v)
	if v.ID == 0 {
		db.Take(&v)
	}
	au := new(template.User)
	us := new(User)
	db.Where("id = ?", v.UserID).Find(&us)

	au.Id = us.ID
	au.Name = us.Name
	au.Avatar = proto.String(us.Avatar)
	au.FollowCount = proto.Int64(us.Follower)
	au.FollowerCount = proto.Int64(us.Fans)
	au.BackgroundImage = proto.String(us.Background)
	au.Signature = proto.String(us.Signature)
	au.TotalFavorited = proto.Int64(us.Favorited)
	au.FavoriteCount = proto.Int64(us.FavoriteCount)
	au.WorkCount = proto.Int64(us.WorkNum)

	return &template.Video{
		Id:            v.ID,
		PlayUrl:       GetFileUri() + v.VideoPath,
		CoverUrl:      GetFileUri() + v.CoverPath,
		Title:         v.Title,
		FavoriteCount: v.Favorite,
		CommentCount:  v.Comment,
		IsFavorite: func() bool {
			if UserID != 0 {
				var IsFavorite IsFavorite
				db.Where("id = ?", UserID).Limit(1).Find(&IsFavorite)
				return IsFavorite.IsFavorite
			}
			return false
		}(),
		Author: au,
	}, v.CreatedAt.Unix()
}
