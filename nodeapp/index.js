var amq = require("amqplib/callback_api");

amq.connect('amqp://localhost', (err1, conn) => {
    if(err1){
        throw err1;
    }
    
    conn.createChannel((err2, channel) => {
        if(err2){
            throw err2;
        }

        channel.assertQueue('hello', {
            durable: false
        });

        channel.sendToQueue("hello", Buffer.from("Hello"));

        console.log("[x] Sent %s", 'Hello');

        channel.assertExchange("logs", "fanout", {
            durable: false
        });

        channel.assertQueue("", {
            exclusive: true
        },
        (error, q) => {
            if(error){ 
                throw error
            }

            console.log(" [*] Waiting for messages in %s. To exit press CTRL+C", q.queue);
            channel.bindQueue(q.queue, 'logs', '');

            channel.consume(q.queue, (message) => {
                console.log(q.queue)
                if(message.content) {
                    console.log(" [x] Receiving %s", message.content.toString());
                  }
            }, {
                noAck: true
            });
        });
    });
});