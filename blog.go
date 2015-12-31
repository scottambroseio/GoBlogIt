package goblogit

import "time"

type Blog struct {
        Author  string
        Content string
        Date    time.Time
        LastUpdated time.Time
}

func newBlog() Blog {
	time := time.Now()

	return Blog{"Test Author", "Test Content", time, time}
}