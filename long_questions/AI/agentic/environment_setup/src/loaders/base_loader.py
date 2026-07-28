"""
Base document loader interface
"""
from abc import ABC, abstractmethod
from typing import Dict, Any
from pathlib import Path


class BaseDocumentLoader(ABC):
    """Base class for all document loaders"""
    
    def __init__(self, file_path: str):
        self.file_path = Path(file_path)
        self._validate_file()
    
    def _validate_file(self):
        """Validate file exists and is readable"""
        if not self.file_path.exists():
            raise FileNotFoundError(f"File not found: {self.file_path}")
        if not self.file_path.is_file():
            raise ValueError(f"Path is not a file: {self.file_path}")
    
    @abstractmethod
    def load(self) -> Dict[str, Any]:
        """Load document and return structured data"""
        pass
    
    @abstractmethod
    def extract_text(self) -> str:
        """Extract text content from document"""
        pass
    
    def get_metadata(self) -> Dict[str, str]:
        """Extract basic metadata"""
        return {
            "file_name": self.file_path.name,
            "file_size": self.file_path.stat().st_size,
            "file_type": self.file_path.suffix
        }
