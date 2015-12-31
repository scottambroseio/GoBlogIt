package goblogit

import (
        "html/template"
        "net/http"
        "appengine"
)

func init() {
    http.HandleFunc("/", createHandler(root))
    http.HandleFunc("/sign", createHandler(sign))
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

func sign(c appengine.Context, w http.ResponseWriter, r *http.Request)  (error) {
    if _, err := createBlog(c, "This is a blog"); err != nil {
    	return err
    }

   	http.Redirect(w, r, "/", http.StatusFound)

   	return nil
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