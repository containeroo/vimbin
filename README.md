# `vimbin`

`vimbin` is a powerful pastebin tool that seamlessly integrates the efficiency of Vim motions with the simplicity of a pastebin service. It offers two main commands:

- **serve**: Start a local web server featuring a textarea for creating, managing, and refining content. All changes made in the textarea are persistently stored to a file, and users can navigate and manipulate text using familiar Vim motions for an enhanced editing experience.

- **push**: Quickly send text to the `vimbin` server from the command line. This allows for easy integration with other tools and scripts, streamlining the process of sharing content through `vimbin`.

## Commands

### Serve

Start the server:

```bash
./vimbin serve
```

#### Options

`--listen-address`, `-a`: The address to listen on for HTTP requests. (Default: `:8080`)
`--theme`, `-t`: The theme to use. Can be auto, light, or dark. (Default: `auto`)
`--directory`, `-d`: The path to the storage directory. Defaults to the current working directory.
`--name`, `-n`: The name of the file to save. (Default: `.vimbin`)

### Push

Push data to the `vimbin` server:

```bash
./vimbin push [text]
```

#### Options

`--url`, `-u`: The URL of the `vimbin` server.
`--append`, `-a`: Append to the existing file on the `vimbin` server.

### Fetch

Fetch the latest data from the `vimbin` server:

```bash
./vimbin fetch
```

#### Options

`--url`, `-u`: The URL of the `vimbin` server.

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

storage:
  name: .vimbin
  directory: $(pwd)
```
