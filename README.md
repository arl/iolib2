# Example of use:

## Python example client:
- cd `client`
- run: `virtualenv -p $(which python2) .`
- enable the virtualenv with `source ./bin/activate`
- install python deps: `pip install -r requirements.txt`
- Even if the daemon is not started yet, we can still use the client to send
  some text, that's not a problem thanks to the use of ZeroMQ :-D.
- Let's choose the `file://` output and write to `/tmp/myfile`

```sh
$ echo 'hello' | python client.py -p 1234 -u file://name=/tmp/world
Connecting to daemon...
('set_port ->', (0, None))
('write ->', (0, None))
('send ->', (0, None))
('reset_port ->', (0, None))
```

## Go daemon:
- install ZeroMQ version 4.0.1 or above on your system, (`libzmq` on linux)
- run: `go install`
- run: `iolib2 1234` to start the daemon on port `1234`. In the `iolib2` python
  module the `IOLin2Daemon` class acts as a wrapper around the Go executable, it
  should be used as a singleton.

```sh
$ iolib2 1234
INFO[0000] waiting for requests or program termination
INFO[0025] received request "SET-PORT|file://name=/tmp/world"
DEBU[0025] received SET-PORT with param file://name=/tmp/world
name=/tmp/world
DEBU[0025] fileport.set(map[name:/tmp/world])
INFO[0025] fileport, set filename to /tmp/world
INFO[0025] received request "WRITE|hello
"
DEBU[0025] received WRITE with param hello

INFO[0025] received request "SEND|"
DEBU[0025] received SEND with param
INFO[0025] received request "RESET|"
DEBU[0025] received RESET with param
INFO[0025] reset current port, "file"
```
