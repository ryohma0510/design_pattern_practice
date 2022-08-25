package facade

import "fmt"

var data = map[string]string{
	"a@a.com": "a",
	"b@b.com": "b",
	"c@b.com": "c",
	"d@b.com": "d",
}

type database struct {
}

func (db *database) getNameByMail(mail string) string {
	return data[mail]
}

type mdWriter struct {
	title   string
	content string
}

func (m mdWriter) Title() string {
	return fmt.Sprintf("# Welcome to %s's !\n", m.title)
}

func (m mdWriter) Content() string {
	return fmt.Sprintf("%s\n", m.content)
}

type PageMaker struct {
}

func (pm *PageMaker) MakeWelcomePage(mail string) string {
	db := database{}
	name := db.getNameByMail(mail)

	writer := mdWriter{
		title:   name,
		content: "This is content",
	}

	page := writer.Title() + writer.Content()
	return page
}
