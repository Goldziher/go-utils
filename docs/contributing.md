# Contributing

To contribute code changes or update the documentation, please follow these steps:

1. Fork the upstream repository and clone the fork locally.
2. Install [prek](https://github.com/RobertCraigie/prek) (a Rust rewrite of pre-commit):

```shell
# Using uv (recommended)
uv tool install prek

# Or using pip
pip install prek
```

3. Navigate to the cloned repository's directory and install the hooks by running:

```shell
prek install && prek install --hook-type commit-msg
```

4. Make whatever changes and additions you wish and commit these - please try to keep your commit history clean.
5. Create a pull request to the main repository with an explanation of your changes.

Note:

- if you add new code or modify existing code - 100% test coverage is mandatory and tests should be well written.
- we follow [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/) and this is enforced using commitlint via a pre-commit hook.
