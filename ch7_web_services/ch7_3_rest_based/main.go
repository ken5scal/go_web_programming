package main

// difference between POST and PUT is
// in order to create a new resource w/o knowing the URL, use POST
// but if you want to replace an existing resource use PUT

// PUT is idempotent and the state of the server
// doesn’t change regardless of the number of times PUT call
// It will create a resource or to modify an existing resource,
// only one resource is being created at the provided URL
// ex PUT /users/1

// But POS isn't idempotent; ; every time you call it, POST will create a resource, with a new URL
// ex POST /users

func main() {
}
