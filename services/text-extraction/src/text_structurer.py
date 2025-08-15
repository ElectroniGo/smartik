import re
import json
from typing import Dict, Any, List


class TextStructurer:
    """
    Helper class to structure extracted text into organized JSON format.
    Handles hierarchical question-answer structures with nested sub-questions.
    """
    
    def __init__(self):
        # Regex patterns to match different question formats
        self.main_question_pattern = r'(?:Question\s*)?(\d+)\.?\s*[:.-]?\s*(.*?)(?=(?:Question\s*)?\d+\.?|$)'
        self.sub_question_pattern = r'(\d+)\.(\d+)\.?\s*[:.-]?\s*(.*?)(?=\d+\.\d+|$)'
        self.letter_sub_pattern = r'([a-z])\)\s*(.*?)(?=[a-z]\)|$)'
    
    def structure_extracted_text(self, extracted_text: str) -> Dict[str, Any]:
        """
        Structure the extracted text into a hierarchical JSON format.
        
        :param extracted_text: The raw extracted text from OCR
        :return: Dictionary with structured question-answer format
        """
        if not extracted_text or not extracted_text.strip():
            return {"error": "No text provided for structuring"}
        
        try:
            # Clean and prepare the text
            cleaned_text = self._clean_text(extracted_text)
            
            # Extract main questions and their content
            structured_data = self._parse_questions(cleaned_text)
            
            # If no structured questions found, return raw text with metadata
            if not structured_data:
                return {
                    "metadata": {
                        "status": "no_structure_detected",
                        "original_text_length": len(extracted_text)
                    },
                    "raw_text": extracted_text
                }
            
            return structured_data
            
        except Exception as e:
            return {
                "error": f"Failed to structure text: {str(e)}",
                "raw_text": extracted_text
            }
    
    def _clean_text(self, text: str) -> str:
        """
        Clean the extracted text by removing extra whitespace and normalizing formatting.
        
        :param text: Raw extracted text
        :return: Cleaned text
        """
        # Remove excessive whitespace and normalize line breaks
        text = re.sub(r'\s+', ' ', text)
        text = re.sub(r'\n\s*\n', '\n', text)
        return text.strip()
    
    def _parse_questions(self, text: str) -> Dict[str, Any]:
        """
        Parse the text to extract questions and their hierarchical structure.
        
        :param text: Cleaned text to parse
        :return: Dictionary with structured questions
        """
        structured_data = {}
        
        # Split text into potential question blocks
        question_blocks = re.split(r'(?=(?:Question\s*)?\d+\.)', text)
        
        for block in question_blocks:
            if not block.strip():
                continue
                
            # Extract main question number and title
            main_match = re.match(r'(?:Question\s*)?(\d+)\.?\s*[:.-]?\s*(.*)', block, re.DOTALL)
            if not main_match:
                continue
                
            question_num = main_match.group(1)
            question_content = main_match.group(2).strip()
            question_key = f"Question {question_num}"
            
            # Parse sub-questions within this block
            sub_questions = self._parse_sub_questions(question_content)
            
            if sub_questions:
                structured_data[question_key] = sub_questions
            else:
                # If no sub-questions, store the entire content as the answer
                structured_data[question_key] = question_content
        
        return structured_data
    
    def _parse_sub_questions(self, content: str) -> Dict[str, Any]:
        """
        Parse sub-questions within a main question block.
        
        :param content: Content of a main question
        :return: Dictionary with sub-questions or None if no sub-questions found
        """
        sub_questions = {}
        
        # Look for numbered sub-questions (e.g., 1.1, 1.2, etc.)
        sub_matches = re.findall(self.sub_question_pattern, content, re.DOTALL)
        
        if sub_matches:
            for main_num, sub_num, sub_content in sub_matches:
                sub_key = f"{main_num}.{sub_num}"
                
                # Check if this sub-question has lettered parts (a, b, c, etc.)
                letter_parts = self._parse_letter_parts(sub_content.strip())
                
                if letter_parts:
                    sub_questions[sub_key] = letter_parts
                else:
                    sub_questions[sub_key] = sub_content.strip()
        
        return sub_questions if sub_questions else None
    
    def _parse_letter_parts(self, content: str) -> Dict[str, str]:
        """
        Parse lettered sub-parts within a sub-question (e.g., a), b), c)).
        
        :param content: Content of a sub-question
        :return: Dictionary with lettered parts or None if no letter parts found
        """
        letter_parts = {}
        
        # Look for lettered parts (a), b), c), etc.)
        letter_matches = re.findall(self.letter_sub_pattern, content, re.DOTALL)
        
        if letter_matches:
            for letter, part_content in letter_matches:
                letter_parts[letter] = part_content.strip()
        
        return letter_parts if letter_parts else None
    
    def format_as_json(self, structured_data: Dict[str, Any], indent: int = 2) -> str:
        """
        Format the structured data as a pretty-printed JSON string.
        
        :param structured_data: The structured data dictionary
        :param indent: Number of spaces for JSON indentation
        :return: Pretty-printed JSON string
        """
        try:
            return json.dumps(structured_data, indent=indent, ensure_ascii=False)
        except Exception as e:
            return json.dumps({"error": f"Failed to format JSON: {str(e)}"}, indent=indent)


def structure_output(extracted_text: str) -> Dict[str, Any]:
    """
    Helper function to structure extracted text output into organized JSON format.
    This is the main function you'll use in your text extraction pipeline.
    
    :param extracted_text: The raw text extracted from OCR
    :return: Structured dictionary with hierarchical question-answer format
    """
    structurer = TextStructurer()
    return structurer.structure_extracted_text(extracted_text)


# Example usage and testing function
def test_structurer():
    """
    Test function to demonstrate how the structurer works with sample text.
    """
    sample_text = """
    Question 1: What is machine learning?
    1.1 Machine learning is a subset of artificial intelligence
    1.2 It involves training algorithms on data
    
    Question 2: Types of machine learning
    2.1 Supervised learning
    a) Uses labeled data
    b) Examples include classification and regression
    2.2 Unsupervised learning finds patterns in unlabeled data
    """
    
    result = structure_output(sample_text)
    structurer = TextStructurer()
    print("Sample structured output:")
    print(structurer.format_as_json(result))


if __name__ == "__main__":
    test_structurer()