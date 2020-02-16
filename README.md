#### Teledger

----

[Ledger-cli](https://www.ledger-cli.org/) is a powerful, double-entry accounting system that is accessed from the UNIX command. 
ledger-cli keeps account transactions in a simple text file and it is easy to use, fast and more powerful.

I am using ledger-cli since 2016 and I love it a lot. 
I am spending money mostly when I am outside and writing ledger-cli transactions to a note app and moving the transactions
to ledger-cli journal file. This is a little hard and time consuming work.

Main goal of teledger is writing ledger-cli transactions easily from mobile phone.
and final goal is to easily do all the work can be done from command-line on the mobile phone with easy special syntax.

#### Future List

- [x] Adding and deleting transactions
- [x] Using aliases
- [x] Balance report
- [ ] Download ledger file
- [ ] Budget reports
- [ ] Execution of custom ledger commands
- [ ] Raw ledger-cli command mode
- [ ] Daily reminder




#### Running Teledger

- Create a telegram bot with [BotFather](http://t.me/BotFather) and get token of created bot.
- Set required environment variables
```
TELEDGER_TELEGRAM_TOKEN=
TELEDGER_JOURNAL_DIR=
```
- Send messages to the bot.