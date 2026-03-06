import os
import re
import json
import urllib.request
import urllib.error
import time

API_KEY = os.environ.get("GEMINI_API_KEY")
# Using Gemini 1.5 Flash as it is fast and excellent for rewriting
API_URL = f"https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-flash:generateContent?key={API_KEY}"

def rewrite_answer(question, answer):
    prompt = f"""You are an expert software engineer interviewing at a top FAANG company.
Rewrite the following interview answer in the "What Perfect Spoken Format Looks Like" style.

This format generally follows a structured speaking pattern:
Immediate Action (Stop the Bleeding) -> Triangulation/Investigation -> Resolution -> Prevention (Post-Mortem).

Keep it concise, professional, and spoken as if answering in a real interview. Speak in the first person ("First, I would...").
Do not include any pleasantries or conversational filler, just the exact spoken answer inside quotes. Format it with markdown blockquotes.

Question:
{question}

Original Answer Notes:
{answer}
"""
    
    data = {
        "contents": [{"parts": [{"text": prompt}]}],
        "generationConfig": {"temperature": 0.7}
    }
    req = urllib.request.Request(
        API_URL, 
        data=json.dumps(data).encode('utf-8'), 
        headers={'Content-Type': 'application/json'}
    )
    
    # Retry mechanism for rate limits
    for _ in range(3):
        try:
            with urllib.request.urlopen(req) as response:
                result = json.loads(response.read().decode('utf-8'))
                content = result['candidates'][0]['content']['parts'][0]['text']
                # Clean up if the model wrapped it in quotes
                content = content.strip()
                if content.startswith('"') and content.endswith('"'):
                    content = content[1:-1]
                return content
        except urllib.error.HTTPError as e:
            if e.code == 429: # Rate Limit
                time.sleep(2)
            else:
                print(f"Error {e.code}: {e.read().decode('utf-8')}")
                break
        except Exception as e:
            print(f"Failed to call API: {e}")
            break
            
    return answer # Fallback to original if API fails

def process_file(filepath, output_filepath):
    print(f"\nReading {filepath}...")
    with open(filepath, 'r', encoding='utf-8') as f:
        content = f.read()
        
    # Split by markdown headers for questions like "### Question 1: ..."
    # This separates the text into pieces we can process
    sections = re.split(r'(### Question \d+:.*?\n)', content)
    
    if len(sections) < 2:
        print("No questions found or format unrecognized.")
        return

    new_content = sections[0]
    
    # Iterate over the pairs of (Header, Body)
    for i in range(1, len(sections), 2):
        question_header = sections[i]
        question_text = question_header.replace('###', '').strip()
        answer_block = sections[i+1]
        
        # separate answer block from the next "---" separator
        answer_parts = answer_block.split('---', 1)
        original_answer = answer_parts[0]
        remainder = '\n---\n' + answer_parts[1] if len(answer_parts) > 1 else ''
        
        print(f"Processing: {question_text}")
        if API_KEY:
            spoken_answer = rewrite_answer(question_text, original_answer)
            # Add a small delay to avoid hitting rate limits on the free tier
            time.sleep(1) 
        else:
            spoken_answer = "*(GEMINI_API_KEY environment variable not set. Original answer retained)*\n\n" + original_answer
            
        new_content += question_header
        new_content += "\n**What Perfect Spoken Format Looks Like:**\n"
        
        # Ensure it's rendered as a blockquote
        spoken_answer_quoted = "\n".join([f"> {line}" for line in spoken_answer.strip().split("\n")])
        new_content += spoken_answer_quoted + "\n\n"
        new_content += remainder
        
    with open(output_filepath, 'w', encoding='utf-8') as f:
        f.write(new_content)
        
    print(f"-> Saved highly-structured spoken answers to {output_filepath}")

if __name__ == "__main__":
    base_dir = r"c:\Users\dhruv\Downloads\personal_projects\golang-java-interview-questions\All_questions\java-golang-interview-questions\long_questions\Scenariobase_questions_bank"
    
    print("==================================================")
    print("  Spoken Interview Format Converter")
    print("==================================================\n")
    
    if not API_KEY:
        print("⚠️ WARNING: GEMINI_API_KEY environment variable is not set.")
        print("You must set your Gemini API key in the terminal before running this script.")
        print("\nTo set it in PowerShell:")
        print("    $env:GEMINI_API_KEY=\"your_actual_api_key_here\"")
        print("\nTo set it in CMD:")
        print("    set GEMINI_API_KEY=your_actual_api_key_here")
        print("\n(You can get a free API key at: https://aistudio.google.com/app/apikey)\n")
        
        choice = input("Do you want to run it anyway without the API key? (It will just copy the original answers) (y/n): ")
        if choice.lower() != 'y':
            print("Exiting. Please set your API key and run again.")
            exit(0)
    else:
        print("✅ API Key found! Beginning transformation...")

    # File 1
    file1 = os.path.join(base_dir, "Service_Based_Scenario_Answers.md")
    out1 = os.path.join(base_dir, "Service_Based_Spoken_Answers.md")
    process_file(file1, out1)
    
    # File 2
    file2 = os.path.join(base_dir, "FAANG_Scenario_Answers.md")
    out2 = os.path.join(base_dir, "FAANG_Spoken_Answers.md")
    process_file(file2, out2)

    print("\n✅ Success! You're ready to rock your mock interviews.")
