# Terraform Practical Projects - Complete Solutions

This file provides step-by-step code implementations and detailed answers for all Terraform practical projects.

---

# 🟢 LEVEL 1: Beginner Projects

## 1. Local File Generator - Complete Solution

### Project Structure:
```
local-file-project/
├── main.tf
├── variables.tf
├── outputs.tf
└── terraform.tfvars
```

### main.tf:
```hcl
terraform {
  required_version = ">= 1.0"
  required_providers {
    local = {
      source  = "hashicorp/local"
      version = "~> 2.0"
    }
  }
}

# Create multiple configuration files
resource "local_file" "app_config" {
  content  = "app_name=${var.app_name}\nversion=${var.version}\nenvironment=${var.environment}"
  filename = "${path.module}/config/app.conf"
}

resource "local_file" "database_config" {
  content  = "host=${var.db_host}\nport=${var.db_port}\ndatabase=${var.db_name}"
  filename = "${path.module}/config/database.conf"
}

resource "local_file" "readme" {
  content  = "# ${var.app_name}\n\nThis is a ${var.environment} deployment.\nVersion: ${var.version}"
  filename = "${path.module}/README.md"
}
```

### variables.tf:
```hcl
variable "app_name" {
  description = "Name of the application"
  type        = string
  default     = "my-app"
}

variable "version" {
  description = "Application version"
  type        = string
  default     = "1.0.0"
}

variable "environment" {
  description = "Deployment environment"
  type        = string
  default     = "development"
  
  validation {
    condition     = contains(["development", "staging", "production"], var.environment)
    error_message = "Environment must be one of: development, staging, production."
  }
}

variable "db_host" {
  description = "Database host"
  type        = string
  default     = "localhost"
}

variable "db_port" {
  description = "Database port"
  type        = number
  default     = 5432
  
  validation {
    condition     = var.db_port > 0 && var.db_port <= 65535
    error_message = "Port must be between 1 and 65535."
  }
}

variable "db_name" {
  description = "Database name"
  type        = string
  default     = "myapp_db"
}
```

### outputs.tf:
```hcl
output "config_files" {
  description = "List of generated configuration files"
  value = [
    local_file.app_config.filename,
    local_file.database_config.filename,
    local_file.readme.filename
  ]
}

output "app_info" {
  description = "Application information"
  value = {
    name        = var.app_name
    version     = var.version
    environment = var.environment
  }
}
```

### terraform.tfvars:
```hcl
app_name    = "production-app"
version     = "2.1.0"
environment = "production"
db_host     = "prod-db.example.com"
db_port     = 5432
db_name     = "production_db"
```

### Commands to run:
```bash
# Initialize Terraform
terraform init

# Plan the changes
terraform plan

# Apply the changes
terraform apply

# Destroy when done
terraform destroy
```

---

## 2. Docker Container Deployment - Complete Solution

### Project Structure:
```
docker-project/
├── main.tf
├── variables.tf
├── outputs.tf
└── docker-compose.yml (for comparison)
```

### main.tf:
```hcl
terraform {
  required_version = ">= 1.0"
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 3.0"
    }
  }
}

# Docker network
resource "docker_network" "app_network" {
  name = "app-network"
}

# Nginx container
resource "docker_container" "nginx" {
  name  = "nginx-server"
  image = "nginx:latest"

  ports {
    internal = 80
    external = var.nginx_port
  }

  networks {
    name = docker_network.app_network.name
  }

  restart = "unless-stopped"
}

# Redis container
resource "docker_container" "redis" {
  name  = "redis-cache"
  image = "redis:latest"

  ports {
    internal = 6379
    external = var.redis_port
  }

  networks {
    name = docker_network.app_network.name
  }

  restart = "unless-stopped"
}

# Custom application container
resource "docker_image" "app" {
  name         = "my-app:latest"
  build_inline {
    context = "."
    dockerfile = <<-EOT
      FROM node:18-alpine
      WORKDIR /app
      COPY package*.json ./
      RUN npm install
      COPY . .
      EXPOSE 3000
      CMD ["npm", "start"]
    EOT
  }
}

resource "docker_container" "app" {
  name  = "my-app"
  image = docker_image.app.latest

  ports {
    internal = 3000
    external = var.app_port
  }

  env = [
    "NODE_ENV=${var.environment}",
    "REDIS_HOST=redis",
    "REDIS_PORT=6379"
  ]

  networks {
    name = docker_network.app_network.name
  }

  depends_on = [
    docker_container.redis
  ]

  restart = "unless-stopped"
}
```

### variables.tf:
```hcl
variable "nginx_port" {
  description = "External port for Nginx"
  type        = number
  default     = 8080
}

variable "redis_port" {
  description = "External port for Redis"
  type        = number
  default     = 6379
}

variable "app_port" {
  description = "External port for application"
  type        = number
  default     = 3000
}

variable "environment" {
  description = "Application environment"
  type        = string
  default     = "development"
}
```

### outputs.tf:
```hcl
output "container_info" {
  description = "Information about deployed containers"
  value = {
    nginx = {
      name = docker_container.nginx.name
      port = "${docker_container.nginx.ports[0].external}:${docker_container.nginx.ports[0].internal}"
    }
    redis = {
      name = docker_container.redis.name
      port = "${docker_container.redis.ports[0].external}:${docker_container.redis.ports[0].internal}"
    }
    app = {
      name = docker_container.app.name
      port = "${docker_container.app.ports[0].external}:${docker_container.app.ports[0].internal}"
    }
  }
}

output "network_name" {
  description = "Docker network name"
  value       = docker_network.app_network.name
}
```

---

# 🔵 LEVEL 2: Basic Cloud Projects

## 3. Single EC2 Instance - Complete Solution

### Project Structure:
```
ec2-project/
├── main.tf
├── variables.tf
├── outputs.tf
├── user_data.sh
└── terraform.tfvars
```

### main.tf:
```hcl
terraform {
  required_version = ">= 1.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
}

# SSH key pair
resource "aws_key_pair" "deployer" {
  key_name   = "deployer-key"
  public_key = file(var.public_key_path)
}

# Security Group
resource "aws_security_group" "web_sg" {
  name        = "web-security-group"
  description = "Allow SSH and HTTP traffic"

  # SSH access
  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # HTTP access
  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # HTTPS access
  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "web-sg"
  }
}

# EC2 Instance
resource "aws_instance" "web_server" {
  ami                    = var.ami_id
  instance_type          = var.instance_type
  key_name               = aws_key_pair.deployer.key_name
  vpc_security_group_ids = [aws_security_group.web_sg.id]
  user_data              = file("user_data.sh")

  root_block_device {
    volume_size           = 20
    volume_type           = "gp3"
    delete_on_termination = true
  }

  tags = {
    Name        = "web-server"
    Environment = var.environment
  }
}
```

### user_data.sh:
```bash
#!/bin/bash
yum update -y
yum install -y httpd
systemctl start httpd
systemctl enable httpd

# Create a simple HTML page
echo "<html>
<head>
    <title>My Terraform Server</title>
</head>
<body>
    <h1>Hello from Terraform!</h1>
    <p>This server was deployed using Terraform on $(date)</p>
    <p>Environment: ${environment}</p>
</body>
</html>" > /var/www/html/index.html

# Create a health check endpoint
echo "OK" > /var/www/html/health
```

### variables.tf:
```hcl
variable "aws_region" {
  description = "AWS region"
  type        = string
  default     = "us-east-1"
}

variable "ami_id" {
  description = "AMI ID for EC2 instance"
  type        = string
  default     = "ami-0c02fb55956c7d316" # Amazon Linux 2
}

variable "instance_type" {
  description = "EC2 instance type"
  type        = string
  default     = "t2.micro"
}

variable "environment" {
  description = "Environment name"
  type        = string
  default     = "development"
}

variable "public_key_path" {
  description = "Path to SSH public key"
  type        = string
  default     = "~/.ssh/id_rsa.pub"
}
```

