## Creating a Project

After installing the forgex CLI tool, you can create a new project with the default settings by running the following command:

```bash
forgex create
```

This command will interactively guide you through the project setup process, allowing you to choose the project name, framework, and database driver.

![ForgeXInteractive](../public/ForgeX_1.png)

## Using Flags for Non-Interactive Setup

For a non-interactive setup, you can use flags to provide the necessary information during project creation. Here's an example:

```
forgex create --name my-project --framework gin --driver postgres --git commit
```

In this example:

- `--name`: Specifies the name of the project (replace "my-project" with your desired project name).
- `--framework`: Specifies the Go framework to be used (e.g., "gin").
- `--driver`: Specifies the database driver to be integrated (e.g., "postgres").
- `--git`: Specifies the git configuration option of the project (e.g., "commit").

Customize the flags according to your project requirements.

## Advanced Flag

By including the `--advanced` flag, users can choose one or all of the advanced features, HTMX, GitHub Actions for CI/CD, Websocket, Docker and TailwindCSS support, during the project creation process. The flag enhances the simplicity of ForgeX while offering flexibility for users who require additional functionality.

```bash
forgex create --advanced
```

To recreate the project using the same configuration semi-interactively, use the following command:
```bash
forgex create --name my-project --framework chi --driver mysql --git commit --advanced
```
This approach opens interactive mode only for advanced features, which allow you to choose the one or combination of available features.

![AdvancedFlag](../public/ForgeX_advanced.png)

## Non-Interactive Setup

Advanced features can be enabled using the `--feature` flag along with the `--advanced` flag:

HTMX:
```bash
forgex create --advanced --feature htmx
```

CI/CD workflow:
```bash
forgex create --advanced --feature githubaction
```

Websocket:
```bash
forgex create --advanced --feature websocket
```
TailwindCSS:
```bash
forgex create --advanced --feature tailwind
```
Docker:
```bash
forgex create --advanced --feature docker
```

Or all features at once:
```bash
forgex create --name my-project --framework chi --driver mysql --git commit --advanced --feature htmx --feature githubaction --feature websocket --feature tailwind --feature docker
```
