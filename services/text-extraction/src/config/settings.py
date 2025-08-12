from dotenv import load_dotenv
import os

class Env:
    def __init__(self):
        self.python_env = os.environ.get("PYTHON_ENV", "development")
        self.rabbitmq_host = os.environ.get("RABBITMQ_HOST", "localhost")
        self.input_queue = os.environ.get("INPUT_QUEUE", "RawScripts")
        self.output_queue = os.environ.get("OUTPUT_QUEUE", "ExtractedText")