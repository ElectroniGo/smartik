import os
import asyncio
import sys
import pika
from config.settings import Env
from text_extractor import TextExtractor

def main() -> None:    
    cfg = Env()

    with pika.BlockingConnection(pika.ConnectionParameters(host=cfg.RABBITMQ_HOST)) as connection:
        channel = connection.channel()
        channel.queue_declare(queue=cfg.INPUT_QUEUE)
        channel.queue_declare(queue=cfg.OUTPUT_QUEUE)
        channel.basic_consume(
            queue=cfg.INPUT_QUEUE,
            on_message_callback=callback,
            auto_ack=True
        )
        channel.start_consuming()


def callback(ch, _method, _properties, body):
    # Convert body to in memory bytes (read pdf file into memory
    extractor = TextExtractor(dpi=Env.DPI)
    
    page_images = extractor.convert_to_image(body)

    extracted_text = ""
  
    # Extract text from the file (Tesseract)
    for image_bytes in page_images:
        # Clean the image (if needed)
        cleaned_image = extractor.clean_image(image_bytes)

        # Extract text from the image
        text = extractor.extract_text(cleaned_image)

        extracted_text += "\n" + text

    # Structure the text
    structured_text = extractor.structure_text(extracted_text)
    print(structured_text)

    # Read text into LangExtract
    # Force text into json structure
    # Send to output queue


if __name__ == "__main__":
    try:
        main()
    except  KeyboardInterrupt:
        try:
            sys.exit(0)
        except SystemExit:
            os._exit(0)
            raise