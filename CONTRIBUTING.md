# Contributing

To contribute code changes or update the documentation, please follow these steps:

1. Fork the upstream repository and clone the fork locally.
2. Install [pre-commit](https://pre-commit.com/)
3. Navigate toe the cloned repository's directory and install the hooks by running:

```shell
pre-commit install \
   && pre-commit install --hook-type commit-msg \
   && pre-commit install-hooks
```

4. Make whatever changes and additions you wish and commit these - please try to keep your commit history clean.
5. Create a pull request to the main repository with an explanation of your changes.

Note:

- if you add new code or modify existing code - 100% test coverage is mandatory and tests should be well written.
- we follow [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/) and this is enforced using a pre-commit hook.
