# PostgreSQL

## Customize Driver

```
import (
    _ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
    "gorm.io/gorm"
)

db, err := gorm.Open(postgres.New(postgre.Config{
    DriverName: "cloudsqlpostgres"
    DDN: "host=project:region:instance user=postgres dbname=postgres password=password sslmode=disable"
}))
```

## Existing database connection

GORM allows to initialize `*gorm.DB` with an existing database connection

```
import (
    "database/sql"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

sqlDB, err := sql.Open("pgx", "mydb_dsn")
gormDB, err := gorm.Open(postgress.New(postgres.Config{
    Conn: sqlDB,
}), &gorm.Config{})
```

