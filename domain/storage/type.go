package storage

// StorageType defines available storage types
type Type int

const (
	// JSON will store data in JSON files saved on disk
	JSONFiles Type = iota
	// Memory will store data in memory
	InMemory
)