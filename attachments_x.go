package slack

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Slackdump extensions.

type AttachmentID string

func (a *AttachmentID) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}

	if data[0] == '"' && data[len(data)-1] == '"' {
		var ret string
		if err := json.Unmarshal(data, &ret); err != nil {
			return err
		}
		*a = AttachmentID(ret)
		return nil
	} else if data[0]-'0' <= 9 {
		var n int64
		if err := json.Unmarshal(data, &n); err != nil {
			return err
		}
		*a = AttachmentID(strconv.FormatInt(n, 10) + "$")
		return nil
	}
	return fmt.Errorf("invalid AttachmentID: %s", string(data))
}

func (a *AttachmentID) MarshalJSON() ([]byte, error) {
	if a == nil {
		return nil, nil
	}
	if len(*a) == 0 {
		return []byte("null"), nil
	}
	if (*a)[len(*a)-1] == '$' {
		// attempt to convert to integer
		var n int64
		if _, err := fmt.Sscanf(string(*a), "%d$", &n); err != nil {
			return nil, err
		}
		return json.Marshal(n)
	}
	// otherwise, return as string
	return json.Marshal(string(*a))
}
