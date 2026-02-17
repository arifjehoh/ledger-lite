# Domain Model - Account
This directory contains the core domain model for the Account entity. It defines the structure of an Account, including its fields and any associated methods.

## Account Struct
The Account struct represents a financial account in the system. It includes the following fields:
- ID: A unique identifier for the account (string).
- Name: The name of the account holder (string).
- CreatedAt: The timestamp when the account was created (time.Time).

## Balance Computation
The balance of an account is not stored directly in the Account struct. Instead, it is computed based on the transactions associated with the account. This allows for a more accurate and up-to-date balance, as it reflects all transactions that have occurred. The balance can be computed by summing the amounts of all transactions linked to the account, taking into account whether they are debits or credits.

```peusdocode
func (a *Account) ComputeBalance(transactions []Transaction) float64 {
    var balance float64
    for _, tx := range transactions {
        if tx.Type == Credit {
            balance += tx.Amount
        } else if tx.Type == Debit {
            balance -= tx.Amount
        }
    }
    return balance
}
```