package factory

// 抽象的な部品

// Abstract class Item
// fields
//   caption string
// methods
//   MakeHTML() string

type ItemInterface interface {
	MakeHTML() string
}

// Itemをcompositionさせてabstract classを表現する

type Item struct {
	Caption string
}

// Abstract class Link extends Item
// fields
//   url string

type LinkInterface interface {
	ItemInterface
}

type Link struct {
	*Item
	Url string
}

type TrayInterface interface {
	ItemInterface
	Add(item ItemInterface)
}

// Abstract class Items extends Item
// fields
//   tray array
// methods
//   Add

type Tray struct {
	*Item
	Items []ItemInterface
}

func (t *Tray) Add(item ItemInterface) {
	t.Items = append(t.Items, item)
}

// 抽象的な製品

type PageInterface interface {
	ItemInterface
	Add(item ItemInterface)
	Output(callerItem ItemInterface) string
}

type Page struct {
	Title, Author string
	Content       []ItemInterface
}

func (p *Page) Add(item ItemInterface) {
	p.Content = append(p.Content, item)
}

// composition しているやつからMakeHTML()したいので引数に渡す

func (p *Page) Output(callerItem ItemInterface) string {
	var result string
	result += callerItem.MakeHTML()

	return result
}

// 抽象的な工場

type Factory interface {
	CreateLink(caption, url string) LinkInterface
	CreateTray(caption string) TrayInterface
	CreatePage(title, author string) PageInterface
}
