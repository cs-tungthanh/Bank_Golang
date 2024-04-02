# User story

## UC 1:
[As a] bank customer  
[I want to] create account by providing my personal information specifying a currency  
[so that] I can start using this app and manage my finances in that currency.  

### AC:
[Given] user provide a request with currency, account name, password  
[When] user submit request  
[Then] 
- the system validates the currency against a supported list  
- the system validates the account is exist or not  
- a new record should be created in the database  

## US 2:
[As a] bank customer  
[I want to] retrieve a record of all transactions for my account  
[so that] I can track my spending and income  

### AC:
[Given] a user is authenticated  
[When] user send request to get transaction history  
[Then]  
- system retrieves transaction data based on provided accountID, and time range  
- if no transaction exist, API should return 204 no content and empty list  

## US 3:
[As a] bank customer  
[I want to] transfer money to other person within a bank in the same currency  
[so that] I can manage my finance and send payments.  

### AC:
[Given] a user is authenticated  
[When] the user initiates a transfer specifying the source account ID, destination account ID, amount transfer  
[Then]  
- system validate that both account have the same currency
- system verify sufficient funds in the source account
- a new transaction record stored in the database
- return 400 if source's amount is not sufficient
- return 404 if account Id is not found
