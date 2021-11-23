# typeform

This package provides a Go interface for the [Typeform API](https://developer.typeform.com/get-started/). 

## Installation

```
go get github.com/roopeshvs/typeform
```

## Getting Started

The first thing is to request an API key from Typeform [here](https://developer.typeform.com/get-started/personal-access-token/). Once you have it, you can create a client to use Typeform programatically.

```go
c := tf.TypeformClient(os.Getenv("TYPEFORM_TOKEN"))
```

Export the API key as an environment variable and use it to create the client.

## Usage

Here are a few examples demonstrating how to create and access resources and data on Typeform.

<b>Create a Form</b>

```go
package main

import (
	"fmt"
	"os"

	tf "github.com/roopeshvs/typeform"
)

func main() {
	c := tf.TypeformClient(os.Getenv("TYPEFORM_TOKEN"))
	newForm := tf.Form{
		Title: "Sample Form",
		Type:  "quiz",
	}
	form, err := c.CreateForm(newForm)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(form.Title)
}
```
<b>Create an Image</b>

```go
func main() {
	c := tf.TypeformClient(os.Getenv("TYPEFORM_TOKEN"))
	image := tf.CreateImageRequestBody{
		FileName: "pattern.png",
		Image:    "iVBORw0KGgoAAAANSUhEUgAAAQAAAAEACAIAAADTED8xAAACaUlEQVR42uzVMRGAAAzAwLSHf8tgAAf95QVkyVNvNRN50FWBl10V6ABa0AFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIB6ADqEAHIB2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdAA6gBZ0ANIBSAcgHYB0ANIBSAcgHYB0ANIBSAcgHYB0ANIBSAcgHYB0ANIBSAcgHYB0ANIBSAcgHYB0ANIBSAcgHYB0ANIBSAcgHYB0ANIBSAcgHYB0ANIBSAcgHYB0ANIBSAcgHYB0ANIBSAcgHYB0ANIB6AAq0AFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgA6gAh2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdADyxy8AAP//YSoDD5pLB7MAAAAASUVORK5CYII=",
	}
	i, err := c.CreateImage(image)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(i.Src)
}
```
<b>Create a Theme</b>

```go
func main() {
	c := tf.TypeformClient(os.Getenv("TYPEFORM_TOKEN"))
	theme := tf.Theme{
		Name: "Go Theme Go",
		Font: "Bangers",
		Background: &tf.Background{
			Brightness: -1,
			Href:       "<USE THE IMAGE SRC OUTPUT FROM THE LAST EXAMPLE HERE>",
			Layout:     "fullscreen",
		},
		Colors: &tf.Colors{
			Answer:     "#0142AC",
			Background: "#FFFFFF",
			Button:     "#0142AC",
			Question:   "#000000",
		},
	}
	t, err := c.CreateTheme(theme)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t.Name)
}
```

## Contribution

Feel free to create a Github Issue regarding any bugs, feature request, or request to include an API method not supported by this package. Also, you are more than welcome to submit a pull request for the same.