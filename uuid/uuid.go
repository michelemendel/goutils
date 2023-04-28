package uuid

import "github.com/segmentio/ksuid"

// Unique IDs
// https://blog.kowalczyk.info/article/JyRZ/generating-good-unique-ids-in-go.html

type UUID string

func GenerateUUID() UUID {
	return UUID(ksuid.New().String())
}
