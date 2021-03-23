# Pipet
GORM wrapper to connect to read and write databases

# Install
```go get github.com/ekyfauzi/pipet```

# Usages
### Connection
```go
import "github.com/ekyfauzi/pipet"

func main() {
  host := "YOUR_DB_WRITE_HOST"
  hostRead := "YOUR_DB_READ_HOST"
  port := "YOUR_DB_PORT"
  user := "YOUR_DB_USER"
  passwd := "YOUR_DB_PASSWORD"
  dbName := "YOUR_DB_NAME"
  
  // Set default connection
  conn := pipet.Init("mysql")
  conn.SetWrite(host, port, user, passwd, dbName)
  
  // Add read db connection
  // You can skip this if your read and write database in the same host
  // You also can set multiple read connection
  conn.SetRead(hostRead, port, user, passwd, dbName)
  
}
```

### Functions
Because this library is a wrapper of [GORM](https://github.com/go-gorm/gorm) so basically usages of functions are same.
```go
conn.Where()
conn.Save()
conn.Create()
conn.Exec()
```

If you want to get instance of connection to database, you can use:
```go
// Get write connection
// Return *gorm.DB
conn.Instance("write")

// Get read connection
// Return *gorm.DB
conn.Instance("read")
```
then you can use as usual GORM
