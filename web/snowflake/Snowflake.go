package snowflake

import "github.com/bwmarrin/snowflake"

// NewSnowFlake -
func NewNode() *snowflake.Node {
	node, _ := snowflake.NewNode(1)

	return node
}
