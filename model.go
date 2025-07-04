// Account Container
// Mirrors TB's account record with user-defined and system fields
type AccountContainer struct {
  ID          num.Int128   `knox:"id,pk"`
  UserData128 num.Int128   `knox:"user_data_128"`
  UserData64  uint64       `knox:"user_data_64"`
  UserData32  uint32       `knox:"user_data_32"`
  Reserved    uint32       `knox:"reserved"`
  Ledger      uint32       `knox:"ledger"`
  Code        uint16       `knox:"code"`
  Flags       AccountFlags `knox:"flags"`
  Timestamp   uint64       `knox:"timestamp"`
}

// Account Flags
type AccountFlags uint16

const (
  AccountFlagLinked AccountFlags = 1 << iota
  AccountFlagDebitsMustNotExceedCredits
  AccountFlagCreditsMustNotExceedDebits
  AccountFlagHistory
  AccountFlagImported
  AccountFlagClosed
)

// Account Balances: volatile fields for in-mem processing
type AccountBalances struct {
  DebitsPosted   num.Int128 `knox:"debits_posted"`
  CreditsPosted  num.Int128 `knox:"credits_posted"`
  DebitsPending  num.Int128 `knox:"debits_pending"`
  CreditsPending num.Int128 `knox:"credits_pending"`
}

// Transfer Container
type TransferContainer struct {
  ID              num.Int128    `knox:"id,pk"`
  DebitAccountID  num.Int128    `knox:"debit_account_id"`
  CreditAccountID num.Int128    `knox:"credit_account_id"`
  Amount          num.Int128    `knox:"amount"`
  PendingID       num.Int128    `knox:"pending_id"`
  UserData128     num.Int128    `knox:"user_data_128"`
  UserData64      uint64        `knox:"user_data_64"`
  UserData32      uint32        `knox:"user_data_32"`
  Timeout         uint32        `knox:"timeout"`
  Ledger          uint32        `knox:"ledger"`
  Code            uint16        `knox:"code"`
  Flags           TransferFlags `knox:"flags"`
  Timestamp       uint64        `knox:"timestamp"`
}

// Transfer Flags
type TransferFlags uint16

const (
  TransferFlagLinked          TransferFlags = 1 << iota // 0x0001 - Chain this transfer with the next
  TransferFlagPending                                   // 0x0002 - Mark as pending (phase 1)
  TransferFlagPostPending                               // 0x0004 - Commit a pending transfer (phase 2)
  TransferFlagVoidPending                               // 0x0008 - Cancel a pending transfer (phase 2)
  TransferFlagBalancingDebit                            // 0x0010 - Transfer up to amount, capped by debit balance
  TransferFlagBalancingCredit                           // 0x0020 - Transfer up to amount, capped by credit balance
  TransferFlagClosingDebit                              // 0x0040 - Close debit account on success (requires pending)
  TransferFlagClosingCredit                             // 0x0080 - Close credit account on success (requires pending)
  TransferFlagImported                                  // 0x0100 - Use user-defined timestamp (historical import)
)