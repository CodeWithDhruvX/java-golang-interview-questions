"""
CSV document loader
"""
from .base_loader import BaseDocumentLoader
from typing import Dict, Any, List
import pandas as pd


class CSVLoader(BaseDocumentLoader):
    """Load and extract text from CSV files"""
    
    def __init__(self, file_path: str, encoding: str = 'utf-8'):
        super().__init__(file_path)
        self.encoding = encoding
        if self.file_path.suffix.lower() != '.csv':
            raise ValueError("File must be a CSV")
    
    def load(self) -> Dict[str, Any]:
        """Load CSV with full metadata"""
        metadata = self.get_metadata()
        df = self._load_dataframe()
        
        return {
            "content": self.extract_text(),
            "metadata": {**metadata, "row_count": len(df), "column_count": len(df.columns)},
            "dataframe": df,
            "columns": list(df.columns),
            "sample_data": df.head().to_dict()
        }
    
    def extract_text(self) -> str:
        """Extract text as formatted table"""
        df = self._load_dataframe()
        return df.to_string(index=False)
    
    def _load_dataframe(self) -> pd.DataFrame:
        """Load CSV into pandas DataFrame"""
        try:
            return pd.read_csv(self.file_path, encoding=self.encoding)
        except UnicodeDecodeError:
            # Try different encodings
            for encoding in ['latin1', 'iso-8859-1', 'cp1252']:
                try:
                    return pd.read_csv(self.file_path, encoding=encoding)
                except UnicodeDecodeError:
                    continue
            raise ValueError("Could not decode CSV with common encodings")
    
    def get_rows(self) -> List[Dict[str, Any]]:
        """Get data as list of dictionaries"""
        df = self._load_dataframe()
        return df.to_dict('records')
