# MR.G.Lab (mrglab)

mrglab is a TUI to manage `merge requests` in Gitlab from the command line.

## Requirements

- Nerd Font (Symbols) v3.2.1 or higher [ download ](https://github.com/ryanoasis/nerd-fonts/releases/download/v3.2.1/NerdFontsSymbolsOnly.zip)

## Install

```bash
go install github.com/felipeospina21/mrglab@latest
```

## Config

config file is read from `~/.config/mrglab/mrglab.toml` by default.

**To access private repos, you will need to set an env variable with a [gitlab personal access token](https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html). This can be set in your shell config file (to persist it) or in your terminal (for the session).**

```bash
export MRGLAB_TOKEN="YOUR_GITLAB_TOKEN"
```

### config properties

<table>
  <thead>
    <th>Option</th>
    <th>Description</th>
    <th>Default</th>
    <th>Example</th>
  </thead>
  
  <tbody>
    <tr>
      <td>base_url</td>
      <td>base api url</td>
      <td></td>
      <td>https://gitlab.com</td>
    </tr>
    <tr>
      <td>filters.projects</td>
      <td>list of <strong>project</strong> objects</td>
      <td></td>
      <td>[
        {
          name="Gitlab Cli",
          id="34675721", 
          fullPath="gitlab-org/cli"
        }
        ]
      </td>
    </tr>
  </tbody>
</table>

`project` is an object with a `name` and `id` properties.

- `name` - `string` is rendered in the project list view
- `id` - `string` is the `gitlab project id` used to fetch selected project `merge requests` & `issues`
- `fullPath` - `string` is the url path to the project after the base_url

#### config example

```toml
base_url = "https://gitlab.com"

[filters]
projects = [
	{ name = "Gitlab Cli", id = "34675721", fullPath = "gitlab-org/cli" },
]
```

## Commands

```bash
mrglab
```

## Disclaimer

The purpose of this project was to learn more about `go` and `bubbletea`. It is by no mean a full replacement of Gitlab UI (and it is not planned to be), but a complementary tool that would fit in some terminal workflows.

## Inspiration

this project is inspired by tools like [gh-dash](https://github.com/dlvhdr/gh-dash).
