# Font to Worksheet

Takes a TTF file and a character (alphabet/number) as input, and generates a PDF work sheet for handwriting practice with that font, for that character.


## Notes

Built almost entirely to test ChatGPT's capabilities, with some direct modifications to the final result to add clarity and fix some silly bugs. Here are the series of prompts that were used:

```
Write Golang Code to read a Truetype Font file and create a PDF handwriting worksheet for a given letter

...

Modify that code to write the character in 12 rows of 10 columns, and every row except the first should have 0.5 opacity

...

The script above simply writes characters on multiple lines, instead of characters in tables. Please fix it.

...

Can you center the table on the page in the script above?

...

With the code above, the first row is not horizontally centered. And everything else is also not vertically centered.

```

At the end of all this, Chat-GPT was spitting out a table, with the characters being written out in a single column, outside of the page boundaries. I was impatient, and just fixed it myself.


**Disclaimer:** The TTF Font (Gashington Classy) included in the repository does not belong to me.
