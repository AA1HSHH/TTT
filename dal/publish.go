package dal

func AddVideo(video DBVideo) error {
	return db.Create(&video).Error
}

func QueryVideosByUserId(user_id int64) ([]DBVideo, error) {

	var data []DBVideo
	result := db.Where(&DBVideo{AuthorId: user_id}).Find(&data)
	//如果出现问题，返回对应到空，并且返回error
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}
