---
layout: default
title: 🔱 Giwow
description: Manage your git projects with one command
---

[Giwow](./../index.html) > [Workspace](./index.html)

# `giwow workspace`

```
Usage:
  giwow workspace
  giwow workspace [command]

Available Commands:
  info        Display workspaces properties
  init        Initialize a new workspace
  set         Manage workspaces properties

Flags:
  -h, --help   help for workspace
```

## Create a new workspace

Usually, you will want to create a new workspace when you want to manage multiple projects.
First, navigate to the folder where you want to store your workspace.
Then, run the following command:

```shell
$ giwow workspace init <workspace-url>
```

It's very important to use a valid host URL if you plan to connect your workspace with `Gitlab` or `Github`. 
Indeed, the workspace url will be used to discover the projects.
Here are some examples of a valid workspace url:

- gitlab.com/my-group
- private.gitlab.com/my-org/my-department/my-group
- github.com/my-org/my-group
- etc 

To check information about your workspace, run the following command:

```shell
$ giwow workspace info
```

> You can add `--verbose` to see more information about your workspace.
