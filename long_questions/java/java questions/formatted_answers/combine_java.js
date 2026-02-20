const fs = require('fs');
const path = require('path');

const dir = process.cwd();
const targetDir = path.join(dir, 'old_format');

if (!fs.existsSync(targetDir)) {
  fs.mkdirSync(targetDir);
}

const levels = {
  '01_Basics.md': ['01_Java_Core_Basics.md', '04_Java_Fundamentals_Core.md', '06_OOP_Basics.md', '14_SOLID_Arrays_Basics.md', '15_Strings_Basics.md', '31_Spring_Boot_Basics_Revision.md'],
  '02_Intermediate.md': ['03_Exceptions_And_IO.md', '07_OOP_Basics_Practice.md', '08_Exceptions_IO_Practice.md', '09_Java_Fundamentals_Practice.md', '16_Data_Structures_Collections.md', '18_Java_Programs_Numbers.md', '19_Java_Programs_Arrays.md', '20_Java_Programs_Strings.md', '21_Java_Programs_Patterns_OOP.md', '25_Arrays_And_Strings_Revision.md', '26_Data_Structures_Intermediate_Revision.md'],
  '03_Advanced.md': ['02_Spring_And_Advanced_Java.md', '05_Modern_Java_And_Patterns.md', '10_Modern_Java_And_Patterns_Practice.md', '17_Data_Structures_Streams_Advanced.md', '22_Java_Programs_Collections_Advanced.md', '23_Java_Programs_Advanced_SQL.md', '27_Data_Structures_Advanced_Algorithms.md', '28_Spring_Boot_3_Advanced.md', '29_Spring_Boot_Config_REST.md', '34_Spring_Boot_Architecture_Config.md'],
  '04_Senior.md': ['24_Mixed_Concepts_Patterns_DB_Testing.md', '30_Spring_Boot_Data_Security.md', '35_Spring_Boot_REST_CLI_MongoDB.md', '37_Spring_Boot_Internals_Testing.md', '38_Spring_Boot_Deployment_Security_JPA.md', '39_Spring_Boot_Security_Testing_JPA_Revision.md', '40_Spring_MVC_Security_WebFlux_Revision.md'],
  '05_Expert.md': ['11_Concurrency_Practice.md', '12_Extra_Concepts_Practice.md', '13_Advanced_Concurrency_JVM_Practice.md', '32_Spring_Core_Monitoring_WebFlux.md', '33_Messaging_Kafka_Docker_Kubernetes.md', '36_Spring_Boot_NoSQL_Integration_Cloud.md']
};

for (const [newFile, oldFiles] of Object.entries(levels)) {
  let content = [];
  const titles = [];
  
  if (newFile === '01_Basics.md') titles.push('# Basic Level Java Interview Questions\n');
  if (newFile === '02_Intermediate.md') titles.push('# Intermediate Level Java Interview Questions\n');
  if (newFile === '03_Advanced.md') titles.push('# Advanced Level Java Interview Questions\n');
  if (newFile === '04_Senior.md') titles.push('# Senior Level Java Interview Questions\n');
  if (newFile === '05_Expert.md') titles.push('# Expert Level Java Interview Questions\n');

  for (const oldFile of oldFiles) {
    const filePath = path.join(dir, oldFile);
    if (!fs.existsSync(filePath)) {
      console.warn('File not found:', oldFile);
      continue;
    }
    const fileContent = fs.readFileSync(filePath, 'utf-8');
    
    // Check if the original file starts with `# ` indicating a title, so we can keep questions inside
    // but the main title will be at the root. We keep all their formatting.
    content.push(`\n## From ${oldFile.replace('.md', '').replaceAll('_', ' ')}\n`);
    content.push(fileContent);
    
    // Move to backup
    const newOldPath = path.join(targetDir, oldFile);
    fs.renameSync(filePath, newOldPath);
  }
  
  fs.writeFileSync(path.join(dir, newFile), titles.join('') + content.join(''));
  console.log(`Created ${newFile}`);
}

console.log('Done organizing files.');
