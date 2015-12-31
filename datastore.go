package goblogit

import (
	"appengine"
	"time"
	"appengine/datastore"
	"appengine/user"
)

const (
	KIND = "Blog"
)

func blogKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, KIND, "all_blogs", 0, nil)
}

func getBlog(c appengine.Context) {
	
}

func getAllBlogs(c appengine.Context) ([]Blog, error) {
	q := datastore.NewQuery(KIND).Ancestor(blogKey(c)).Order("-Date")

    blogs := make([]Blog, 0)

    if _, err := q.GetAll(c, &blogs); err != nil {
        return nil, err
    }

    return blogs, nil
}

func updateBlog(c appengine.Context) {

}

func deleteBlog(c appengine.Context) {

}

func createBlog(c appengine.Context, content string) (*datastore.Key, error){
	blog := &Blog{
        Content: content,
        Date:  time.Now(),
	}

	if u := user.Current(c); u != nil {
        blog.Author = u.String()
    } else {
    	blog.Author = "Anonymous"
    }

	key := datastore.NewIncompleteKey(c, KIND, blogKey(c))

	if generatedKey, err := datastore.Put(c, key, blog); err != nil {
		return nil, err
	} else {
		return generatedKey, nil
	}
}