"""
TXT document loader
"""
from .base_loader import BaseDocumentLoader
from typing import Dict, Any


class TXTLoader(BaseDocumentLoader):
    """Load and extract text from TXT files"""
    
    def __init__(self, file_path: str, encoding: str = 'utf-8'):
        super().__init__(file_path)
        self.encoding = encoding
        if self.file_path.suffix.lower() != '.txt':
            raise ValueError("File must be a TXT file")
    
    def load(self) -> Dict[str, Any]:
        """Load TXT with full metadata"""
        metadata = self.get_metadata()
        text = self.extract_text()
        
        return {
            "content": text,
            "metadata": {**metadata, "line_count": len(text.split('\n')), "char_count": len(text)},
            "encoding": self.encoding
        }
    
    def extract_text(self) -> str:
        """Extract text with encoding handling"""
        try:
            with open(self.file_path, 'r', encoding=self.encoding) as file:
                return file.read()
        except UnicodeDecodeError:
            # Try different encodings
            for encoding in ['latin1', 'iso-8859-1', 'cp1252', 'utf-16']:
                try:
                    with open(self.file_path, 'r', encoding=encoding) as file:
                        return file.read()
                except UnicodeDecodeError:
                    continue
            raise ValueError("Could not decode TXT with common encodings")
