# translate
A small translate CLI tool using Google Translate without requiring an API key. It relies on Google Translate's internal web endpoint so it will probably break at some point.

## some examples

1. default -from=ro; default -to=en
```sh
> translate salutare tuturor
hello everyone
```
2. some non-defaults
```sh
> translate -from en -to ja hello there
こんにちは

> translate -from en -to sl "Are you there?"
Ste tam?
```

3. piping input
```sh
> echo "Hello. I like your hairstyle." | translate -from en -to es
Hola. Me gusta tu peinado.
```
4. input redirection
```sh
> echo "hello" > file.txt
> translate -from en -to es < file.txt
Hola
```
