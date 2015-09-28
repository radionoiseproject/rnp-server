Radio Noise Project (server component)
======================================

This is the server component required for the Radio Noise Project. It
serves as an introduction point for peers in the same room (WebRTC
signalling relay), and is responsible for some coordination of user roles
and song synchronization in rooms.

Building rnp-server
-------------------

Create a go project root to use:

```
mkdir radionoiseproject; cd radionoiseproject
export GOPATH=`pwd`
export PATH=$GOPATH/bin:$PATH
```

fetch and build rnp-server:

```
go get github.com/radionoiseproject/rnp-server
go install github.com/radionoiseproject/rnp-server
```

and then you can run the `rnp-server` executable.
