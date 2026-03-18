# File Operations: Practical Programs

**Goal**: Master comprehensive file handling including CSV/Excel processing, properties files, and NIO.2 operations.

## Prerequisites

Add these dependencies to your `pom.xml` for Excel processing:

```xml
<dependencies>
    <dependency>
        <groupId>org.apache.poi</groupId>
        <artifactId>poi</artifactId>
        <version>5.2.4</version>
    </dependency>
    <dependency>
        <groupId>org.apache.poi</groupId>
        <artifactId>poi-ooxml</artifactId>
        <version>5.2.4</version>
    </dependency>
    <dependency>
        <groupId>commons-io</groupId>
        <artifactId>commons-io</artifactId>
        <version>2.11.0</version>
    </dependency>
    <dependency>
        <groupId>org.apache.commons</groupId>
        <artifactId>commons-csv</artifactId>
        <version>1.10.0</version>
    </dependency>
</dependencies>
```

## 1. CSV File Processing

### Comprehensive CSV Operations

```java
import org.apache.commons.csv.*;
import java.io.*;
import java.nio.file.*;
import java.util.*;
import java.util.stream.*;

// Data model for CSV records
class Employee {
    private int id;
    private String name;
    private String department;
    private double salary;
    private String email;
    private LocalDate hireDate;
    
    public Employee(int id, String name, String department, double salary, String email, LocalDate hireDate) {
        this.id = id;
        this.name = name;
        this.department = department;
        this.salary = salary;
        this.email = email;
        this.hireDate = hireDate;
    }
    
    // Getters and setters
    public int getId() { return id; }
    public void setId(int id) { this.id = id; }
    
    public String getName() { return name; }
    public void setName(String name) { this.name = name; }
    
    public String getDepartment() { return department; }
    public void setDepartment(String department) { this.department = department; }
    
    public double getSalary() { return salary; }
    public void setSalary(double salary) { this.salary = salary; }
    
    public String getEmail() { return email; }
    public void setEmail(String email) { this.email = email; }
    
    public LocalDate getHireDate() { return hireDate; }
    public void setHireDate(LocalDate hireDate) { this.hireDate = hireDate; }
    
    @Override
    public String toString() {
        return String.format("Employee{id=%d, name='%s', department='%s', salary=%.2f, email='%s', hireDate=%s}",
                           id, name, department, salary, email, hireDate);
    }
}

// CSV processor class
class CsvProcessor {
    
    // Write employees to CSV
    public void writeEmployeesToCSV(List<Employee> employees, String filePath) throws IOException {
        try (BufferedWriter writer = Files.newBufferedWriter(Paths.get(filePath));
             CSVPrinter csvPrinter = new CSVPrinter(writer, CSVFormat.DEFAULT
                 .withHeader("ID", "Name", "Department", "Salary", "Email", "HireDate"))) {
            
            for (Employee employee : employees) {
                csvPrinter.printRecord(
                    employee.getId(),
                    employee.getName(),
                    employee.getDepartment(),
                    employee.getSalary(),
                    employee.getEmail(),
                    employee.getHireDate()
                );
            }
            
            csvPrinter.flush();
            System.out.println("CSV file written successfully: " + filePath);
        }
    }
    
    // Read employees from CSV
    public List<Employee> readEmployeesFromCSV(String filePath) throws IOException {
        List<Employee> employees = new ArrayList<>();
        
        try (BufferedReader reader = Files.newBufferedReader(Paths.get(filePath));
             CSVParser csvParser = new CSVParser(reader, CSVFormat.DEFAULT
                 .withFirstRecordAsHeader()
                 .withIgnoreHeaderCase()
                 .withTrim())) {
            
            for (CSVRecord record : csvParser) {
                Employee employee = new Employee(
                    Integer.parseInt(record.get("ID")),
                    record.get("Name"),
                    record.get("Department"),
                    Double.parseDouble(record.get("Salary")),
                    record.get("Email"),
                    LocalDate.parse(record.get("HireDate"))
                );
                employees.add(employee);
            }
        }
        
        System.out.println("Read " + employees.size() + " employees from CSV");
        return employees;
    }
    
    // Filter and write specific departments
    public void filterByDepartment(List<Employee> employees, String department, String outputPath) throws IOException {
        List<Employee> filtered = employees.stream()
            .filter(emp -> emp.getDepartment().equalsIgnoreCase(department))
            .collect(Collectors.toList());
        
        writeEmployeesToCSV(filtered, outputPath);
        System.out.println("Filtered " + filtered.size() + " employees from department: " + department);
    }
    
    // Calculate statistics and write to CSV
    public void generateDepartmentStatistics(List<Employee> employees, String outputPath) throws IOException {
        Map<String, DoubleSummaryStatistics> stats = employees.stream()
            .collect(Collectors.groupingBy(
                Employee::getDepartment,
                Collectors.summarizingDouble(Employee::getSalary)
            ));
        
        try (BufferedWriter writer = Files.newBufferedWriter(Paths.get(outputPath));
             CSVPrinter csvPrinter = new CSVPrinter(writer, CSVFormat.DEFAULT
                 .withHeader("Department", "Count", "Average Salary", "Min Salary", "Max Salary", "Total Salary"))) {
            
            for (Map.Entry<String, DoubleSummaryStatistics> entry : stats.entrySet()) {
                String department = entry.getKey();
                DoubleSummaryStatistics departmentStats = entry.getValue();
                
                csvPrinter.printRecord(
                    department,
                    departmentStats.getCount(),
                    String.format("%.2f", departmentStats.getAverage()),
                    String.format("%.2f", departmentStats.getMin()),
                    String.format("%.2f", departmentStats.getMax()),
                    String.format("%.2f", departmentStats.getSum())
                );
            }
            
            csvPrinter.flush();
            System.out.println("Department statistics written to: " + outputPath);
        }
    }
    
    // Advanced CSV processing with custom format
    public void processCustomCSV(String inputPath, String outputPath) throws IOException {
        CSVFormat customFormat = CSVFormat.DEFAULT
            .withQuote('"')
            .withEscape('\\')
            .withIgnoreEmptyLines(true)
            .withIgnoreSurroundingSpaces(true);
        
        List<String[]> processedRows = new ArrayList<>();
        
        try (BufferedReader reader = Files.newBufferedReader(Paths.get(inputPath));
             CSVParser parser = new CSVParser(reader, customFormat)) {
            
            for (CSVRecord record : parser) {
                // Process each record
                String[] processedRow = new String[record.size()];
                for (int i = 0; i < record.size(); i++) {
                    String value = record.get(i);
                    // Example processing: uppercase names and format salary
                    if (i == 1) { // Name column
                        processedRow[i] = value.toUpperCase();
                    } else if (i == 3) { // Salary column
                        try {
                            double salary = Double.parseDouble(value);
                            processedRow[i] = String.format("$%.2f", salary);
                        } catch (NumberFormatException e) {
                            processedRow[i] = value;
                        }
                    } else {
                        processedRow[i] = value;
                    }
                }
                processedRows.add(processedRow);
            }
        }
        
        // Write processed data
        try (BufferedWriter writer = Files.newBufferedWriter(Paths.get(outputPath));
             CSVPrinter csvPrinter = new CSVPrinter(writer, customFormat)) {
            
            for (String[] row : processedRows) {
                csvPrinter.printRecord((Object[]) row);
            }
            
            csvPrinter.flush();
            System.out.println("Custom CSV processing completed: " + outputPath);
        }
    }
}

public class CsvProcessingDemo {
    public static void main(String[] args) {
        CsvProcessor processor = new CsvProcessor();
        
        // Create sample data
        List<Employee> employees = Arrays.asList(
            new Employee(1, "John Doe", "Engineering", 75000, "john@company.com", LocalDate.of(2020, 1, 15)),
            new Employee(2, "Jane Smith", "Marketing", 65000, "jane@company.com", LocalDate.of(2019, 3, 22)),
            new Employee(3, "Bob Johnson", "Engineering", 80000, "bob@company.com", LocalDate.of(2021, 6, 10)),
            new Employee(4, "Alice Brown", "HR", 55000, "alice@company.com", LocalDate.of(2018, 11, 5)),
            new Employee(5, "Charlie Wilson", "Engineering", 85000, "charlie@company.com", LocalDate.of(2020, 8, 18)),
            new Employee(6, "Diana Davis", "Marketing", 62000, "diana@company.com", LocalDate.of(2021, 2, 28)),
            new Employee(7, "Eve Miller", "Finance", 70000, "eve@company.com", LocalDate.of(2019, 7, 12)),
            new Employee(8, "Frank Garcia", "Engineering", 90000, "frank@company.com", LocalDate.of(2017, 4, 3))
        );
        
        try {
            // 1. Write to CSV
            System.out.println("=== Writing Employees to CSV ===");
            processor.writeEmployeesToCSV(employees, "employees.csv");
            
            // 2. Read from CSV
            System.out.println("\n=== Reading Employees from CSV ===");
            List<Employee> readEmployees = processor.readEmployeesFromCSV("employees.csv");
            readEmployees.forEach(System.out::println);
            
            // 3. Filter by department
            System.out.println("\n=== Filtering by Department ===");
            processor.filterByDepartment(employees, "Engineering", "engineering_employees.csv");
            
            // 4. Generate statistics
            System.out.println("\n=== Generating Department Statistics ===");
            processor.generateDepartmentStatistics(employees, "department_stats.csv");
            
            // 5. Custom processing
            System.out.println("\n=== Custom CSV Processing ===");
            processor.processCustomCSV("employees.csv", "processed_employees.csv");
            
            // 6. Display statistics
            System.out.println("\n=== Department Statistics ===");
            List<String> statsLines = Files.readAllLines(Paths.get("department_stats.csv"));
            statsLines.forEach(System.out::println);
            
        } catch (IOException e) {
            System.err.println("Error processing CSV files: " + e.getMessage());
        } finally {
            // Cleanup files
            cleanupCSVFiles();
        }
    }
    
    private static void cleanupCSVFiles() {
        try {
            Files.deleteIfExists(Paths.get("employees.csv"));
            Files.deleteIfExists(Paths.get("engineering_employees.csv"));
            Files.deleteIfExists(Paths.get("department_stats.csv"));
            Files.deleteIfExists(Paths.get("processed_employees.csv"));
        } catch (IOException e) {
            System.err.println("Error cleaning up files: " + e.getMessage());
        }
    }
}
```

