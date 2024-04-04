package tools

import "github.com/bwmarrin/snowflake"

func GenId() (int64, error) {
	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(1)
	if err != nil {
		return -1, err
	}

	// Generate a snowflake ID.
	id := node.Generate()
	return id.Int64(), nil
}
