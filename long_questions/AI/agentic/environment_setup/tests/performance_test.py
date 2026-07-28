"""
Performance testing for document processors
"""
import time
from src.loaders.document_processor import DocumentProcessor
from pathlib import Path


def performance_test():
    """Test processing performance"""
    processor = DocumentProcessor()
    
    print("🚀 Performance Test")
    print("=" * 30)
    
    # Test single file
    start = time.time()
    result = processor.process_document("sample_docs/sample.txt")
    elapsed = time.time() - start
    
    print(f"TXT Processing: {elapsed:.4f}s")
    
    # Test directory
    start = time.time()
    results = processor.process_directory("sample_docs")
    elapsed = time.time() - start
    
    print(f"Directory Processing: {elapsed:.4f}s")
    print(f"Documents processed: {len(results)}")


if __name__ == "__main__":
    performance_test()
