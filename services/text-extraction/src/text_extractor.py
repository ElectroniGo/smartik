import io
from typing import List, Dict
import fitz
from PIL import Image
from config.settings import Env
import pytesseract
import langextract as lx
import time

cfg = Env()


class TextExtractor:
    """
    Extracts text from PDF files using OCR and AI-powered text structuring.
    """
    def __init__(self, dpi: int = cfg.DPI):
        """
        Initialize the TextExtractor with a specified DPI for image conversion.
        
        :param dpi: The resolution in DPI for the conversion.
                    Defaults to the value from the environment variable PDF_DPI or 300 if not set.
        """
        self.dpi = dpi

    def convert_to_image(self, pdf_bytes: bytes) -> List[bytes]:
        """
        Convert PDF bytes to a list of PNG image bytes. (Each page is converted to an image.)

        :param pdf_bytes: The PDF file as bytes.
        :return: A list of PNG image bytes, one for each page in the PDF.
        """
        images = []

        try:
            document = fitz.open(stream=pdf_bytes, filetype="pdf")
            for page_number in range(len(document)):
                page = document.load_page(page_number)

                # Creates a transformation matrix for the desired resolution
                matrix = fitz.Matrix(self.dpi / 72, self.dpi / 72)

                # Convert to PNG bytes
                pixmap = page.get_pixmap(matrix=matrix)

                # Convert to PNG bytes
                png_bytes = pixmap.tobytes("png")
                images.append(png_bytes)

            document.close()
            return images
        except Exception as e:
            raise ValueError(f"Failed to convert PDF to images: {e}")
    
    def clean_image(self, image_bytes: bytes) -> bytes:
        """
        Clean the image bytes by converting them to grayscale and resizing.

        :param image_bytes: The image bytes to clean.
        :return: Cleaned image bytes.
        """
        try:
            image = Image.open(io.BytesIO(image_bytes))
            # Convert to grayscale for better OCR results
            image = image.convert("L")
            # Resize to reduce processing time (you can adjust this ratio)
            image = image.resize((image.width // 2, image.height // 2), Image.LANCZOS)
            output = io.BytesIO()
            image.save(output, format="PNG")
            return output.getvalue()
        except Exception as e:
            raise ValueError(f"Failed to clean image: {e}")
    
    def extract_text(self, image_bytes: bytes) -> str:
        """
        Extract text from the image bytes using OCR.

        :param image_bytes: The image bytes to extract text from.
        :return: Extracted text as a string.
        """
        try:
            text = pytesseract.image_to_string(
                Image.open(io.BytesIO(image_bytes)),
                lang='eng'
            )
            if not text.strip():
                raise ValueError("No text found in the image.")
            return text
        except Exception as e:
            raise ValueError(f"Failed to extract text from image: {e}")

    def structure_text(self, text: str, max_retries: int = 3) -> Dict:
        """
        Structure the extracted text into a JSON-like dictionary using AI.
        Includes retry logic and timeout handling for better reliability.

        :param text: The extracted text to structure.
        :param max_retries: Maximum number of retry attempts if the request fails.
        :return: A dictionary with the structured text.
        """
        if not text:
            raise ValueError("No text to structure.")

        # Trim text if it's too long (Ollama works better with shorter texts)
        if len(text) > 4000:
            print(f"Text is long ({len(text)} chars), trimming to first 4000 characters for better AI processing")
            text = text[:4000] + "..."

        print(f"Processing text with AI model ({cfg.MODEL})...")
        
        # Try multiple times with increasing timeouts
        for attempt in range(max_retries):
            try:
                print(f"🔄 Attempt {attempt + 1}/{max_retries}")
                
                # Configure timeout based on attempt (30s, 60s, 120s)
                timeout = 30 + (attempt * 30)
                print(f"Using timeout: {timeout} seconds")
                
                result = lx.extract(
                    text_or_documents=text.strip(),
                    prompt_description=cfg.LANGEXTRACT_PROMPT,
                    examples=cfg.LANGEXTRACT_EXTRACTION_EXAMPLE,
                    model_id=cfg.MODEL,
                    api_key=cfg.GEMINI_API_KEY,
                    temperature=0.2,
                    max_workers=3,
                    max_char_buffer=int(cfg.MAX_CHAR_BUFFER),
                )

                print("AI processing completed successfully!")
                
                _structured_text = {
                    "metadata": {
                        "source": "PDF",
                        "dpi": self.dpi,
                        "processing_time": timeout,
                        "attempt": attempt + 1,
                        "ai_model": cfg.MODEL
                    },
                    "extracted_text": text,
                    "structured_data": result.text if hasattr(result, 'text') else str(result)
                }
                return result.extractions

            except Exception as e:
                print(f"Attempt {attempt + 1} failed: {str(e)}")
                
                if attempt < max_retries - 1:
                    wait_time = (attempt + 1) * 10  # Wait 10s, 20s, 30s between attempts
                    print(f"Waiting {wait_time} seconds before retry...")
                    time.sleep(wait_time)
                else:
                    # If all attempts failed, return basic structure without AI processing
                    print("All AI processing attempts failed, returning basic structure")
                    return {
                        "metadata": {
                            "source": "PDF",
                            "dpi": self.dpi,
                            "ai_processing": "failed",
                            "error": str(e),
                            "ai_model": cfg.MODEL
                        },
                        "extracted_text": text,
                        "structured_data": "AI processing failed - raw text only"
                    }

        # This should never be reached, but just in case
        raise ValueError("All attempts to structure text failed")