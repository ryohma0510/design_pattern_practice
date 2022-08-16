package adapter

import "fmt"

// 元々あるやつ
type Banner struct {
	content string
}

func (b Banner) ShowWithParen() string {
	return fmt.Sprintf("(%s)", b.content)
}

func (b Banner) ShowWithAster() string {
	return fmt.Sprintf("*%s*", b.content)
}

type Print interface {
	PrintWeak() string
	PrintStrong() string
}

// Adapter
// Bannerがいい感じにinterfaceを満たしていないときに調整する
type PrintBanner struct {
	Banner
}

func (pb PrintBanner) PrintWeak() string {
	return pb.ShowWithParen()
}

func (pb PrintBanner) PrintStrong() string {
	return pb.ShowWithAster()
}
