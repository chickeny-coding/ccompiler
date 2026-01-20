# ccompiler

## Description
A compiler which compiles `.cl` files.

## Syntax

The program has several blocks.

Every block is one of the following:

* A comment, which starts with `[` and ends with `]`. You can write down the comment content between `[` and `]`.

```cl
[e.g.]
[This is a comment]
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

> An identity is a string which includes neither whitespaces (which means its ascii isn't greater than $32$) nor the following tokens: `()[]{}`.

```
This-sentence-can-print-itself.
And this sentence will print itself separated by linefeed.
```

> You can also see the `syntax.txt` file in the project.

## Usage

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

```
./your_exe_code
```
