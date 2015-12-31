package goblogit

import "time"

type Blog struct {
        Author  string
        Content string
        Date    time.Time
}

func newBlog() Blog {
	return Blog{"Test Author", "Test Content", time.Now()}
}