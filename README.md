# Blockchain Learningtool by Mocacinno

## tipjar

<a href="https://coindrop.to/mocacinno" target="_blank"><img src="https://coindrop.to/embed-button.png" style="border-radius: 10px; height: 57px !important;width: 229px !important;" alt="Coindrop.to me"></img></a>

i basically accept any kind of tips, eventough i only posted a bitcoin address... If you want to support this project but want to tip some altcoin, or using a L2 sollution: just contact me on bitcointalk, or open an issue on this repo

## intro

This tool was written to give some ELI12 insight in what a blockchain is, what a transaction is, what a wallet is,...
The tool writes a csv "blockchain" that gives some insight in the basic principles applied in blockchain technology.

## warning

It's an ELI12 tool (explain it like i'm 12), it's not meanth to create a "real" blockchain. Many concepts have been omitted (did somebody say merkle tree, scripting, mempool, fees, coinbase reward,...). Understanding a "blockchain" created by this tool will not give you a full understanding of the bitcoin blockchain, merely a basic insight!!!

## technical

block block0001.csv (the genesis block) only creates one transaction out of thin air (in the bitcoin blockchain, you can compare this to the coinbase transaction, but instead of adding a coinbase transaction in every block, it's only added once). From here on out, this one unspent transaction will be used as an input for subsequent transactions in the following blocks

### block layout

line 0: the block header in format "<sha256 hash of previous block>,<block number>"  
line 1-unlimited: transactions in format "<keywords INPUTS>,<blocknumber of input transaction>,<line inside blocknumber of incoming transaction>,<input value><repeat two previous fields if necessary>,<keywords SENDER>,<name of the sender>,<keywords OUTPUTS><transmitted value>,<name of the receiver>,<public key of the receiver>,<repeat previous 3 fields if necessary>,<keyword SIGNATURE>,<pgp signature of tx minus the signature field itself, using private key of sender>"

Just like in a "normal" blockchain, many inputs and outputs are possible in one transaction, however, only Only one private key can be used to sign the transaction, so all spent unspent outputs had to be funding the same public key.

## verification

### verify the headers

We have a blockchain of (default) 10 blocks... Now what's stopping me from editing an existing block??? It's the header (Line 0, the first line in each block). We opted for a learning-header that's grossly simplified... But it should work nonetheless...

The very first block is the genesis block and follows different rules, but starting from the second block, the first line starts with a sha256 hash of the previous block... Let's examine my own block 4, if i open it, i see the header contains hash "f6ee2792e5075bd1fe127fbcad82a977fa53ea0b6c037b39d782b84ad5e730a6". This means that if i execute `sha256sum block0003.csv`, the output HAS to be `f6ee2792e5075bd1fe127fbcad82a977fa53ea0b6c037b39d782b84ad5e730a6` (this is the case). If you modify block 3 in *any* way (even by adding or removing a space, blank line or converting lower/upper cases, the hash will no longer match, and the chain becomes invalid... This way, you know nobody can tamper with the chain.

### verify the transaction chain

if you open the last block (the one with the highest number), you should be able to pick *any* line (except the very first one, containing the block header). If you read this line, you should pay attention to the inputs (everything between the keyword INPUTS and SENDER). The inputs are in groups of 3: the block number, the line number and the value. If, for example, the beginning of your transaction line looks like this: INPUTS,7,3,163,SENDER,david,OUTPUTS, you know that in block 7, line 4, there should be a transaction funding david's public key with 163 coins. Why line 4? Because my script starts counting lines at 0 :). Line 0 = the first line, Line 1 = the second line, Line 2 = the thirth line and Line 3 = the fourth line :)

You should be able to track back any transaction all the way back to the initial block. No transaction stands on it's own... They all use a pre-existing input as input and create new outputs that can be used in the same or the next blocks! 

This way, nobody can make coins out of "thin air". You cannot just add a transaction funding yourself to the chain... The transaction has to come from somewhere... And the person who owns the inputs has to sign the transaction to make it valid (so you can't spend somebody else's coins... See the next chapter for more info... This way, nobody can tamper with the transaction chain to generate money out of thin air.

### verify the transactions

if you use an unspent output from a transaction belonging to me, you need my private key to sign said transaction, otherwise it's invalid. For this learning tool, we know all private keys from all users, but in a "real" blockchain, every user would only know his or her own private key. The public keys from each user can be found in the output/keys folder, but you could theoretically also get them from the transactions in the blocks...

* open a block, pick a line, look at the sender (in my case, it was james)
* go to the output/keys folder, create two new files: signature.txt and message.txt
* the signature is the part AFTER the last comma of the line picked in the first verification step... copy it, this is your signature
* the message is is the part BEFORE the last comma of the line picked in the first verification step... copy it, this is your message
* convert the public key of james: `openssl rsa -RSAPublicKey_in -in james_public.pem -pubout -outform PEM -out james_public_spki.pem`
* use a tool like [this](https://kjur.github.io/jsrsasign/sample/sample-rsasign.html). Use the right side of the page. In the "verify signature" field, paste the signature, in the "Text message to be verified" field, paste the message and in the "Signer's Public Key Certificate" field, paste the content of the james_public_spki.pem file. Click on "verify this signature" button, and the text "This signature is *VALID*" should appear...
* you can also do the opposite and use the left side of the verification website... Paste james's private key (the content of james_private.pem), leave the sha256d dropdown selection and click on "sign to this message" button. The signature that appears has to match the signature you have copied...

This way, nobody can tamper with the transactions by stealing somebody else's funds...




