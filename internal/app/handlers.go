package app

import (
	"net/http"
)

func HomepageHandler(c Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := c.dbGateway.Connection()
		if err != nil {
			c.Logger.Println("Unable to connect to database:", err)
			ErrorDbConnection(w)
			return
		}

		pages, err := LoadPages(db, 1, 10)
		if err != nil {
			c.Logger.Println(err);
			http.Error(w, "Unable to load pages", http.StatusInternalServerError)
			return
		}

		data := struct{ Content interface{} }{Content: pages}

		if err = c.tplRenderer.Render(w, "index.html", data); err != nil {
			c.Logger.Println(err);
		}
	}
}

func ViewHandler(c Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := c.dbGateway.Connection()
		if err != nil {
			c.Logger.Println("Unable to connect to database:", err)
			ErrorDbConnection(w)
			return
		}

		slug := r.Context().Value("slug").(string)
		if slug == "" {
			http.NotFound(w, r)
			return
		}

		page, err := LoadPage(db, slug)

		if err != nil {
			c.Logger.Println(err);
			http.Error(w, "Unable to load page", http.StatusInternalServerError)
			return
		}

		data := struct{ Content interface{} }{Content: page}

		if err = c.tplRenderer.Render(w, "view.html", data); err != nil {
			c.Logger.Println(err);
		}
	}
}

func EditHandler(c Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := c.dbGateway.Connection()
		if err != nil {
			c.Logger.Println("Unable to connect to database:", err)
			ErrorDbConnection(w)
			return
		}

		slug := r.Context().Value("slug").(string)
		if slug == "" {
			http.NotFound(w, r)
			return
		}

		page, err := LoadPage(db, slug)

		if err != nil {
			c.Logger.Println(err);
			http.Error(w, "Unable to load page", http.StatusInternalServerError)
			return
		}

		data := struct{ Content interface{} }{Content: page}

		if err = c.tplRenderer.Render(w, "edit.html", data); err != nil {
			c.Logger.Println(err);
		}
	}
}

func CreateHandler(c Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := c.tplRenderer.Render(w, "create.html", nil); err != nil {
			c.Logger.Println(err);
		}
	}
}

func SaveHandler(c Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := c.dbGateway.Connection()
		if err != nil {
			c.Logger.Println("Unable to connect to database:", err)
			ErrorDbConnection(w)
			return
		}

		slug := r.Context().Value("slug").(string)
		if slug == "" { // create action
			// @todo add check if page with save slug already exists
			slug = GenerateSlug(r.PostFormValue("Title"))
			page := NewPage(r.PostFormValue("Title"), slug, r.PostFormValue("Content"))
			if err := SavePage(db, page); err != nil {
				c.Logger.Println(err);
				http.Error(w, "Unable to save page", http.StatusInternalServerError)
				return
			}
		} else { // update action
			page := NewPage(r.PostFormValue("Title"), slug, r.PostFormValue("Content"))
			c.Logger.Println(slug)
			if err := UpdatePage(db, page); err != nil {
				c.Logger.Println(err);
				http.Error(w, "Unable to update page", http.StatusInternalServerError)
				return
			}
		}

		http.Redirect(w, r, "/view/"+slug, http.StatusSeeOther)
	}
}

func ErrorDbConnection(w http.ResponseWriter) {
	http.Error(w, "Unable to connect to database", http.StatusInternalServerError)
}
