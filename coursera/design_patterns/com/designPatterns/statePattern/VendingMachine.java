package com.designPatterns.statePattern;

// public class VendingMachine {
//   private State currentState;
//   // allow it to be called from outside the class
//   private int count;
//
//   public VendingMachine(int count) {
//     if (count > 0) {
//       currentState = State.Idle;
//       this.count = count;
//     } else {
//       currentState = State.OutOfStock;
//       this.count = 0;
//     }
//   }
//
//   // handle user events
//   public void insertDollar() {
//     if (currentState == State.Idle) {
//       currentState = State.HasOneDollar;
//
//     } else if (currentState == State.HasOneDollar) {
//       doReturnMoney();
//       currentState = State.Idle;
//
//     } else if (currentState == State.OutOfStock) {
//       doReturnMoney();
//     }
//   }
//
//   public void ejectMoney() {}
//
//   public void dispense() {}
//
//   public void doReturnMoney() {}
// }

public class VendingMachine {

  private IState idleState;
  private IState hasOneDollarState;
  private IState outOfStockState;

  private IState currentState;
  private int count;

  public VendingMachine() {
    idleState = new IdleState();
    hasOneDollarState = new HasOneDollarState();
    outOfStockState = new OutOfStockState();
  }

  public IState getHasOneDollarState() {
    return this.hasOneDollarState;
  }

  public IState getIdleState() {
    return this.idleState;
  }

  public IState getOutOfStockState() {
    return this.outOfStockState;
  }

  public int getCount() {
    return this.count;
  }

  public IState getCurrentState() {
    return this.currentState;
  }

  public void setState(IState state) {
    if (count > 0) {
      this.currentState = idleState;
    } else {
      this.currentState = outOfStockState;
      this.count = 0;
    }
  }

  public void ejectMoney() {}

  public void dispense() {}

  public void doReturnMoney() {}

  public void doReleaseProduct() {}
}