## 2. Excel File Processing

### Excel Operations with Apache POI

```java
import org.apache.poi.ss.*;
import org.apache.poi.xssf.usermodel.*;
import org.apache.poi.hssf.usermodel.*;
import org.apache.poi.ss.usermodel.*;
import org.apache.poi.ss.util.*;
import java.io.*;
import java.time.LocalDate;
import java.util.*;

// Excel processor class
class ExcelProcessor {
    
    // Create a new Excel workbook with sample data
    public void createEmployeeWorkbook(String filePath) throws IOException {
        try (Workbook workbook = new XSSFWorkbook()) {
            // Create sheet
            Sheet sheet = workbook.createSheet("Employees");
            
            // Create header row
            Row headerRow = sheet.createRow(0);
            String[] headers = {"ID", "Name", "Department", "Salary", "Email", "Hire Date", "Performance"};
            
            CellStyle headerStyle = createHeaderStyle(workbook);
            
            for (int i = 0; i < headers.length; i++) {
                Cell cell = headerRow.createCell(i);
                cell.setCellValue(headers[i]);
                cell.setCellStyle(headerStyle);
            }
            
            // Create data rows
            List<Object[]> employeeData = getSampleEmployeeData();
            CellStyle dateStyle = createDateStyle(workbook);
            CellStyle currencyStyle = createCurrencyStyle(workbook);
            
            for (int i = 0; i < employeeData.size(); i++) {
                Row row = sheet.createRow(i + 1);
                Object[] data = employeeData.get(i);
                
                for (int j = 0; j < data.length; j++) {
                    Cell cell = row.createCell(j);
                    
                    if (data[j] instanceof Integer) {
                        cell.setCellValue((Integer) data[j]);
                    } else if (data[j] instanceof String) {
                        cell.setCellValue((String) data[j]);
                    } else if (data[j] instanceof Double) {
                        cell.setCellValue((Double) data[j]);
                        cell.setCellStyle(currencyStyle);
                    } else if (data[j] instanceof LocalDate) {
                        cell.setCellValue((LocalDate) data[j]);
                        cell.setCellStyle(dateStyle);
                    }
                }
            }
            
            // Auto-size columns
            for (int i = 0; i < headers.length; i++) {
                sheet.autoSizeColumn(i);
            }
            
            // Create summary sheet
            createSummarySheet(workbook, employeeData);
            
            // Write to file
            try (FileOutputStream outputStream = new FileOutputStream(filePath)) {
                workbook.write(outputStream);
                System.out.println("Excel workbook created: " + filePath);
            }
        }
    }
    
    // Read data from Excel workbook
    public List<Employee> readEmployeesFromExcel(String filePath) throws IOException {
        List<Employee> employees = new ArrayList<>();
        
        try (Workbook workbook = WorkbookFactory.create(new File(filePath))) {
            Sheet sheet = workbook.getSheet("Employees");
            
            for (int i = 1; i <= sheet.getLastRowNum(); i++) { // Skip header row
                Row row = sheet.getRow(i);
                if (row != null) {
                    Employee employee = new Employee(
                        (int) row.getCell(0).getNumericCellValue(),
                        row.getCell(1).getStringCellValue(),
                        row.getCell(2).getStringCellValue(),
                        row.getCell(3).getNumericCellValue(),
                        row.getCell(4).getStringCellValue(),
                        row.getCell(5).getLocalDateTimeCellValue().toLocalDate()
                    );
                    employees.add(employee);
                }
            }
        }
        
        System.out.println("Read " + employees.size() + " employees from Excel");
        return employees;
    }
    
    // Add formulas and charts
    public void enhanceExcelWorkbook(String inputPath, String outputPath) throws IOException {
        try (Workbook workbook = WorkbookFactory.create(new File(inputPath))) {
            Sheet sheet = workbook.getSheet("Employees");
            
            // Add formula columns
            addFormulaColumns(sheet);
            
            // Create conditional formatting
            addConditionalFormatting(sheet, workbook);
            
            // Create chart sheet
            createChartSheet(workbook);
            
            // Write enhanced workbook
            try (FileOutputStream outputStream = new FileOutputStream(outputPath)) {
                workbook.write(outputStream);
                System.out.println("Enhanced Excel workbook created: " + outputPath);
            }
        }
    }
    
    // Process multiple sheets
    public void processMultipleSheets(String filePath) throws IOException {
        try (Workbook workbook = WorkbookFactory.create(new File(filePath))) {
            
            for (int i = 0; i < workbook.getNumberOfSheets(); i++) {
                Sheet sheet = workbook.getSheetAt(i);
                System.out.println("Processing sheet: " + sheet.getSheetName());
                
                // Process each sheet
                processSheet(sheet);
            }
        }
    }
    
    private void createHeaderStyle(Workbook workbook) {
        CellStyle style = workbook.createCellStyle();
        style.setFillForegroundColor(IndexedColors.GREY_25_PERCENT.getIndex());
        style.setFillPattern(FillPatternType.SOLID_FOREGROUND);
        style.setBorderTop(BorderStyle.THIN);
        style.setBorderBottom(BorderStyle.THIN);
        style.setBorderLeft(BorderStyle.THIN);
        style.setBorderRight(BorderStyle.THIN);
        
        Font font = workbook.createFont();
        font.setBold(true);
        font.setFontHeightInPoints((short) 12);
        style.setFont(font);
        
        return style;
    }
    
    private CellStyle createDateStyle(Workbook workbook) {
        CellStyle style = workbook.createCellStyle();
        CreationHelper createHelper = workbook.getCreationHelper();
        style.setDataFormat(createHelper.createDataFormat().getFormat("yyyy-mm-dd"));
        return style;
    }
    
    private CellStyle createCurrencyStyle(Workbook workbook) {
        CellStyle style = workbook.createCellStyle();
        style.setDataFormat((short) 8); // Currency format
        return style;
    }
    
    private List<Object[]> getSampleEmployeeData() {
        return Arrays.asList(
            new Object[]{1, "John Doe", "Engineering", 75000.0, "john@company.com", LocalDate.of(2020, 1, 15), "Excellent"},
            new Object[]{2, "Jane Smith", "Marketing", 65000.0, "jane@company.com", LocalDate.of(2019, 3, 22), "Good"},
            new Object[]{3, "Bob Johnson", "Engineering", 80000.0, "bob@company.com", LocalDate.of(2021, 6, 10), "Excellent"},
            new Object[]{4, "Alice Brown", "HR", 55000.0, "alice@company.com", LocalDate.of(2018, 11, 5), "Good"},
            new Object[]{5, "Charlie Wilson", "Engineering", 85000.0, "charlie@company.com", LocalDate.of(2020, 8, 18), "Outstanding"},
            new Object[]{6, "Diana Davis", "Marketing", 62000.0, "diana@company.com", LocalDate.of(2021, 2, 28), "Average"},
            new Object[]{7, "Eve Miller", "Finance", 70000.0, "eve@company.com", LocalDate.of(2019, 7, 12), "Good"},
            new Object[]{8, "Frank Garcia", "Engineering", 90000.0, "frank@company.com", LocalDate.of(2017, 4, 3), "Outstanding"}
        );
    }
    
    private void createSummarySheet(Workbook workbook, List<Object[]> employeeData) {
        Sheet summarySheet = workbook.createSheet("Summary");
        
        // Create summary statistics
        Row headerRow = summarySheet.createRow(0);
        String[] headers = {"Department", "Count", "Average Salary", "Min Salary", "Max Salary"};
        
        for (int i = 0; i < headers.length; i++) {
            headerRow.createCell(i).setCellValue(headers[i]);
        }
        
        // Group by department and calculate statistics
        Map<String, List<Double>> departmentSalaries = new HashMap<>();
        
        for (Object[] data : employeeData) {
            String department = (String) data[2];
            double salary = (Double) data[3];
            
            departmentSalaries.computeIfAbsent(department, k -> new ArrayList<>()).add(salary);
        }
        
        int rowNum = 1;
        for (Map.Entry<String, List<Double>> entry : departmentSalaries.entrySet()) {
            Row row = summarySheet.createRow(rowNum++);
            List<Double> salaries = entry.getValue();
            
            row.createCell(0).setCellValue(entry.getKey());
            row.createCell(1).setCellValue(salaries.size());
            row.createCell(2).setCellValue(salaries.stream().mapToDouble(Double::doubleValue).average().orElse(0));
            row.createCell(3).setCellValue(salaries.stream().mapToDouble(Double::doubleValue).min().orElse(0));
            row.createCell(4).setCellValue(salaries.stream().mapToDouble(Double::doubleValue).max().orElse(0));
        }
        
        // Auto-size columns
        for (int i = 0; i < headers.length; i++) {
            summarySheet.autoSizeColumn(i);
        }
    }
    
    private void addFormulaColumns(Sheet sheet) {
        // Add bonus column (10% of salary)
        Cell bonusCell = sheet.getRow(1).createCell(7);
        bonusCell.setCellFormula("D2*0.1");
        bonusCell.setCellValue("Bonus");
        
        // Copy formula to other rows
        for (int i = 2; i <= sheet.getLastRowNum(); i++) {
            Row row = sheet.getRow(i);
            if (row != null) {
                Cell cell = row.createCell(7);
                cell.setCellFormula("D" + (i + 1) + "*0.1");
            }
        }
        
        // Add total compensation column
        Cell totalCell = sheet.getRow(1).createCell(8);
        totalCell.setCellFormula("D2+H2");
        totalCell.setCellValue("Total Compensation");
        
        for (int i = 2; i <= sheet.getLastRowNum(); i++) {
            Row row = sheet.getRow(i);
            if (row != null) {
                Cell cell = row.createCell(8);
                cell.setCellFormula("D" + (i + 1) + "+H" + (i + 1));
            }
        }
    }
    
    private void addConditionalFormatting(Sheet sheet, Workbook workbook) {
        SheetConditionalFormatting sheetCF = sheet.getSheetConditionalFormatting();
        
        // Create conditional formatting rule for high salaries
        ConditionalFormattingRule rule1 = sheetCF.createConditionalFormattingRule(
            ComparisonOperator.GT, "80000"
        );
        PatternFormatting fill1 = rule1.createPatternFormatting();
        fill1.setFillBackgroundColor(IndexedColors.LIGHT_GREEN.getIndex());
        fill1.setFillPattern(PatternFormatting.SOLID_FOREGROUND);
        
        // Create conditional formatting rule for low salaries
        ConditionalFormattingRule rule2 = sheetCF.createConditionalFormattingRule(
            ComparisonOperator.LT, "60000"
        );
        PatternFormatting fill2 = rule2.createPatternFormatting();
        fill2.setFillBackgroundColor(IndexedColors.LIGHT_YELLOW.getIndex());
        fill2.setFillPattern(PatternFormatting.SOLID_FOREGROUND);
        
        // Apply conditional formatting to salary column (column D)
        CellRangeAddress[] ranges = {
            CellRangeAddress.valueOf("D2:D" + (sheet.getLastRowNum() + 1))
        };
        
        sheetCF.addConditionalFormatting(ranges, rule1, rule2);
    }
    
    private void createChartSheet(Workbook workbook) {
        // Note: Creating charts in Apache POI is complex and requires additional setup
        // This is a simplified example
        Sheet chartSheet = workbook.createSheet("Chart Info");
        
        Row row = chartSheet.createRow(0);
        row.createCell(0).setCellValue("Chart Information");
        
        row = chartSheet.createRow(1);
        row.createCell(0).setCellValue("Department Salary Distribution");
        
        row = chartSheet.createRow(2);
        row.createCell(0).setCellValue("Employee Performance Analysis");
        
        row = chartSheet.createRow(3);
        row.createCell(0).setCellValue("Salary Trends Over Time");
    }
    
    private void processSheet(Sheet sheet) {
        int rowCount = sheet.getLastRowNum();
        int columnCount = 0;
        
        if (sheet.getRow(0) != null) {
            columnCount = sheet.getRow(0).getLastCellNum();
        }
        
        System.out.println("Rows: " + rowCount + ", Columns: " + columnCount);
        
        // Process data
        for (int i = 0; i <= rowCount; i++) {
            Row row = sheet.getRow(i);
            if (row != null) {
                for (int j = 0; j < columnCount; j++) {
                    Cell cell = row.getCell(j);
                    if (cell != null) {
                        // Process cell based on type
                        switch (cell.getCellType()) {
                            case STRING:
                                // Process string
                                break;
                            case NUMERIC:
                                // Process number
                                break;
                            case FORMULA:
                                // Process formula
                                break;
                            default:
                                break;
                        }
                    }
                }
            }
        }
    }
}

public class ExcelProcessingDemo {
    public static void main(String[] args) {
        ExcelProcessor processor = new ExcelProcessor();
        
        try {
            // 1. Create Excel workbook
            System.out.println("=== Creating Excel Workbook ===");
            processor.createEmployeeWorkbook("employees.xlsx");
            
            // 2. Read from Excel
            System.out.println("\n=== Reading from Excel ===");
            List<Employee> employees = processor.readEmployeesFromExcel("employees.xlsx");
            employees.forEach(System.out::println);
            
            // 3. Enhance Excel with formulas and formatting
            System.out.println("\n=== Enhancing Excel Workbook ===");
            processor.enhanceExcelWorkbook("employees.xlsx", "employees_enhanced.xlsx");
            
            // 4. Process multiple sheets
            System.out.println("\n=== Processing Multiple Sheets ===");
            processor.processMultipleSheets("employees_enhanced.xlsx");
            
        } catch (IOException e) {
            System.err.println("Error processing Excel files: " + e.getMessage());
        } finally {
            // Cleanup files
            cleanupExcelFiles();
        }
    }
    
    private static void cleanupExcelFiles() {
        try {
            Files.deleteIfExists(Paths.get("employees.xlsx"));
            Files.deleteIfExists(Paths.get("employees_enhanced.xlsx"));
        } catch (IOException e) {
            System.err.println("Error cleaning up files: " + e.getMessage());
        }
    }
}
```

