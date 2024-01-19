package com.designPatterns.facadePattern;

import java.math.BigDecimal;

public class Savings implements IAccount {
  private BigDecimal initAmount;

  public Savings(BigDecimal initAmount) {
    this.initAmount = initAmount;
  }

  public BigDecimal getInitAmount() {
    return this.initAmount;
  }

  public void deposit(BigDecimal amount) {}

  public void withdraw(BigDecimal amount) {}

  public void transfer(IAccount toAccount, BigDecimal amount) {}

  public int getAccountNumber() {
    return 2; // Some number account
  }
}
