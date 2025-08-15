"""
Exam Answer Extractor Module

This module provides functionality to extract and structure exam answers from 
LangExtract extraction results into a hierarchical JSON format suitable for 
exam paper processing.
"""

import json
import re
from typing import Dict, List, Any, Optional
from collections import defaultdict

# Constants to avoid string duplication
QUESTION_NUMBER_CLASS = "question number"
ANSWER_SECTION_CLASS = "answer section"
EXAM_NUMBER_CLASS = "exam number"


class ExamAnswerExtractor:
    """
    Extracts and structures exam answer data from LangExtract extraction results.
    
    Processes a list of Extraction objects and organizes them into a hierarchical
    JSON structure with questions, sub-questions, and nested answers.
    """
    
    def __init__(self):
        """Initialize the ExamAnswerExtractor."""
        # Regex patterns to match different question formats
        self.main_question_pattern = r'(?:Question\s*)?(\d+)\.?\s*'
        self.sub_question_pattern = r'(\d+)\.(\d+)\.?\s*'
        self.letter_pattern = r'([a-z])\)\s*'
        self.roman_pattern = r'([ivxlcdm]+)\)\s*'
    
    def extract_exam_data(self, extractions: List[Any]) -> Dict[str, Any]:
        """
        Main method to extract and structure exam data from LangExtract results.
        
        Args:
            extractions: List of LangExtract Extraction objects
            
        Returns:
            Dict containing structured exam data with metadata and organized answers
        """
        try:
            # Initialize the result structure
            result = {
                "metadata": {
                    "exam_number": None,
                    "total_questions": 0,
                    "extraction_stats": {
                        "total_extractions": len(extractions),
                        "questions_found": 0,
                        "answers_found": 0,
                        "sections_found": 0
                    }
                },
                "structured_answers": {}
            }
            
            # Process extractions and organize them
            organized_data = self._organize_extractions(extractions)
            
            # Update metadata
            result["metadata"]["exam_number"] = organized_data.get("exam_number")
            result["metadata"]["extraction_stats"] = organized_data.get("stats", {})
            
            # Structure the answers hierarchically
            result["structured_answers"] = self._structure_answers(organized_data.get("raw_data", []))
            
            # Calculate total questions
            result["metadata"]["total_questions"] = len(result["structured_answers"])
            
            return result
            
        except Exception as e:
            return {
                "error": f"Failed to extract exam data: {str(e)}",
                "metadata": {
                    "extraction_stats": {
                        "total_extractions": len(extractions) if extractions else 0,
                        "processing_failed": True
                    }
                },
                "raw_extractions": [self._extraction_to_dict(ext) for ext in extractions] if extractions else []
            }
    
    def _organize_extractions(self, extractions: List[Any]) -> Dict[str, Any]:
        """
        Organize extractions by their classification and extract metadata.
        
        Args:
            extractions: List of Extraction objects
            
        Returns:
            Dict with organized extraction data and metadata
        """
        organized = {
            "exam_number": None,
            "raw_data": [],
            "stats": {
                "total_extractions": len(extractions),
                "questions_found": 0,
                "answers_found": 0,
                "sections_found": 0,
                "by_class": defaultdict(int)
            }
        }
        
        for extraction in extractions:
            # Convert extraction to a workable dictionary
            ext_data = self._extraction_to_dict(extraction)
            
            # Track statistics
            organized["stats"]["by_class"][ext_data["extraction_class"]] += 1
            
            # Extract exam number if found
            if ext_data["extraction_class"] == EXAM_NUMBER_CLASS:
                organized["exam_number"] = ext_data["extraction_text"]
            
            # Count different types
            if "question" in ext_data["extraction_class"].lower():
                organized["stats"]["questions_found"] += 1
            elif "answer" in ext_data["extraction_class"].lower():
                organized["stats"]["answers_found"] += 1
            elif "section" in ext_data["extraction_class"].lower():
                organized["stats"]["sections_found"] += 1
            
            organized["raw_data"].append(ext_data)
        
        return organized
    
    def _extraction_to_dict(self, extraction: Any) -> Dict[str, Any]:
        """
        Convert an Extraction object to a dictionary for easier processing.
        
        Args:
            extraction: LangExtract Extraction object
            
        Returns:
            Dict representation of the extraction
        """
        try:
            return {
                "extraction_class": getattr(extraction, 'extraction_class', 'unknown'),
                "extraction_text": getattr(extraction, 'extraction_text', ''),
                "attributes": getattr(extraction, 'attributes', {}),
                "char_interval": getattr(extraction, 'char_interval', None)
            }
        except Exception as e:
            # Fallback for different extraction object structures
            return {
                "extraction_class": str(extraction) if extraction else 'unknown',
                "extraction_text": '',
                "attributes": {},
                "error": f"Failed to parse extraction: {str(e)}"
            }
    
    def _structure_answers(self, raw_data: List[Dict[str, Any]]) -> Dict[str, Any]:
        """
        Structure the answers into the desired hierarchical format.
        
        Args:
            raw_data: List of extraction dictionaries
            
        Returns:
            Dict with structured answers in the format:
            {
                "Question 1": {
                    "1.1": "answer for 1.1",
                    "1.2": {
                        "a": "answer for 1.2a",
                        "b": "answer for 1.2b"
                    }
                }
            }
        """
        structured = {}
        # Group extractions by question context
        question_groups = self._group_by_question(raw_data)
        
        for question_key, question_data in question_groups.items():
            structured[question_key] = self._process_question_group(question_data)
        
        return structured
    
    def _group_by_question(self, raw_data: List[Dict[str, Any]]) -> Dict[str, List[Dict[str, Any]]]:
        """
        Group extractions by question number or context.
        
        Args:
            raw_data: List of extraction dictionaries
            
        Returns:
            Dict grouped by question
        """
        groups = defaultdict(list)
        current_question = "Question 1"  # Default fallback
        
        for item in raw_data:
            current_question = self._determine_question_context(item, current_question)
            groups[current_question].append(item)
        
        return dict(groups)
    
    def _determine_question_context(self, item: Dict[str, Any], current_question: str) -> str:
        """
        Determine which question context an item belongs to.
        
        Args:
            item: Extraction item dictionary
            current_question: Current question context
            
        Returns:
            Updated question context
        """
        text = item.get("extraction_text", "")
        class_name = item.get("extraction_class", "")
        
        # Check if this is a question number or question section
        if "question" in class_name.lower():
            question_num = self._extract_question_number(text)
            if question_num:
                return f"Question {question_num}"
            
            # Use the text as is if it looks like a question identifier
            if text.strip() and (text.strip().startswith("Question") or text.strip().isdigit()):
                if text.strip().isdigit():
                    return f"Question {text.strip()}"
                return text.strip()
        
        return current_question
    
    def _extract_question_number(self, text: str) -> Optional[str]:
        """
        Extract question number from text.
        
        Args:
            text: Text to extract question number from
            
        Returns:
            Question number as string or None
        """
        # Try different patterns
        patterns = [
            r'Question\s*(\d+)',
            r'^\s*(\d+)\s*\.?$',
            r'^\s*(\d+)\s*:',
            r'Q\.?\s*(\d+)'
        ]
        
        for pattern in patterns:
            match = re.search(pattern, text, re.IGNORECASE)
            if match:
                return match.group(1)
        
        return None
    
    def _process_question_group(self, question_data: List[Dict[str, Any]]) -> Any:
        """
        Process a group of extractions for a single question.
        
        Args:
            question_data: List of extractions for one question
            
        Returns:
            Structured answer data for the question
        """
        if not question_data:
            return {}
        
        # Categorize the content
        content_categories = self._categorize_question_content(question_data)
        
        # Structure based on content type
        return self._structure_based_on_content(content_categories, question_data)
    
    def _categorize_question_content(self, question_data: List[Dict[str, Any]]) -> Dict[str, List[Dict[str, Any]]]:
        """
        Categorize question content into different types.
        
        Args:
            question_data: List of extractions for one question
            
        Returns:
            Dict with categorized content
        """
        categories = {
            "answers": [],
            "sections": [],
            "sub_questions": []
        }
        
        for item in question_data:
            class_name = item.get("extraction_class", "").lower()
            
            if "answer" in class_name:
                categories["answers"].append(item)
            elif "section" in class_name:
                categories["sections"].append(item)
            elif class_name != QUESTION_NUMBER_CLASS:
                categories["sub_questions"].append(item)
        
        return categories
    
    def _structure_based_on_content(self, categories: Dict[str, List[Dict[str, Any]]], question_data: List[Dict[str, Any]]) -> Any:
        """
        Structure content based on its categories and patterns.
        
        Args:
            categories: Categorized content
            question_data: Original question data
            
        Returns:
            Structured content
        """
        answers = categories["answers"]
        sub_questions = categories["sub_questions"]
        
        # Check for hierarchical structure first
        if self._has_sub_question_structure(question_data):
            return self._structure_sub_questions(question_data)
        
        # Handle simple answer structures
        if len(answers) == 1:
            return answers[0].get("extraction_text", "")
        elif len(answers) > 1:
            return {str(i): answer.get("extraction_text", "") for i, answer in enumerate(answers, 1)}
        elif sub_questions:
            if len(sub_questions) == 1:
                return sub_questions[0].get("extraction_text", "")
            return {str(i): sq.get("extraction_text", "") for i, sq in enumerate(sub_questions, 1)}
        
        return "No answer content found"
    
    def _has_sub_question_structure(self, data: List[Dict[str, Any]]) -> bool:
        """
        Check if the data contains sub-question numbering (like 1.1, 1.2, etc.).
        
        Args:
            data: List of extraction dictionaries
            
        Returns:
            True if sub-question structure is detected
        """
        for item in data:
            text = item.get("extraction_text", "")
            if re.search(r'\d+\.\d+', text):
                return True
        return False
    
    def _structure_sub_questions(self, data: List[Dict[str, Any]]) -> Dict[str, Any]:
        """
        Structure data that contains sub-questions (1.1, 1.2, etc.).
        
        Args:
            data: List of extraction dictionaries
            
        Returns:
            Hierarchically structured sub-questions
        """
        result = {}
        
        for item in data:
            text = item.get("extraction_text", "")
            
            # Skip question numbers
            if item.get("extraction_class", "").lower() == QUESTION_NUMBER_CLASS:
                continue
            
            # Try to find sub-question patterns
            sub_match = re.search(r'(\d+)\.(\d+)', text)
            if sub_match:
                sub_key = f"{sub_match.group(1)}.{sub_match.group(2)}"
                
                # Check for letter sub-parts (a), b), etc.)
                if self._has_letter_parts(text):
                    result[sub_key] = self._extract_letter_parts(text)
                else:
                    # Clean the text by removing the sub-question number
                    clean_text = re.sub(r'^\d+\.\d+\.?\s*', '', text).strip()
                    result[sub_key] = clean_text if clean_text else text
            else:
                # No clear sub-question pattern, add as is
                # Try to create a reasonable key
                key = self._generate_key(text, len(result) + 1)
                result[key] = text
        
        return result if result else {"1": "No structured content found"}
    
    def _has_letter_parts(self, text: str) -> bool:
        """Check if text contains lettered parts (a), b), c), etc.)."""
        return bool(re.search(r'[a-z]\)', text))
    
    def _extract_letter_parts(self, text: str) -> Dict[str, str]:
        """Extract lettered parts from text."""
        parts = {}
        
        # Split text by letter markers and process each part
        letter_pattern = r'([a-z])\)'
        sections = re.split(letter_pattern, text)
        
        # Process pairs of (letter, content)
        for i in range(1, len(sections), 2):
            if i + 1 < len(sections):
                letter = sections[i]
                content = sections[i + 1].strip()
                if content:
                    parts[letter] = content
        
        return parts if parts else {"a": text}
    
    def _generate_key(self, text: str, index: int) -> str:
        """Generate a reasonable key for unstructured content."""
        # Try to extract a number from the beginning
        match = re.match(r'(\d+)', text.strip())
        if match:
            return match.group(1)
        
        # Use index as fallback
        return str(index)
    
    def format_as_json(self, structured_data: Dict[str, Any], indent: int = 2) -> str:
        """
        Format the structured data as pretty-printed JSON.
        
        Args:
            structured_data: The structured exam data
            indent: Number of spaces for indentation
            
        Returns:
            Pretty-printed JSON string
        """
        try:
            return json.dumps(structured_data, indent=indent, ensure_ascii=False)
        except Exception as e:
            return json.dumps({"error": f"Failed to format JSON: {str(e)}"}, indent=indent)


def process_langextract_results(extractions: List[Any]) -> Dict[str, Any]:
    """
    Helper function to process LangExtract extraction results into structured exam format.
    
    This is the main function you'll use in your text extraction pipeline.
    
    Args:
        extractions: List of LangExtract Extraction objects
        
    Returns:
        Dictionary with structured exam data in the format:
        {
            "metadata": {
                "exam_number": "12345",
                "total_questions": 3,
                "extraction_stats": {...}
            },
            "structured_answers": {
                "Question 1": {
                    "1.1": "answer for 1.1",
                    "1.2": "answer for 1.2"
                },
                "Question 2": {
                    "2.1": {
                        "a": "answer for 2.1 a)",
                        "b": "answer for 2.1 b)"
                    },
                    "2.2": "answer for 2.2"
                }
            }
        }
    """
    extractor = ExamAnswerExtractor()
    return extractor.extract_exam_data(extractions)
