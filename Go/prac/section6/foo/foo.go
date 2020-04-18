package foo

const (
	// 最大値
	MAX = 100
)

const (
	A = iota
	B
	C
)

// Tの定義
type T struct {
	Field1 int
	field2 string
}

// *T型に定義された1番目のメソッド
func (t *T) Method2() {

}

// 型コンストラクタ
func New() *T {
	return &T{}
}
