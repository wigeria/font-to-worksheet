# Font to Worksheet

Takes a TTF file and a character (alphabet/number) as input, and generates a PDF work sheet for handwriting practice with that font, for that character.


## Example Usage

```
font-to-worksheet --font test.ttf --character a
```


## Notes

Built mainly with ChatGPT-3


## TODO

- Use https://github.com/jung-kurt/gofpdf
- Load in the TTF file, write the font in even spaces
    - Find horizontal spacing by getting bounding box of character and dividing it by page width
    - Bounding box can be found using https://pkg.go.dev/github.com/golang/freetype/truetype and/or https://pkg.go.dev/github.com/golang/freetype/truetype#Font
- After row 1, next rows should have lower opacity (for practice)
