package cursor

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Cursor helps identify resource by its timestamp and UUID.
type Cursor struct {
	Timestamp time.Time
	UUID      uuid.UUID
}

// Decode decodes information from Cursor.
func Decode(encodedCursor string) (*Cursor, error) {
	byt, err := base64.StdEncoding.DecodeString(encodedCursor)
	if err != nil {
		return nil, err
	}

	arrStr := strings.Split(string(byt), "#")
	if len(arrStr) != 2 {
		err = errors.New("cursor is invalid")
		return nil, err
	}

	t, err := time.Parse(time.RFC3339Nano, arrStr[0])
	if err != nil {
		return nil, err
	}
	id, err := uuid.Parse(arrStr[1])
	if err != nil {
		return nil, err
	}
	return &Cursor{
		Timestamp: t,
		UUID:      id,
	}, nil
}

// Encode encodes information into cursor.
func Encode(c *Cursor) string {
	key := fmt.Sprintf("%s#%s", c.Timestamp.Format(time.RFC3339Nano), c.UUID.String())
	return base64.StdEncoding.EncodeToString([]byte(key))
}
