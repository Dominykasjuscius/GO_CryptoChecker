# GO_CryptoChecker

This program uses CoinLore API to get the newest data of Crypto currency
and compares them with the rules in the rules.json file. This is done every 30 seconds
If a currency satisfies a rule then the program prints a notification of the change, after
comparison the rule is deleted. If no rules were satisfied then the program prints "No changes in price".
