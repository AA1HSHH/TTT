package dal

import (
	"fmt"
	"github.com/AA1HSHH/TTT/model/ws"
	"sort"
	"time"
)

type SendSortMsg struct {
	Content  string `json:"content"`
	Read     uint   `json:"read"`
	CreateAt int64  `json:"create_at"`
}

func InsertMsg(database string, id string, content string, read uint, expire int64) (err error) {
	comment := ws.Trainer{
		Content:   content,
		StartTime: time.Now().Unix(),
		EndTime:   time.Now().Unix() + expire,
		Read:      read,
	}
	rst := db.Create(&comment)
	return rst.Error
}

func FindMany(database string, sendId string, id string, time int64, pageSize int) (results []ws.Result, err error) {
	var resultsMe []ws.Trainer
	var resultsYou []ws.Trainer

	sendIdCollection := db.Where("from = ?", sendId).Find(&resultsYou)
	if sendIdCollection.Error != nil || sendIdCollection.RowsAffected == 0 {
		return
	}
	IdCollection := db.Where("to = ?", sendId).Find(&resultsMe)
	if IdCollection.Error != nil || IdCollection.RowsAffected == 0 {
		return
	}
	results, _ = AppendAndSort(resultsMe, resultsYou)
	return
}

func AppendAndSort(resultsMe, resultsYou []ws.Trainer) (results []ws.Result, err error) {
	for _, r := range resultsMe {
		sendSort := SendSortMsg{
			Content:  r.Content,
			Read:     r.Read,
			CreateAt: r.StartTime,
		}
		result := ws.Result{
			StartTime: r.StartTime,
			Msg:       fmt.Sprintf("%v", sendSort),
			From:      "me",
		}
		results = append(results, result)
	}
	for _, r := range resultsYou {
		sendSort := SendSortMsg{
			Content:  r.Content,
			Read:     r.Read,
			CreateAt: r.StartTime,
		}
		result := ws.Result{
			StartTime: r.StartTime,
			Msg:       fmt.Sprintf("%v", sendSort),
			From:      "you",
		}
		results = append(results, result)
	}
	// 最后进行排序
	sort.Slice(results, func(i, j int) bool { return results[i].StartTime < results[j].StartTime })
	return results, nil
}
