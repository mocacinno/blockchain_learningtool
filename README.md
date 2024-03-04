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
line 1-unlimited: transactions in format "<blocknumber of input transaction>,<line inside blocknumber of incoming transaction>,<name of the sender>,<transmitted value>,<name of the receiver>,<public key of the receiver>,<pgp signature of tx minux the signature itself, using private key of sender>"

Just like in a "normal" blockchain, one unspent output needs to be fully spent, one unspent output CAN be spent in as many lines as you want. The difference is that this "learning tool" will show these splits as multiple transactions, the concept of vout is not implemented since it would make it to hard for new users to grasp what's going on

