##  使用说明
 
```go
package main

import (
    "fmt"
    "djangoauth"
)

func main() {
   //Verify（密码，数据库保存的密码字符串）
    b :=auth.Verify("vemuservemuser","pbkdf2_sha256$30000$TFKTX5dwlI77$onGTTQkvwGujhrctWUK/XKXFvkAoQfO4NRY6MpPK+Bw=")
    fmt.Println(b)
}

```

