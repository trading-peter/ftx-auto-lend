# FTX AUTO LENDING

This is a CLI tool that allows to automatically compound payouts earned from lending coins.
It will check for newly available funds every hour (5 min after hour elapsed to be precise) and automatically update the lending offer to the max. size that can be lend out on the account.

## Example
ftx-auto-lend --key xxxxxxx --secret yyyyyyyy --subaccount mylendingsubacc --coin USD

