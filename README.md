# `vimbin`

`vimbin` is a powerful pastebin tool that seamlessly integrates the efficiency of Vim motions with the simplicity of a pastebin service. It offers two main commands:

- **serve**: Start a local web server featuring a textarea for creating, managing, and refining content. All changes made in the textarea are persistently stored to a file, and users can navigate and manipulate text using familiar Vim motions for an enhanced editing experience.

- **push**: Quickly send text to the `vimbin` server from the command line. This allows for easy integration with other tools and scripts, streamlining the process of sharing content through `vimbin`.

## Themes

For the editor you can chooe between the catppuchin themes `latte`, `mocha`, `frappe` or `macchiato`. If theme is set to `auto`, it gets system preference.

Themes are borrowed from [here](https://github.com/catppuccin/codemirror).

## Commands

### Global Flags

| Flag                    | Description                                                                                                                                 |
| :---------------------- | :------------------------------------------------------------------------------------------------------------------------------------------ |
| `-c`, `--config` `PATH` | Path to the configuration file.                                                                                                             |
| `--debug`               | Activates debug output for detailed logging. Can also be set with the environment variable `VIMBIN_DEBUG`                                   |
| `-t`, `--token` `TOKEN` | Token to use for authentication. If not set, a random token will be generated. Can also be set with the environment variable `VIMBIN_TOKEN` |
| `--trace`               | Enables trace mode. This will show the content in the logs! Can also be set with the environment variable `VIMBIN_TRACE`                    |
| `-v`, `--version`       | Print version and exit.                                                                                                                     |

### Serve

Start the server:

```bash
./vimbin serve
```

**Flags:**

| Flag                                    | Description                                                                                                                                                                      |
| :-------------------------------------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `-d`, `--directory` `DIRECTORY`         | The path to the storage directory. (default `$(pwd)`)                                                                                                                            |
| `-a`, `--listen-address` `ADDRESS:PORT` | The address to listen on for HTTP requests. (default `:8080`)                                                                                                                    |
| `-n`, `--name` string                   | The name of the file to save. (default ".vimbin")                                                                                                                                |
| `--theme` THEME                         | The theme to use. Can be `auto`, `light` or `dark`. (default `auto`). Can also be set with the environment variable `VIMBIN_THEME`                                               |
| `--dark-theme` THEME                    | When `theme` set to `auto`, use this as dark theme. Can be `mocha`, `frappe`, `macchiato`. (default `frappe`). Can also be set with the environment variable `VIMBIN_DARK_THEME` |
| `-h`, `--help`                          | help for serve                                                                                                                                                                   |

### Push

Push data to the `vimbin` server:

```bash
./vimbin push [text]
```

**Flags:**

| Flag                           | Description                            |
| :----------------------------- | :------------------------------------- |
| `-a`, `--append`               | Append content to the existing content |
| `-i`, `--insecure-skip-verify` | Skip TLS certificate verification      |
| `-u`, `--url` `URL`            | The URL of the vimbin server           |
| `-h`, `--help`                 | help for push                          |

### Pull

Pull the latest data from the `vimbin` server:

```bash
./vimbin pull
```

**Flags:**

| Flag                           | Description                       |
| :----------------------------- | :-------------------------------- |
| `-i`, `--insecure-skip-verify` | Skip TLS certificate verification |
| `-u`, `--url` `URL`            | The URL of the vimbin server      |
| `-h`, `--help`                 | help for fetch                    |

## Configuration

`vimbin` can be configured using a YAML configuration file. By default, it looks for a file named `.vimbin.yaml` where the `vimbin` binary is started.

Example configuration:

```yaml
server:
  web:
    address: ":8080"
    theme: auto
  api:
    address: "http://vimbin.example.com"
    token: secure token

storage:
  name: .vimbin
  directory: $(pwd)
```
