package deliveryparam

import "encoding/json"

type Request struct {
	command string
}

func NewRequest(command string) *Request {
	return &Request{command: command}
}

func (r *Request) GetCommand() string {
	return r.command
}

func (r *Request) SetCommand(command string) {
	r.command = command
}

func (r *Request) UnmarshalJSON(data []byte) error {
	var aux struct {
		Command string `json:"command"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	r.SetCommand(aux.Command)
	return nil
}

func (r *Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"command": r.GetCommand(),
	})
}
