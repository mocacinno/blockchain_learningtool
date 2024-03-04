# Blockchain Learningtool by Mocacinno

## intro

This tool was written to give some ELI12 insight in what a blockchain is, what a transaction is, what a wallet is,...
The tool writes a csv "blockchain" that gives some insight in the basic principles applied in blockchain technology

## warning

It's an ELI12 tool (explain it like i'm 12), it's not meanth to create a "real" blockchain. Many concepts have been omitted (did somebody say merkle tree, scripting, mempool, fees, coinbase reward,...). Understanding a "blockchain" created by this tool will not give you a full understanding of the bitcoin blockchain, merely a basic insight!!!

## technical

block block0001.csv only creates one transaction out of thin air (in the bitcoin blockchain, you can compare this to the coinbase transaction, but instead of adding a coinbase transaction in every block, it's only added once). From here on out, this one unspent transaction will be used as an input for subsequent transactions in the following blocks

### block layout

line 0: the block header in format "<sha256 hash of previous block>,<block number>"
line 1-unlimited: transactions in format "<keywords INPUTS>,<blocknumber of input transaction>,<line inside blocknumber of incoming transaction>,<input value><repeat two previous fields if necessary>,<keywords SENDER>,<name of the sender>,<keywords OUTPUTS><transmitted value>,<name of the receiver>,<public key of the receiver>,<repeat previous 3 fields if necessary>,<keyword SIGNATURE>,<pgp signature of tx minus the signature field itself, using private key of sender>"

Just like in a "normal" blockchain, many inputs and outputs are possible in one transaction, however, only Only one private key can be used to sign the transaction, so all spent unspent outputs had to be funding the same public key.

