# Contributing to Client-Side Prometheus Metrics API

We welcome contributions to the Client-Side Prometheus Metrics API project! If you're looking to contribute, please take a moment to read this guide to understand our process and expectations.

## Contribution Process

1. **Fork the Repository**: Start by forking the repository to your own GitHub account.

2. **Clone Your Fork**: Clone your forked repository to your local machine.

3. **Create a Branch**: Create a new branch for your work. Branch names should be descriptive and reflect the nature of the change.

4. **Make Your Changes**: Implement your changes or fixes in your branch. Be sure to keep your branch up to date with the upstream repository to avoid merge conflicts.

5. **Test Your Changes**: Ensure that your changes do not break any existing functionality. Add any necessary tests if applicable.

6. **Commit Your Changes**: Follow the [Conventional Commits](https://www.conventionalcommits.org/) method for your commit messages. This standardized format helps us maintain a clear and consistent commit history.

7. **Push Your Changes**: Push your changes to your forked repository.

8. **Create a Pull Request**: Submit a pull request from your branch to the main repository. Clearly describe the changes you've made and their purpose.

## Conventional Commits

We use the Conventional Commits format for our commit messages. This format provides an easy set of rules for creating an explicit commit history, which makes it easier to write automated tools on top of.

A Conventional Commit message should be structured as follows:

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

### Types

- **feat**: A new feature
- **fix**: A bug fix
- **docs**: Documentation only changes
- **style**: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
- **refactor**: A code change that neither fixes a bug nor adds a feature
- **perf**: A code change that improves performance
- **test**: Adding missing tests or correcting existing tests
- **chore**: Changes to the build process or auxiliary tools and libraries such as documentation generation

### Example

```
feat: add histogram metric support

* Added new histogram metric processing in clientmetrics package
* Updated unit tests for new functionality

```
