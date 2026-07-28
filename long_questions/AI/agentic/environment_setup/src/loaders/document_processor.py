"""
Unified document processor
"""
from .base_loader import BaseDocumentLoader
from .pdf_loader import PDFLoader
from .docx_loader import DOCXLoader
from .csv_loader import CSVLoader
from .txt_loader import TXTLoader
from pathlib import Path
from typing import Dict, Any


class DocumentProcessor:
    """Unified processor for multiple document formats"""
    
    def __init__(self):
        self.loaders = {
            '.pdf': PDFLoader,
            '.docx': DOCXLoader,
            '.doc': DOCXLoader,
            '.csv': CSVLoader,
            '.txt': TXTLoader
        }
    
    def process_document(self, file_path: str) -> Dict[str, Any]:
        """Process document based on file type"""
        file_path_obj = Path(file_path)
        file_extension = file_path_obj.suffix.lower()
        
        if file_extension not in self.loaders:
            raise ValueError(f"Unsupported file type: {file_extension}")
        
        loader_class = self.loaders[file_extension]
        loader = loader_class(str(file_path_obj))
        
        return loader.load()
    
    def process_directory(self, directory_path: str) -> Dict[str, Dict[str, Any]]:
        """Process all supported documents in a directory"""
        directory = Path(directory_path)
        results = {}
        
        for file_path in directory.iterdir():
            if file_path.is_file() and file_path.suffix.lower() in self.loaders:
                try:
                    results[file_path.name] = self.process_document(str(file_path))
                except Exception as e:
                    results[file_path.name] = {"error": str(e)}
        
        return results
    
    def get_supported_formats(self) -> list:
        """Get list of supported file formats"""
        return list(self.loaders.keys())
