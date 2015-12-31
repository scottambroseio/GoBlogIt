package goblogit

import "time"

type Comment struct {
        Author  string
        Content string
        Date    time.Time
        LastUpdated time.Time
}

func newComment() Comment {
	time := time.Now()

	return Comment{"Test Comment Author", "Test Comment Content", time, time}
}