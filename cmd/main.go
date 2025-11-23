package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env")
	}

	conn, err := connect()
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	ctx := context.Background()

	// ===== Добавление данных в таблицу =====
	// batch, err := conn.PrepareBatch(ctx, "INSERT INTO helloworld.my_first_table VALUES (user_id, message, timestamp, metric)")
	// if err != nil {
	// 	log.Fatal("Prepare batch failed:", err)
	// }
	// defer batch.Close()

	// //Insert any data
	// if err := batch.Append(150, "Change metric", time.Now(), 730); err != nil {
	// 	log.Fatal("Append failed:", err)
	// }
	// if err := batch.Append(109, "Change metric", time.Now(), 200); err != nil {
	// 	log.Fatal("Append failed:", err)
	// }
	// if err := batch.Append(115, "Change metric", time.Now(), 350); err != nil {
	// 	log.Fatal("Append failed:", err)
	// }
	// if err := batch.Send(); err != nil {
	// 	log.Fatal("Send batch failed:", err)
	// }
	//========================================

	rows, err := conn.Query(ctx, "SELECT user_id, message, timestamp, metric FROM helloworld.my_first_table")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			user_id   uint32
			message   string
			timestamp time.Time
			metric    float32
		)
		if err := rows.Scan(&user_id, &message, &timestamp, &metric); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, message: %s\n, Time: %s, Metric: %.2f\n", user_id, message, timestamp, metric)
	}
}

func connect() (driver.Conn, error) {
	userPass := os.Getenv("USER_PASSWORD")
	clickhouseHost := os.Getenv("CLICKHOUSE_HOST")
	clickhousePort := os.Getenv("CLICKHOUSE_PORT")

	var (
		ctx       = context.Background()
		conn, err = clickhouse.Open(&clickhouse.Options{
			Addr: []string{fmt.Sprintf("%s:%s", clickhouseHost, clickhousePort)},
			Auth: clickhouse.Auth{
				Database: "default",
				Username: "default",
				Password: userPass,
			},
		})
	)
	if err != nil {
		return nil, err
	}
	if err := conn.Ping(ctx); err != nil {
		if exeption, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("Expection [%d] %s \n%s\n", exeption.Code, exeption.Message, exeption.StackTrace)
		}
		return nil, err
	}
	return conn, nil

}

// Grafana madesvi.grafana.net
