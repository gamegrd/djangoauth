##使用说明
 
```go
package main

import (
    "fmt"
    "djangoauth"
)

func main() {
    b :=auth.Verify("vemuservemuser","pbkdf2_sha256$30000$TFKTX5dwlI77$onGTTQkvwGujhrctWUK/XKXFvkAoQfO4NRY6MpPK+Bw=")
    fmt.Println(b)
}

```

