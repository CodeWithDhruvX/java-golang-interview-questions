import sys

file_path = r'g:\My Drive\All Documents\java-golang-interview-questions\long_questions\Python\code_questions.md'

with open(file_path, 'r', encoding='utf-8') as f:
    lines = f.readlines()

levels = []

for i, line in enumerate(lines):
    if line.startswith('You said:'):
        prompt = ""
        if i + 1 < len(lines):
            prompt += lines[i+1].strip()
        levels.append(prompt)

with open(r'g:\My Drive\All Documents\java-golang-interview-questions\long_questions\Python\output_utf8.txt', 'w', encoding='utf-8') as f:
    for prompt in levels:
        f.write(f"Prompt found: {prompt}\n")
    f.write(f"Total prompts: {len(levels)}\n")
