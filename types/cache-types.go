package types

type Cache interface {
	Get(key string, objStruct any) (any, error)
	Set(key string, value any) error
}
