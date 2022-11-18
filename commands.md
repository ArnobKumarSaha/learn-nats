
## Stream 
nats stream add   # Creating a stream (with name `orders`, with subject `orders.*`)
nats stream ls
nats pub orders.us "{{.Count}}" --count 1000  # Publish 1000 messages in orders.us subject

nats sub --stream orders  # Read from the stream;  Options: --all, --new, --last, --last-per-subject <subject-name>, --start-sequence=<nth message & onward>
Above is an example of ephemeral Consumer. When some clients get connected with the stream, the stream gets cleaned up.  We need a Durable Consumer.

nats sub "orders.*" --durable my_consumer  # Stream will be cleaned after running this command.

### Sourcing
nats stream add --source <STREAM1> --source <STREAM2>



## Pull Consumer

All the above consumers were Push Consumer, where the server pushed the messages as fast as is possibly can.

nats consumer create  # Lets create a Pull Consumer with ecplicit Ack policy, & Start policy all. `orders` stream need to be selected also.
nats consumer next orders pull_consumer --count=1000  # pull 1000 message from orders stream to pull_consumer
nats consumer info




## KeyValue

nats kv add <bucket>
nats kv put <bucket> key value
nats kv put <bucket> key
nats object add <bucket>

nats stream ls -a  # Shows stream, Key-value & Object stores. As they are special kind of stream

