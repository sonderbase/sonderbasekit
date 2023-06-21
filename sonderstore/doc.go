package sonderstore

import "google.golang.org/protobuf/proto"

type Doc interface {
	proto.Message
	GetCollectionName() string
	GetID() string
}
