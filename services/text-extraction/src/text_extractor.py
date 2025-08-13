import io
from typing import List, Dict
import fitz
from PIL import Image
from config.settings import Env
import pytesseract
import langextract as lx

cfg = Env()


class TextExtractor:
    """
    Extracts text from PDF files.
    """
    def __init__(self, dpi: int = cfg.DPI):
        """
        Initialize the TextExtractor with a specified DPI for image conversion.
        
        :param pdf_bytes: A PDF file as bytes.
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

                # Convert  to PNG bytes
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
            # Convert to grayscale
            image = image.convert("L")
            # Resize if necessary (optional, can be adjusted)
            image = image.resize((image.width // 2, image.height // 2))
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

    def structure_text(self, text: str) -> Dict:
        """
        Structure the extracted text into a JSON-like dictionary.

        :param text: The extracted text to structure.
        :return: A dictionary with the structured text.
        """
        if not text:
            raise ValueError("No text to structure.")

        result = lx.extract(
            text_or_documents=text.strip(),
            prompt_description=cfg.LANGEXTRACT_PROMPT,
            examples=cfg.LANGEXTRACT_EXTRACTION_EXAMPLE,
            model_id="gemini-2.5-flash",
            # TODO: Add ollama model support
        )

        structured_text = {
            "metadata": {
                "source": "PDF",
                "dpi": self.dpi
            },
            "extracted_text": result.text 
        }
        return structured_text