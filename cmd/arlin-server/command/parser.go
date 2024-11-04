package command

import (
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

	// Process each key=value pair after the action
	for _, part := range parts[1:] {
		kv := strings.SplitN(part, "=", 2)
		if len(kv) != 2 {
			return nil, fmt.Errorf("invalid parameter format: %s", part)
		}
		key, value := kv[0], kv[1]

		// Remove any surrounding quotes (single or double) from the value
		if (strings.HasPrefix(value, `"`) && strings.HasSuffix(value, `"`)) ||
			(strings.HasPrefix(value, `'`) && strings.HasSuffix(value, `'`)) {
			value = value[1 : len(value)-1]
		}

		// Store the key and value in the Params map
		cmd.Params[key] = value
	}

	return cmd, nil
}
