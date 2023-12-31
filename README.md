# SOL Shotty

![](./docs/sol-shotty.png)

SOL Shotty is a Solana RPC proxy with a little something "extra".

When running sol-shotty, it will take any RPC request and "shotgun blast" them at every RPC provider configured, only returning the result from the first to respond.

## Why would I want this?

When you need to ensure your transaction goes through *no matter what*, you can send the same request to many different providers and let the validators sort out which one got there first.

All RPC providers have bad days, sol-shotty helps get faster responses when a single RPC provider is degraded, or even fully down. No reconfiguration required!

# Quick Start

Edit the `config.yaml` file to include the endpoints you wish to use. Some providers use Cloudflare to provide anycast routing when they only have infrastructure in a few locations.

A lot of providers allow you to sign up with just a web3 wallet, if you have multiple you might be able to sign up multiple times. 

Make sure you have `golang` installed and then run `go run ./cmd/sol-shotty`

sol-shotty will now be listening on `http://127.0.0.1:420` and can be used anywhere an RPC URL is accepted.

# Development

This tool was a quick rip of some internal tooling we've built to increase reliability in [solan.ai](https://solan.ai) RPC requests.

While we will update it from time to time when we can, we *highly encourage* PR's if you see a bug or would like additional functionality that isn't here.

If you'd like a feature that isn't here, donations greatly encourage us to prioritize those requests.

Donations are accepted at `23Q5e33JnmKWACqmmYW1owRwfs7ToD4SuCTzfDGxMcfA`.