# HyperUtility

## Data structure

Simple implementation of **Stack** and **Queue** data structure.
Both structure were backend with array of `[]interface{}`

### Stack

Implement the Stack data structure FILO (First In Last Out).
In stack, the first data added into the stack can only
be retrieved the latest, and the last added data into the stack
can be retrieved immediately.

```text
+---+---+---+---+  <--- push
| 1 | 2 | 3 | n |
+---+---+---+---+  ---> pop
```


### Queue

Implement the Queue data structure FIFO (First In First Out).
In Queue, the first data added into the Queue can be
retrieved immediately, and the last added data into the stack
can be retrieved the latest.

```text
           +---+---+---+---+
pop <---   | 1 | 2 | 3 | n |  <--- push
           +---+---+---+---+ 
```

## io.Writer

### MultipleWriter

MultipleWriter easily write byte stream into multiple write stream.
Its very useful for the case like writing multiple log,
for example, one to the console `os.StdErr` and one to 
some external file writer.

```text
                       +--> writer_a.Write()
                       |
multiWriter.Write() ---+
                       |
                       +--> writer_b.Write()
```
