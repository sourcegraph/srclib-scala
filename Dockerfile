FROM ubuntu:14.04

RUN apt-get update -q
RUN apt-get install -qq curl

RUN apt-get install -qq openjdk-7-jdk
RUN apt-get install -qq scala

# Add this toolchain
ADD . /srclib/srclib-scala/
WORKDIR /srclib/srclib-scala
ENV PATH /srclib/srclib-scala/.bin:$PATH

# Add srclib (unprivileged) user
RUN useradd -ms /bin/bash srclib
RUN mkdir /src
RUN chown -R srclib /src /srclib
USER srclib

# The repository being processed will be mounted as a read-only volume
# at /src.
WORKDIR /src

ENTRYPOINT ["srclib-scala"]
