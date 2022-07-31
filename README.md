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

## 🚀 Getting started

### 📦 Installation

```shell
$ go install github.com/nicolasdscp/giwow
```
or 

```shell
$ git clone https://github.com/nicolasdscp/giwow.git
$ make install
```

## 🖥 Setting up your workspace

Choose a folder where your project's workspace will be stored. 
Then init a new workspace.

```shell
$ giwow workspace init my-workspace
```

### 🕵🏼‍ Working with private repositories

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

### 🦊 Gitlab connection

Giwow will use your `.netrc` file to connect to your Gitlab account.
First you need to authenticate with your Gitlab account. 
It's highly recommended to use a personal access token instead of your password.
Then configure giwow to use your Gitlab account.
Note that if tou already have a personal access token in your `.netrc`, you can skip this step.

```shell
$ giwow token add gitlab.com -u gitlab-username -p access-token
```
> Note that you can replace `gitlab.com` with your Gitlab domain.

> To generate a personal access token see [Gitlab personal access tokens](https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html)

Now set the type of your workspace to `gitlab`.

```shell
$ giwow workspace set --type gitlab
```

**Working with gitlab allows you to automatically discover and clone all your projects.**

## 📚 Working with projects

Now you can manage your projects with the `giwow projects` command.

### Adding your projects

```shell
$ giwow projects add <project1> <project2> ...
```

> You can add the `--clone` flag to clone projects directly.

If you are using a Gitlab workspace you can automatically discover all your projects in your workspace.

```
$ giwow projects discover
> 🔎 Discovering projects in the current workspace ...
> 🦊 Using Gitlab API to discover projects ...
> 🎉 Discovered 13 projects
> ✅ 13 project(s) added to the workspace
```

> Note that `discover` will also add projects in sub groups. 
> Giwow will automatically clone these projects in the good hierarchy.

### Cloning your projects

Now you probably want to clone your projects.

```shell
$ giwow projects clone
```

## 🛟 Useful commands

If you want to know more about a command, you can use the `giwow [command] -h` command.

```shell
$ giwow workspace info # Show workspace info
$ giwow projects prune # Delete all projects present in the workspace file (Not projects by themself)
$ giwow prune # Remove all configuration file from all workspaces including $HOME/.giwow
```