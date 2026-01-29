# Ccompiler

## Description
A compiler which compiles `.cl` files.

## Syntax

### Program

Each program has several funcs. The format of each func is:

```cl
(
  func-name
  {
    func-code
  }
)
```

```
[ e.g. ]
[ A simple func ]
(
  f
  {
    func-f
  }
)
[ Or something more interesting ]
(
  f
  {
    f
    f
    (
    )
  }
)
```

### Block

Each block has several items and each item can be one of the following:

* An element, which is defined below.

* A block with `{` and `}` around.

```cl
[ e.g. ]
[ Multiple nesting ]
{
  {
    {
      {
        nested-text
      }
    }
  }
}
[ Nesting for a sentence ]
{
  Tip
  {
    use
    {
      2-space
      indentation
      which
      {
        symbolizes
        {
          binary
        }
        {
          the
          symbol
          of
          {
            computer
            science
          }
        }
      }
      instead
      of
      {
        {
          other-length-space
          indentation
        }
        or
        {
          tab
          indentation
        }
      }
    }
  }
}
```

### Element

Each element is one of the following:

* A comment, which starts with `[` and ends with `]`. You can write down the comment content between `[` and `]`.

```cl
[ e.g. ]
[ This is a comment ]
[
Multi
line
comment
is
also
allowed
]
```

* An identity or a number, which is actually an output statement that prints the identity or the number.

> An identity is a string which includes neither whitespaces (which means its ascii code isn't greater than $32$) nor the following tokens: `()[]{}`.

```
[ e.g. ]
This-sentence-can-print-itself.
And this sentence will print itself separated by linefeed.
```

* A func call, which has the func name and `()`.

```
[ e.g. ]
(
  f
  {
    [ f calls g ]
    g
    (
    )
  }
)
```

> You can also see the `syntax.txt` file in the project.

## Common Styles

### Comments

To make the comment more beautiful (it is just made for this purpose), always add a space between `[` and the comment text, and also between the text and `]`, just like add a space between the text and `//` or `#` in other languages.

### Indentations

Always use 2-space indentations, which symbolizes binary, the symbol of computer science, instead of other-length-space indentations or tab indentations.

## Usage

### Steps

* Clone the project with `Git`.

```bash
git clone git@github.com:chickeny-coding/ccompiler.git
```

* Compile the compiler codes with `Makefile`.

```bash
# Please guarantee you're in the ccompiler/ directory now.
make
```

* Added your own source code with suffix `.cl`.

* Use `ccompiler` to compile your source code into `.s`.

```bash
./ccompiler your_source_code.cl [--some_flags]
```

* Use `gcc` to compiler the `.s` code into `.exe`.

```bash
gcc your_source_code.s -o your_exe_code.exe
```

* And you can run it now.

```bash
./your_exe_code
```

### Flags

Here is all the flags to pass to ccompiler:

* `--info`

* `-i`

  Make ccompiler output the debug infos.
