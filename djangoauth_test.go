package djangoauth

import (
    "testing"
)

func Test_Verify(t *testing.T) {
    num := Verify("vemuservemuser","pbkdf2_sha256$30000$TFKTX5dwlI77$onGTTQkvwGujhrctWUK/XKXFvkAoQfO4NRY6MpPK+Bw=")
    if !num{
        t.Error()
    }
}
