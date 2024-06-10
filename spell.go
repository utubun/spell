package spell

import (
	"fmt"

	words "github.com/utubun/spell/internal"
)

func Spell() string {
	return fmt.Sprintf("A %s %s of the %s", words.Adjective(), words.Noun(), words.Noun())
}
