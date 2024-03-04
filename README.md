# Spinspirational API

This is an example Web API built using Go for a fictional Record Store, Spinspirational.

This exercise is my first time writing Go. A the start of this project, I used Gin as a framework to handle routing. I appreciated the convenience of the JSOn handling built in, but I wante to experiement with creating a web API with 0 external dependencies, so I refactored to do eveything using Go's internal `net/http` and `encode.json` libraries.

ince this is a small project with a handful of routes that needed to be handled, I could see the desire to use internal libraries to keep dependencies at 0 for a small application. However, if this were a much larger library, I would might consider a package like `gin` or `chi` to cut down on boilerplate and increase productivity.

I am not using a persistence layer with this API, but I plan on building an entirely functional Go WebAPI in a future project.
