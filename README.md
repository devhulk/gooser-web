# gooser-web

Go with HTMX and Templ UI based on the cmd line version of this project "gooser". Searches top ten sites to see if username is being used. Basic OSINT tool based on WhatsMyName project.

After cloning you can run the project locally with
```go run *.go```

You will need Templ installed if you want to edit and regen templ files.


Running with Docker:
```
go mod tidy
docker build -t gooser-web:latest .
docker run -p 8080:8080 gooser-web:latest
```