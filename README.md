# srclib-scala [![Build Status](https://travis-ci.org/sourcegraph/srclib-scala.png?branch=master)](https://travis-ci.org/sourcegraph/srclib-scala)

Work-in-progress [srclib](http://srclib.org) toolchain for
[Scala](http://scala-lang.org/) source code and dependency analysis.


## Usage

First, install [srclib](https://srclib.org).

Then, to set up this toolchain, clone this repository. Then from the directory
you cloned it to, run:

```
src toolchain add sourcegraph.com/sourcegraph/srclib-scala
```

That adds a symlink to this directory in your SRCLIBPATH
(`$HOME/.srclib` by default).

Now, running `src toolchain list` should show this toolchain

```
$ src toolchain list
PATH                                       TYPE
...
sourcegraph.com/sourcegraph/srclib-scala  program, docker
```

Next, build this toolchain's Docker image:

```
src toolchain build sourcegraph.com/sourcegraph/srclib-scala
```

Now try running the tests, in both variants: invoking this toolchain as a
program and invoking it in a Docker container.

```
git submodule update --init # Initializes the testing submodules.
src test -m docker
src test -m program
```