## 3. Properties File Handling

### Configuration Management with Properties

```java
import java.io.*;
import java.nio.file.*;
import java.util.*;
import java.util.stream.*;

// Configuration manager class
class ConfigurationManager {
    private Properties properties;
    private String configFilePath;
    
    public ConfigurationManager(String configFilePath) {
        this.configFilePath = configFilePath;
        this.properties = new Properties();
        loadProperties();
    }
    
    // Load properties from file
    private void loadProperties() {
        try (InputStream input = Files.newInputStream(Paths.get(configFilePath))) {
            properties.load(input);
            System.out.println("Properties loaded from: " + configFilePath);
        } catch (IOException e) {
            System.err.println("Error loading properties: " + e.getMessage());
            // Create default properties if file doesn't exist
            createDefaultProperties();
        }
    }
    
    // Create default properties
    private void createDefaultProperties() {
        properties.setProperty("app.name", "MyApplication");
        properties.setProperty("app.version", "1.0.0");
        properties.setProperty("database.url", "jdbc:mysql://localhost:3306/mydb");
        properties.setProperty("database.username", "root");
        properties.setProperty("database.password", "password");
        properties.setProperty("database.pool.size", "10");
        properties.setProperty("server.port", "8080");
        properties.setProperty("server.host", "localhost");
        properties.setProperty("logging.level", "INFO");
        properties.setProperty("logging.file.path", "logs/app.log");
        properties.setProperty("cache.size", "1000");
        properties.setProperty("cache.ttl", "3600");
        properties.setProperty("feature.new_ui", "true");
        properties.setProperty("feature.analytics", "false");
        
        saveProperties();
        System.out.println("Default properties created");
    }
    
    // Save properties to file
    public void saveProperties() {
        try (OutputStream output = Files.newOutputStream(Paths.get(configFilePath))) {
            properties.store(output, "Application Configuration");
            System.out.println("Properties saved to: " + configFilePath);
        } catch (IOException e) {
            System.err.println("Error saving properties: " + e.getMessage());
        }
    }
    
    // Get property with default value
    public String getProperty(String key, String defaultValue) {
        return properties.getProperty(key, defaultValue);
    }
    
    // Get property as integer
    public int getIntProperty(String key, int defaultValue) {
        try {
            return Integer.parseInt(properties.getProperty(key, String.valueOf(defaultValue)));
        } catch (NumberFormatException e) {
            return defaultValue;
        }
    }
    
    // Get property as boolean
    public boolean getBooleanProperty(String key, boolean defaultValue) {
        String value = properties.getProperty(key, String.valueOf(defaultValue));
        return Boolean.parseBoolean(value);
    }
    
    // Get property as double
    public double getDoubleProperty(String key, double defaultValue) {
        try {
            return Double.parseDouble(properties.getProperty(key, String.valueOf(defaultValue)));
        } catch (NumberFormatException e) {
            return defaultValue;
        }
    }
    
    // Set property
    public void setProperty(String key, String value) {
        properties.setProperty(key, value);
    }
    
    // Set property with type conversion
    public void setProperty(String key, int value) {
        properties.setProperty(key, String.valueOf(value));
    }
    
    public void setProperty(String key, boolean value) {
        properties.setProperty(key, String.valueOf(value));
    }
    
    public void setProperty(String key, double value) {
        properties.setProperty(key, String.valueOf(value));
    }
    
    // Get all properties
    public Properties getAllProperties() {
        return new Properties(properties);
    }
    
    // Get properties by prefix
    public Map<String, String> getPropertiesByPrefix(String prefix) {
        return properties.entrySet().stream()
            .filter(entry -> ((String) entry.getKey()).startsWith(prefix))
            .collect(Collectors.toMap(
                entry -> (String) entry.getKey(),
                entry -> (String) entry.getValue()
            ));
    }
    
    // Update properties from map
    public void updateProperties(Map<String, String> updates) {
        updates.forEach(properties::setProperty);
        saveProperties();
    }
    
    // Remove property
    public void removeProperty(String key) {
        properties.remove(key);
    }
    
    // Clear all properties
    public void clearProperties() {
        properties.clear();
    }
    
    // Display all properties
    public void displayProperties() {
        System.out.println("\n=== Configuration Properties ===");
        properties.forEach((key, value) -> 
            System.out.println(key + " = " + value));
    }
    
    // Validate required properties
    public boolean validateRequiredProperties(List<String> requiredKeys) {
        List<String> missingKeys = new ArrayList<>();
        
        for (String key : requiredKeys) {
            if (!properties.containsKey(key) || properties.getProperty(key).trim().isEmpty()) {
                missingKeys.add(key);
            }
        }
        
        if (!missingKeys.isEmpty()) {
            System.err.println("Missing required properties: " + missingKeys);
            return false;
        }
        
        return true;
    }
    
    // Export properties to XML
    public void exportToXML(String xmlFilePath) {
        try (OutputStream output = Files.newOutputStream(Paths.get(xmlFilePath))) {
            properties.storeToXML(output, "Application Configuration");
            System.out.println("Properties exported to XML: " + xmlFilePath);
        } catch (IOException e) {
            System.err.println("Error exporting to XML: " + e.getMessage());
        }
    }
    
    // Import properties from XML
    public void importFromXML(String xmlFilePath) {
        try (InputStream input = Files.newInputStream(Paths.get(xmlFilePath))) {
            properties.loadFromXML(input);
            System.out.println("Properties imported from XML: " + xmlFilePath);
            saveProperties();
        } catch (IOException e) {
            System.err.println("Error importing from XML: " + e.getMessage());
        }
    }
}

// Advanced configuration with environment-specific settings
class AdvancedConfigurationManager {
    private Properties baseProperties;
    private Properties environmentProperties;
    private String environment;
    
    public AdvancedConfigurationManager(String environment) {
        this.environment = environment;
        this.baseProperties = new Properties();
        this.environmentProperties = new Properties();
        loadConfigurations();
    }
    
    private void loadConfigurations() {
        try {
            // Load base configuration
            try (InputStream input = Files.newInputStream(Paths.get("config/base.properties"))) {
                baseProperties.load(input);
            }
            
            // Load environment-specific configuration
            String envConfigPath = "config/" + environment + ".properties";
            try (InputStream input = Files.newInputStream(Paths.get(envConfigPath))) {
                environmentProperties.load(input);
            }
            
            System.out.println("Loaded configuration for environment: " + environment);
            
        } catch (IOException e) {
            System.err.println("Error loading configurations: " + e.getMessage());
            createDefaultConfigurations();
        }
    }
    
    private void createDefaultConfigurations() {
        // Create base properties
        baseProperties.setProperty("app.name", "MyApplication");
        baseProperties.setProperty("app.version", "1.0.0");
        baseProperties.setProperty("logging.level", "INFO");
        
        // Create environment-specific properties
        if ("development".equals(environment)) {
            environmentProperties.setProperty("database.url", "jdbc:mysql://localhost:3306/devdb");
            environmentProperties.setProperty("server.port", "8080");
            environmentProperties.setProperty("debug.enabled", "true");
        } else if ("production".equals(environment)) {
            environmentProperties.setProperty("database.url", "jdbc:mysql://prod-server:3306/proddb");
            environmentProperties.setProperty("server.port", "80");
            environmentProperties.setProperty("debug.enabled", "false");
        }
        
        System.out.println("Created default configurations for: " + environment);
    }
    
    public String getProperty(String key) {
        // Environment-specific properties override base properties
        String value = environmentProperties.getProperty(key);
        return value != null ? value : baseProperties.getProperty(key);
    }
    
    public void displayConfiguration() {
        System.out.println("\n=== Configuration for " + environment.toUpperCase() + " ===");
        
        System.out.println("\n--- Base Properties ---");
        baseProperties.forEach((key, value) -> 
            System.out.println(key + " = " + value));
        
        System.out.println("\n--- Environment-Specific Properties ---");
        environmentProperties.forEach((key, value) -> 
            System.out.println(key + " = " + value));
        
        System.out.println("\n--- Effective Configuration ---");
        Properties allProps = new Properties();
        allProps.putAll(baseProperties);
        allProps.putAll(environmentProperties);
        
        allProps.forEach((key, value) -> 
            System.out.println(key + " = " + value));
    }
}

public class PropertiesDemo {
    public static void main(String[] args) {
        // 1. Basic properties management
        System.out.println("=== Basic Properties Management ===");
        basicPropertiesDemo();
        
        // 2. Advanced configuration management
        System.out.println("\n=== Advanced Configuration Management ===");
        advancedConfigurationDemo();
        
        // 3. Properties validation and manipulation
        System.out.println("\n=== Properties Validation and Manipulation ===");
        propertiesValidationDemo();
        
        // 4. Environment-specific configuration
        System.out.println("\n=== Environment-Specific Configuration ===");
        environmentConfigurationDemo();
    }
    
    private static void basicPropertiesDemo() {
        ConfigurationManager config = new ConfigurationManager("app.properties");
        
        // Display current properties
        config.displayProperties();
        
        // Get properties with different types
        System.out.println("\n--- Getting Properties ---");
        System.out.println("App Name: " + config.getProperty("app.name", "Unknown"));
        System.out.println("Server Port: " + config.getIntProperty("server.port", 8080));
        System.out.println("Debug Mode: " + config.getBooleanProperty("debug.enabled", false));
        System.out.println("Cache Size: " + config.getIntProperty("cache.size", 100));
        
        // Update properties
        System.out.println("\n--- Updating Properties ---");
        config.setProperty("app.version", "1.1.0");
        config.setProperty("server.port", 9090);
        config.setProperty("debug.enabled", true);
        config.setProperty("cache.size", 2000);
        
        // Save updated properties
        config.saveProperties();
        
        // Get properties by prefix
        System.out.println("\n--- Database Properties ---");
        Map<String, String> dbProps = config.getPropertiesByPrefix("database.");
        dbProps.forEach((key, value) -> System.out.println(key + " = " + value));
    }
    
    private static void advancedConfigurationDemo() {
        ConfigurationManager config = new ConfigurationManager("app.properties");
        
        // Export to XML
        config.exportToXML("app_config.xml");
        
        // Create new configuration manager and import from XML
        ConfigurationManager newConfig = new ConfigurationManager("new_app.properties");
        newConfig.importFromXML("app_config.xml");
        
        newConfig.displayProperties();
    }
    
    private static void propertiesValidationDemo() {
        ConfigurationManager config = new ConfigurationManager("app.properties");
        
        // Validate required properties
        List<String> requiredProps = Arrays.asList(
            "app.name", "database.url", "server.port"
        );
        
        boolean isValid = config.validateRequiredProperties(requiredProps);
        System.out.println("Configuration validation: " + (isValid ? "PASSED" : "FAILED"));
        
        // Add missing property and validate again
        if (!isValid) {
            config.setProperty("database.url", "jdbc:mysql://localhost:3306/testdb");
            config.saveProperties();
            
            isValid = config.validateRequiredProperties(requiredProps);
            System.out.println("After adding missing property: " + (isValid ? "PASSED" : "FAILED"));
        }
    }
    
    private static void environmentConfigurationDemo() {
        // Development environment
        System.out.println("--- Development Environment ---");
        AdvancedConfigurationManager devConfig = new AdvancedConfigurationManager("development");
        devConfig.displayConfiguration();
        
        // Production environment
        System.out.println("\n--- Production Environment ---");
        AdvancedConfigurationManager prodConfig = new AdvancedConfigurationManager("production");
        prodConfig.displayConfiguration();
    }
    
    // Cleanup method
    private static void cleanupFiles() {
        try {
            Files.deleteIfExists(Paths.get("app.properties"));
            Files.deleteIfExists(Paths.get("new_app.properties"));
            Files.deleteIfExists(Paths.get("app_config.xml"));
            Files.deleteIfExists(Paths.get("config/base.properties"));
            Files.deleteIfExists(Paths.get("config/development.properties"));
            Files.deleteIfExists(Paths.get("config/production.properties"));
        } catch (IOException e) {
            System.err.println("Error cleaning up files: " + e.getMessage());
        }
    }
}
```

