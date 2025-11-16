package property

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/jedib0t/go-pretty/table"
)

// Description: retrieves a property from a target.
func GetProperty(targetName string, property string, params ...string) (string, error) {

	// get function that manages that property
	fnHandler, ok := propertyMap[property]
	if !ok {
		return "", fmt.Errorf("unknown property requested: %s", property)
	}

	// play that function and get it output
	output, err := fnHandler(params...)
	if err != nil {
		return "", fmt.Errorf("error getting %s: %w", property, err)
	}

	return strings.TrimSpace(output), nil
}

func TodoShowMapProperty() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	// Simple header
	t.AppendHeader(table.Row{"Property Name"})

	// sort keys
	var listPropertyName []string
	for name := range propertyMap {
		listPropertyName = append(listPropertyName, name)
	}
	sort.Strings(listPropertyName)

	// Add rows
	for _, name := range listPropertyName {
		t.AppendRow(table.Row{
			name,
		})
	}

	// Render with default style
	t.Render()
}
