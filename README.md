## vercelgate

vercelgate is a command-line tool designed to streamline the process of managing multiple Vercel client accounts. It eliminates the need for repetitive login and logout actions, allowing users to switch between accounts and teams seamlessly.

## WHY

As developers, we often encounter situations where clients provide access to their Vercel accounts without subscribing to a Team plan. The Vercel CLI restricts usage to one account at a time, requiring you to log out from one account before logging into another. This can be a significant inconvenience, especially since the Vercel CLI does not natively support multi-account management, likely encouraging users to opt for a paid Team plan.

vercelgate offers a practical solution by enabling users to switch between multiple personal Vercel hobby plan accounts without the need to upgrade to a Vercel Pro Team plan.

## Installation

```go
go install github.com/khanakia/vercelgate@main
```

### Install with Homebrew:

```sh
brew tap khanakia/vercelgate
brew install vercelgate
vercelgate --help
vercelgate init
```

## Usage

To begin using **vercelgate**, first perform the initialization:

```bash
vercelgate init
```

This only needs to be done once. For further guidance on how to use **vercelgate** efficiently, use the help command:

```bash
vercelgate --help
```

This will display a list of available commands and options to tailor your account management experience.

## Commands

### Version

```bash
vercelgate --version
```

Displays the current version of vercelgate installed on your system.

### Initialize

```bash
vercelgate init
```

Initializes vercelgate for first-time use. This command sets up the necessary configuration and database.

## Ubuntu

For ubuntu vercel configuration is stored in `~/.local/share/com.vercel.cli/` instead of `~/.config/com.vercel.cli/`

Make sure you create a Symlink: `mkdir -p ~/.config/com.vercel.cli && ln -s ~/.local/share/com.vercel.cli/* ~/.config/com.vercel.cli/`

### Sync

```bash
vercelgate sync
```

Syncs your current logged-in Vercel account with vercelgate.

### New Account

```bash
vercelgate new
```

Clears the current Vercel account configuration. After running this, you can add a new Vercel account using `vercel login` and then run `vercelgate sync` to add it to vercelgate.

### Switch Account

```bash
vercelgate switch
```

Displays a list of all synced Vercel accounts, allowing you to choose which account to set as active.

### Switch Team

```bash
vercelgate switchteam
```

Shows a list of all synced Vercel accounts and their teams, allowing you to select a team to set as the current team.

### Show Config Path

```bash
vercelgate path
```

Displays the Vercel global configuration path where settings and authentication data are stored.

### Reset

```bash
vercelgate reset
```

Resets all synced Vercel accounts from vercelgate.

## Example

```sh
➜ ✗ vercelgate switch
Use the arrow keys to navigate: ↓ ↑ → ←
? Select Account:
  ▸ Jane Doe (jandoe@gmail.com)

Switched to user Jane Doe


➜ ✗ vercelgate new
you can now add new account using `vercel login` and then run `vercelgate sync` again


➜ ✗ vercel login
Vercel CLI 34.0.0
? Log in to Vercel Continue with Email
? Enter your email address: dummy@gmail.com
We sent an email to dummy@gmail.com. Please follow the steps provided inside it and make sure the security code matches Eager Bornean Orang-utan.
> Success! Email authentication complete for dummy@gmail.com


➜ ✗ vercelgate sync
synced successfully

➜ ✗ vercelgate switch
Use the arrow keys to navigate: ↓ ↑ → ←
? Select Account:
  ▸ Jane Doe (jandoe@gmail.com)
    dummy (dummy@gmail.com)

```

## Features

- **Simple Account Switching**: Quickly switch between different Vercel accounts without logging out
- **Support for Multiple Accounts**: Manage multiple personal and hobby plan accounts
- **Team Management**: Switch between teams within accounts
- **Configuration Access**: View global configuration paths for troubleshooting
- **Command-Line Interface**: Easy-to-use CLI for all operations

## TODO

- [ ] Auto detect Ubuntu config file path
