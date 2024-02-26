## Mutex Review

The principle problem that mutexes help us avoid is the <em>concurrent
read/write problem</em>. This problem arises when one thread is writing to a
variable while another thread is reading from that same variable <em>at the same
time</em>.

When this happens, a Go program will panic because the reader could be reading
bad data while it's being mutated in place.

```
----------         --------------------         ----------
|        |  wait   |       Mutex      |  wait   |        |
| THREAD |<------------->|______|<------------->| THREAD |
|        | release |                  | release |        |
|        |         |  shared resource |         |        |
----------         --------------------         ----------
```
