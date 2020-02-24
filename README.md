#### Teledger

----

[Ledger-cli](https://www.ledger-cli.org/) is a powerful, double-entry accounting system that is accessed from the UNIX command-line. 
ledger-cli keeps account transactions in a simple text file and it is easy to use, fast and more powerful.

I am using ledger-cli since 2016 and I love it a lot. 

I have been spending money mostly when I am outside. I used to keep costs by writing them to a note-keeping app on a mobile phone. 
Writing expenses from mobile and rewriting them to ledger-cli journal file a little hard and time-consuming work. So I wrote the teledger bot which you can use it to create and report ledger-cli transactions easily.

The final goal is to easily do all the works that can be done with ledger-cli.

#### Future List

- [x] Adding and deleting transactions
- [x] Using aliases
- [x] Balance report
- [ ] Download ledger file
- [ ] Budget reports
- [ ] Execution of custom ledger commands
- [ ] Raw ledger-cli command mode
- [ ] Daily expense reminder


#### Running Teledger

- Create a telegram bot with [BotFather](http://t.me/BotFather) and get token of created bot.
- Set required environment variables
```
TELEDGER_TELEGRAM_TOKEN=
TELEDGER_JOURNAL_DIR=
```

----

### Teledger Usage

#### Show accounts

You can show accounts which used when adding transactions.
Account names reading from ledger journal file.

![](https://user-images.githubusercontent.com/823338/75116846-5e3a4800-567d-11ea-84b7-b22c1a22ed69.jpeg)

#### Adding transaction

You can add transactions easily with simple syntax.

`[Day.Month] fromAccount,toAccount amount [payee]`

![](https://user-images.githubusercontent.com/823338/75116981-9f7f2780-567e-11ea-98c8-faf2cd8c45ae.jpeg)


Adding transaction syntax examples:

```bash
Checkings,food 20.15
Checkings,food 20.15 dinner
20.12 Checkings,food 20.15
20.12 Checkings,food wp.qt (using qwertyuiop keys for amount)
```

#### Deleting transaction

You can delete added transaction using transaction id:

![](https://user-images.githubusercontent.com/823338/75117104-a22e4c80-567f-11ea-88b8-aadcb85d8992.jpeg)

#### Adding aliases

Adding new alias syntax is simple: 

```A AliasName AccountNake```

Alias names adding to journal files also account names will be written to journal file instead alias names when you add a new transaction.   

![](https://user-images.githubusercontent.com/823338/75117181-457f6180-5680-11ea-9207-a287350433df.jpeg)

#### Showing added aliases

![](https://user-images.githubusercontent.com/823338/75117161-f5a09a80-567f-11ea-8141-bebf952b9bdd.jpeg)


#### Showing balance report

You can report account balances with `b` command.

```
B [aliasName|accountName]
```

![](https://user-images.githubusercontent.com/823338/75117510-14ecf700-5683-11ea-971b-40324d2d9c5b.jpeg)


----

#### License

MIT
