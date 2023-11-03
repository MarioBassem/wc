# MWC

mwc is a tool to count words, lines, bytes, and characeters in a file.

## Usage

> `mwc [OPTION] FILE`

### Available options

- `-c`: print the byte count
- `-w`: print the word count
- `-l`: print the line count
- `-m`: print the character count

If no flag is provided, it works as if the flags `-l`, `-w`, `-c` are provided, in that order.