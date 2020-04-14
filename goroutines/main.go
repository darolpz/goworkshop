package main

import (
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	insertNumber := 3000
	db, err := gorm.Open("mysql", "dummy:12345@/dummy?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	db.Exec("TRUNCATE TABLE dummy")

	// Se ejecuta concurrentemente
	withGoroutines(db, insertNumber)
	// Se ejecuta en serie
	withoutGoroutines(db, insertNumber)

}

func insertDummy(db *gorm.DB, dummyString string, insertNumber int, flag bool, waitGroup *sync.WaitGroup) {
	for i := 0; i < insertNumber; i++ {
		db.Exec("INSERT INTO dummy (dummy) VALUES (?)", dummyString)
		if flag {
			defer waitGroup.Done()
		}
	}
}

func withGoroutines(db *gorm.DB, insertNumber int) {
	start := time.Now()
	var waitGroup1, waitGroup2, waitGroup3 sync.WaitGroup
	waitGroup1.Add(insertNumber)
	waitGroup2.Add(insertNumber)
	waitGroup3.Add(insertNumber)
	go insertDummy(db, "Goroutine", insertNumber, true, &waitGroup1)
	go insertDummy(db, "Goroutine", insertNumber, true, &waitGroup2)
	go insertDummy(db, "Goroutine", insertNumber, true, &waitGroup3)
	fmt.Printf("%v \n", "Esperando")
	waitGroup1.Wait()
	waitGroup2.Wait()
	waitGroup3.Wait()
	fmt.Printf("Tiempo transcurrido con goroutines %v \n", time.Since(start))
}

func withoutGoroutines(db *gorm.DB, insertNumber int) {
	start := time.Now()

	insertDummy(db, "Normal", insertNumber*3, false, new(sync.WaitGroup))

	fmt.Printf("Tiempo transcurrido sin goroutine %v \n", time.Since(start))
}
