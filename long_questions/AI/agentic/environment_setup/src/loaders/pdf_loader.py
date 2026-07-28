"""
PDF document loader
"""
from .base_loader import BaseDocumentLoader
from typing import Dict, Any
import pypdf


class PDFLoader(BaseDocumentLoader):
    """Load and extract text from PDF files"""
    
    def __init__(self, file_path: str):
        super().__init__(file_path)
        if self.file_path.suffix.lower() != '.pdf':
            raise ValueError("File must be a PDF")
    
    def load(self) -> Dict[str, Any]:
        """Load PDF with full metadata"""
        metadata = self.get_metadata()
        text = self.extract_text()
        
        # Additional PDF-specific metadata
        pdf_metadata = self._extract_pdf_metadata()
        
        return {
            "content": text,
            "metadata": {**metadata, **pdf_metadata},
            "page_count": self._get_page_count()
        }
    
    def extract_text(self) -> str:
        """Extract text from all pages"""
        text_content = []
        
        with open(self.file_path, 'rb') as file:
            pdf_reader = pypdf.PdfReader(file)
            
            for page_num, page in enumerate(pdf_reader.pages):
                try:
                    page_text = page.extract_text()
                    text_content.append(f"--- Page {page_num + 1} ---\n{page_text}")
                except Exception as e:
                    text_content.append(f"--- Page {page_num + 1} ---\n[Error: {str(e)}]")
        
        return "\n\n".join(text_content)
    
    def _extract_pdf_metadata(self) -> Dict[str, str]:
        """Extract PDF-specific metadata"""
        with open(self.file_path, 'rb') as file:
            pdf_reader = pypdf.PdfReader(file)
            info = pdf_reader.metadata
            
            return {
                "title": info.get('/Title', 'Unknown') if info else 'Unknown',
                "author": info.get('/Author', 'Unknown') if info else 'Unknown',
                "creator": info.get('/Creator', 'Unknown') if info else 'Unknown',
            }
    
    def _get_page_count(self) -> int:
        """Get total page count"""
        with open(self.file_path, 'rb') as file:
            pdf_reader = pypdf.PdfReader(file)
            return len(pdf_reader.pages)
