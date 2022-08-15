package pretty

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func PrettyPrint(data string) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(data), "", "    "); err != nil {
		fmt.Println("{}")
	} else {
		fmt.Println(prettyJSON.String())
	}
}
