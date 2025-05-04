# 📚 Bookworm

A simple Go program that tracks readers and their book collections.

## Current Features
- Loads reader (bookworm) data from JSON files
- Stores basic book information (author and title)
- Supports multiple readers and their book collections

## Structure
```go
type Bookworm struct {
    Name  string `json:"name"`
    Books []Book `json:"books"`
}

type Book struct {
    Author string `json:"author"`
    Title  string `json:"title"`
}
```


## Contributing
Feel free to open issues and submit PRs.

## License
MIT
