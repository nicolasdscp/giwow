---
layout: default
title: 🔱 Giwow
description: Manage your git projects with one command
---

[Giwow](/) > [Token](/token)

# giwow token

```
Usage:
  giwow token
  giwow token [command]

Available Commands:
  add         Add a new machine to your .netrc file
  ls          List your auth tokens
  rm          Remove a machine from your .netrc file
  set         Set values for a machine

Flags:
  -h, --help           help for token
      --netrc string   The path to the netrc file, default is $HOME/.netrc
```

## List your auth tokens

```shell
$ giwow token ls
```

## Adding a new machine to your .netrc file

Usually, you will want to add a new machine to your .netrc file manually. Giwow can do it for you.

```shell
$ giwow token add git.private.com
```
or
```shell
$ giwow token add git.private.com -u <username> -p <token>
```

> Giwow is supporting `default' machine in your .netrc file.

## Changing values for a machine

```shell
$ giwow token set git.private.com -u <username> -p <token>
```

Obviously, you can change only the username or the password.

## Removing a machine from your .netrc file

```shell
$ giwow token rm git.private.com
```