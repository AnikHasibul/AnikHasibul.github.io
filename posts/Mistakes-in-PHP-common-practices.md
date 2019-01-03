# Mistakes in PHP's 2 most common practices.

> 1 January, 2018

## Introduction

Maybe still PHP is the most used language for web development. And it also has the top record for it's bad practices.

So in this post I am going to write about two commonly practiced php snippet.

## Start Exploiting

### $_SERVER[PHP_SELF]

We've seen this snippet for form action page reflection.

```php
<form action="<?php $_SERVER['PHP_SELF'] ?>"></form>
```

And this gives `<form action="/login.php"></form>` as output.

But what if we go to this url `example.com/login.php/1234`?

It will output like this `<form action="/login.php/1234"></form>`

So, if we go to this url `example.com/login.php/" onload="alert(1)`, it will reflect `<form action="/login.php/" onload="alert(1)"></form>`

XSS! ;)

### $_SERVER['HTTP_USER_AGENT']

Ever think that user agent may contain malicious inputs like SQL injection, HTML injection even sometimes system command injections!

```txt
APPLE' or '1'='1
APPLE <span onclick=alert(1)>Orange</span>
APPLE \r\r\r
```

So,

* Before you store an user agent to database must sanitize the input.
* Before you reflect an user agent to another user, be sure it's sanitized.
* Even before you `cat access_log.log` be sure it's safe to print in terminal.

## Conclusion

Nothing is secure. But you should try your best to keep your things as safe as possible. I hope, you've found this post informative. And as always, thanks for reading.

> Drop a comment if this post has any typo or other mistakes, let me fix it ASAP.
