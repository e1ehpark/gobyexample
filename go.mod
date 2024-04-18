module gobyexample

go 1.22.0

require (
	cloud.google.com/go/compute v1.20.1 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	golang.org/x/oauth2 v0.19.0 // indirect
	gobyexample.com/googleoauth2/handlers
)

replace (
	obyexample.com/googleoauth2/handlers v0.0.0 => ./googleoauth2/handlers
)
