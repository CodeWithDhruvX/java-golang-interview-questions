const fs = require('fs');
const path = require('path');

const dir = process.cwd();
const targetDir = path.join(dir, 'old_format');

if (!fs.existsSync(targetDir)) {
    fs.mkdirSync(targetDir);
}

const levels = {
    '01_Basics.md': [
        '01_Basics.md',
        '02_Arrays_Slices_Maps.md',
        '03_Pointers_Interfaces_Methods.md',
        '12_Files_OS_System.md',
        '51_Modern_Go_Features.md'
    ],
    '02_Intermediate.md': [
        '04_Concurrency.md',
        '08_Networking_WebDev.md',
        '09_Databases_ORMs.md',
        '10_Tools_Testing_Ecosystem.md',
        '23_CLI_Automation.md',
        '28_Testing.md',
        '40_Tooling_DevExp.md',
        '42_Testing_Part2.md',
        '47_Databases_Part2.md'
    ],
    '03_Advanced.md': [
        '05_Advanced_BestPractices.md',
        '06_ProjectStructure_DesignPatterns.md',
        '07_Generics_AdvancedTypes.md',
        '11_Performance_Optimization.md',
        '14_Security_BestPractices.md',
        '15_Testing_Strategy.md',
        '22_ErrorHandling_Observability.md',
        '26_Security.md',
        '27_Performance_Optimization.md',
        '29_API_Design_REST_gRPC.md',
        '30_DesignPatterns_Part2.md',
        '31_Advanced_Concurrency.md',
        '38_ErrorHandling_Part2.md',
        '41_Security_Part2.md',
        '43_Performance_Part2.md',
        '45_Refactoring_Design.md',
        '49_Concurrency_Patterns_Part2.md'
    ],
    '04_Senior.md': [
        '13_Microservices_gRPC.md',
        '17_DevOps_Containers.md',
        '18_Streaming_Async.md',
        '19_Architecture_SystemDesign.md',
        '20_Troubleshooting_Debugging.md',
        '32_EventDriven_Messaging.md',
        '33_DevOps_Infrastructure.md',
        '34_Caching_Storage.md',
        '48_API_Microservices_Part2.md'
    ],
    '05_Expert.md': [
        '16_Go_Internals.md',
        '21_Networking_LowLevel.md',
        '24_AI_MachineLearning.md',
        '25_WASM_Blockchain.md',
        '35_RealTime_IoT.md',
        '36_Go_Internals.md',
        '37_Network_Protocol_DeepDive.md',
        '39_Streaming_DataPipelines.md',
        '44_Compiler_Theory.md',
        '46_AI_ML_Part2.md',
        '50_Tooling_Maintenance_Part2.md',
        '52_Niche_Patterns.md'
    ]
};

for (const [newFile, oldFiles] of Object.entries(levels)) {
    let content = [];
    const titles = [];

    if (newFile === '01_Basics.md') titles.push('# Basic Level Golang Interview Questions\n');
    if (newFile === '02_Intermediate.md') titles.push('# Intermediate Level Golang Interview Questions\n');
    if (newFile === '03_Advanced.md') titles.push('# Advanced Level Golang Interview Questions\n');
    if (newFile === '04_Senior.md') titles.push('# Senior Level Golang Interview Questions\n');
    if (newFile === '05_Expert.md') titles.push('# Expert Level Golang Interview Questions\n');

    for (const oldFile of oldFiles) {
        const filePath = path.join(dir, oldFile);
        if (!fs.existsSync(filePath)) {
            console.warn('File not found:', oldFile);
            continue;
        }
        const fileContent = fs.readFileSync(filePath, 'utf-8');

        content.push(`\n## From ${oldFile.replace('.md', '').replaceAll('_', ' ')}\n`);
        content.push(fileContent);

        // Move to backup
        const newOldPath = path.join(targetDir, oldFile);
        // don't overwrite if we run multiple times by accident, or if source and target are same (should not happen)
        fs.renameSync(filePath, newOldPath);
    }

    fs.writeFileSync(path.join(dir, newFile), titles.join('') + content.join('\n'));
    console.log(`Created ${newFile}`);
}

console.log('Done organizing files.');
