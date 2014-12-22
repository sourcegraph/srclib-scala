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

The `src test` command writes to the `testdata/actual` tree.

You can also run `src` by itself in a checked-out copy of, e.g.,
`scala-pants-sample` (which is basically what `src test` does).

```
git clone https://github.com/sgtest/scala-pants-sample.git
cd scala-pants-sample
src config
find .srclib-cache # notice that the Scala source unit file is there now
src make
find .srclib-cache # notice that the graph and depresolve files also there now
cat .srclib-cache/*/org.srclib.scala/welcome/scala.graph.json # inspect the Scala graph output
```

If you run these commands in the `testdata/case/scala-pants-sample`
directory, be sure to `rm -rf .srclib-cache` before running `src test`
next time, or it'll complain.


## TODOs

1. [ ] Add a real scanner to detect Pants/sbt projects (can hardcode them for now, doesn't block building a grapher).
1. [ ] Create a real Scala grapher and make `.bin/srclib-scala graph` call it and pass-through its JSON output.
1. [ ] Make an actual pants-enabled test repo ([scala-pants-sample](https://github.com/sgtest/scala-pants-sample) only has a Scala and BUILD file, not a full pants installation).
1. [ ] Add a dep resolver (can use code from srclib-java, which can determine the group/artifact given a JAR filename and can attempt to look up the original VCS repository clone URL for a given group/artifact).
