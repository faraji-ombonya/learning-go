# MySQL

## Custom Driver

GORM allows to customize the MySQL driver with the  `DriverName` option, for example:

```
import (
    _ "example.com/my_mysql_driver"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

db, err := gorm.Open(mysql.New(mysql.Config{
    DriverName: "my_mysql_driver",
    DSN: "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
}), &gorm.Config{})

```


## Existing database connection

GORM allows to initialize `*gorm.DB` with an existing database connection

```
import (
    "database/sql"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

sqlDB, err := sql.Open("mysql", "mydb_dsn")
gormDB, err := gorm.Open(mysql.New(mysql.Config{
    Conn: sqlDB,
}), &gorm.Config{})

```


