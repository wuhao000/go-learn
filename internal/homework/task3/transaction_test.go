package task3

import (
  "fmt"
  "testing"
)

// 测试事务回滚
func TestRollback(t *testing.T) {
  db := createDb()
  err := db.AutoMigrate(Transaction{})
  if err != nil {
    return
  }
  err = db.AutoMigrate(Account{})
  if err != nil {
    return
  }
  db.Where("1 = 1").Delete(&Transaction{})
  db.Delete(&Account{}, 1)
  db.Delete(&Account{}, 2)
  db.Create(&Account{
    Id:      1,
    Balance: 50,
  })
  db.Create(&Account{
    Id:      2,
    Balance: 100,
  })
  TransferMoney(db, 1, 2)
  var account1, account2 Account
  db.Find(&account1, 1)
  db.Find(&account2, 2)
  if account1.Balance != 50 {
    t.Error("balance of account 1 should be 50")
  }
  if account2.Balance != 100 {
    t.Error("balance of account 2 should be 100")
  }
  fmt.Println(account1)
  fmt.Println(account2)
  var transactions []Transaction
  db.Where("1=1").Find(&transactions)
  fmt.Println(transactions)
  if len(transactions) > 0 {
    t.Error("transactions should  be empty")
  }
}

// 测试事务不回滚
func TestTransactionNotRollback(t *testing.T) {
  db := createDb()
  err := db.AutoMigrate(Transaction{})
  if err != nil {
    return
  }
  err = db.AutoMigrate(Account{})
  if err != nil {
    return
  }
  db.Where("1 = 1").Delete(&Transaction{})
  db.Delete(&Account{}, 1)
  db.Delete(&Account{}, 2)
  db.Create(&Account{
    Id:      1,
    Balance: 500,
  })
  db.Create(&Account{
    Id:      2,
    Balance: 100,
  })
  TransferMoney(db, 1, 2)
  var account1, account2 Account
  db.Find(&account1, 1)
  db.Find(&account2, 2)
  if account1.Balance != 400 {
    t.Error("balance of account 1 should be 400")
  }
  if account2.Balance != 200 {
    t.Error("balance of account 2 should be 200")
  }
  fmt.Println(account1)
  fmt.Println(account2)
  var transactions []Transaction
  db.Where("1=1").Find(&transactions)
  fmt.Println(transactions)
  if len(transactions) == 0 {
    t.Error("transactions should not be empty")
  }
}
