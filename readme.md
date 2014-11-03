Radix tree in Go
==========

Quick experiment to speed up searchs form my http router.

http://en.wikipedia.org/wiki/Radix_tree


![Patricia trie](https://upload.wikimedia.org/wikipedia/commons/thumb/a/ae/Patricia_trie.svg/320px-Patricia_trie.svg.png)


Example:

```Go
	n := radix.New()
	n.Insert("romane", "romane")
	n.Insert("romanus", "romanus")
	n.Insert("romulus", "romulus")
	n.Insert("rubens", "rubens")
	n.Insert("ruber", "ruber")
	n.Insert("rubicon", "rubicon")
	n.Insert("rubicundus", "rubicundus")
````

Builds:

	r <nil>
	-om <nil>
	--an <nil>
	---e romane
	---us romanus
	--ulus romulus
	-ub <nil>
	--e <nil>
	---ns rubens
	---r ruber
	--ic <nil>
	---on rubicon
	---undus rubicundus
	
Search a path:

```Go
	v, ok := n.Lookup("rubicundus")
	if !ok {
		t.Fatal("path not found")
	}
````
