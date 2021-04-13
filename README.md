# FTX AUTO LENDING

This is a CLI tool that allows to automatically compound payouts earned from lending coins.
It will check for newly available funds every hour (5 min after hour elapsed to be precise) and automatically update the lending offer to the max. size that can be lend out on the account.

**Warning:**  
Do note that any coin lent out is removed from your collateral pool, so have a lending only subaccount... Otherwise its very easy to find yourself overleveraged on drawdowns even if margin levels previously looked healthy.

## Example
This will compound lending for USD and ETH once per hour (always xx:05:00). You can add more coins by repeating the `--coin [coin name]` parameter.

ftx-auto-lend-win.exe --key xxxxxxx --secret yyyyyyyy --subaccount mylendingsubacc --coin USD --coin ETH

