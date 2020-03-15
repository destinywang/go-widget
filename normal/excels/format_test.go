package excels

import "testing"

func TestDiff(t *testing.T) {
	Diff("/Users/destiny/go/src/github.com/DestinyWang/go-widget/normal/excels/src.xlsx", "/Users/destiny/go/src/github.com/DestinyWang/go-widget/normal/excels/target.xlsx")
}

func TestExtractFromTxt(t *testing.T) {
	targetSet := ExtractFromTxt("/Users/destiny/go/src/github.com/DestinyWang/go-widget/normal/excels/target.txt")
	t.Log(targetSet)
}
