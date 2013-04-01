---
layout: default
---
## DESCRIPTION

The `direnv` project is a bash/zsh shell extension that allows you to load/unload environment variables depending on your path.

It has many uses but mine it to have project-specific settings so as to not clutter my `~/.profile`. I'm using it to specify ruby version, set AWS or SSH keys, ...

The `direnv` command-line is a proxy to the sub-commands and it work like the git(1) wrapper. You don't usually use that. Once the shell extension is loaded, most of the interaction is done in your `.envrc` files.

## EXAMPLE

    $ cd ~/code/my_project
    $ ls
    bin/ lib/ Rakefile README.md
    $ echo $PATH
    /usr/local/bin:/usr/local/sbin:/usr/bin:/bin:/usr/sbin:/sbin
    $ echo PATH_add bin > .envrc
    direnv: loading /Users/zimbatm/code/my_project
    $ echo $PATH
    /Users/zimbatm/code/my_project/bin:/usr/local/bin:/usr/local/sbin:/usr/bin:/bin:/usr/sbin:/sbin
    $ cd ..
    direnv: unloading /Users/zimbatm/code/my_project
    $ echo $PATH
    /usr/local/bin:/usr/local/sbin:/usr/bin:/bin:/usr/sbin:/sbin


## USAGE

Place a `.envrc` file in the folder you want and add some export(1) in it. If you cd(1) in that directory or child, the exported variables are available. If you cd(1) out of it, they are unloaded.

A range of utility functions are available in the `.envrc` context and are documented in the direnv-stdlib(1) page.

If the `.envrc` file is added/changed/removed or you is no longer in the path, the environment is reloaded.

## INSTALL

### 1. Install the code

    git clone http://github.com/zimbatm/direnv
    cd direnv
    make install
    # or symlink bin/direnv into your $PATH

For Mac users, you can also use `brew install direnv`

### 2. Add the hook for your shell

Add this line at the very end of your .bashrc or .zshrc:

    eval "$(direnv hook $0)"

Make sure it's even after rvm, git-prompt and other shell extensions that manipulate your prompt.

## HOW IT WORKS

The first thing, `direnv hook` is going tell your shell to eval the output of `direnv export` before any prompt is displayed. This is the magic sauce. Notice how we didn't call `direnv-hook`. This is because `direnv` is a command dispatcher like `git`.

Now when `direnv export` is executed, it checks your $PWD and parents for an .envrc file. If one is found, it loads it in bash and constructs a diff of the environment variables. It then prints the difference so that your shell can adjust, along with direnv-specific environment variables. This is how it stores it states and knows how to revert the changes when you cd to a different path.

Notice that the .envrc is just a bash script from which direnv extracts the environment variables that are `export`-ed. It also sources some utility functions that may be handy to you. See the direnv-stdlib(1) man page for these.

That's about it ! I try to keep the tool simple and as an enabler.

## CONTRIBUTE

Bug reports, contributions and forks are welcome.

For bugs, report them on <http://github.com/zimbatm/direnv/issues>

Or discuss by email <direnv@librelist.com>

Or if you have some cool usages of direnv that you want to share, feel free to put them in the wiki
<https://github.com/zimbatm/direnv/wiki>

## COPYRIGHT

Thank you for making direnv better

* Alexander Kobel for his patches
* Brian M. Clapper (aka. bmc) for his patch
* Joshua Peek (aka. josh) for his patch and support
* Laurie Young (aka. wildfalcon) for fixing my engrish
* Magnus Holm (aka. judofyr) for his patches and ideas
* Martin Aumüller (aka. aumuell) for his patches and awesomeness
* Peter Waller (aka. pwaller) for his patches and insights
* Sam Stephenson (aka. sstephenson) for his expand_path code that I stole from https://github.com/sstephenson/bats

Copyright (C) 2011 Jonas Pfenniger and contributors under the MIT licence.
