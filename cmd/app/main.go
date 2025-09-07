package main

import (
	"bytes"
	"context"
	"fmt"
	"image/gif"
	"math/rand"
	"time"

	"sample/assets"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	app := app.New()
	window := app.NewWindow("sample GIF Viewer")
	window.Resize(fyne.NewSize(500, 400))

	gifImg, err := gif.DecodeAll(bytes.NewReader(assets.GifData))
	if err != nil {
		fmt.Println("Failed to decode GIF:", err)
		return
	}

	img := canvas.NewImageFromImage(gifImg.Image[0])
	img.FillMode = canvas.ImageFillContain
	stack := container.NewStack(img)
	window.SetContent(stack)

	// アプリケーション終了時にgoルーチンを停止させるためのcontext
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// goルーチンでアニメーションを再生
	go func() {
		for {
			println("gif start")
			for i, frame := range gifImg.Image {
				select {
				case <-ctx.Done():
					return
				default:
				}

				// UIの更新はメインスレッドで実行
				fyne.Do(func() {
					img.Image = frame
					img.Refresh()
				})
				// フレームの表示時間を設定（GIFの遅延時間 × 10ミリ秒）
				delay := time.Duration(gifImg.Delay[i]) * 10 * time.Millisecond
				// 設定した時間だけ待機
				time.Sleep(delay)
			}
			println("gif end")
		}
	}()

	// 別のgoルーチンで5秒ごとにメッセージを表示
	go func() {
		total := 0
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(5 * time.Second):
				total += rand.Intn(100)
				println(total)
			}
		}
	}()

	// ウィンドウを表示してアプリケーションを実行
	window.ShowAndRun()
}
