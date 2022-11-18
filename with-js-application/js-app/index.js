const { connect, StringCodec } = require('nats');

const subject = 'my_subject';
const servers = 'localhost:4222';

async function demo() {
  const codec = StringCodec();
  const nc = await connect({ servers });

  nc.publish(subject, codec.encode('Hello there!'));

  console.log('Sent...');

  await nc.drain();
}

demo();
