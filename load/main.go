package load

import (
	"encoding/json"
	"markdown/targetstructures"

	"github.com/spf13/afero"
)

func UnmarshalFile(content []byte) targetstructures.Output {
	loaded := targetstructures.Output{}
	json.Unmarshal(content, &loaded)
	return loaded
}

func Load(fs afero.Fs, file string) targetstructures.Output {
	content, _ := afero.ReadFile(fs, file)
	loaded := UnmarshalFile(content)

	return loaded
}