### outputs.tf:
```hcl
output "instance_public_ip" {
  description = "Public IP address of EC2 instance"
  value       = aws_instance.web_server.public_ip
}

output "instance_id" {
  description = "ID of the EC2 instance"
  value       = aws_instance.web_server.id
}

output "security_group_id" {
  description = "ID of the security group"
  value       = aws_security_group.web_sg.id
}
```

---

## 4. S3 Bucket + Static Website - Complete Solution

### main.tf:
```hcl
terraform {
  required_version = ">= 1.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
}

# S3 Bucket
resource "aws_s3_bucket" "website_bucket" {
  bucket = var.bucket_name

  tags = {
    Name        = "website-bucket"
    Environment = var.environment
  }
}

# Bucket ownership controls
resource "aws_s3_bucket_ownership_controls" "website_bucket_ownership" {
  bucket = aws_s3_bucket.website_bucket.id

  rule {
    object_ownership = "BucketOwnerPreferred"
  }
}

# Public access block configuration
resource "aws_s3_bucket_public_access_block" "website_bucket_pab" {
  bucket = aws_s3_bucket.website_bucket.id

  block_public_acls       = false
  block_public_policy     = false
  ignore_public_acls      = false
  restrict_public_buckets = false
}

# Bucket ACL
resource "aws_s3_bucket_acl" "website_bucket_acl" {
  depends_on = [
    aws_s3_bucket_ownership_controls.website_bucket_ownership,
    aws_s3_bucket_public_access_block.website_bucket_pab,
  ]

  bucket = aws_s3_bucket.website_bucket.id
  acl    = "public-read"
}

# Static website configuration
resource "aws_s3_bucket_website_configuration" "website_config" {
  bucket = aws_s3_bucket.website_bucket.id

  index_document {
    suffix = "index.html"
  }

  error_document {
    key = "error.html"
  }
}

# Upload website files
resource "aws_s3_object" "index_html" {
  bucket       = aws_s3_bucket.website_bucket.id
  key          = "index.html"
  content_type = "text/html"
  source       = "${path.module}/website/index.html"
  etag         = filemd5("${path.module}/website/index.html")
}

resource "aws_s3_object" "error_html" {
  bucket       = aws_s3_bucket.website_bucket.id
  key          = "error.html"
  content_type = "text/html"
  source       = "${path.module}/website/error.html"
  etag         = filemd5("${path.module}/website/error.html")
}

resource "aws_s3_object" "css_style" {
  bucket       = aws_s3_bucket.website_bucket.id
  key          = "style.css"
  content_type = "text/css"
  source       = "${path.module}/website/style.css"
  etag         = filemd5("${path.module}/website/style.css")
}
```

### website/index.html:
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>My Static Website</title>
    <link rel="stylesheet" href="style.css">
</head>
<body>
    <header>
        <h1>Welcome to My Static Website</h1>
        <p>Deployed with Terraform on AWS S3</p>
    </header>
    
    <main>
        <section>
            <h2>About This Project</h2>
            <p>This is a static website deployed using Terraform infrastructure as code.</p>
        </section>
        
        <section>
            <h2>Technologies Used</h2>
            <ul>
                <li>Terraform for infrastructure</li>
                <li>AWS S3 for hosting</li>
                <li>HTML/CSS for the website</li>
            </ul>
        </section>
    </main>
    
    <footer>
        <p>&copy; 2024 My Static Website. Environment: ${environment}</p>
    </footer>
</body>
</html>
```

### website/style.css:
```css
* {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
}

body {
    font-family: Arial, sans-serif;
    line-height: 1.6;
    color: #333;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    min-height: 100vh;
}

header {
    background: rgba(255, 255, 255, 0.1);
    padding: 2rem;
    text-align: center;
    color: white;
}

main {
    max-width: 800px;
    margin: 2rem auto;
    padding: 0 1rem;
}

