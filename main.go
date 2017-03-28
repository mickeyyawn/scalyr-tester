package main

func main() {

	//var attrs map[string]string
	//attrs = make(map[string]string)

	//var attrs interface{}
	//err := json.Unmarshal([]byte(""), &attrs)

	//otherAttrs := attrs.(map[string]interface{})

	//m := attrs.(map[string]interface{})
	//attrs["completelynewattribute"] = "NEW ATTR"
	//attrs["this will be a number"] = 42

	/*
		b := []byte("")

		var attrs interface{}
		err := json.Unmarshal(b, &attrs)
		if err != nil {
			panic("trouble unmarshalling")
		}

		m := attrs.(map[string]interface{})

		m["completelynewattribute"] = "NEW ATTR HERE AND ...."
		m["this will be a number"] = 55

	*/

	m := map[string]interface{}{"a": "apple", "b": 2}
	// or
	z := map[string]interface{}{}
	z["c"] = "this is c"
	z["d"] = "pear"
	// or
	h := map[string]interface{}{}
	h["f"] = "this is f"
	h["message"] = "my message is here"

	// or

	l := map[string]interface{}{"a": "banana", "b": 22, "message": "my most awesome message is here...."}

	Event(Warning, m)
	Event(Info, z)
	Event(Info, h)
	Event(Info, l)

	for i := 0; i < 3; i++ {
		Event(Info, l)
	}

}
