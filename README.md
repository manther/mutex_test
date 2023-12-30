# mutex_test
 This is a test to solve a concurrency race condition using a mutex. 

 The test uses a pretty common banking scenario. There is one bank, several merchants and several customer accounts held by both the bank and the merchants. 

Each customer and their starting accound balances: 
```
"Dave": 400, "Susan": 1200, "Mike": 1000, "Steve": 300, "James": -32
```

The bank is just a hashmap. 
```
map[string]int
```

There should be around 5 merchants. 
Every merchant should have a data structure of customers accounts that get billed in arrears. 

The structure containing the customer information just needs the customer id (name), and the amount they are to be billed this month. (It's ok if this amount is the exact same for every merchant)

Define an interface to describe what every merchant should be able to do. (They shouod be able to run a biling cycle againsted their customer accounts).

Setup a test that loops over merchants and calls the function to run billing charges for each merchant serially. The final balance in the bank for each customer is your correct answer. Hard code a data structure to hold the "correct" ending bank balance.

Now setup a test that is very similar except call each merchant in a seperate go routing. Create a wait group and wait for all to finish. Check the final answer against the correct answer you hardcoded earlier. They may not match. If they do match we haven't created the failure case yet. Make it so that each merchant reads each customers balance, sleeps for a bit, then sets the customers balance to what you just read plus the amount you plan to bill them this month. This *should* cause some dirty reads and incorrect results. 
This will represent your fail test case.

Now setup another test that does all the same, but make a new function for your merchants that uses a mutext during the process of billing each customer. (It may be tempting to put the mutext at the bank level, but then your fail case would not work correctly).

*Note there are many ways to solve this problem. This is setup to solve with a mutex in a certain situation. 
 
