package goblogit

import (
	"appengine"
	"time"
	"appengine/datastore"
	"appengine/user"
)

func blogKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, KIND, "all_blogs", 0, nil)
}

func getBlog(c appengine.Context, id int64) (*Blog, error){
	key := datastore.NewKey(c, KIND, "", id, blogKey(c))

	blog := &Blog{}

	if err := datastore.Get(c, key, blog); err != nil {
		return nil, err
	}

	return blog, nil
}

func getAllBlogs(c appengine.Context) (*[]*Blog, error) {
	q := datastore.NewQuery(KIND).Ancestor(blogKey(c)).Order("-Date")

    blogs := make([]*Blog, 0)

    if _, err := q.GetAll(c, &blogs); err != nil {
        return nil, err
    }

    return &blogs, nil
}

func updateBlog(c appengine.Context, content string, id int64) error {

	key := datastore.NewKey(c, KIND, "", id, blogKey(c))

	blog := &Blog{}

	if err := datastore.Get(c, key, blog); err != nil {
		return err
	}

	blog.Content = content
	blog.LastUpdated = time.Now()

	if _, err := datastore.Put(c, key, blog); err != nil {
		return err
	}

	return nil
}

func deleteBlog(c appengine.Context, id int64) error {
	key := datastore.NewKey(c, KIND, "", id, blogKey(c))

	err := datastore.Delete(c, key)

	return err;
}

func createBlog(c appengine.Context, content string) (*datastore.Key, error){
	time := time.Now()

	blog := &Blog{
        Content: content,
        Date:  time,
        LastUpdated: time,
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