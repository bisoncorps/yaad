// +build !windows

package integration

import (
	"fmt"
	"testing"

	"github.com/bisoncorps/saido/driver"
	"github.com/bisoncorps/saido/inspector"
)

func TestDFonLocal(t *testing.T) {
	d := driver.Local{}
	// can either use NewDF() or get the interface and perform type assertion
	i := (inspector.GetInspector(`disk`)).(*inspector.DF)
	output, err := d.RunCommand(i.String())
	if err != nil {
		t.Error(err)
	}
	i.Parse(output)
	fmt.Printf(`%#v`, i.Values)
}
