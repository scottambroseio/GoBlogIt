package goblogit

import (
        "html/template"
        "net/http"
        "appengine"
        "fmt"
)

func init() {
    http.HandleFunc("/", createHandler(root))
    http.HandleFunc("/get", createHandler(get))
    http.HandleFunc("/delete", createHandler(delete))
    http.HandleFunc("/update", createHandler(update))
    http.HandleFunc("/create", createHandler(create))
}

func get(c appengine.Context, w http.ResponseWriter, r *http.Request) (error) {
	var id int64 = 5064350557536256

	b, err := getBlog(c, id)

	fmt.Fprint(w, b.Content)

	return err
}

func update(c appengine.Context, w http.ResponseWriter, r *http.Request) (error) {
	var id int64 = 5064350557536256

	err := updateBlog(c, "This has been updated", id)

	return err
}


func delete(c appengine.Context, w http.ResponseWriter, r *http.Request) (error) {
	var id int64 = 5064350557536256

	err := deleteBlog(c, id)

	return err
}

func create(c appengine.Context, w http.ResponseWriter, r *http.Request) (error) {
	_, err := createBlog(c, "This is a blog")
    
    return err
}

func root(c appengine.Context, w http.ResponseWriter, r *http.Request) (error) {
	blogs, err := getAllBlogs(c)

    if err != nil {
    	return err;
    }

    if err := blogsTemplate.Execute(w, blogs); err != nil {
        return err;
    }

    return nil;
}

func createHandler(fn func(appengine.Context, http.ResponseWriter, *http.Request) (error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	    c := appengine.NewContext(r)

		if err := fn(c, w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

var blogsTemplate = template.Must(template.New("book").Parse(`
<html>
  <head>
    <title>Go Guestbook</title>
  </head>
  <body>
    {{range .}}
      {{with .Author}}
        <p><b>{{.}}</b> wrote:</p>
      {{else}}
        <p>An anonymous person wrote:</p>
      {{end}}
      <pre>{{.Content}}</pre>
    {{end}}
    <form action="/sign" method="post">
      <div><textarea name="content" rows="3" cols="60"></textarea></div>
      <div><input type="submit" value="Sign Guestbook"></div>
    </form>
  </body>
</html>
`))