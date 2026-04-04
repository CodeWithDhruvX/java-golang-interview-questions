#!/usr/bin/env python3
"""
Script to convert Go files to Java files maintaining the same structure and format.
This script reads all Go files from the golang folder and creates corresponding Java files.
"""

import os
import re
from pathlib import Path

def go_to_java_type(go_type):
    """Convert Go type to Java type"""
    type_mapping = {
        'int': 'int',
        'string': 'String',
        'bool': 'boolean',
        '[]int': 'int[]',
        '[]string': 'String[]',
        '[][]int': 'int[][]',
        '[][]string': 'String[][]',
        'float64': 'double',
        '[]float64': 'double[]'
    }
    return type_mapping.get(go_type, go_type)

def convert_go_to_java(go_content, filename):
    """Convert Go code to Java code"""
    lines = go_content.split('\n')
    java_lines = []
    
    # Extract class name from filename
    class_name = filename.replace('.go', '').replace('_', '').replace('-', '')
    
    # Add imports
    imports = set()
    if 'fmt' in go_content:
        imports.add('java.util.Arrays')
    if 'map' in go_content.lower():
        imports.add('java.util.HashMap')
        imports.add('java.util.Map')
    if 'sort' in go_content.lower():
        imports.add('java.util.ArrayList')
        imports.add('java.util.List')
        imports.add('java.util.Collections')
    
    # Add imports
    for imp in sorted(imports):
        java_lines.append(f'import {imp};')
    
    java_lines.append('')
    java_lines.append(f'public class {class_name} {{')
    java_lines.append('')
    
    # Convert functions
    in_function = False
    function_lines = []
    brace_count = 0
    
    for line in lines:
        line = line.strip()
        
        # Skip package declaration
        if line.startswith('package '):
            continue
            
        # Skip import statements (already handled)
        if line.startswith('import '):
            continue
            
        # Convert function declaration
        if line.startswith('func ') and not in_function:
            in_function = True
            function_lines = []
            
            # Extract function name and parameters
            func_match = re.match(r'func (\w+)\(([^)]*)\)([^{]*)', line)
            if func_match:
                func_name = func_match.group(1)
                params = func_match.group(2)
                return_type = func_match.group(3).strip()
                
                # Convert parameters
                java_params = []
                if params.strip():
                    for param in params.split(','):
                        param = param.strip()
                        if ' ' in param:
                            parts = param.split()
                            java_params.append(f'{go_to_java_type(parts[1])} {parts[0]}')
                        else:
                            java_params.append(param)
                
                # Convert return type
                java_return = 'void'
                if return_type:
                    if 'int' in return_type:
                        java_return = 'int'
                    elif 'string' in return_type:
                        java_return = 'String'
                    elif 'bool' in return_type:
                        java_return = 'boolean'
                    elif '[]' in return_type:
                        java_return = go_to_java_type(return_type.strip())
                
                # Add function signature
                java_lines.append(f'    // {line.split("//")[1] if "//" in line else ""}')
                java_lines.append(f'    public static {java_return} {func_name}({", ".join(java_params)}) {{')
                brace_count = 1
            continue
        
        if in_function:
            # Track braces
            brace_count += line.count('{')
            brace_count -= line.count('}')
            
            # Convert Go syntax to Java
            converted_line = line
            
            # Convert range loops
            if 'range ' in converted_line:
                converted_line = re.sub(r'(\w+),\s*(\w+)\s*:=\s*range\s+(\w+)', 
                                       r'for (int \2 = 0; \2 < \3.length; \2++) {\n            int \1 = \3[\2];', converted_line)
                converted_line = re.sub(r'(\w+)\s*:=\s*range\s+(\w+)', 
                                       r'for (int \1 : \2)', converted_line)
            
            # Convert make() to new
            converted_line = re.sub(r'make\(([^)]+)\)', r'new \1', converted_line)
            
            # Convert map operations
            if 'exists :=' in converted_line:
                converted_line = re.sub(r'(\w+),\s*exists\s*:=\s*(\w+)\[(\w+)\]', 
                                       r'boolean exists = \2.containsKey(\3);\n            int \1 = \2.get(\3);', converted_line)
            
            # Convert len() to .length
            converted_line = re.sub(r'len\((\w+)\)', r'\1.length', converted_line)
            
            # Convert append() to add()
            converted_line = re.sub(r'append\(([^,]+),\s*([^)]+)\)', r'\1.add(\2)', converted_line)
            
            # Convert fmt.Printf to System.out.printf
            converted_line = converted_line.replace('fmt.Printf', 'System.out.printf')
            converted_line = converted_line.replace('fmt.Println', 'System.out.println')
            
            # Convert %v to %s for arrays
            converted_line = converted_line.replace('%v', '%s')
            
            # Convert true/false to Java boolean
            converted_line = converted_line.replace('true', 'true')
            converted_line = converted_line.replace('false', 'false')
            
            # Convert nil to null
            converted_line = converted_line.replace('nil', 'null')
            
            function_lines.append('    ' + converted_line)
            
            if brace_count == 0:
                java_lines.extend(function_lines)
                java_lines.append('    }')
                java_lines.append('')
                in_function = False
    
    java_lines.append('}')
    
    return '\n'.join(java_lines)

def main():
    """Main function to convert Go files to Java"""
    golang_root = Path("golang")
    java_root = Path("java")
    
    if not golang_root.exists():
        print("Golang folder not found!")
        return
    
    # Create Java root directory
    java_root.mkdir(exist_ok=True)
    
    # Process each directory
    for go_dir in golang_root.iterdir():
        if go_dir.is_dir():
            java_dir = java_root / go_dir.name
            java_dir.mkdir(exist_ok=True)
            
            # Process each Go file
            for go_file in go_dir.glob("*.go"):
                java_file = java_dir / (go_file.stem + ".java")
                
                # Read Go file
                with open(go_file, 'r', encoding='utf-8') as f:
                    go_content = f.read()
                
                # Convert to Java
                java_content = convert_go_to_java(go_content, go_file.name)
                
                # Write Java file
                with open(java_file, 'w', encoding='utf-8') as f:
                    f.write(java_content)
                
                print(f"Created: {java_file}")

if __name__ == "__main__":
    main()
