# Gogs Cli

[![Join the chat at https://gitter.im/andreynering/gogscli](https://badges.gitter.im/andreynering/gogscli.svg)](https://gitter.im/andreynering/gogscli?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

> This is an experimental project. By now it's just a proof of concept.

An attempt to create a command line tool to interact with a [Gogs][gogs]
instance.

## Installation

This assumes you have a [Golang][golang] installation working, and `GOPATH/bin`
in you `PATH` environment variable.

```bash
go get -u https://github.com/andreynering/gogscli/...
```

## Configuration

Configuration can be made by command line, and it's saved on `$HOME/gogscli.ini`.

### Remote Gogs instance

Set the HTTP/HTTPS URL of your Gogs instance:

```bash
gogscli config --remote=https://git.yourdomain.com
```

### Token

In order to have access to the Gogs API, you have to create and configure the
access token. To create it, navigate to *User* -> *Configuration* ->
*Applications* (`/user/settings/applications`) in your Gogs instance.
Add it to Gogs Cli by doing:

```bash
gogscli config --token=YOUR_TOKEN_HERE
```

## Issue

### List issues

```bash
gogscli issue list user/repo
```

### Add issue

```bash
# syntax
gogscli issue add user/repo "Issue title" ["Optional issue body"] [--assignee=optionalassignee]

# example
gogscli issue add andreynering/gogscli "Please fix X" "etc, etc..." --assignee=andreynering
```

## More

More to be added. Contributions are welcome.

[gogs]: https://gogs.io/
[golang]: https://golang.org/
