package command

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Command struct to hold parsed data
type Command struct {
	Action string
	Params map[string]string
}

// ParseCommand parses the command string into a Command struct
func ParseCommand(input string) (*Command, error) {
	// Split the input by spaces to isolate the action from the parameters
	parts := strings.Fields(input)
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid command format")
	}

	// Initialize the Command with the action and an empty Params map
	cmd := &Command{
		Action: strings.ToUpper(parts[0]),
		Params: make(map[string]string),
	}

	// Process each key=value pair
	for _, part := range parts[1:] {
		kv := strings.SplitN(part, "=", 2)
		if len(kv) != 2 {
			return nil, fmt.Errorf("invalid parameter format: %s", part)
		}
		key, value := kv[0], kv[1]

		// Special handling for the PAIR command with JSON data
		if cmd.Action == "PAIR" && key == "data" {
			// Deserialize JSON to validate it
			var jsonData map[string]interface{}
			if err := json.Unmarshal([]byte(value), &jsonData); err != nil {
				return nil, fmt.Errorf("invalid JSON format for data: %s", err)
			}
			// Store JSON string as-is in Params
			cmd.Params[key] = value
		} else {
			cmd.Params[key] = value
		}
	}

	return cmd, nil
}
