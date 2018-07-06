# gitcdpath

helps to populate your CDPATH variable so you can `cd` directly to git repositories under one or more base paths (defaulted to current dir).

eg:

```sh
$ export CDPATH=".:$(gitcdpath ~/git)"
$ cd gitcdpath
/Users/stuart-warren/git/src/github.com/stuart-warren/gitcdpath
$
```

You may want/need to setup bash completion on OSX

```sh
$ brew install bash
$ chsh -s /usr/local/bin/bash
$ brew install bash-completion
$ source "$(brew --prefix bash-completion)/etc/bash_completion"
$ export CDPATH=".:$(gitcdpath ~/git)"
```