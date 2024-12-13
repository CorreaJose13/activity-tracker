# Deployment Instructions

## Prerequisites

To deploy changes to the bot, ensure you have the following:

- **AWS CLI** installed and configured.
- **Terraform** installed.
- **IAM** user credentials provided by the team.

## Steps to Deploy

### 1. Configure Your AWS CLI Profile

- Create a new AWS CLI profile using the command:

```bash
aws configure --profile my-profile
```

Replace <code>my-profile</code> with your preferred profile name.

- Enter the required details:

  - AWS Access Key ID and AWS Secret Access Key (available in the Security Credentials section of your AWS account by creating an access key).
  - Default Region Name: Use <code>us-east-1</code>.

### 2. Set AWS CLI Profile (If Needed)

- To switch to your configured profile, execute (in Linux/macOS):

```bash
export AWS_PROFILE=my-profile
```

### 3. Prepare for Deployment

- Navigate to the lambda terraform directory:

```bash
cd terraform/staging/lambda
```

- Initialize Terraform:

```bash
terraform init
```

### 4. Set Sensitive Values

To protect the project sensitive values (e.g., Telegram bot API token, MongoDB token, etc.) you can choose one of two methods:

a. Using a <code>.tfvars</code> file

- Create a file with extension \*<code>.tfvars</code> in the lambda directory:

```bash
bot_api_token = "BOT-TOKEN"
mongo_token = "MONGO-TOKEN"
gemini_key = "GEM_API_KEY"
```

- Deploy using the command:

```bash
terraform apply -var-file="your_file_name.tfvars"
```

b. Using Environment Variables

- Set environment variables with the command:

```bash
export TF_VAR_bot_api_token=BOT-TOKEN
export TF_VAR_mongo_token=MONGO-TOKEN
export TF_VAR_gemini_key=GEM_API_KEY
```

- Deploy using:

```bash
terraform apply
```

### 5. Revert AWS CLI Profile (If Needed)

If you changed your AWS CLI profile, switch it back to your default settings as needed.

## Contributors

<a href="https://github.com/CorreaJose13/activity-tracker/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=CorreaJose13/activity-tracker" />
</a>
