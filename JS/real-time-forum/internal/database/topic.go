package database

import (
	"fmt"
	"net/http"
	"real-time-forum/internal/variables"
)

func GetTopicByID(id int) *variables.Topic {
	topic := &variables.Topic{}
	rows, err := variables.A.DB.Query(`SELECT * FROM topics WHERE id = ?`, id)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		rows.Scan(&topic.ID, &topic.Title, &topic.Content)
	}
	return topic
}

func GetTopics() ([]*variables.Topic, int) {
	var topics []*variables.Topic
	rows, _ := variables.A.DB.Query(`SELECT id from topics`)
	for rows.Next() {
		var topic_id int
		err := rows.Scan(&topic_id)
		if err != nil {
			fmt.Println(err)
			return nil, http.StatusInternalServerError
		}
		topics = append(topics, GetTopicByID(topic_id))
	}
	return topics, http.StatusOK
}
