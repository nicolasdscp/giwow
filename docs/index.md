---
layout: default
title: 🔱 Giwow
description: Manage your git projects with one command
---

<a href="https://go.dev/blog/go1.18">
    <img alt="Go version" src="https://img.shields.io/badge/go-1.18-blue" />
</a>
<a href="https://github.com/nicolasdscp/giwow/actions">
    <img alt="Lint passing" src="https://github.com/nicolasdscp/giwow/actions/workflows/golangci-lint.yml/badge.svg" />
</a>
<a href="https://giwow.run">
    <img alt="Documentation" src="https://img.shields.io/badge/documentation-giwow.run-blue" />
</a>

Giwow is a simple `git` workspace manager.
It allows you to manage all projects in a Gitlab group for example.
Clone, pull and push your projects in a single command.
You can connect `giwow` with Jira to sync your backlogs and issue.

## 🗺 Website map

- [Giwow](/)
    - [Getting Started](/getting-started.html)
    - [Workspace](/workspace)
    - [Token](/token)
    - [Projects](/projects)

## ⚙️ Usage

```
$ giwow
Git workspace manager

Usage:
  giwow [command]

Available Commands:
  help        Help about any command
  projects    Manage projects in the current workspace
  prune       Delete all configuration files and all workspaces
  token       Manage your platform tokens. This will interact with your $HOME/.netrc file
  workspace   Manage workspaces

Flags:
      --debug     Enable debug and verbose messages, use in development only
  -h, --help      help for giwow
  -t, --toggle    Help message for toggle
      --verbose   Enable verbose messages
```

## 🌍 How to use this wiki

If you want information about a command, you can directly access to its documentation by reaching `giwow.run/command/subcommand`.



