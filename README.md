# cloudbuild-actions-test

Test with Cloud Build and GitHub Actions.

## Install

Download the binary for your platform:

```bash
VERSION=$(curl -s https://api.github.com/repos/halvards/cloudbuild-actions-test/releases/latest | jq -r '.tag_name')

curl -sLo cloudbuild-actions-test "https://github.com/halvards/cloudbuild-actions-test/releases/download/$VERSION/cloudbuild-actions-test_$(uname -s)_$(uname -m)"

chmod +x cloudbuild-actions-test
```

## Run

Display the version number:

```bash
./cloudbuild-actions-test
```
