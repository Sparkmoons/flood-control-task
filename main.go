package main

import (
	"context"
)

func main() {
	//Адрес, пароль, временной промежуток, порог
	fc := NewFloodControl("localhost:6379", "", 5*time.Second, 3)

	userID := int64(159)
	for i := 0; i < 5; i++ {
		allowed, err := fc.Check(context.Background(), userID)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Check passed: %v\n", allowed)
		time.Sleep(1 * time.Second)
	}
}

// FloodControl интерфейс, который нужно реализовать.
// Рекомендуем создать директорию-пакет, в которой будет находиться реализация.
type FloodControl interface {
	// Check возвращает false если достигнут лимит максимально разрешенного
	// кол-ва запросов согласно заданным правилам флуд контроля.
	Check(ctx context.Context, userID int64) (bool, error)
}
