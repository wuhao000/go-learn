package task3

import "gorm.io/gorm"

//题目2：事务语句
//假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键，
//from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
//要求 ：
//编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，
//需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，
//并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。

type Account struct {
  Id      int
  Balance int
}

type Transaction struct {
  Id            int
  FromAccountId int
  ToAccountId   int
  Amount        int
}

// TransferMoney 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，
//需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，
//并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
func TransferMoney(db *gorm.DB, fromId int, toId int) {
  tx := db.Begin()
  trans := Transaction{
    FromAccountId: fromId,
    ToAccountId:   toId,
    Amount:        100,
  }
  tx.Create(&trans)
  var fromAccount, toAccount Account
  tx.Find(&fromAccount, fromId)
  tx.First(&toAccount, toId)

  if fromAccount.Balance < 100 {
    tx.Rollback()
    return
  }
  fromAccount.Balance -= 100
  toAccount.Balance += 100
  tx.Updates(&fromAccount)
  tx.Updates(&toAccount)
  tx.Commit()
}
