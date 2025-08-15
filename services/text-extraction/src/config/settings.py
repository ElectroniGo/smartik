from dotenv import load_dotenv
import textwrap
import os
import langextract as lx

load_dotenv()

class Env:
    DPI = int(os.environ.get("PDF_DPI", 300))
    PYTHON_ENV = os.environ.get("PYTHON_ENV", "development")
    RABBITMQ_HOST = os.environ.get("RABBITMQ_HOST", "localhost")
    INPUT_QUEUE = os.environ.get("INPUT_QUEUE", "RawScripts")
    OUTPUT_QUEUE = os.environ.get("OUTPUT_QUEUE", "ExtractedText")
    MODEL = os.environ.get("OLLAMA_MODEL", "gemini-1.5-flash")
    MAX_CHAR_BUFFER = int(os.environ.get("MAX_CHAR_BUFFER", 1000))
    GEMINI_API_KEY = os.environ.get("GEMINI_API_KEY", None)

    # Example text for langextract
    LANGEXTRACT_EXAMPLE_TEXT = textwrap.dedent("""\
        Name: John Smith
        Exam Number: 12345678
        Subject: English Paper 1
        Question 1
        1.1 The digital revolution means the big change in technology and how we use computers and phones.
        Question 2
        2.1 Adjective
        2.2 A new social media platform has been launched by the company.
        2.8 a) Unheard of
            b) Improved
            c) Move through
            d) Obvious""")

    # Langextract prompt for extracting structured text
    LANGEXTRACT_PROMPT = textwrap.dedent("""\
        Extract the exam number, answer section, question number and the answer to the question
        as they appear in the text.
        Do not paraphrase or change anything in the text.""")

    # Example data for langextract
    LANGEXTRACT_EXTRACTION_EXAMPLE = [
        lx.data.ExampleData(
            text=LANGEXTRACT_EXAMPLE_TEXT,
            extractions=[
                lx.data.Extraction(
                    extraction_class="exam number",
                    extraction_text="12345678",
                    attributes={
                        "type": "metadata",
                        "category": "student details"
                    }
                ),
                lx.data.Extraction(
                    extraction_class="answer section",
                    extraction_text="Question 1",
                    attributes={
                        "type": "section",
                        "category": "exam questions segment"
                    }
                ),
                lx.data.Extraction(
                    extraction_class="question number",
                    extraction_text="1.1",
                    attributes={
                        "type": "question",
                        "category": "exam question order"
                    }
                ),
                lx.data.Extraction(
                    extraction_class="answer",
                    extraction_text="The digital revolution means the big change in technology and how we use computers and phones.",
                    attributes={
                        "type": "answer",
                        "category": "exam question answer"
                    }
                )
            ]
        )
    ]
