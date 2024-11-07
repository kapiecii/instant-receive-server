package main

import (
    "fmt"
    "html/template"
    "io"
    "net/http"
    "os"
    "time"
)

var tmpl = template.Must(template.New("upload").Parse(`
<!DOCTYPE html>
<html>
<head>
    <title>Upload File</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f4f4f4;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            margin: 0;
        }
        .container {
            background-color: #fff;
            padding: 20px;
            border-radius: 8px;
            box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
        }
        h1 {
            margin-bottom: 20px;
        }
        input[type="file"] {
            margin-bottom: 10px;
        }
        input[type="submit"] {
            background-color: #4CAF50;
            color: white;
            border: none;
            padding: 10px 20px;
            text-align: center;
            text-decoration: none;
            display: inline-block;
            font-size: 16px;
            margin: 4px 2px;
            cursor: pointer;
            border-radius: 4px;
        }
        input[type="submit"]:hover {
            background-color: #45a049;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Upload File</h1>
        <form enctype="multipart/form-data" action="/upload" method="post">
            <input type="file" name="file" />
            <input type="submit" value="Upload" />
        </form>
    </div>
</body>
</html>
`))

func uploadHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        tmpl.Execute(w, nil)
        return
    }

    file, handler, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "Error retrieving the file: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer file.Close()

    // Create uploads directory if it doesn't exist
    if err := os.MkdirAll("./uploads", os.ModePerm); err != nil {
        http.Error(w, "Error creating uploads directory: "+err.Error(), http.StatusInternalServerError)
        return
    }

    dst, err := os.Create("./uploads/" + handler.Filename)
    if err != nil {
        http.Error(w, "Error saving the file: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer dst.Close()

    if _, err := io.Copy(dst, file); err != nil {
        http.Error(w, "Error copying the file: "+err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "File uploaded successfully: %s", handler.Filename)
}

func main() {
    server := &http.Server{Addr: ":8080"}

    http.HandleFunc("/upload", uploadHandler)
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl.Execute(w, nil)
    })

    go func() {
        fmt.Println("Server started at :8080")
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            fmt.Printf("ListenAndServe(): %s\n", err)
        }
    }()

    time.Sleep(180 * time.Second)
    fmt.Println("Shutting down the server...")
    if err := server.Close(); err != nil {
        fmt.Printf("Server Close: %s\n", err)
    }
}