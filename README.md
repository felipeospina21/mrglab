# MR.G.Lab (mrglab)

mrglab is a TUI to manage `merge requests` in Gitlab from the command line.

## install

// TODO

## config

config file is read from `~/.config/mrglab/mrglab.toml` by default

### config properties

| Option        | Description               | Example                                | Default                |
| ------------- | ------------------------- | -------------------------------------- | ---------------------- |
| `base_url`    | base api url              | `"https://gitlab.com"`                 | `"https://gitlab.com"` |
| `token`       | gitlab access token       | `""`                                   |                        |
| `api_version` | gitlab api version        | `"v4"`                                 |                        |
| `projects`    | list of `project` objects | `[{name="Gitlab Cli", id="34675721"}]` |                        |

`project` is an object with a `name` and `id` properties.

- `name` - `string` is rendered in the project list view
- `id` - `string` is the `gitlab project id` used to fetch selected project `merge requests` & `issues`

// TODO: Add config example

## commands

// TODO
