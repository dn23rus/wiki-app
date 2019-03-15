package app

import (
	"html/template"
	"io"
	"sync"
)

type TplRenderer struct {
	sync.RWMutex
	tplCache map[string]*template.Template
	tplDir   string
}

func NewTplRenderer(tplDir string) *TplRenderer {
	r := &TplRenderer{
		tplCache: make(map[string]*template.Template),
		tplDir:   tplDir,
	}
	return r
}

func (r *TplRenderer) Render(w io.Writer, filename string, data interface{}) error {
	r.Lock()
	defer r.Unlock()

	if _, has := r.tplCache[filename]; !has {
		files := r.prependLayout(filename)
		r.tplCache[filename] = template.Must(template.ParseFiles(files...))
	}
	if err := r.tplCache[filename].ExecuteTemplate(w, "layout", data); err != nil {
		return err
	}
	return nil
}

func (r *TplRenderer) prependLayout(filename string) []string {
	files := _map([]string{
		"layout/main.html",
		filename,
	}, func(s string) string {
		return r.resolveFile(s)
	})

	return files
}

func (r *TplRenderer) resolveFile(filename string) string {
	return r.tplDir + "/" + filename
}

func _map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
