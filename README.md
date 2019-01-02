# datauri
datauri converts images into Base64 encoded string.

## Usage
```golang
import (
    fmt
    os
    github.com/gekkotokio/datauri
)

func main() {
    if encodedString, err := datauri.Encoder("path/to/image.jpg"); err != nil {
        # do something 
    }

    # print out Base64 encoded string.
    fmt.Println(encodedString)
}
```

## TODO
- Add Decoder method
- Add the function to indetify image format