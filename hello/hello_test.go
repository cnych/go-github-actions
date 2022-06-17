package hello

import "testing"

func TestGreet(t *testing.T) {
	result := Greet()
	if result != "Hello GitHub Actions. qikqiak.com is awesomefwef" {
		t.Errorf("Greet() = %s; Expected Hello GitHub Actions. qikqiak.com is awesome fwef", result)
	}

}
