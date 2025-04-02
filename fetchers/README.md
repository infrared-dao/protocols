The adapters at the base level of this repo are meant for onchain interaction with the pools/vaults to get TVL and staking token prices per token.

Fetchers in this subfolder implement logic for any additional offchain APIs needed to get protocol data.
The first and main purpose is to fetch protocol "base" APRs per pool/vault specific to each protocol.

The adapters are meant to interact with a single pool/vault but the fetchers can get all at once if this is possible based on the API.
