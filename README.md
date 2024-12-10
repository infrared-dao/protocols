# Protocols Adapters Repo
This repo implements adapters to different berachain protocols needed to compute LP token prices / TVL for vaults of interest to the infrared project.

See https://github.com/infrared-dao/infrared-default-list for a complete list.


## Inclusion as an Infrared Partner

In order for a partner project to be shown in the infrared frontend and attract the economic energy of infrared users, there are two needed steps:
1. Add token, protocol, and gauge information to our default list
2. Implement an adapter to compute the staking token price

All related tokens (tokens in the pool and reward tokens) must first be added to the default token lists, and the protocol and gauge information must be added to the default gauge list (see [default list repo](https://github.com/infrared-dao/infrared-default-list)). Do not add the staking token to these lists unless it is a token commonly traded on DEXes. 

Anyone can open a PR to add this data for their protocol, but it is advised to be in contact with the infrared business development team in advance for the PR to be merged quickly.

Additionally infrared needs a way to track the staking token price to estimate infrared vault TVL and APR.  

If the staking token of the gauge is a token we already track in the default token list, *then you do not need to implement any custom adapter in this repo*.  For example, BERPS bHONEY and Infrared iBGT vaults have staking tokens bHONEY and iBGT respectively which are in the tokens list because they commonly trade on DEXes

However, the more common case is that the staking token for the gauge is an LP "claim" token which is not traded commonly on any DEX and is only used to track positions in your protocol.  In this case you will also need to create an adapter which computes the current staking token price before infrared will show your vault information to our users.


## Implementing an Adapter

Please follow these best practices when implementing an adapter for your protocol.  

When you have implemented your adapter, please write a simple test script which demonstrates the Config(), LPTokenPrice(), and TVL() functions in a file called main.go at /cmd/test/protoName/main.go


##### Golang Types and Stateful Structs

Use the golang `big.Int` type to represent token amounts, use the `decimal.Decimal` type to represent prices and rates.  Do not hardcode any contract addresses but allow them to be passed in to a function which creates a new stateful data struct to store them.  

If your protocol has multiple pools/vaults with infrared, then we will initialize and run separate instances of the adapter for each pool/vault with different initialization data.  Please write your adapter and data structs to represent a single pool/vault in your protocol with all differentiation passed as inputs.


##### New Struct and Initialize Functions

Create a function which returns a new data structure for any stateful information, but make sure that this 'New' function does not contain any code which has the potential to throw an error.  Examples would be the `NewKodiakLPPriceProvider` or `NewBexLPPriceProvider`.  

You can expect that this 'New' function will be given both a price map for all underlying tokens and the config data as a json encoded string in a byte array (discussed below).  The price map will link all underlying token addresses as string keys to their corresponding Price struct, which contains token name, the number of decimals in amount representations of the token, and a decimal.Decimal price value in USD.

In the Initialize() function on this type, you should perform any setup which has the potential to actually throw errors, and make sure that all potential errors are correctly handled and returned.  Common examples are errors could be thrown when unmarshalling json config data, if needed token prices were not supplied, or if there are problems with the RPC node when initializing smart contract connections, etc.

##### Required Adapter Functions

The functions which you must implement for the adapter to be used in the infrared backend system are the LPTokenPrice() and TVL() functions which should return either a string representing the USD amount, or an error which happened during execution.  The LPTokenPrice() should return the price of a single staking token (if it has 18 decimals then the price of 10^18 sub units) in USD.  The TVL() should return the total amount in USD represented by the entire pool/vault in your protocol. 

It is best practice to internally work with all price information in these functions as decimal.Decimal types until the final serialization to a string before output.


##### Config Function and Data

The Protocol interface exposes a Config function which is meant to collect any data which is unlikely to change between calls but would be needed input to all onchain function calls in the process of computing the LP token price.  This will be custom to each adapter dependent on what data is needed to uniquely describe the pool/vault.  

For example, with Kodiak, the order of the tokens in a pool matters, so the Config function performs an onchain query to the Kodiak vault to find which underlying token is token0 and which is token1.  Another example is with BEX, where which token is the quote token and which is the base token is important information to store, as well as the IDX or pool type for the given pool.  

If your specific protocol does not require any such config data to later compute the LP Token price or TVL then you can return an empty byte array or simply "" since it will not need to be parsed back into json.

This config information will only be generated once by the system by calling to Config() and storing its output, and then the output will be present to all future calls.  It is recommended that you make a custom type to store your config data and annotate it with appropriate json serialization directives eg. the `KodiakConfig` and `BexConfig` type structs.
