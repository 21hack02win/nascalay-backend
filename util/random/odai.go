package random

import "math/rand"

var (
	prefix2Data = [...]string{
		"ねこの", "異世界", "アイドル", "天然", "真っ黒な", "トラと", "空から", "伝説の", "昭和", "平成", "大盛り",
		"宇宙の", "古代",
	}
	suffixData = [...]string{
		"リンゴ", "おばけ", "なす", "トラ", "博士", "スイカ", "パイナップル", "メロン", "バナナ", "ピーマン", "ブドウ",
		"爆弾", 
	}
)

//nolint:unused,deadcode
func odaiExample() string {
	prefix := prefix2Data[rand.Intn(len(prefix2Data))]
	suffix := suffixData[rand.Intn(len(suffixData))]
	return prefix + suffix
}
