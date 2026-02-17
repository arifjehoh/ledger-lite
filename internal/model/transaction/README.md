# Domain Model - Transaction
This directory contains the core domain model for the Transaction entity. It defines the structure of a Transaction, including its fields and any associated methods.

## Transaction Struct
The TransactionType represents the type of a transaction, which can be either "debit" or "credit". It is defined as a custom type with constants for each possible value.
- Credit: Represents a credit transaction, which increases the account balance.
- Debit: Represents a debit transaction, which decreases the account balance.

The Transaction struct represents a financial transaction in the system. It includes the following fields:
- ID: A unique identifier for the transaction (string).
- Type: The type of transaction, either "debit" or "credit" (TransactionType).
- Description: A brief description of the transaction (string).
- Amount: The amount of the transaction (float64).
- CreatedAt: The timestamp when the transaction was created (time.Time).
- CreatedBy: The identifier of the user or system that created the transaction (string).
- UpdatedAt: The timestamp when the transaction was last updated (time.Time).
- UpdatedBy: The identifier of the user or system that last updated the transaction (string).