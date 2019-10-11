package googlecal

import (
	"fmt"
	"os"
	"testing"
)

func TestUpdate(t *testing.T) {
	NewGoogleCalendarModule(os.Getenv("MAGIC_MIRROR_SERVICE_LOCATION"))
	fmt.Println("I DID NOT CRASH!")
}
