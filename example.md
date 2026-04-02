mako will be a command that's run from the command line

whatever is defined in the mako header will be used to replace whatever is in the file

main command:
```mako <input path> <output path>```

prevent the same input and output path unless `-force` flag exists in the command

alternative (dedicated mako file)
```mako run <mako header file> <target file>```

```txt
<mako>

// inline numerical functions
def double(x): 2 * x;
def add(x,y): x + y;

// patterns that call functions
pattern "double [x:num]" = double(x);
pattern "add [x:num] [y:num]" = add(x,y);

// inline string functions with inline pattern matching
def concat(x,y) from "concat [x:str] [y:str]": x y;

// simple replacement
"name" = "mako";

// pattern replacement
pattern "number [x:num]" = x x

</mako>

double 10

add 10 20

number 10

concat he llo

Hello from name
```
would compile to
```txt
20

30

10 10

hello

Hello from mako
```
