package utils

import (
	"crypto/rand"
	"fmt"
)

func GenerateId() string{
	b:=make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
//func GenerateId() bson.ObjectId{
//	x:=bson.NewObjectId()
//	return  x
//}