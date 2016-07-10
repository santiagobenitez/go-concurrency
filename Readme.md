#Golang concurrency example
I am getting started with Go language and as a first example I wanted to try the concurrency model of Go. The idea of this code is to have a shared resource, in this case is the people variable which is a slice of Person. It's simulating the access to a database to get all of the persons only once and then caching them in memory. Next calls to get the all the persons will read from memory instead of db.

In order to simulate concurrency I created several goroutines running in parallel and asking continously for all of the persons. The output will print the number of operations done, the number of db reads done and the number of cache reads.
