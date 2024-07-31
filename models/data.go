package models

import "time"

var (
	// 初期化時に文字列を time.Time に変換
	Todo1 = Todo{
		ID:         1,
		Title:      "first title",
		Body:       "first body",
		DueDate:    parseTime("2023-12-28T23:08:18.417583+09:00"),
		CompleteAt: parseTime("2023-12-28T23:08:18.417583+09:00"),
		CreatedAt:  parseTime("2023-12-28T23:08:18.417583+09:00"),
		UpdateAt:   parseTime("2023-12-28T23:08:18.417583+09:00"),
	}
	Todo2 = Todo{
		ID:         2,
		Title:      "second title",
		Body:       "second body",
		DueDate:    parseTime("2023-12-28T23:08:18.417583+09:00"),
		CompleteAt: parseTime("2023-12-28T23:08:18.417583+09:00"),
		CreatedAt:  parseTime("2023-12-28T23:08:18.417583+09:00"),
		UpdateAt:   parseTime("2023-12-28T23:08:18.417583+09:00"),
	}
)

// parseTime 関数を追加
func parseTime(timeStr string) time.Time {
	layout := time.RFC3339Nano
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		panic(err) // エラーが発生した場合はパニックにします（実際のアプリケーションでは適切にエラーハンドリングするべきです）
	}
	return t
}
