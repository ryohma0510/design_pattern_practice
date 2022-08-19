package listfactory

import (
	"design_pattern_practice/abstract_factory/factory"
	"fmt"
)

type ListLink struct {
	*factory.Link
}

func newListLink(caption, url string) *ListLink {
	return &ListLink{
		Link: &factory.Link{
			Item: &factory.Item{Caption: caption},
			Url:  url,
		},
	}
}

func (ll *ListLink) MakeHTML() string {
	return fmt.Sprintf("  <li><a href=\"%s\">%s</a></li>\n", ll.Url, ll.Caption)
}

type ListTray struct {
	*factory.Tray
}

func newListTray(caption string) *ListTray {
	return &ListTray{
		Tray: &factory.Tray{
			Item: &factory.Item{
				Caption: caption,
			},
		},
	}
}

func (lt *ListTray) MakeHTML() string {
	buf := "<li>\n"
	buf += fmt.Sprintf("%s\n", lt.Caption)
	buf += "<ul>\n"
	for _, item := range lt.Items {
		buf += item.MakeHTML()
	}
	buf += "</ul>\n"
	buf += "</li>\n"
	return buf
}

type ListPage struct {
	*factory.Page
}

func newListPage(title, author string) *ListPage {
	return &ListPage{
		Page: &factory.Page{
			Title:  title,
			Author: author,
		},
	}
}

func (lp *ListPage) MakeHTML() string {
	buf := "<html>\n"
	buf += fmt.Sprintf("  <head><title>%s</title></head>\n", lp.Title)
	buf += "<body>\n"
	buf += fmt.Sprintf("<h1>%s</h1>", lp.Title)
	buf += "<ul>"
	for _, item := range lp.Content {
		buf += item.MakeHTML()
	}
	buf += "</ul>"
	buf += fmt.Sprintf("<hr><adress>%s</adress>", lp.Author)
	buf += "</body>\n</html>\n"
	return buf
}

// factory

type ListFactory struct {
}

func (lf *ListFactory) CreateLink(caption, url string) factory.ItemInterface {
	return newListLink(caption, url)
}

func (lf *ListFactory) CreateTray(caption string) factory.TrayInterface {
	return newListTray(caption)
}

func (lf *ListFactory) CreatePage(title, author string) factory.PageInterface {
	return newListPage(title, author)
}
