package radix

import "testing"

func TestBuildTree(t *testing.T) {
	n := New()
	n.Insert("romane", "romane")
	n.Insert("romanus", "romanus")
	n.Insert("romulus", "romulus")
	n.Insert("rubens", "rubens")
	n.Insert("ruber", "ruber")
	n.Insert("rubicon", "rubicon")
	n.Insert("rubicundus", "rubicundus")
	n.Debug()

	v, ok := n.Lookup("rubicundus")
	if !ok {
		t.Fatal("path not found")
	}

	if v.(string) != "rubicundus" {
		t.Fatal("wrong value")
	}
}
