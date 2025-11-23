// package backup

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/ClickHouse/clickhouse-go/v2"
// )

// func main() {
// 	conn, err := clickhouse.Open(&clickhouse.Options{
// 		Addr: []string{"127.0.0.1:9000"},
// 		Auth: clickhouse.Auth{
// 			Database: "default",
// 			Username: "default",
// 			Password: "home",
// 		},
// 	})
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer conn.Close()

// 	// err = conn.Exec(context.Background(), "INSERT INTO helloworld.my_first_table VALUES (user)")
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }

// 	ctx := context.Background()
// 	if err := conn.Ping(ctx); err != nil {
// 		log.Fatal("Ping failed", err)
// 	}

// 	// batch, err := conn.PrepareBatch(ctx, "INSERT INTO helloworld.my_first_table VALUES (user_id, message, timestamp, metric)")
// 	// if err != nil {
// 	// 	log.Fatal("Prepare batch failed:", err)
// 	// }
// 	// defer batch.Close()

// 	// //Insert any data
// 	// if err := batch.Append(105, "Hello, from Go", time.Now(), 555); err != nil {
// 	// 	log.Fatal("Append failed:", err)
// 	// }
// 	// if err := batch.Send(); err != nil {
// 	// 	log.Fatal("Send batch failed:", err)
// 	// }

// 	rows, err := conn.Query(ctx, "SELECT user_id, message, timestamp, metric FROM helloworld.my_first_table")

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var (
// 			user_id   uint32
// 			message   string
// 			timestamp time.Time
// 			metric    float32
// 		)
// 		if err := rows.Scan(&user_id, &message, &timestamp, &metric); err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Printf("ID: %d, message: %s\n, Time: %s, Metric: %.2f\n", user_id, message, timestamp, metric)
// 	}
// }

// // Grafana madesvi.grafana.net
