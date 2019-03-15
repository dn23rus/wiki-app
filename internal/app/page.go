package app

import (
	"database/sql"
	"errors"
	"github.com/gosimple/slug"
)

type Page struct {
	Title, Slug, Content string
}

func NewPage(title, slug, content string) *Page {
	page := new(Page)
	page.Title = title
	page.Slug = slug
	page.Content = content

	return page
}

func GenerateSlug(title string) string {
	return slug.Make(title)
}

func LoadPages(db *sql.DB, page, count uint) ([]Page, error) {
	limit, offset := count, (page-1)*count
	stmt, err := db.Prepare("select slug, title, content from pages order by created_at desc limit ? offset ?")

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(limit, offset)
	if err != nil {
		return nil, err
	}

	var result = []Page{}
	for rows.Next() {
		var slug, title, content string
		err = rows.Scan(&slug, &title, &content)
		if err != nil {
			return result, err
		}
		result = append(result, *NewPage(title, slug, content))
	}
	err = rows.Err()

	return result, err
}

func LoadPage(db *sql.DB, key string) (*Page, error) {
	stmt, err := db.Prepare("select slug, title, content from pages where slug = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var slug, title, content string
	err = stmt.QueryRow(key).Scan(&slug, &title, &content)
	if err != nil {
		return nil, err
	}

	return NewPage(title, slug, content), nil
}

func SavePage(db *sql.DB, page *Page) error {

	stmt, err := db.Prepare("INSERT INTO pages(slug, title, content) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(page.Slug, page.Title, page.Content)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("new page was not created")
	}

	return nil
}

func UpdatePage(db *sql.DB, page *Page) error {
	stmt, err := db.Prepare("UPDATE pages set content = ? WHERE slug = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(page.Content, page.Slug)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("page was not updated" )
	}

	return nil
}
