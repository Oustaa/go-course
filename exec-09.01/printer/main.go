package printer

import (
	"fmt"

	"github.com/google/uuid"
)

func PrintNewUUID() string {
	id := uuid.New()

	return fmt.Sprintf("you generated UUID is: %s\n", id)
}
