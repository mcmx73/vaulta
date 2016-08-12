package bson

import (
	"gopkg.in/mgo.v2/bson"
)

var Codec = new(bsonCodec)

type bsonCodec int

func (c bsonCodec) Encode(v interface{}) (data []byte, err error) {
	data, err = bson.Marshal(v)
	return
}

func (c bsonCodec) Decode(b []byte, v interface{}) (err error) {
	err = bson.Unmarshal(b, &v)
	return
}
