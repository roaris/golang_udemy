package main

import (
	"fmt"
	"sync"
)

// sync
// ミューテックスによる同期処理

// ミューテックスを保持するパッケージ変数
// var mutex *sync.Mutex

// var st struct{ A, B, C int }

// func UpdateAndPrint(n int) {
// 	/*
// 		st.A = n
// 		time.Sleep(time.Microsecond)
// 		st.B = n
// 		time.Sleep(time.Microsecond)
// 		st.C = n
// 		time.Sleep(time.Microsecond)
// 		fmt.Println(st)
// 	*/

// 	// ロック
// 	mutex.Lock()
// 	st.A = n
// 	time.Sleep(time.Microsecond)
// 	st.B = n
// 	time.Sleep(time.Microsecond)
// 	st.C = n
// 	time.Sleep(time.Microsecond)
// 	fmt.Println(st)

// 	// アンロック
// 	mutex.Unlock()
// }

func main() {
	/*
		mutex = new(sync.Mutex)

		for i := 0; i < 5; i++ {
			go func() {
				for i := 0; i < 1000; i++ {
					UpdateAndPrint(i)
				}
			}()
		}

		for {

		}
	*/

	// ゴルーチンの終了を待ち受ける

	// sync.WaitGroupを生成
	wg := new(sync.WaitGroup)
	// 待ち受けするゴルーチンの数は3
	wg.Add(3)

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("1st Goroutine")
		}
		wg.Done() // 完了
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("2nd Goroutine")
		}
		wg.Done() // 完了
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("3rd Goroutine")
		}
		wg.Done() // 完了
	}()

	// ゴルーチンの完了を待ち受ける
	// Doneが3つ完了するまで待つ
	wg.Wait()
}
