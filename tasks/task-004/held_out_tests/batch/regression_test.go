package batch

import "testing"

func TestRegressionOrdering(t *testing.T) {
	encoder := &Encoder{}
	output, _ := encoder.Encode([]Event{{ID: "first"}, {ID: "second"}})
	if string(output) != "{\"id\":\"first\",\"payload\":\"\"}\n{\"id\":\"second\",\"payload\":\"\"}\n" {
		t.Fatalf("%s", output)
	}
}
