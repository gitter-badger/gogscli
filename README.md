# Gogs Cli

> This is an experimental project. It is not ready for production use, and may
never be. It's just a proof of concept.

An attempt to create a command line tool to interact with a [Gogs][gogs]
instance.

## Issue

### List issues

```bash
./gogscli issue list user/repo
```

### Add issue

```bash
# syntax
./gogscli issue add user/repo "Issue title" ["Optional issue body"] [--assignee=optionalassignee]

# example
./gogscli issue add andreynering/gogscli "Please fix X" "etc, etc..." --assignee=andreynering
```

[gogs]: https://gogs.io/
