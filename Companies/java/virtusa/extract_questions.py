#!/usr/bin/env python3
import os
import re

# Try to import PDF libraries
try:
    import PyPDF2
    PDF_AVAILABLE = True
except ImportError:
    try:
        import fitz  # PyMuPDF
        PDF_AVAILABLE = True
    except ImportError:
        PDF_AVAILABLE = False

def extract_text_from_pdf_pypdf2(pdf_path):
    """Extract text from PDF using PyPDF2"""
    text = ""
    try:
        with open(pdf_path, 'rb') as file:
            reader = PyPDF2.PdfReader(file)
            for page in reader.pages:
                text += page.extract_text() + "\n"
    except Exception as e:
        print(f"Error reading {pdf_path} with PyPDF2: {e}")
    return text

def extract_text_from_pdf_pymupdf(pdf_path):
    """Extract text from PDF using PyMuPDF"""
    text = ""
    try:
        doc = fitz.open(pdf_path)
        for page in doc:
            text += page.get_text() + "\n"
        doc.close()
    except Exception as e:
        print(f"Error reading {pdf_path} with PyMuPDF: {e}")
    return text

def extract_questions_from_text(text):
    """Extract questions from text using regex patterns"""
    questions = []
    
    # Pattern 1: Lines ending with question mark
    pattern1 = r'.*\?'
    
    # Pattern 2: Lines starting with question words
    pattern2 = r'^(What|Where|When|Why|How|Which|Who|Whom|Whose|Describe|Explain|Define|List|Name|Give|State|Compare|Contrast|Analyze|Evaluate|Discuss|Demonstrate|Show|Calculate|Compute|Find|Determine|Identify|Differentiate|Summarize|Outline|Trace|Illustrate|Interpret|Justify|Prove|Design|Develop|Create|Implement|Write|Code|Program|Build|Construct|Formulate|Generate|Produce|Make|Set up|Configure|Install|Deploy|Test|Debug|Troubleshoot|Optimize|Refactor|Maintain|Update|Modify|Enhance|Improve|Fix|Resolve|Solve|Address|Handle|Manage|Process|Execute|Perform|Carry out|Conduct|Run|Start|Begin|Initialize|Setup|Prepare|Plan|Organize|Structure|Arrange|Layout|Format|Present|Display|Show|Render|Generate|Create|Build|Develop|Design|Implement|Write|Code|Program|Construct|Formulate|Design|Develop|Create|Build|Implement|Write|Code|Program|Construct|Formulate)[^.!?]*[.!?]?$'
    
    # Pattern 3: Numbered questions
    pattern3 = r'^\d+[\.\)]\s.*[.!?]?$'
    
    # Pattern 4: Questions with parentheses
    pattern4 = r'^\([a-zA-Z]\)\s.*[.!?]?$'
    
    lines = text.split('\n')
    
    for line in lines:
        line = line.strip()
        if not line:
            continue
            
        # Check if it's a question
        if (re.search(pattern1, line, re.IGNORECASE) or 
            re.search(pattern2, line, re.IGNORECASE) or
            re.search(pattern3, line) or
            re.search(pattern4, line)):
            
            # Clean up the question
            question = re.sub(r'\s+', ' ', line).strip()
            
            # Skip if it's too short or looks like a header/footer
            if len(question) > 10 and not re.match(r'^Page \d+', question, re.IGNORECASE):
                questions.append(question)
    
    return questions

def main():
    pdf_files = [
        "1752131246981.pdf",
        "1753162718665.pdf", 
        "1756109960440.pdf",
        "1758986655329.pdf"
    ]
    
    all_questions = []
    
    for pdf_file in pdf_files:
        if os.path.exists(pdf_file):
            print(f"Processing {pdf_file}...")
            
            if not PDF_AVAILABLE:
                print(f"Cannot process {pdf_file}: No PDF processing library available")
                continue
                
            # Try PyPDF2 first, then PyMuPDF
            try:
                text = extract_text_from_pdf_pypdf2(pdf_file)
                if not text:
                    text = extract_text_from_pdf_pymupdf(pdf_file)
            except:
                try:
                    text = extract_text_from_pdf_pymupdf(pdf_file)
                except Exception as e:
                    print(f"Failed to extract text from {pdf_file}: {e}")
                    continue
            
            if text:
                questions = extract_questions_from_text(text)
                all_questions.extend(questions)
                print(f"Found {len(questions)} questions in {pdf_file}")
            else:
                print(f"No text extracted from {pdf_file}")
        else:
            print(f"File not found: {pdf_file}")
    
    # Save questions to file
    if all_questions:
        output_file = "virtusa_interview_questions.txt"
        with open(output_file, 'w', encoding='utf-8') as f:
            f.write("Virtusa Java Interview Questions\n")
            f.write("=" * 50 + "\n\n")
            
            for i, question in enumerate(all_questions, 1):
                f.write(f"{i}. {question}\n\n")
        
        print(f"\nTotal questions extracted: {len(all_questions)}")
        print(f"Questions saved to: {output_file}")
    else:
        print("No questions were extracted.")

if __name__ == "__main__":
    main()
