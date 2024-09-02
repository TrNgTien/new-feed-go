package configs

type StorageType string

const (
	StorageTypeMinio StorageType = "minio"
)

type Storage struct {
	Type     StorageType `yaml:"type"`
	Host     string      `yaml:"host"`
	Port     int         `yaml:"port"`
	Username string      `yaml:"username"`
	Password string      `yaml:"password"`
}
