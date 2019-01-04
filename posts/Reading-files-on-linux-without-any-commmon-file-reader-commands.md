# Reading files on linux without any common file reader commands!

> 2 January, 2018

## Introduction:
Hello there, here I'm gonna write about few ways to read files on linux without any common file reader commands.

## In action:

### base64

```shell
$ base64 /etc/passwd | base64 -d
```
#### How it works?

This command first encode a file to base64, then we passed the output to decode the base64 data. And we will get the actual file data.

### bash

```shell
$ bash -c 'echo "$(</etc/passwd)"'
```

#### How it works?

we asked `bash` to echo the `stdin`, and the `stdin` is containing the body of `/etc/passwd`.


### cut 

```shell
$ cut -d "" -f1 /etc/passwd
```

#### How it works?

It reads data from files. Its a file reader but a bit uncommon.


### date

```shell
$ date -f /etc/passwd
```

## Conclusion:

I hope you found this post helpful. And as always, thanks for reading.


> Drop a comment if this post has any typo or other mistakes, let me fix it ASAP.
