import os
import sys
import pika
from config.settings import Env

def main() -> None:    
    cfg = Env()

    with pika.BlockingConnection(pika.ConnectionParameters(host=cfg.rabbitmq_host)) as connection:
        channel = connection.channel()
        channel.queue_declare(queue=cfg.input_queue)
        channel.queue_declare(queue=cfg.output_queue)
        channel.basic_consume(
            queue=cfg.input_queue,
            on_message_callback=callback,
            auto_ack=True
        )
        channel.start_consuming()


def callback(ch, _method, _properties, body):
    # Convert body to in memory bytes
    # Extract text from the file (Tesseract)
    
    # Read text into LangExtract
    # Force text into json structure
    # Send to output queue
    print(f"Received message: {body}")


if __name__ == "__main__":
    try:
        main()
    except  KeyboardInterrupt:
        try:
            sys.exit(0)
        except SystemExit:
            os._exit(0)
            raise