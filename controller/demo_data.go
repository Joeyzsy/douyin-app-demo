package controller

var DemoVideos = []Video{
	{
		Id:      1,
		Author:  DemoUser,
		PlayUrl: "http://localhost:80/bear.mp4",
		//PlayUrl:       "https://joy1.videvo.net/videvo_files/video/free/video0477/large_watermarked/_import_624e6fe259ac93.74720681_preview.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
}

var DemoComments = []Comment{
	{
		Id:         1,
		User:       DemoUser,
		Content:    "Test Comment",
		CreateDate: "05-01",
	},
}

var DemoUser = User{
	Id:            1,
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}
