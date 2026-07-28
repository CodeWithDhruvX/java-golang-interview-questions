"""
Loaders package initialization
"""
from .base_loader import BaseDocumentLoader
from .pdf_loader import PDFLoader
from .docx_loader import DOCXLoader
from .csv_loader import CSVLoader
from .txt_loader import TXTLoader
from .document_processor import DocumentProcessor

__all__ = [
    'BaseDocumentLoader',
    'PDFLoader', 
    'DOCXLoader',
    'CSVLoader',
    'TXTLoader',
    'DocumentProcessor'
]
