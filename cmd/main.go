package main

import "fmt"
import "github.com/BonyChops/osgreet"

func main() {
	client, err := osgreet.NewOSGreet()
	if err != nil {
		fmt.Println(`main.go (this package is seems to be not supported this os...)`)
		return
	}

	err = client.Greet()
	if err != nil {
		// 適切に処理されていれば、そもそもここは呼ばれない
		// そういう意味では`Greet()`の実装はerrorのreturnなしのpanic()でも良いのかも
		fmt.Println(`main.go (this package is seems to be not supported this os...)`)
		return
	}
}
