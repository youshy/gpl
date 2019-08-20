# Strings

ASCII escape characters:

```golang
\a      "alert" or bell
\b      backspace
\f      form feed
\n      newline
\r      carriage return
\t      tab
\v      vertical tab
\'      single quote (only in the rune literal '\'')
\"      double quote (only within "..." literals)
\\      backslash
```

String literals by default use `"`. To use a *raw string literal* use ```` (backquotes). Within a raw string literal, not escape sequences are processed, everything within is taken literally.
