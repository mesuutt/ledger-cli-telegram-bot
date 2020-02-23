#### Teledger

----

[Ledger-cli](https://www.ledger-cli.org/) is a powerful, double-entry accounting system that is accessed from the UNIX command-line. 
ledger-cli keeps account transactions in a simple text file and it is easy to use, fast and more powerful.

I am using ledger-cli since 2016 and I love it a lot. 

I have been spending money mostly when I am outside. I used to keep expenses with writing them to a note keeping app on mobile phone.
Writing expenses from mobile phone and rewriting them to ledger-cli journal file little hard and time consuming work.
So I wrote the teledger bot which you can use it for create and report ledger-cli transactions easily. 

Main goal of teledger is writing ledger-cli transactions with easy special syntax.
and final goal is to easily do all the work can be done with ledger-cli with easily with the bot.


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