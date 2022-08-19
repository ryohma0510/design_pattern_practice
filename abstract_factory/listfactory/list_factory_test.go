package listfactory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFactory(t *testing.T) {
	factoryObject := ListFactory{}
	asahi := factoryObject.CreateLink("Asahi", "http://www.asahi.com")
	yomiuri := factoryObject.CreateLink("Yomiuri", "http://www.yomiuri.co.jp")
	usYahoo := factoryObject.CreateLink("Yahoo", "http://www.yahoo.com")
	jaYahoo := factoryObject.CreateLink("Yahoo!Japan", "http://www.yahoo.co.jp")
	google := factoryObject.CreateLink("Google", "http://www.google.com")
	excite := factoryObject.CreateLink("Excite", "http://www.excite.co.jp")

	traynews := factoryObject.CreateTray("Newspaper")
	traynews.Add(asahi)
	traynews.Add(yomiuri)

	trayyahoo := factoryObject.CreateTray("Yahoo!")
	trayyahoo.Add(usYahoo)
	trayyahoo.Add(jaYahoo)

	traysearch := factoryObject.CreateTray("Search Engine")
	traysearch.Add(trayyahoo)
	traysearch.Add(excite)
	traysearch.Add(google)

	page := factoryObject.CreatePage("LinkPage", "Hiroshi Yuki")
	page.Add(traynews)
	page.Add(traysearch)

	assert.Equal(
		t,
		"<html>\n  <head><title>LinkPage</title></head>\n<body>\n<h1>LinkPage</h1><ul><li>\nNewspaper\n<ul>\n  <li><a href=\"http://www.asahi.com\">Asahi</a></li>\n  <li><a href=\"http://www.yomiuri.co.jp\">Yomiuri</a></li>\n</ul>\n</li>\n<li>\nSearch Engine\n<ul>\n<li>\nYahoo!\n<ul>\n  <li><a href=\"http://www.yahoo.com\">Yahoo</a></li>\n  <li><a href=\"http://www.yahoo.co.jp\">Yahoo!Japan</a></li>\n</ul>\n</li>\n  <li><a href=\"http://www.excite.co.jp\">Excite</a></li>\n  <li><a href=\"http://www.google.com\">Google</a></li>\n</ul>\n</li>\n</ul><hr><adress>Hiroshi Yuki</adress></body>\n</html>\n",
		page.Output(page),
	)
}
