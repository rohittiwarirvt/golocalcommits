// You can edit this code!
// Click here and start typing.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	ct "git.druva.org/enterprise-wl/ekey/constants"
	jwt "github.com/golang-jwt/jwt/v4"
)

type myClaims struct {
	*jwt.StandardClaims
	Name string `json:"name,omitempty"`
}

func (c myClaims) Valid() error {
	return nil
}

func main() {
	//itemInfoR := "{\"sub\":\"1234567890\",\"name\":\"John Doe\",\"iat\":1516239022}"

	_ = jwt.NewParser(jwt.WithoutClaimsValidation())
	// token, _, err := newParser.ParseUnverified(ct.JWT, myClaims{})
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// //v := token.Claims.(myClaims)
	// //fmt.Println(myClaims)
	// //fmt.Println(myClaims.Name)
	// // itemInfoBytes := []byte(itemInfoR)
	// // var ItemInfo myClaims
	// // er := json.Unmarshal(itemInfoBytes, &ItemInfo)
	// // if er != nil {
	// // 	panic(er)
	// // }
	// // fmt.Println(ItemInfo.Subject)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// if claims, ok := token.Claims.(jwt.MapClaims); ok {
	// 	fmt.Println(claims)
	// } else {
	// 	fmt.Println(err)
	// }

	// test1 := `{\"sub\":\"1234567890\","name":"John Doe","iat":1516239022}`
	// str := []byte(test1)

	// dec := json.NewDecoder(bytes.NewBuffer(str))
	// // // if c, ok := token.Claims.(MapClaims); ok {
	// // // 	err = dec.Decode(&c)
	// // // } else {
	// a := &myClaims{}
	// err = dec.Decode(a)
	// // // }
	// // fmt.Println(a.Name)
	// // Handle decode error
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }


	// ekey generations

	var encProductSecret, tokenSecret string
	encProductSecret = "1fXoIBWVtnpH6ivZLFF//8Y2BXMZipanrT0bB9JSQ8g22pp9U7YbhCp8K12eFSO1ijvnxNcZ3PJ4cTbbkQuvJwZGwuuvoqeeM/4qtUJ7xEmA/SVdRzzNgkaSxYk13S5ewlBemHfnLoNqxhLn0j2cZL10Mz+I6q7c0HdoCKJYlrhQSdi7p72ZIgZR0Ipo2yJkyv/RSUeJ/RAEJl4YCBP+t6Hcuu/PZ/d30ygY2iWnq8xe5ukd4csU6Gal6gVaYykLbzSDvp6a3A7t2WqDx2WnuJw7KLns/AF7TJ+l8goobMXT1bD1LnXzc8UE5SJBrw2N"

  tokenSecret = "iypTVqUMo4w3Ye9cETzq6A=="
	
}
