const AccountTypes = {
  bankAccount: 'BANK_ACCOUNT',
  // investment: 'INVESTMENT',
  // realState: 'REAL_STATE',
  // vehicle: 'VEHICLE',
  creditCard: 'CREDIT_CARD',
  // loan: 'LOAN',
  // otherAsset: 'OTHER_ASSET',
  // otherLiability: 'OTHER_LIABILITY',
}

export type TAccountTypes = (typeof AccountTypes)[keyof typeof AccountTypes]

const BankAccountTypes = {
  savings: 'SAVINGS',
  checking: 'CHECKING',
} as const

export type TBankAccountTypes =
  (typeof BankAccountTypes)[keyof typeof BankAccountTypes]

const CreditCardTypes = {
  visa: 'VISA',
  mastercard: 'MASTERCARD',
}

export type TCreditCardTypes =
  (typeof CreditCardTypes)[keyof typeof CreditCardTypes]

export type TClassification = 'ASSET' | 'LIABILITY'

export type TCurrency = 'COP'
