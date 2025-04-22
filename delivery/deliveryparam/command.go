package deliveryparam

import "encoding/json"

type Command struct {
	command string
}

func NewCommand(command string) *Command {
	return &Command{command: command}
}

func (c *Command) GetCommand() string {
	return c.command
}
func (c *Command) SetCommand(command string) {
	c.command = command
}
func (c *Command) MarshalJSON() ([]byte, error) {

	return json.Marshal(map[string]any{
		"command": c.GetCommand(),
	})
}
func (c *Command) UnmarshalJSON(data []byte) error {
	var aux struct {
		Command string `json:"command"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {

		return err
	}

	c.SetCommand(aux.Command)

	return nil
}
