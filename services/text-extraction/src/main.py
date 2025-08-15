import os
import asyncio
import sys
import json
import pika
from config.settings import Env
from text_extractor import TextExtractor
from text_structurer import structure_output  # Import our new helper function

cfg = Env()


def main() -> None:    
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
    """
    Process incoming PDF files and extract structured text.
    """
    try:
        # Initialize the text extractor
        extractor = TextExtractor(dpi=cfg.DPI)
        
        # Convert PDF to images (one per page)
        page_images = extractor.convert_to_image(body)

        # Extract text from each page
        extracted_text = ""
        for image_bytes in page_images:
            cleaned_image = extractor.clean_image(image_bytes)
            text = extractor.extract_text(cleaned_image)
            extracted_text += "\n" + text

        # Structure the extracted text using AI (from text_extractor.py)
        ai_structured_text = extractor.structure_text(extracted_text)
        for item in ai_structured_text:
            print(f"AI Structured Text: {item}\n")
        
        # Use our helper function to organize the data into neat JSON structure
        # organized_output = structure_output(extracted_text)
        
        # Combine AI processing with our structured format
        # final_output = {
        #     "ai_structured_data": ai_structured_text,
        #     "organized_questions": organized_output,
        #     "metadata": {
        #         "total_pages": len(page_images),
        #         "processing_status": "completed",
        #         "text_length": len(extracted_text)
        #     }
        # }
        
        # Print the structured output (you can see this in the terminal)
        print("=" * 50)
        print("STRUCTURED OUTPUT:")
        print("=" * 50)
        print(json.dumps(final_output, indent=2, ensure_ascii=False))
        
        # TODO: Send final_output to the RabbitMQ output queue
        # ch.basic_publish(
        #     exchange='',
        #     routing_key=cfg.OUTPUT_QUEUE,
        #     body=json.dumps(final_output)
        # )
        
    except Exception as e:
        error_output = {
            "error": str(e),
            "status": "processing_failed"
        }
        print("ERROR:", json.dumps(error_output, indent=2))


if __name__ == "__main__":
    try:
        main()
    except KeyboardInterrupt:
        try:
            sys.exit(0)
        except SystemExit:
            os._exit(0)
            raise