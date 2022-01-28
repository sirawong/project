package utils

import "github.com/bwmarrin/snowflake"

//go:generate mockery --name=UUID
type UUID interface {
	Generate() (uuid string)
}

type SnowflaskUUID struct {
	generator *snowflake.Node
}

func NewUUID() (uuid UUID, err error) {
	node, err := snowflake.NewNode(1023)
	if err != nil {
		return nil, err
	}
	return &SnowflaskUUID{generator: node}, nil
}

func (u *SnowflaskUUID) Generate() (uuid string) {
	return u.generator.Generate().String()
}
