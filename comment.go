package goblogit

import "time"

type Comment struct {
        Author  string
        Content string
        Date    time.Time
}

func newComment() Comment {
	return Comment{"Test Comment Author", "Test Comment Content", time.Now()}
}