# Multi-Format Document Processor

**Stream 1: Environment Setup & Document Loading**

A comprehensive document processing system that handles multiple file formats (PDF, DOCX, CSV, TXT) with unified interface and robust error handling.

## 🚀 Features

- **Multi-format support**: PDF, DOCX, CSV, TXT
- **Unified interface**: Consistent API across all formats
- **Metadata extraction**: Automatic metadata collection
- **Error handling**: Graceful degradation and encoding fallback
- **Batch processing**: Process entire directories
- **Type hints**: Full type safety
- **Extensible design**: Easy to add new formats

## 📋 Prerequisites

- Python 3.9+
- pip (Python package manager)
- Git (optional, for version control)

## 🔧 Installation

### 1. Clone or Setup Repository
```bash
cd long_questions/AI/agentic/environment_setup
```

### 2. Create Virtual Environment
```bash
# Windows
python -m venv venv
venv\Scripts\activate

# Mac/Linux
python -m venv venv
source venv/bin/activate
```

### 3. Install Dependencies
```bash
pip install --upgrade pip
pip install -r requirements.txt
```

## 📁 Project Structure

```
environment_setup/
├── src/
│   ├── loaders/          # Document loading implementations
│   │   ├── base_loader.py
│   │   ├── pdf_loader.py
│   │   ├── docx_loader.py
│   │   ├── csv_loader.py
│   │   ├── txt_loader.py
│   │   └── document_processor.py
│   ├── utils/            # Helper functions
│   └── main.py           # Main application
├── data/
│   ├── input/            # Input documents
│   └── output/           # Processed data
├── sample_docs/          # Sample documents for testing
├── tests/                # Unit tests
├── docs/                 # Documentation
├── requirements.txt
├── .gitignore
└── README.md
```

## 🎯 Usage

### Basic Usage

```python
from src.loaders.document_processor import DocumentProcessor

# Initialize processor
processor = DocumentProcessor()

# Process single document
result = processor.process_document("sample.pdf")
print(result['content'])
print(result['metadata'])

# Process entire directory
results = processor.process_directory("sample_docs/")
for filename, data in results.items():
    print(f"{filename}: {data['metadata']['file_type']}")
```

### Command Line

```bash
# Run main application
python src/main.py

# Create sample documents
python tests/create_sample_docs.py

# Run performance test
python tests/performance_test.py
```

## 📚 Supported Formats

| Format | Extension | Loader Class | Features |
|--------|-----------|--------------|----------|
| PDF | `.pdf` | `PDFLoader` | Page extraction, metadata, page count |
| Word | `.docx`, `.doc` | `DOCXLoader` | Paragraphs, tables, metadata |
| CSV | `.csv` | `CSVLoader` | DataFrame, encoding handling |
| Text | `.txt` | `TXTLoader` | Encoding fallback, statistics |

## 🔍 Examples

### PDF Processing
```python
from src.loaders.pdf_loader import PDFLoader

loader = PDFLoader("document.pdf")
result = loader.load()
print(f"Pages: {result['page_count']}")
print(f"Title: {result['metadata']['title']}")
```

### DOCX Processing
```python
from src.loaders.docx_loader import DOCXLoader

loader = DOCXLoader("document.docx")
result = loader.load()
print(f"Paragraphs: {result['paragraph_count']}")
print(f"Tables: {result['table_count']}")
```

### CSV Processing
```python
from src.loaders.csv_loader import CSVLoader

loader = CSVLoader("data.csv")
result = loader.load()
print(f"Rows: {result['metadata']['row_count']}")
print(f"Columns: {result['metadata']['column_count']}")
```

## 🧪 Testing

```bash
# Run all tests
pytest tests/

# Run specific test
pytest tests/test_loaders.py

# Run with coverage
pytest --cov=src tests/
```

## 🛠️ Development

### Code Style
```bash
# Format code with black
black src/

# Check code style
flake8 src/
```

### Adding New Loaders

1. Create new loader class inheriting from `BaseDocumentLoader`
2. Implement required methods: `load()` and `extract_text()`
3. Add to `DocumentProcessor.loaders` dictionary
4. Add tests in `tests/`

Example:
```python
from src.loaders.base_loader import BaseDocumentLoader

class MarkdownLoader(BaseDocumentLoader):
    def load(self):
        # Implementation
        pass
    
    def extract_text(self):
        # Implementation
        pass
```

## 📖 API Reference

### DocumentProcessor

#### `process_document(file_path: str) -> Dict[str, Any]`
Process a single document and return structured data.

#### `process_directory(directory_path: str) -> Dict[str, Dict[str, Any]]`
Process all supported documents in a directory.

#### `get_supported_formats() -> list`
Return list of supported file extensions.

### BaseDocumentLoader

#### `load() -> Dict[str, Any]`
Load document with full metadata.

#### `extract_text() -> str`
Extract text content from document.

#### `get_metadata() -> Dict[str, str]`
Extract basic file metadata.

## 🐛 Troubleshooting

### Common Issues

**Issue**: `ModuleNotFoundError: No module named 'pypdf'`
**Solution**: Ensure virtual environment is activated and dependencies installed

**Issue**: `UnicodeDecodeError: 'utf-8' codec can't decode`
**Solution**: Loaders automatically try multiple encodings. Check file encoding.

**Issue**: `FileNotFoundError`
**Solution**: Verify file path is correct and file exists.

## 📝 License

This project is part of the AI Interview Questions repository.

## 🤝 Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## 📞 Support

For issues and questions:
- GitHub Issues: [Repository Issues]
- Discord: [Community Server]
- Twitter: [@YourHandle]

## 🎓 Learning Resources

- [Stream Guide](../../stream_guides/stream1_environment_setup_document_loading.md)
- [YouTube Stream](https://youtube.com/yourchannel)
- [Agentic AI Roadmap](../Agentic_AI_Interview_Questions.md)

## 🗺️ Next Steps

- **Stream 2**: Web Scraping Fundamentals
- **Stream 3**: Cloud Storage Integration
- **Stream 4**: Text Splitting & Chunking

---

**Built for AI development journey - from fundamentals to production**
