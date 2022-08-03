ls.md---
layout: default
title: 🔱 Giwow
description: Manage your git projects with one command
---

[Giwow](/) > [Token](/token) > [add](/token/add.html)

# giwow token add [machine]

```
[machine] is the name of the machine to add to your .netrc file.

Usually, it is the name of the git host you want to access eg: private.gitlab.com 
or github.com/my-private.
This will basically add a new line in your $HOME/.netrc file. 

You can avoid the interactive terminal by using the --login (-u) and --password (-p) flags.
Note that it's highly recommended to generate a personal access token on your git host
instead of using your password.

If a similar entry already exists in your .netrc file, it will be overwritten.

Usage:
  giwow token add [machine] [flags]

Flags:
  -h, --help              help for add
  -u, --login string      Directly set the login user
  -p, --password string   Directly set the password
```