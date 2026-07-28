"""
Main application entry point
"""
from loaders.document_processor import DocumentProcessor
import json
from pathlib import Path


def main():
    """Main application entry point"""
    print("🚀 Multi-Format Document Processor")
    print("=" * 50)
    
    # Initialize processor
    processor = DocumentProcessor()
    
    print(f"Supported formats: {', '.join(processor.get_supported_formats())}")
    print()
    
    # Process single file
    print("Processing single document...")
    sample_txt = "sample_docs/sample.txt"
    if Path(sample_txt).exists():
        result = processor.process_document(sample_txt)
        print(f"✅ Processed: {sample_txt}")
        print(f"   Lines: {result.get('metadata', {}).get('line_count', 'N/A')}")
        print(f"   Content preview: {result['content'][:200]}...")
    else:
        print(f"❌ File not found: {sample_txt}")
    
    print()
    
    # Process directory
    print("Processing directory...")
    input_dir = "sample_docs"
    if Path(input_dir).exists():
        results = processor.process_directory(input_dir)
        print(f"✅ Processed {len(results)} documents")
        
        for filename, result in results.items():
            if 'error' in result:
                print(f"   ❌ {filename}: {result['error']}")
            else:
                print(f"   ✅ {filename}: {result['metadata']['file_type']}")
    else:
        print(f"❌ Directory not found: {input_dir}")
    
    print()
    print("🎉 Processing complete!")


if __name__ == "__main__":
    main()
