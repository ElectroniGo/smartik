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

    # Example text for langextract
    LANGEXTRACT_EXAMPLE_TEXT = textwrap.dedent("""\
        Name: Lebogang Phoshoko
        Student Number: 12345678
        Subject: English Paper 1
        Question 1
        1.1 The digital revolution means the big change in technology and how we use computers
        and phones.
        1.2 Global connection and instant access to information
        1.3 Privacy concerns, cyberbullying, and misinformation
        1.4 Digital natives are young people who grew up with technology and are good at using it.
        1.5 They lack critical evaluation skills when reading things online.
        1.6 The author suggests that balancing technology with ethical considerations is
        important.
        1.7 The passage is about how technology has good and bad effects on society and we need
        to be careful about how we use it.
        Question 2
        2.1 Adjective
        2.2 A new social media platform has been launched by the company.
        2.3 Information
        2.4 Metaphor
        2.5 Mary said that she would attend the digital literacy workshop the next day.
        2.6 The student who completed the project received top marks.
        2.7 Error: "was" should be "were" because when using "neither...nor" with plural subjects,
        the verb should be plural.
        2.8 a) Unheard of b) Improved c) Move through d) Obvious
        Question 3
        3.1 a) Accept means to receive something, except means to leave out something. b) Affect
        is a verb meaning to influence, effect is a noun meaning a result.
        3.2 a) Her perseverance helped her finish the marathon. b) The new policy caused much
        controversy in the school. c) The teacher's explanation helped to illuminate the difficult
        concept.
        3.3 a) To make people feel comfortable in a new situation b) To be exactly right about
        something c) To work late into the night
        Question 4
        4.1 ABAB
        4.2 "The wind whispers secrets" - personification gives the wind human qualities of being
        able to whisper and have secrets. It makes nature seem alive and mysterious.
        4.3 The mood is peaceful and calm. Evidence: "gentle ease," "peaceful place," and
        "silence is found."
        4.4 "Nature's symphony" is a metaphor comparing the sounds of nature to a musical
        symphony. It means all the natural sounds work together harmoniously.
        4.5 The poet creates peace through gentle imagery like "gentle ease" and "peaceful place."
        Also through soft sounds like "whispers" and describing \"silence.\"""")

    # Langextract prompt for extracting structured text
    LANGEXTRACT_PROMPT = textwrap.dedent("""\
            Extract the answers and their question number.
            Use exact text extractions. Do not paraphrase or change any of the text you extract.""")

    # Example data for langextract
    LANGEXTRACT_EXTRACTION_EXAMPLE = [
        lx.data.ExampleData(
            text=LANGEXTRACT_EXAMPLE_TEXT,
            extractions=[
                lx.data.Extraction(
                    extraction_class="question number",
                    extraction_text="1.1",
                    attributes={"type": "numberical"}
                ),
                lx.data.Extraction(
                    extraction_class="answer",
                    extraction_text="The digital revolution means the big change in technology and how we use computers and phones.",
                    attributes={"type": "actual"}
                ),
            ]
        )
    ]
