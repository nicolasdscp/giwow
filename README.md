# 🔱 Giwow [WIP]

Giwow is a simple `git` workspace manager. 
It allows you to manage all projects in a Gitlab group for example.
Clone, pull and push your projects in a single command. 
You can connect `giwow` with Jira to sync your backlogs and issue.

```
$ giwow
Git workspace manager

Usage:
  giwow [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  token       Manage your platform tokens. This will interact with your $HOME/.netrc file
  workspace   Manage workspaces

Flags:
      --debug     Enable debug and verbose messages, use in development only
  -h, --help      help for giwow
  -t, --toggle    Help message for toggle
      --verbose   Enable verbose messages
```

## 🚀 Getting started

Start by cloning this repository:

```shell
$ git clone https://github.com/nicolasdscp/giwow.git
```

Then install giwow

```shell
$ make install
```

## 🖥 Setting up your workspace

Choose a folder where your project's workspace will be stored. 
Then init a new workspace.

```shell
$ giwow workspace init my-workspace
```

### Working with private repositories

In order to work with private repositories, you need to configure `~/.netrc`.
You can let giwow manage your credentials for you, or you can configure it manually.
If you don't want to configure it manually, you can use the `giwow token add` command.

```shell
$ giwow token add <repo-url>
or
$ giwow token add <repo-url> -u my-user -p my-token
```

This will basically add a new entry to `~/.netrc` with the following content:

```shell
machine <repo-url> login my-user password my-token
```

You can list all your tokens with the `giwow token ls` command.