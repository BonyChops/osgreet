// OSごとに挨拶が変わるかわいいパッケージ
package osgreet

import "errors"

type Greet interface {
	IsImplemented() bool // 実装されているかの判定
	Greet() error        // 実装するメソッド
}

// オーバーライド(風なこと)をするため、親に当たる構造体を作る
// client.goでのみ用いる
type OSGreetBase struct{}

// 実際の実装にはこちらを用いる
type OSGreet struct {
	OSGreetBase // OSGreetBaseを継承(風なことをする)
}

// OSGreetのコンストラクタ
func NewOSGreet() (Greet, error) {
	var greet Greet = &OSGreet{}

	if ok := greet.IsImplemented(); !ok {
		// 実装されていない = このOSではサポートされていない場合
		return nil, errors.New("this os is not supported")
	}

	return greet, nil
}

// 実装されているかの判定につかう
// 親構造体をfalseにしておいて、実装した際にはtrueを返すメソッドをオーバーライド(風なことを)する
func (*OSGreetBase) IsImplemented() bool {
	return false
}

// ダミーのメソッド
func (*OSGreetBase) Greet() error {
	return errors.New("this os is not supported")
}