## 4. NIO.2 File Operations

### Modern File System Operations

```java
import java.nio.file.*;
import java.nio.file.attribute.*;
import java.io.*;
import java.util.*;
import java.util.stream.*;
import java.nio.charset.*;

class NIO2FileManager {
    
    // Create directories with attributes
    public void createDirectoriesWithAttributes(String dirPath) throws IOException {
        Path path = Paths.get(dirPath);
        
        // Create directory with specific attributes
        Set<PosixFilePermission> permissions = PosixFilePermissions.fromString("rwxr-x---");
        FileAttribute<Set<PosixFilePermission>> attr = PosixFilePermissions.asFileAttribute(permissions);
        
        Files.createDirectories(path, attr);
        System.out.println("Directory created with permissions: " + dirPath);
        
        // Display directory attributes
        displayFileAttributes(path);
    }
    
    // Copy file with progress monitoring
    public void copyFileWithProgress(String sourcePath, String targetPath) throws IOException {
        Path source = Paths.get(sourcePath);
        Path target = Paths.get(targetPath);
        
        // Get file size for progress tracking
        long fileSize = Files.size(source);
        System.out.println("Copying file of size: " + fileSize + " bytes");
        
        // Copy with options
        CopyOption[] options = {
            StandardCopyOption.REPLACE_EXISTING,
            StandardCopyOption.COPY_ATTRIBUTES
        };
        
        long startTime = System.currentTimeMillis();
        Files.copy(source, target, options);
        long endTime = System.currentTimeMillis();
        
        System.out.println("File copied successfully in " + (endTime - startTime) + " ms");
        System.out.println("Target file: " + target);
    }
    
    // Move file with atomic operation
    public void moveFileAtomically(String sourcePath, String targetPath) throws IOException {
        Path source = Paths.get(sourcePath);
        Path target = Paths.get(targetPath);
        
        try {
            Files.move(source, target, StandardCopyOption.ATOMIC_MOVE);
            System.out.println("File moved atomically: " + source + " -> " + target);
        } catch (AtomicMoveNotSupportedException e) {
            System.out.println("Atomic move not supported, falling back to regular move");
            Files.move(source, target, StandardCopyOption.REPLACE_EXISTING);
        }
    }
    
    // Watch directory for changes
    public void watchDirectory(String dirPath) throws IOException, InterruptedException {
        Path path = Paths.get(dirPath);
        WatchService watchService = FileSystems.getDefault().newWatchService();
        
        // Register directory for events
        path.register(watchService, 
            StandardWatchEventKinds.ENTRY_CREATE,
            StandardWatchEventKinds.ENTRY_DELETE,
            StandardWatchEventKinds.ENTRY_MODIFY);
        
        System.out.println("Watching directory: " + dirPath);
        System.out.println("Press Ctrl+C to stop watching...");
        
        while (true) {
            WatchKey key = watchService.take();
            
            for (WatchEvent<?> event : key.pollEvents()) {
                WatchEvent.Kind<?> kind = event.kind();
                Path fileName = (Path) event.context();
                
                System.out.println(kind.name() + ": " + fileName);
                
                if (kind == StandardWatchEventKinds.OVERFLOW) {
                    continue;
                }
                
                // Handle the event
                Path fullPath = path.resolve(fileName);
                handleFileEvent(kind, fullPath);
            }
            
            boolean valid = key.reset();
            if (!valid) {
                break;
            }
        }
    }
    
    // Search files with advanced criteria
    public List<Path> searchFiles(String rootDir, String fileNamePattern, int maxDepth) throws IOException {
        Path startPath = Paths.get(rootDir);
        List<Path> foundFiles = new ArrayList<>();
        
        try (Stream<Path> stream = Files.walk(startPath, maxDepth)) {
            stream.filter(Files::isRegularFile)
                 .filter(path -> path.getFileName().toString().matches(fileNamePattern))
                 .forEach(foundFiles::add);
        }
        
        System.out.println("Found " + foundFiles.size() + " files matching pattern: " + fileNamePattern);
        return foundFiles;
    }
    
    // Find duplicate files
    public Map<String, List<Path>> findDuplicateFiles(String directory) throws IOException {
        Map<String, List<Path>> duplicates = new HashMap<>();
        
        try (Stream<Path> stream = Files.walk(Paths.get(directory))) {
            stream.filter(Files::isRegularFile)
                 .collect(Collectors.groupingBy(path -> {
                     try {
                         return Files.size(path) + "_" + Files.getLastModifiedTime(path);
                     } catch (IOException e) {
                         return "";
                     }
                 }))
                 .entrySet().stream()
                 .filter(entry -> entry.getValue().size() > 1)
                 .forEach(entry -> duplicates.put(entry.getKey(), entry.getValue()));
        }
        
        System.out.println("Found " + duplicates.size() + " groups of duplicate files");
        return duplicates;
    }
    
    // Create symbolic link
    public void createSymbolicLink(String linkPath, String targetPath) throws IOException {
        Path link = Paths.get(linkPath);
        Path target = Paths.get(targetPath);
        
        try {
            Files.createSymbolicLink(link, target);
            System.out.println("Symbolic link created: " + link + " -> " + target);
        } catch (UnsupportedOperationException e) {
            System.out.println("Symbolic links not supported on this platform");
        }
    }
    
    // Read file with encoding
    public String readFileWithEncoding(String filePath, String encoding) throws IOException {
        Path path = Paths.get(filePath);
        byte[] bytes = Files.readAllBytes(path);
        return new String(bytes, Charset.forName(encoding));
    }
    
    // Write file with encoding
    public void writeFileWithEncoding(String filePath, String content, String encoding) throws IOException {
        Path path = Paths.get(filePath);
        byte[] bytes = content.getBytes(Charset.forName(encoding));
        Files.write(path, bytes);
        System.out.println("File written with encoding " + encoding + ": " + filePath);
    }
    
    // Get file space information
    public void displayDiskSpace(String path) throws IOException {
        Path dir = Paths.get(path);
        FileStore store = Files.getFileStore(dir);
        
        long total = store.getTotalSpace();
        long free = store.getUnallocatedSpace();
        long usable = store.getUsableSpace();
        
        System.out.println("Disk Space Information for: " + path);
        System.out.printf("Total: %,d bytes (%.2f GB)\n", total, total / (1024.0 * 1024 * 1024));
        System.out.printf("Free: %,d bytes (%.2f GB)\n", free, free / (1024.0 * 1024 * 1024));
        System.out.printf("Usable: %,d bytes (%.2f GB)\n", usable, usable / (1024.0 * 1024 * 1024));
        System.out.printf("Used: %,d bytes (%.2f GB)\n", total - usable, (total - usable) / (1024.0 * 1024 * 1024));
    }
    
    // Display file attributes
    private void displayFileAttributes(Path path) throws IOException {
        System.out.println("\n--- File Attributes for: " + path + " ---");
        
        BasicFileAttributes attrs = Files.readAttributes(path, BasicFileAttributes.class);
        System.out.println("Size: " + attrs.size() + " bytes");
        System.out.println("Creation Time: " + attrs.creationTime());
        System.out.println("Last Modified: " + attrs.lastModifiedTime());
        System.out.println("Last Accessed: " + attrs.lastAccessTime());
        System.out.println("Is Directory: " + attrs.isDirectory());
        System.out.println("Is Regular File: " + attrs.isRegularFile());
        System.out.println("Is Symbolic Link: " + attrs.isSymbolicLink());
        
        if (Files.isRegularFile(path)) {
            System.out.println("Content Type: " + Files.probeContentType(path));
        }
    }
    
    private void handleFileEvent(WatchEvent.Kind<?> kind, Path fullPath) {
        try {
            System.out.println("  -> Full path: " + fullPath);
            if (Files.isRegularFile(fullPath)) {
                System.out.println("  -> Size: " + Files.size(fullPath) + " bytes");
            }
        } catch (IOException e) {
            System.err.println("Error handling file event: " + e.getMessage());
        }
    }
}

public class NIO2Demo {
    public static void main(String[] args) {
        NIO2FileManager fileManager = new NIO2FileManager();
        
        try {
            // 1. Create directories with attributes
            System.out.println("=== Creating Directories ===");
            fileManager.createDirectoriesWithAttributes("test_directory/subdir");
            
            // 2. Create sample files for testing
            createSampleFiles();
            
            // 3. Copy file with progress
            System.out.println("\n=== Copying Files ===");
            fileManager.copyFileWithProgress("sample.txt", "sample_copy.txt");
            
            // 4. Move file atomically
            System.out.println("\n=== Moving Files ===");
            fileManager.moveFileAtomically("sample_copy.txt", "moved_sample.txt");
            
            // 5. Search files
            System.out.println("\n=== Searching Files ===");
            List<Path> foundFiles = fileManager.searchFiles(".", ".*\\.txt", 2);
            foundFiles.forEach(System.out::println);
            
            // 6. Find duplicates
            System.out.println("\n=== Finding Duplicates ===");
            Map<String, List<Path>> duplicates = fileManager.findDuplicateFiles(".");
            duplicates.forEach((key, paths) -> {
                System.out.println("Duplicate group: " + key);
                paths.forEach(path -> System.out.println("  " + path));
            });
            
            // 7. File encoding operations
            System.out.println("\n=== File Encoding Operations ===");
            String content = "Hello, 世界! This is a test file.";
            fileManager.writeFileWithEncoding("utf8_sample.txt", content, "UTF-8");
            String readContent = fileManager.readFileWithEncoding("utf8_sample.txt", "UTF-8");
            System.out.println("Read content: " + readContent);
            
            // 8. Disk space information
            System.out.println("\n=== Disk Space Information ===");
            fileManager.displayDiskSpace(".");
            
            // 9. Symbolic link creation
            System.out.println("\n=== Creating Symbolic Links ===");
            fileManager.createSymbolicLink("sample_link.txt", "sample.txt");
            
            // 10. File attributes display
            System.out.println("\n=== File Attributes ===");
            fileManager.displayFileAttributes(Paths.get("sample.txt"));
            
        } catch (IOException | InterruptedException e) {
            System.err.println("Error in NIO.2 operations: " + e.getMessage());
        } finally {
            // Cleanup
            cleanupFiles();
        }
    }
    
    private static void createSampleFiles() throws IOException {
        // Create sample files
        Files.write(Paths.get("sample.txt"), "This is a sample file for testing.".getBytes());
        Files.write(Paths.get("another_sample.txt"), "This is another sample file.".getBytes());
        Files.write(Paths.get("test.txt"), "Test file content.".getBytes());
        
        // Create duplicate files
        Files.write(Paths.get("duplicate1.txt"), "Same content".getBytes());
        Files.write(Paths.get("duplicate2.txt"), "Same content".getBytes());
        
        System.out.println("Sample files created for testing");
    }
    
    private static void cleanupFiles() {
        try {
            Files.deleteIfExists(Paths.get("sample.txt"));
            Files.deleteIfExists(Paths.get("sample_copy.txt"));
            Files.deleteIfExists(Paths.get("moved_sample.txt"));
            Files.deleteIfExists(Paths.get("another_sample.txt"));
            Files.deleteIfExists(Paths.get("test.txt"));
            Files.deleteIfExists(Paths.get("duplicate1.txt"));
            Files.deleteIfExists(Paths.get("duplicate2.txt"));
            Files.deleteIfExists(Paths.get("utf8_sample.txt"));
            Files.deleteIfExists(Paths.get("sample_link.txt"));
            Files.deleteIfExists(Paths.get("test_directory"));
            Files.deleteIfExists(Paths.get("test_directory/subdir"));
        } catch (IOException e) {
            System.err.println("Error cleaning up files: " + e.getMessage());
        }
    }
}
```

## Practice Exercises

1. **CSV Processing**: Create a data import/export system for a contact management application
2. **Excel Operations**: Build a financial reporting system with charts and formulas
3. **Properties Management**: Design a configuration system with environment-specific settings
4. **NIO.2 Operations**: Implement a file synchronization utility

## Interview Questions

1. What's the difference between traditional I/O and NIO.2?
2. How do you handle different character encodings when reading files?
3. What are the benefits of using try-with-resources for file operations?
4. How does the WatchService work for monitoring file system changes?
5. What's the difference between Files.copy() and Files.move() in NIO.2?
