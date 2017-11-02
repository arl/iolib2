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
$ echo 'hello' | python client.py -p 1234 -u file://name=/tmp/myfile
Connecting to server...
Sending command  SET-PORT|file://name=/tmp/myfile ...
```

## Go daemon:
- install ZeroMQ version 4.0.1 or above on your system `libzmq` on linux.
- run: `go install`
- run: `iolib2 1234`

```sh
$ iolib2 1234
INFO[0000] waiting for requests or program termination  
INFO[0000] received request "SET-PORT|file://name=/tmp/myfile" 
DEBU[0000] received SET-PORT with param file://name=/tmp/myfile 
name=/tmp/myfile
DEBU[0000] fileport.set(map[name:/tmp/myfile])
INFO[0000] fileport, set filename to /tmp/myfile
Received reply  SET-PORT|file://name=/tmp/myfile [ 0| ]
Sending command  WRITE|hello

...
INFO[0000] received request "WRITE|hello

"
DEBU[0000] received WRITE with param hello


Received reply  WRITE|hello

[ 0| ]
Sending command  SEND| ...
INFO[0000] received request "SEND|"
DEBU[0000] received SEND with param
Received reply  SEND| [ 0| ]
Sending command  RESET| ...
INFO[0000] received request "RESET|"
DEBU[0000] received RESET with param
INFO[0000] reset current port, "file"
Received reply  RESET| [ 0| ]
```
