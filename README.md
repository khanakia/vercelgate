## vercelgate
vercelgate is a command-line tool designed to streamline the process of managing multiple Vercel client accounts. It eliminates the need for repetitive login and logout actions, allowing users to switch between accounts and teams seamlessly.


## WHY
As developers, we often encounter situations where clients provide access to their Vercel accounts without subscribing to a Team plan. The Vercel CLI restricts usage to one account at a time, necessitating a logout from one account before logging into another. This can be a significant inconvenience, especially since the Vercel CLI does not natively support multi-account management, likely encouraging users to opt for a paid Team plan.

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


```bash
vercelgate sync
```
Sync current logged in vercel account with vercelgate

```bash
vercelgate new
```
Once synced this command will empty the current vercel account and now you add new vercel account `vercel login` and then `vercel sycn` again now the previously logged in account and the new one both are sycned with vercelgate



```bash
vercelgate switch
```
It will show you list of all the synced vercel accounts and you can choose from to set any of account as actie account


```bash
vercelgate switchteam
```
It will show you list of all the synced vercel accounts with their teams and you can choose from any team to set the selected team as current team

```bash
vercelgate reset
```
It will reset all the synced vercel accounts with vercelgate.


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

- **Simple Account Switching**: Quickly switch between different Vercel accounts without logging out.
- **Support for Multiple Accounts**: Manage multiple personal and hobby plan accounts.
- **Command-Line Interface**: Easy-to-use CLI for all interactions.