section {
    background: rgba(255, 255, 255, 0.9);
    padding: 2rem;
    margin-bottom: 2rem;
    border-radius: 10px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

h1, h2 {
    color: #2c3e50;
    margin-bottom: 1rem;
}

ul {
    margin-left: 2rem;
}

footer {
    background: rgba(255, 255, 255, 0.1);
    padding: 1rem;
    text-align: center;
    color: white;
    position: fixed;
    bottom: 0;
    width: 100%;
}
```

### website/error.html:
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Page Not Found</title>
    <link rel="stylesheet" href="style.css">
</head>
<body>
    <main style="text-align: center; margin-top: 10rem;">
        <section>
            <h1>404 - Page Not Found</h1>
            <p>The page you're looking for doesn't exist.</p>
            <a href="/" style="color: #667eea;">Go back to homepage</a>
        </section>
    </main>
</body>
</html>
```

---

# 🟡 LEVEL 3: Intermediate Projects

## 5. VPC + Subnet + EC2 - Complete Solution

### Project Structure:
```
vpc-project/
├── main.tf
├── variables.tf
├── outputs.tf
├── user_data.sh
└── terraform.tfvars
```

### main.tf:
```hcl
terraform {
  required_version = ">= 1.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
}

# VPC
resource "aws_vpc" "main" {
  cidr_block           = var.vpc_cidr
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = {
    Name        = "main-vpc"
    Environment = var.environment
  }
}

# Internet Gateway
resource "aws_internet_gateway" "main" {
  vpc_id = aws_vpc.main.id

  tags = {
    Name = "main-igw"
  }
}

# Public Subnets
resource "aws_subnet" "public" {
  count                   = length(var.public_subnet_cidrs)
  vpc_id                  = aws_vpc.main.id
  cidr_block              = var.public_subnet_cidrs[count.index]
  availability_zone       = var.availability_zones[count.index]
  map_public_ip_on_launch = true

  tags = {
    Name        = "public-subnet-${count.index + 1}"
    Environment = var.environment
    Type        = "Public"
  }
}

# Route Table for Public Subnets
resource "aws_route_table" "public" {
  vpc_id = aws_vpc.main.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.main.id
  }

  tags = {
    Name = "public-rt"
  }
}

# Route Table Associations
resource "aws_route_table_association" "public" {
  count          = length(aws_subnet.public)
  subnet_id      = aws_subnet.public[count.index].id
  route_table_id = aws_route_table.public.id
}

# Private Subnets
resource "aws_subnet" "private" {
  count             = length(var.private_subnet_cidrs)
  vpc_id            = aws_vpc.main.id
  cidr_block        = var.private_subnet_cidrs[count.index]
  availability_zone = var.availability_zones[count.index]

  tags = {
    Name        = "private-subnet-${count.index + 1}"
    Environment = var.environment
    Type        = "Private"
  }
}

# Elastic IP for NAT Gateway
resource "aws_eip" "nat" {
  domain = "vpc"
  
  tags = {
    Name = "nat-eip"
  }
}

# NAT Gateway
resource "aws_nat_gateway" "main" {
  allocation_id = aws_eip.nat.id
  subnet_id     = aws_subnet.public[0].id

  tags = {
    Name = "main-nat"
  }

  depends_on = [aws_internet_gateway.main]
}

# Route Table for Private Subnets
resource "aws_route_table" "private" {
  vpc_id = aws_vpc.main.id

  route {
    cidr_block     = "0.0.0.0/0"
    nat_gateway_id = aws_nat_gateway.main.id
  }

  tags = {
    Name = "private-rt"
  }
}

# Route Table Associations for Private Subnets
resource "aws_route_table_association" "private" {
  count          = length(aws_subnet.private)
  subnet_id      = aws_subnet.private[count.index].id
  route_table_id = aws_route_table.private.id
}

# Security Group
resource "aws_security_group" "web_sg" {
  name        = "web-security-group"
  description = "Allow SSH and HTTP traffic"
  vpc_id      = aws_vpc.main.id

  # SSH access
  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # HTTP access
  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "web-sg"
  }
}

# EC2 Instance in Public Subnet
resource "aws_instance" "web_server" {
  ami                    = var.ami_id
  instance_type          = var.instance_type
  key_name               = aws_key_pair.deployer.key_name
  vpc_security_group_ids = [aws_security_group.web_sg.id]
  subnet_id              = aws_subnet.public[0].id
  user_data              = file("user_data.sh")

  tags = {
    Name        = "web-server"
    Environment = var.environment
  }
}

# SSH Key Pair
resource "aws_key_pair" "deployer" {
  key_name   = "deployer-key"
  public_key = file(var.public_key_path)
}
```

### variables.tf:
```hcl
variable "aws_region" {
  description = "AWS region"
  type        = string
  default     = "us-east-1"
}

variable "environment" {
  description = "Environment name"
  type        = string
  default     = "development"
}

variable "vpc_cidr" {
  description = "CIDR block for VPC"
  type        = string
  default     = "10.0.0.0/16"
}

variable "public_subnet_cidrs" {
  description = "CIDR blocks for public subnets"
  type        = list(string)
  default     = ["10.0.1.0/24", "10.0.2.0/24"]
}

variable "private_subnet_cidrs" {
  description = "CIDR blocks for private subnets"
  type        = list(string)
  default     = ["10.0.10.0/24", "10.0.20.0/24"]
}

variable "availability_zones" {
  description = "Availability zones"
  type        = list(string)
  default     = ["us-east-1a", "us-east-1b"]
}

variable "ami_id" {
  description = "AMI ID for EC2 instance"
  type        = string
  default     = "ami-0c02fb55956c7d316"
}

variable "instance_type" {
  description = "EC2 instance type"
  type        = string
  default     = "t2.micro"
}

variable "public_key_path" {
  description = "Path to SSH public key"
  type        = string
  default     = "~/.ssh/id_rsa.pub"
}
```

### outputs.tf:
```hcl
output "vpc_id" {
  description = "ID of the VPC"
  value       = aws_vpc.main.id
}

output "public_subnet_ids" {
  description = "IDs of public subnets"
  value       = aws_subnet.public[*].id
}

output "private_subnet_ids" {
  description = "IDs of private subnets"
  value       = aws_subnet.private[*].id
}

output "internet_gateway_id" {
  description = "ID of the Internet Gateway"
  value       = aws_internet_gateway.main.id
}

output "nat_gateway_id" {
  description = "ID of the NAT Gateway"
  value       = aws_nat_gateway.main.id
}

output "instance_public_ip" {
  description = "Public IP address of EC2 instance"
  value       = aws_instance.web_server.public_ip
}

output "instance_private_ip" {
  description = "Private IP address of EC2 instance"
  value       = aws_instance.web_server.private_ip
}
```

---

## 6. Modular Terraform Project - Complete Solution

### Project Structure:
```
modular-project/
├── main.tf
├── variables.tf
├── outputs.tf
├── terraform.tfvars
├── modules/
│   ├── vpc/
│   │   ├── main.tf
│   │   ├── variables.tf
│   │   └── outputs.tf
│   ├── security/
│   │   ├── main.tf
│   │   ├── variables.tf
│   │   └── outputs.tf
│   └── ec2/
│       ├── main.tf
│       ├── variables.tf
│       └── outputs.tf
└── user_data.sh
```

### Root main.tf:
```hcl
terraform {
  required_version = ">= 1.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
}

# VPC Module
module "vpc" {
  source = "./modules/vpc"

  environment           = var.environment
  vpc_cidr             = var.vpc_cidr
  public_subnet_cidrs  = var.public_subnet_cidrs
  private_subnet_cidrs = var.private_subnet_cidrs
  availability_zones   = var.availability_zones
}

# Security Module
module "security" {
  source = "./modules/security"

  environment = var.environment
  vpc_id      = module.vpc.vpc_id
}

# EC2 Module
module "ec2" {
  source = "./modules/ec2"

  environment           = var.environment
  ami_id               = var.ami_id
  instance_type        = var.instance_type
  public_key_path      = var.public_key_path
  public_subnet_id     = module.vpc.public_subnet_ids[0]
  security_group_id    = module.security.web_security_group_id
  user_data_script     = file("user_data.sh")
}
```

### modules/vpc/main.tf:
```hcl
# VPC
resource "aws_vpc" "main" {
  cidr_block           = var.vpc_cidr
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = {
    Name        = "${var.environment}-vpc"
    Environment = var.environment
  }
}

# Internet Gateway
resource "aws_internet_gateway" "main" {
  vpc_id = aws_vpc.main.id

  tags = {
    Name = "${var.environment}-igw"
  }
}

# Public Subnets
resource "aws_subnet" "public" {
  count                   = length(var.public_subnet_cidrs)
  vpc_id                  = aws_vpc.main.id
  cidr_block              = var.public_subnet_cidrs[count.index]
  availability_zone       = var.availability_zones[count.index]
  map_public_ip_on_launch = true

  tags = {
    Name        = "${var.environment}-public-subnet-${count.index + 1}"
    Environment = var.environment
    Type        = "Public"
  }
}

# Private Subnets
resource "aws_subnet" "private" {
  count             = length(var.private_subnet_cidrs)
  vpc_id            = aws_vpc.main.id
  cidr_block        = var.private_subnet_cidrs[count.index]
  availability_zone = var.availability_zones[count.index]

  tags = {
    Name        = "${var.environment}-private-subnet-${count.index + 1}"
    Environment = var.environment
    Type        = "Private"
  }
}

# NAT Gateway
resource "aws_eip" "nat" {
  domain = "vpc"
  tags   = {
    Name = "${var.environment}-nat-eip"
  }
}

resource "aws_nat_gateway" "main" {
  allocation_id = aws_eip.nat.id
  subnet_id     = aws_subnet.public[0].id

  tags = {
    Name = "${var.environment}-nat"
  }

  depends_on = [aws_internet_gateway.main]
}

# Route Tables
resource "aws_route_table" "public" {
  vpc_id = aws_vpc.main.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.main.id
  }

  tags = {
    Name = "${var.environment}-public-rt"
  }
}

resource "aws_route_table" "private" {
  vpc_id = aws_vpc.main.id

  route {
    cidr_block     = "0.0.0.0/0"
    nat_gateway_id = aws_nat_gateway.main.id
  }

  tags = {
    Name = "${var.environment}-private-rt"
  }
}

# Route Table Associations
resource "aws_route_table_association" "public" {
  count          = length(aws_subnet.public)
  subnet_id      = aws_subnet.public[count.index].id
  route_table_id = aws_route_table.public.id
}

resource "aws_route_table_association" "private" {
  count          = length(aws_subnet.private)
  subnet_id      = aws_subnet.private[count.index].id
  route_table_id = aws_route_table.private.id
}
```

### modules/security/main.tf:
```hcl
# Web Security Group
resource "aws_security_group" "web" {
  name        = "${var.environment}-web-sg"
  description = "Allow SSH and HTTP traffic"
  vpc_id      = var.vpc_id

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name        = "${var.environment}-web-sg"
    Environment = var.environment
  }
}

# Database Security Group
resource "aws_security_group" "database" {
  name        = "${var.environment}-db-sg"
  description = "Allow database traffic"
  vpc_id      = var.vpc_id

  ingress {
    from_port       = 3306
    to_port         = 3306
    protocol        = "tcp"
    security_groups = [aws_security_group.web.id]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name        = "${var.environment}-db-sg"
    Environment = var.environment
  }
}
```

### modules/ec2/main.tf:
```hcl
# SSH Key Pair
resource "aws_key_pair" "deployer" {
  key_name   = "${var.environment}-deployer-key"
  public_key = file(var.public_key_path)
}

# EC2 Instance
resource "aws_instance" "web_server" {
  ami                    = var.ami_id
  instance_type          = var.instance_type
  key_name               = aws_key_pair.deployer.key_name
  vpc_security_group_ids = [var.security_group_id]
  subnet_id              = var.public_subnet_id
  user_data              = var.user_data_script

  root_block_device {
    volume_size           = 20
    volume_type           = "gp3"
    delete_on_termination = true
  }

  tags = {
    Name        = "${var.environment}-web-server"
    Environment = var.environment
  }
}
```

---

## 7. Remote State Setup - Complete Solution

### Project Structure:
```
remote-state-project/
├── backend.tf
├── main.tf
├── variables.tf
├── outputs.tf
└── terraform.tfvars
```

### backend.tf:
```hcl
terraform {
  backend "s3" {
    bucket         = "my-terraform-state-bucket"
    key            = "terraform/terraform.tfstate"
    region         = "us-east-1"
    encrypt        = true
    dynamodb_table = "terraform-locks"
  }
}
```

### main.tf:
```hcl
terraform {
  required_version = ">= 1.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
}

# S3 Bucket for State Storage
resource "aws_s3_bucket" "terraform_state" {
  bucket = var.state_bucket_name

  tags = {
    Name        = "terraform-state-bucket"
    Environment = var.environment
  }
}

# Bucket Versioning
resource "aws_s3_bucket_versioning" "terraform_state" {
  bucket = aws_s3_bucket.terraform_state.id
  versioning_configuration {
    status = "Enabled"
  }
}

# Bucket Encryption
resource "aws_s3_bucket_server_side_encryption_configuration" "terraform_state" {
  bucket = aws_s3_bucket.terraform_state.id

  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm = "AES256"
    }
  }
}

# Block Public Access
resource "aws_s3_bucket_public_access_block" "terraform_state" {
  bucket = aws_s3_bucket.terraform_state.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

# DynamoDB Table for State Locking
resource "aws_dynamodb_table" "terraform_locks" {
  name         = var.dynamodb_table_name
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "LockID"

  attribute {
    name = "LockID"
    type = "S"
  }

  tags = {
    Name        = "terraform-locks"
    Environment = var.environment
  }
}

# Example Infrastructure
resource "aws_vpc" "example" {
  cidr_block = "10.0.0.0/16"

  tags = {
    Name = "example-vpc"
  }
}

resource "aws_subnet" "example" {
  vpc_id     = aws_vpc.example.id
  cidr_block = "10.0.1.0/24"

  tags = {
    Name = "example-subnet"
  }
}
```

### variables.tf:
```hcl
variable "aws_region" {
  description = "AWS region"
  type        = string
  default     = "us-east-1"
}

variable "environment" {
  description = "Environment name"
  type        = string
  default     = "development"
}

variable "state_bucket_name" {
  description = "Name of the S3 bucket for state storage"
  type        = string
}

variable "dynamodb_table_name" {
  description = "Name of the DynamoDB table for state locking"
  type        = string
}
```

### terraform.tfvars:
```hcl
state_bucket_name   = "my-unique-terraform-state-bucket-12345"
dynamodb_table_name = "terraform-locks"
```

---

# 🟠 LEVEL 4: Advanced Projects

## 8. Load Balancer + Auto Scaling - Complete Solution

### Project Structure:
```
alb-asg-project/
├── main.tf
├── variables.tf
├── outputs.tf
├── user_data.sh
└── terraform.tfvars
```

### main.tf:
```hcl
terraform {
  required_version = ">= 1.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
}

# VPC and Networking (simplified version)
resource "aws_vpc" "main" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = {
    Name = "main-vpc"
  }
}

resource "aws_internet_gateway" "main" {
  vpc_id = aws_vpc.main.id

  tags = {
    Name = "main-igw"
  }
}

resource "aws_subnet" "public" {
  count                   = 2
  vpc_id                  = aws_vpc.main.id
  cidr_block              = "10.0.${count.index + 1}.0/24"
  availability_zone       = data.aws_availability_zones.available.names[count.index]
  map_public_ip_on_launch = true

  tags = {
    Name = "public-subnet-${count.index + 1}"
  }
}

resource "aws_route_table" "public" {
  vpc_id = aws_vpc.main.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.main.id
  }

  tags = {
    Name = "public-rt"
  }
}

resource "aws_route_table_association" "public" {
  count          = 2
  subnet_id      = aws_subnet.public[count.index].id
  route_table_id = aws_route_table.public.id
}

# Security Groups
resource "aws_security_group" "alb" {
  name        = "alb-security-group"
  description = "Allow HTTP traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "alb-sg"
  }
}

resource "aws_security_group" "web" {
  name        = "web-security-group"
  description = "Allow HTTP from ALB"
  vpc_id      = aws_vpc.main.id

  ingress {
    from_port       = 80
    to_port         = 80
    protocol        = "tcp"
    security_groups = [aws_security_group.alb.id]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "web-sg"
  }
}

# Application Load Balancer
resource "aws_lb" "web" {
  name               = "web-alb"
  internal           = false
  load_balancer_type = "application"
  security_groups    = [aws_security_group.alb.id]
  subnets            = aws_subnet.public[*].id

  tags = {
    Name = "web-alb"
  }
}

# Target Group
resource "aws_lb_target_group" "web" {
  name     = "web-target-group"
  port     = 80
  protocol = "HTTP"
  vpc_id   = aws_vpc.main.id

  health_check {
    enabled             = true
    healthy_threshold   = 2
    interval            = 30
    matcher             = "200"
    path                = "/health"
    port                = "traffic-port"
    protocol            = "HTTP"
    timeout             = 5
    unhealthy_threshold = 2
  }

  tags = {
    Name = "web-tg"
  }
}

# ALB Listener
resource "aws_lb_listener" "web" {
  load_balancer_arn = aws_lb.web.arn
  port              = "80"
  protocol          = "HTTP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.web.arn
  }
}

# Launch Template
resource "aws_launch_template" "web" {
  name_prefix   = "web-"
  image_id      = var.ami_id
  instance_type = var.instance_type

  key_name = aws_key_pair.deployer.key_name

  vpc_security_group_ids = [aws_security_group.web.id]

  user_data = base64encode(file("user_data.sh"))

  tag_specifications {
    resource_type = "instance"
    tags = {
      Name = "web-instance"
    }
  }

  tags = {
    Name = "web-launch-template"
  }
}

# Auto Scaling Group
resource "aws_autoscaling_group" "web" {
  desired_capacity    = 2
  max_size           = 4
  min_size           = 1
  vpc_zone_identifier = aws_subnet.public[*].id

  target_group_arns = [aws_lb_target_group.web.arn]

  launch_template {
    id      = aws_launch_template.web.id
    version = "$Latest"
  }

  tag {
    key                 = "Name"
    value               = "web-instance"
    propagate_at_launch = true
  }

  depends_on = [aws_lb_listener.web]
}

# Auto Scaling Policies
resource "aws_autoscaling_policy" "scale_up" {
  name                   = "web-scale-up"
  scaling_adjustment     = 1
  adjustment_type        = "ChangeInCapacity"
  cooldown               = 300
  autoscaling_group_name = aws_autoscaling_group.web.name
}

resource "aws_autoscaling_policy" "scale_down" {
  name                   = "web-scale-down"
  scaling_adjustment     = -1
  adjustment_type        = "ChangeInCapacity"
  cooldown               = 300
  autoscaling_group_name = aws_autoscaling_group.web.name
}

# CloudWatch Alarms
resource "aws_cloudwatch_metric_alarm" "cpu_high" {
  alarm_name          = "web-cpu-high"
  comparison_operator = "GreaterThanThreshold"
  evaluation_periods  = "2"
  metric_name         = "CPUUtilization"
  namespace           = "AWS/EC2"
  period              = "120"
  statistic           = "Average"
  threshold           = "70"
  alarm_description   = "This metric monitors ec2 cpu utilization"
  alarm_actions       = [aws_autoscaling_policy.scale_up.arn]

  dimensions = {
    AutoScalingGroupName = aws_autoscaling_group.web.name
  }
}

resource "aws_cloudwatch_metric_alarm" "cpu_low" {
  alarm_name          = "web-cpu-low"
  comparison_operator = "LessThanThreshold"
  evaluation_periods  = "2"
  metric_name         = "CPUUtilization"
  namespace           = "AWS/EC2"
  period              = "120"
  statistic           = "Average"
  threshold           = "20"
  alarm_description   = "This metric monitors ec2 cpu utilization"
  alarm_actions       = [aws_autoscaling_policy.scale_down.arn]

  dimensions = {
    AutoScalingGroupName = aws_autoscaling_group.web.name
  }
}

# SSH Key Pair
resource "aws_key_pair" "deployer" {
  key_name   = "deployer-key"
  public_key = file(var.public_key_path)
}

# Data source for availability zones
data "aws_availability_zones" "available" {}
```

### user_data.sh:
```bash
#!/bin/bash
yum update -y
yum install -y httpd
systemctl start httpd
systemctl enable httpd

# Create a simple HTML page with instance ID
cat > /var/www/html/index.html << EOF
<!DOCTYPE html>
<html>
<head>
    <title>Web Server</title>
</head>
<body>
    <h1>Hello from Auto Scaling Group!</h1>
    <p>Instance ID: $(curl -s http://169.254.169.254/latest/meta-data/instance-id)</p>
    <p>Availability Zone: $(curl -s http://169.254.169.254/latest/meta-data/placement/availability-zone)</p>
    <p>Time: $(date)</p>
</body>
</html>
EOF

# Create health check endpoint
echo "OK" > /var/www/html/health
```

---

## 9. Multi-Environment Setup - Complete Solution

### Project Structure:
```
multi-env-project/
├── environments/
│   ├── dev/
│   │   ├── main.tf
│   │   ├── variables.tf
│   │   └── terraform.tfvars
│   ├── staging/
│   │   ├── main.tf
│   │   ├── variables.tf
│   │   └── terraform.tfvars
│   └── prod/
│       ├── main.tf
│       ├── variables.tf
│       └── terraform.tfvars
├── modules/
│   ├── vpc/
│   │   ├── main.tf
│   │   ├── variables.tf
│   │   └── outputs.tf
│   ├── ec2/
│   │   ├── main.tf
│   │   ├── variables.tf
│   │   └── outputs.tf
│   └── security/
│       ├── main.tf
│       ├── variables.tf
│       └── outputs.tf
└── global.tf
```

### global.tf (Common configuration):
```hcl
terraform {
  required_version = ">= 1.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }

  # Backend configuration (example)
  backend "s3" {
    bucket         = "my-terraform-state-bucket"
    key            = "terraform/terraform.tfstate"
    region         = "us-east-1"
    encrypt        = true
    dynamodb_table = "terraform-locks"
  }
}

# Configure AWS provider
provider "aws" {
  region = var.aws_region
}
```

### environments/dev/main.tf:
```hcl
# Development Environment

# VPC Module
module "vpc" {
  source = "../../modules/vpc"

  environment           = var.environment
  vpc_cidr             = var.vpc_cidr
  public_subnet_cidrs  = var.public_subnet_cidrs
  private_subnet_cidrs = var.private_subnet_cidrs
  availability_zones   = var.availability_zones
}

# Security Module
module "security" {
  source = "../../modules/security"

  environment = var.environment
  vpc_id      = module.vpc.vpc_id
}

# EC2 Module
module "ec2" {
  source = "../../modules/ec2"

  environment           = var.environment
  ami_id               = var.ami_id
  instance_type        = var.instance_type
  public_key_path      = var.public_key_path
  public_subnet_id     = module.vpc.public_subnet_ids[0]
  security_group_id    = module.security.web_security_group_id
  instance_count       = var.instance_count
}
```

### environments/dev/variables.tf:
```hcl
variable "aws_region" {
  description = "AWS region"
  type        = string
  default     = "us-east-1"
}

variable "environment" {
  description = "Environment name"
  type        = string
  default     = "dev"
}

variable "vpc_cidr" {
  description = "CIDR block for VPC"
  type        = string
  default     = "10.10.0.0/16"
}

variable "public_subnet_cidrs" {
  description = "CIDR blocks for public subnets"
  type        = list(string)
  default     = ["10.10.1.0/24", "10.10.2.0/24"]
}

variable "private_subnet_cidrs" {
  description = "CIDR blocks for private subnets"
  type        = list(string)
  default     = ["10.10.10.0/24", "10.10.20.0/24"]
}

variable "availability_zones" {
  description = "Availability zones"
  type        = list(string)
  default     = ["us-east-1a", "us-east-1b"]
}

variable "ami_id" {
  description = "AMI ID for EC2 instance"
  type        = string
  default     = "ami-0c02fb55956c7d316"
}

variable "instance_type" {
  description = "EC2 instance type"
  type        = string
  default     = "t2.micro"
}

variable "public_key_path" {
  description = "Path to SSH public key"
  type        = string
  default     = "~/.ssh/id_rsa.pub"
}

variable "instance_count" {
  description = "Number of EC2 instances"
  type        = number
  default     = 1
}
```

### environments/prod/main.tf:
```hcl
# Production Environment

# VPC Module
module "vpc" {
  source = "../../modules/vpc"

  environment           = var.environment
  vpc_cidr             = var.vpc_cidr
  public_subnet_cidrs  = var.public_subnet_cidrs
  private_subnet_cidrs = var.private_subnet_cidrs
  availability_zones   = var.availability_zones
}

# Security Module
module "security" {
  source = "../../modules/security"

  environment = var.environment
  vpc_id      = module.vpc.vpc_id
}

# EC2 Module
module "ec2" {
  source = "../../modules/ec2"

  environment           = var.environment
  ami_id               = var.ami_id
  instance_type        = var.instance_type
  public_key_path      = var.public_key_path
  public_subnet_id     = module.vpc.public_subnet_ids[0]
  security_group_id    = module.security.web_security_group_id
  instance_count       = var.instance_count
}
```

### environments/prod/variables.tf:
```hcl
variable "aws_region" {
  description = "AWS region"
  type        = string
  default     = "us-east-1"
}

variable "environment" {
  description = "Environment name"
  type        = string
  default     = "prod"
}

variable "vpc_cidr" {
  description = "CIDR block for VPC"
  type        = string
  default     = "10.30.0.0/16"
}

variable "public_subnet_cidrs" {
  description = "CIDR blocks for public subnets"
  type        = list(string)
  default     = ["10.30.1.0/24", "10.30.2.0/24", "10.30.3.0/24"]
}

variable "private_subnet_cidrs" {
  description = "CIDR blocks for private subnets"
  type        = list(string)
  default     = ["10.30.10.0/24", "10.30.20.0/24", "10.30.30.0/24"]
}

variable "availability_zones" {
  description = "Availability zones"
  type        = list(string)
  default     = ["us-east-1a", "us-east-1b", "us-east-1c"]
}

variable "ami_id" {
  description = "AMI ID for EC2 instance"
  type        = string
  default     = "ami-0c02fb55956c7d316"
}

variable "instance_type" {
  description = "EC2 instance type"
  type        = string
  default     = "t3.medium"
}

variable "public_key_path" {
  description = "Path to SSH public key"
  type        = string
  default     = "~/.ssh/id_rsa.pub"
}

variable "instance_count" {
  description = "Number of EC2 instances"
  type        = number
  default     = 3
}
```

### Commands for environment management:
```bash
# Deploy development environment
cd environments/dev
terraform init
terraform plan
terraform apply

# Deploy staging environment
cd ../staging
terraform init
terraform plan
terraform apply

# Deploy production environment
cd ../prod
terraform init
terraform plan
terraform apply
```

---

## 10. CI/CD Integration - Complete Solution

### GitHub Actions Workflow (.github/workflows/terraform.yml):
```yaml
name: 'Terraform'

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main ]

env:
  TF_VERSION: '1.5.7'
  AWS_REGION: 'us-east-1'

jobs:
  terraform:
    name: 'Terraform'
    runs-on: ubuntu-latest
    
    steps:
    - name: Checkout code
      uses: actions/checkout@v3

    - name: Setup Terraform
      uses: hashicorp/setup-terraform@v2
      with:
        terraform_version: ${{ env.TF_VERSION }}

    - name: Configure AWS Credentials
      uses: aws-actions/configure-aws-credentials@v2
      with:
        aws-access-key-id: ${{ secrets.AWS_ACCESS_KEY_ID }}
        aws-secret-access-key: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
        aws-region: ${{ env.AWS_REGION }}

    - name: Terraform Format Check
      run: terraform fmt -check
      working-directory: ./environments/dev

    - name: Terraform Init
      run: terraform init
      working-directory: ./environments/dev

    - name: Terraform Validate
      run: terraform validate
      working-directory: ./environments/dev

    - name: Terraform Plan
      run: terraform plan -out=terraform.tfplan
      working-directory: ./environments/dev
      if: github.event_name == 'pull_request'

    - name: Terraform Plan Summary
      run: terraform show -json terraform.tfplan
      working-directory: ./environments/dev
      if: github.event_name == 'pull_request'

    - name: Terraform Apply
      run: terraform apply -auto-approve
      working-directory: ./environments/dev
      if: github.ref == 'refs/heads/develop' && github.event_name == 'push'

    - name: Terraform Destroy (for cleanup)
      run: terraform destroy -auto-approve
      working-directory: ./environments/dev
      if: github.ref == 'refs/heads/main' && github.event_name == 'push'
```

### Jenkins Pipeline (Jenkinsfile):
```groovy
pipeline {
    agent any
    
    environment {
        TF_VERSION = '1.5.7'
        AWS_REGION = 'us-east-1'
        TF_WORKING_DIR = 'environments/dev'
    }
    
    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
        
        stage('Setup Terraform') {
            steps {
                sh 'wget https://releases.hashicorp.com/terraform/${TF_VERSION}/terraform_${TF_VERSION}_linux_amd64.zip'
                sh 'unzip terraform_${TF_VERSION}_linux_amd64.zip'
                sh 'sudo mv terraform /usr/local/bin/'
                sh 'terraform --version'
            }
        }
        
        stage('Terraform Format') {
            steps {
                sh 'terraform fmt -check'
                dir("${TF_WORKING_DIR}") {
                    sh 'terraform fmt -check'
                }
            }
        }
        
        stage('Terraform Init') {
            steps {
                dir("${TF_WORKING_DIR}") {
                    sh 'terraform init'
                }
            }
        }
        
        stage('Terraform Validate') {
            steps {
                dir("${TF_WORKING_DIR}") {
                    sh 'terraform validate'
                }
            }
        }
        
        stage('Terraform Plan') {
            steps {
                dir("${TF_WORKING_DIR}") {
                    sh 'terraform plan -out=terraform.tfplan'
                }
                archiveArtifacts artifacts: "${TF_WORKING_DIR}/terraform.tfplan", fingerprint: true
            }
        }
        
        stage('Terraform Apply') {
            when {
                branch 'develop'
            }
            steps {
                input message: 'Do you want to apply Terraform changes?', ok: 'Apply'
                dir("${TF_WORKING_DIR}") {
                    sh 'terraform apply -auto-approve terraform.tfplan'
                }
            }
        }
        
        stage('Terraform Destroy') {
            when {
                branch 'main'
            }
            steps {
                input message: 'Do you want to destroy Terraform infrastructure?', ok: 'Destroy'
                dir("${TF_WORKING_DIR}") {
                    sh 'terraform destroy -auto-approve'
                }
            }
        }
    }
    
    post {
        always {
            cleanWs()
        }
        success {
            echo 'Pipeline succeeded!'
        }
        failure {
            echo 'Pipeline failed!'
        }
    }
}
```

---

# 🔴 LEVEL 5: Real-World / Product-Level Projects

## 11. Multi-Region Deployment - Complete Solution

### Project Structure:
```
multi-region-project/
├── main.tf
├── variables.tf
├── outputs.tf
├── provider.tf
├── terraform.tfvars
└── modules/
    ├── vpc/
    ├── ec2/
    └── route53/
```

### provider.tf:
```hcl
terraform {
  required_version = ">= 1.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

# Primary Region Provider
provider "aws" {
  alias  = "primary"
  region = var.primary_region
}

# Secondary Region Provider
provider "aws" {
  alias  = "secondary"
  region = var.secondary_region
}
```

### main.tf:
```hcl
# Primary Region Infrastructure
module "primary_vpc" {
  source = "./modules/vpc"
  
  providers = {
    aws = aws.primary
  }
  
  environment = "${var.environment}-primary"
  vpc_cidr   = var.primary_vpc_cidr
}

module "primary_ec2" {
  source = "./modules/ec2"
  
  providers = {
    aws = aws.primary
  }
  
  environment        = "${var.environment}-primary"
  vpc_id            = module.primary_vpc.vpc_id
  public_subnet_id  = module.primary_vpc.public_subnet_ids[0]
  security_group_id = module.primary_vpc.security_group_id
  ami_id            = var.primary_ami_id
  instance_type     = var.instance_type
}

# Secondary Region Infrastructure
module "secondary_vpc" {
  source = "./modules/vpc"
  
  providers = {
    aws = aws.secondary
  }
  
  environment = "${var.environment}-secondary"
  vpc_cidr   = var.secondary_vpc_cidr
}

module "secondary_ec2" {
  source = "./modules/ec2"
  
  providers = {
    aws = aws.secondary
  }
  
  environment        = "${var.environment}-secondary"
  vpc_id            = module.secondary_vpc.vpc_id
  public_subnet_id  = module.secondary_vpc.public_subnet_ids[0]
  security_group_id = module.secondary_vpc.security_group_id
  ami_id            = var.secondary_ami_id
  instance_type     = var.instance_type
}

# Route53 DNS Failover
module "route53" {
  source = "./modules/route53"
  
  providers = {
    aws = aws.primary
  }
  
  domain_name         = var.domain_name
  primary_instance_id = module.primary_ec2.instance_id
  secondary_instance_id = module.secondary_ec2.instance_id
  primary_region      = var.primary_region
  secondary_region    = var.secondary_region
}
```

### variables.tf:
```hcl
variable "environment" {
  description = "Environment name"
  type        = string
  default     = "production"
}

variable "primary_region" {
  description = "Primary AWS region"
  type        = string
  default     = "us-east-1"
}

variable "secondary_region" {
  description = "Secondary AWS region"
  type        = string
  default     = "us-west-2"
}

variable "primary_vpc_cidr" {
  description = "CIDR block for primary VPC"
  type        = string
  default     = "10.0.0.0/16"
}

variable "secondary_vpc_cidr" {
  description = "CIDR block for secondary VPC"
  type        = string
  default     = "10.1.0.0/16"
}

variable "primary_ami_id" {
  description = "AMI ID for primary region"
  type        = string
  default     = "ami-0c02fb55956c7d316"
}

variable "secondary_ami_id" {
  description = "AMI ID for secondary region"
  type        = string
  default     = "ami-0d1cd67126c60ad51"
}

variable "instance_type" {
  description = "EC2 instance type"
  type        = string
  default     = "t2.micro"
}

variable "domain_name" {
  description = "Domain name for Route53"
  type        = string
}
```

---

## 12. Secure Infrastructure Project - Complete Solution

### main.tf:
```hcl
terraform {
  required_version = ">= 1.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
}

# KMS Key for encryption
resource "aws_kms_key" "main" {
  description             = "KMS key for application encryption"
  deletion_window_in_days = 7
  enable_key_rotation     = true

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid    = "Enable IAM User Permissions"
        Effect = "Allow"
        Principal = {
          AWS = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
        }
        Action   = "kms:*"
        Resource = "*"
      },
      {
        Sid    = "Allow access for Key Administrators"
        Effect = "Allow"
        Principal = {
          AWS = data.aws_iam_role.terraform.arn
        }
        Action = [
          "kms:Create*",
          "kms:Describe*",
          "kms:Enable*",
          "kms:List*",
          "kms:Put*",
          "kms:Update*",
          "kms:Revoke*",
          "kms:Disable*",
          "kms:Get*",
          "kms:Delete*",
          "kms:TagResource",
          "kms:UntagResource",
          "kms:ScheduleKeyDeletion",
          "kms:CancelKeyDeletion"
        ]
        Resource = "*"
      }
    ]
  })

  tags = {
    Name = "application-kms-key"
  }
}

# KMS Key Alias
resource "aws_kms_alias" "main" {
  name          = "alias/application-key"
  target_key_id = aws_kms_key.main.key_id
}

# IAM Role for EC2 instances
resource "aws_iam_role" "ec2_role" {
  name = "${var.environment}-ec2-role"
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "ec2.amazonaws.com"
        }
      }
    ]
  })

  tags = {
    Name = "${var.environment}-ec2-role"
  }
}

# IAM Policy for EC2 role
resource "aws_iam_role_policy" "ec2_policy" {
  name = "${var.environment}-ec2-policy"
  role = aws_iam_role.ec2_role.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "kms:Decrypt",
          "kms:Encrypt",
          "kms:GenerateDataKey*",
          "kms:DescribeKey"
        ]
        Resource = aws_kms_key.main.arn
      },
      {
        Effect = "Allow"
        Action = [
          "ssm:GetParameters",
          "ssm:GetParameter"
        ]
        Resource = "arn:aws:ssm:${var.aws_region}:${data.aws_caller_identity.current.account_id}:parameter/${var.environment}/*"
      }
    ]
  })
}

# Instance Profile
resource "aws_iam_instance_profile" "ec2_profile" {
  name = "${var.environment}-ec2-profile"
  role = aws_iam_role.ec2_role.name
}

# Store secrets in Parameter Store (encrypted)
resource "aws_ssm_parameter" "database_password" {
  name  = "/${var.environment}/database/password"
  type  = "SecureString"
  value = var.database_password
  key_id = aws_kms_key.main.key_id

  tags = {
    Name = "database-password"
  }
}

resource "aws_ssm_parameter" "api_key" {
  name  = "/${var.environment}/api/key"
  type  = "SecureString"
  value = var.api_key
  key_id = aws_kms_key.main.key_id

  tags = {
    Name = "api-key"
  }
}

# VPC with private subnets only
resource "aws_vpc" "main" {
  cidr_block           = var.vpc_cidr
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = {
    Name        = "${var.environment}-vpc"
    Environment = var.environment
  }
}

# Private subnets only
resource "aws_subnet" "private" {
  count             = length(var.private_subnet_cidrs)
  vpc_id            = aws_vpc.main.id
  cidr_block        = var.private_subnet_cidrs[count.index]
  availability_zone = var.availability_zones[count.index]

  tags = {
    Name        = "${var.environment}-private-subnet-${count.index + 1}"
    Environment = var.environment
    Type        = "Private"
  }
}

# NAT Gateway for outbound access
resource "aws_eip" "nat" {
  domain = "vpc"
  tags   = {
    Name = "${var.environment}-nat-eip"
  }
}

resource "aws_nat_gateway" "main" {
  allocation_id = aws_eip.nat.id
  subnet_id     = aws_subnet.private[0].id

  tags = {
    Name = "${var.environment}-nat"
  }
}

# Route table for private subnets
resource "aws_route_table" "private" {
  vpc_id = aws_vpc.main.id

  route {
    cidr_block     = "0.0.0.0/0"
    nat_gateway_id = aws_nat_gateway.main.id
  }

  tags = {
    Name = "${var.environment}-private-rt"
  }
}

resource "aws_route_table_association" "private" {
  count          = length(aws_subnet.private)
  subnet_id      = aws_subnet.private[count.index].id
  route_table_id = aws_route_table.private.id
}

# Security Group with restrictive rules
resource "aws_security_group" "secure" {
  name        = "${var.environment}-secure-sg"
  description = "Restrictive security group"
  vpc_id      = aws_vpc.main.id

  # No inbound rules - instances are not accessible from internet
  # Only outbound traffic to specific services

  egress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    description = "HTTPS outbound"
  }

  egress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    description = "HTTP outbound"
  }

  tags = {
    Name        = "${var.environment}-secure-sg"
    Environment = var.environment
  }
}

# EC2 Instance with secure configuration
resource "aws_instance" "secure_app" {
  ami                    = var.ami_id
  instance_type          = var.instance_type
  subnet_id              = aws_subnet.private[0].id
  vpc_security_group_ids = [aws_security_group.secure.id]
  iam_instance_profile   = aws_iam_instance_profile.ec2_profile.name

  # No public IP - instance is in private subnet
  associate_public_ip_address = false

  # Encrypted root volume
  root_block_device {
    volume_size           = 20
    volume_type           = "gp3"
    encrypted            = true
    kms_key_id           = aws_kms_key.main.arn
    delete_on_termination = true
  }

  # Additional encrypted data volume
  ebs_block_device {
    device_name           = "/dev/sdf"
    volume_size           = 50
    volume_type           = "gp3"
    encrypted            = true
    kms_key_id           = aws_kms_key.main.arn
    delete_on_termination = true
  }

  user_data = base64encode(templatefile("${path.module}/secure_user_data.sh", {
    kms_key_id = aws_kms_key.main.key_id
    env        = var.environment
  }))

  tags = {
    Name        = "${var.environment}-secure-app"
    Environment = var.environment
  }
}

# CloudTrail for auditing
resource "aws_cloudtrail" "main" {
  name                          = "${var.environment}-cloudtrail"
  s3_bucket_name                = aws_s3_bucket.cloudtrail_bucket.bucket
  s3_key_prefix                 = "cloudtrail-logs/"
  include_global_service_events = true
  is_multi_region_trail         = true
  enable_logging                = true

  tags = {
    Name = "${var.environment}-cloudtrail"
  }
}

# S3 bucket for CloudTrail logs (encrypted)
resource "aws_s3_bucket" "cloudtrail_bucket" {
  bucket = "${var.environment}-cloudtrail-logs-${random_id.bucket_suffix.hex}"

  tags = {
    Name = "${var.environment}-cloudtrail-logs"
  }
}

resource "aws_s3_bucket_versioning" "cloudtrail_bucket" {
  bucket = aws_s3_bucket.cloudtrail_bucket.id
  versioning_configuration {
    status = "Enabled"
  }
}

resource "aws_s3_bucket_server_side_encryption_configuration" "cloudtrail_bucket" {
  bucket = aws_s3_bucket.cloudtrail_bucket.id

  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm = "AES256"
    }
  }
}

resource "aws_s3_bucket_public_access_block" "cloudtrail_bucket" {
  bucket = aws_s3_bucket.cloudtrail_bucket.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

# Config for compliance monitoring
resource "aws_config_configuration_recorder" "main" {
  name     = "${var.environment}-config-recorder"
  role_arn = aws_iam_role.config.arn

  recording_group {
    all_supported = true
  }
}

resource "aws_config_delivery_channel" "main" {
  name           = "${var.environment}-config-delivery"
  s3_bucket_name = aws_s3_bucket.config_bucket.bucket
}

# IAM role for AWS Config
resource "aws_iam_role" "config" {
  name = "${var.environment}-config-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "config.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "config_policy" {
  role       = aws_iam_role.config.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSConfigRulesExecutionRole"
}

# S3 bucket for Config
resource "aws_s3_bucket" "config_bucket" {
  bucket = "${var.environment}-config-logs-${random_id.config_bucket_suffix.hex}"

  tags = {
    Name = "${var.environment}-config-logs"
  }
}

# Random ID for unique bucket names
resource "random_id" "bucket_suffix" {
  byte_length = 4
}

resource "random_id" "config_bucket_suffix" {
  byte_length = 4
}

# Data source to get current account ID
data "aws_caller_identity" "current" {}

# Data source for availability zones
data "aws_availability_zones" "available" {}
```

### secure_user_data.sh:
```bash
#!/bin/bash
yum update -y

# Install AWS CLI
curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install

# Install application dependencies
yum install -y python3 python3-pip
pip3 install boto3

# Create application script
cat > /opt/app.py << 'EOF'
#!/usr/bin/env python3
import boto3
import os

def get_secret(parameter_name):
    """Get secret from AWS Parameter Store"""
    ssm = boto3.client('ssm')
    response = ssm.get_parameter(
        Name=parameter_name,
        WithDecryption=True
    )
    return response['Parameter']['Value']

def main():
    env = os.getenv('ENVIRONMENT', 'development')
    
    # Retrieve secrets
    db_password = get_secret(f'/{env}/database/password')
    api_key = get_secret(f'/{env}/api/key')
    
    print(f"Application running in {env} environment")
    print("Successfully retrieved encrypted secrets")
    
    # Your application logic here
    # Use the secrets securely

if __name__ == "__main__":
    main()
EOF

chmod +x /opt/app.py

# Create systemd service
cat > /etc/systemd/system/secure-app.service << EOF
[Unit]
Description=Secure Application
After=network.target

[Service]
Type=simple
User=ec2-user
Environment=ENVIRONMENT=${env}
ExecStart=/usr/bin/python3 /opt/app.py
Restart=always

[Install]
WantedBy=multi-user.target
EOF

systemctl daemon-reload
systemctl enable secure-app
systemctl start secure-app

# Log rotation for application logs
cat > /etc/logrotate.d/secure-app << EOF
/var/log/secure-app.log {
    daily
    rotate 7
    compress
    missingok
    notifempty
    create 644 ec2-user ec2-user
}
EOF
```

---

## 13. Import Existing Infrastructure - Complete Solution

### Steps to Import Existing Infrastructure:

#### 1. Discover Existing Resources

```bash
# List existing EC2 instances
aws ec2 describe-instances --region us-east-1

# List existing VPCs
aws ec2 describe-vpcs --region us-east-1

# List existing security groups
aws ec2 describe-security-groups --region us-east-1
```

#### 2. Create Terraform Configuration

### import-project/main.tf:
```hcl
terraform {
  required_version = ">= 1.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
}

# Existing VPC (to be imported)
resource "aws_vpc" "existing" {
  cidr_block = var.vpc_cidr
  
  tags = {
    Name = "existing-vpc"
    Environment = "imported"
  }
}

# Existing Subnet (to be imported)
resource "aws_subnet" "existing" {
  vpc_id     = aws_vpc.existing.id
  cidr_block = var.subnet_cidr
  availability_zone = var.availability_zone
  
  tags = {
    Name = "existing-subnet"
    Environment = "imported"
  }
}

# Existing Security Group (to be imported)
resource "aws_security_group" "existing" {
  name        = "existing-sg"
  description = "Existing security group"
  vpc_id      = aws_vpc.existing.id

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "existing-sg"
    Environment = "imported"
  }
}

# Existing EC2 Instance (to be imported)
resource "aws_instance" "existing" {
  ami                    = var.ami_id
  instance_type          = var.instance_type
  vpc_security_group_ids = [aws_security_group.existing.id]
  subnet_id              = aws_subnet.existing.id
  
  tags = {
    Name = "existing-instance"
    Environment = "imported"
  }
}
```

### import-project/variables.tf:
```hcl
variable "aws_region" {
  description = "AWS region"
  type        = string
  default     = "us-east-1"
}

variable "vpc_cidr" {
  description = "CIDR block of existing VPC"
  type        = string
}

variable "subnet_cidr" {
  description = "CIDR block of existing subnet"
  type        = string
}

variable "availability_zone" {
  description = "Availability zone of existing subnet"
  type        = string
}

variable "ami_id" {
  description = "AMI ID of existing EC2 instance"
  type        = string
}

variable "instance_type" {
  description = "Instance type of existing EC2"
  type        = string
}
```

### import-project/terraform.tfvars:
```hcl
vpc_cidr        = "172.31.0.0/16"  # Replace with actual VPC CIDR
subnet_cidr     = "172.31.32.0/20"  # Replace with actual subnet CIDR
availability_zone = "us-east-1a"     # Replace with actual AZ
ami_id          = "ami-0c02fb55956c7d316"  # Replace with actual AMI
instance_type   = "t2.micro"         # Replace with actual instance type
```

#### 3. Import Commands

```bash
# Initialize Terraform
terraform init

# Import existing resources (replace IDs with actual resource IDs)
terraform import aws_vpc.existing vpc-12345678
terraform import aws_subnet.existing subnet-12345678
terraform import aws_security_group.existing sg-12345678
terraform import aws_instance.existing i-1234567890abcdef0

# Verify the import
terraform plan

# If everything looks good, apply
terraform apply
```

#### 4. Import Script (for multiple resources)

### import-script.sh:
```bash
#!/bin/bash

# Configuration
REGION="us-east-1"
VPC_ID="vpc-12345678"
SUBNET_ID="subnet-12345678"
SG_ID="sg-12345678"
INSTANCE_ID="i-1234567890abcdef0"

# Get resource details
VPC_CIDR=$(aws ec2 describe-vpcs --vpc-ids $VPC_ID --query 'Vpcs[0].CidrBlock' --output text --region $REGION)
SUBNET_CIDR=$(aws ec2 describe-subnets --subnet-ids $SUBNET_ID --query 'Subnets[0].CidrBlock' --output text --region $REGION)
AZ=$(aws ec2 describe-subnets --subnet-ids $SUBNET_ID --query 'Subnets[0].AvailabilityZone' --output text --region $REGION)
AMI_ID=$(aws ec2 describe-instances --instance-ids $INSTANCE_ID --query 'Reservations[0].Instances[0].ImageId' --output text --region $REGION)
INSTANCE_TYPE=$(aws ec2 describe-instances --instance-ids $INSTANCE_ID --query 'Reservations[0].Instances[0].InstanceType' --output text --region $REGION)

# Create terraform.tfvars
cat > terraform.tfvars << EOF
vpc_cidr        = "$VPC_CIDR"
subnet_cidr     = "$SUBNET_CIDR"
availability_zone = "$AZ"
ami_id          = "$AMI_ID"
instance_type   = "$INSTANCE_TYPE"
EOF

echo "Created terraform.tfvars with resource details"
echo "VPC CIDR: $VPC_CIDR"
echo "Subnet CIDR: $SUBNET_CIDR"
echo "Availability Zone: $AZ"
echo "AMI ID: $AMI_ID"
echo "Instance Type: $INSTANCE_TYPE"

# Initialize Terraform
terraform init

# Import resources
echo "Importing VPC..."
terraform import aws_vpc.existing $VPC_ID

echo "Importing Subnet..."
terraform import aws_subnet.existing $SUBNET_ID

echo "Importing Security Group..."
terraform import aws_security_group.existing $SG_ID

echo "Importing EC2 Instance..."
terraform import aws_instance.existing $INSTANCE_ID

echo "Import completed! Run 'terraform plan' to verify."
```

---

# 🎯 Interview Questions and Answers

## Common Interview Questions with Answers:

### Q1: How do you handle secrets in Terraform?

**Answer:**
- Use AWS Parameter Store/Secrets Manager with KMS encryption
- Use Terraform variables with sensitive flag
- Use environment variables
- Use HashiCorp Vault
- Never hardcode secrets in code

```hcl
resource "aws_ssm_parameter" "database_password" {
  name  = "/${var.environment}/database/password"
  type  = "SecureString"
  value = var.database_password
  key_id = aws_kms_key.main.key_id
}
```

### Q2: What is the difference between `terraform plan` and `terraform apply`?

**Answer:**
- `terraform plan`: Shows what changes will be made without executing them
- `terraform apply`: Actually makes the changes to infrastructure
- Always run plan first to review changes before applying

### Q3: How do you manage state in a team environment?

**Answer:**
- Use remote state storage (S3, Terraform Cloud)
- Enable state locking with DynamoDB
- Use workspaces for environment separation
- Implement proper access controls

### Q4: What are modules and why use them?

**Answer:**
- Modules are reusable Terraform configurations
- Promote code reuse and organization
- Enable consistent infrastructure patterns
- Make configurations more maintainable

### Q5: How do you handle dependencies between resources?

**Answer:**
- Terraform automatically handles dependencies
- Use explicit dependencies with `depends_on`
- Use implicit dependencies through references
- Use `for_each` and `count` for dynamic dependencies

---

# 💡 Pro Tips for Interviews

## Technical Interview Preparation:

1. **Know the Core Concepts**
   - State management
   - Resource lifecycle
   - Providers and modules
   - Workspaces and environments

2. **Practice Common Scenarios**
   - Multi-environment deployments
   - State locking and remote state
   - Infrastructure imports
   - Security best practices

3. **Be Ready for Hands-on**
   - Write Terraform code on whiteboard
   - Debug common errors
   - Explain your architectural decisions

4. **Know the Why**
   - Why use Terraform over other tools
   - Why choose specific configurations
   - Why implement certain security measures

---

This comprehensive answer file provides complete solutions for all Terraform practical projects from beginner to advanced levels, ready for interview preparation!
