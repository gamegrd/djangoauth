package auth
/*
#############################################
# File Name: djangopwd.go
# Author: gamegrd
# Mail: gamegrd@gmail.com
# Created Time:  2017-06-22 12:06:22
#############################################



                         _oo8oo_
                        o8888888o
                        88" . "88
                        (| -_- |)
                        0\  =  /0
                      ___/'==='\___
                    .' \\|     |// '.
                   / \\|||  :  |||// \
                  / _||||| -:- |||||_ \
                 |   | \\\  -  /// |   |
                 | \_|  ''\---/''  |_/ |
                 \  .-\__  '-'  __/-.  /
               ___'. .'  /--.--\  '. .'___
            ."" '<  '.___\_<|>_/___.'  >' "".
           | | :  `- \`.:`\ _ /`:.`/ -`  : | |
           \  \ `-.   \_ __\ /__ _/   .-` /  /
       =====`-.____`.___ \_____/ ___.`____.-`=====
                         `=---=`

                佛祖保佑        永无BUG

*/

import (
    "crypto/sha256"
    "encoding/base64"
    "golang.org/x/crypto/pbkdf2"
    "regexp"
    _ "fmt"
    "strconv"
)

func EncodePassword(password,salt string, iterations int)(string){
    digest := sha256.New
    dk := pbkdf2.Key([]byte( password ),[]byte( salt), iterations, 32, digest)
    str := base64.StdEncoding.EncodeToString(dk)
    rt := "pbkdf2_sha256" + "$" + strconv.FormatInt(int64(iterations), 10) + "$" + string(salt) + "$" + str
    // fmt.Println(rt)
    return rt
}

var encodedRegexp = regexp.MustCompile(`(.+)\$(\d+)\$(.+)\$(.+)`)
func Verify(pwd,encoded string)(bool){
    dict := encodedRegexp.FindStringSubmatch(encoded)
    iterations,salt:=dict[2],dict[3]
    iiterations,_ := strconv.Atoi(iterations)
    return EncodePassword(pwd,salt,iiterations) == encoded
}

