# got
Golang result building

# Examples
```go
func IsOdd(num int) got.Result {
	if num%2 == 0 {
		return got.Err(fmt.Sprintf("%d is not odd", num))
	}

	return got.Ok(true)
}

func GetAllContacts(num int) got.Result {
	result := make(chan got.Result)

	go func() {
		resp, err := goreq.Get("https://my-json-server.typicode.com/hadihammurabi/flutter-webservice/contacts", nil)
		if err != nil {
			result <- got.Err(err.Error())
		}
		defer resp.Body.Close()

		var data interface{}
		err = resp.Json(&data)
		if err != nil {
			result <- got.Err(err.Error())
		}

		result <- got.Ok(data)
	}()

	res := <-result
	return res
}
```
