# Spinspirational API

This is an example Web API built using Go for a fictional Record Store, Spinspirational.

This exercise is my first time writing Go. At the start of this exercise, I used the Gin package to handle request routing. I appreciated the convenience of the built-in JSON handling of that package, but I wanted to experiement with creating a web API with 0 external dependencies. So I refactored to do eveything using Go's internal `net/http` and `encode.json` libraries. With Go 1.22, there are new implementations that easily provide the routing and route parsing required for this exercise.

This was a great size exercise to play with using the built-in standard libraries. If this were a bigger project, I would spend some time attempting to create some very focused abstractions that simplify JSON marshalling and error handling. If there was functionality that I wasn't able to easily achieve with my own abstractions, I might consider integrating a package like `gin` or `chi`.

I am not using a persistence layer with this API, but I plan on building an entirely functional Go WebAPI in a future project.
