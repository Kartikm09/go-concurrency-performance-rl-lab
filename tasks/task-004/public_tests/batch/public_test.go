package batch

import "testing"

func TestPublicOutputCompatibility(t *testing.T) {
	encoder := &Encoder{}
	output, err := encoder.Encode([]Event{{ID: "1", Payload: "a"}, {ID: "2", Payload: "b"}})
	if err != nil {
		t.Fatal(err)
	}
	expected := "{\"id\":\"1\",\"payload\":\"a\"}\n{\"id\":\"2\",\"payload\":\"b\"}\n"
	if string(output) != expected {
		t.Fatalf("%q", output)
	}
}
