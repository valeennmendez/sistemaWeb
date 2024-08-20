/*package connection

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const DSN = "root:@tcp(127.0.0.1:3306)/Acneclinic?charset=utf8mb4&parseTime=True&loc=UTC"

func ConnectionDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Base de Datos corriendo...")

}*/

package connection

import (
    "fmt"
    "os"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDB() {
    dbURL := os.Getenv("MYSQL_URL")
    if dbURL == "" {
        panic("No se ha definido la variable de entorno MYSQL_URL")
    }

    var err error
    DB, err = gorm.Open(mysql.Open(dbURL), &gorm.Config{})
    if err != nil {
        panic(err)
    }

    fmt.Println("Base de Datos corriendo...")
}

