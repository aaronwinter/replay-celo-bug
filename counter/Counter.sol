// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

contract Counter {
    uint256 private _clock;

  constructor(uint256 init) public {
      _clock = init;
  }

  function increment() public returns(bool) {
      _clock = _clock + 42;
  }

  function value() public view returns(uint256) {
      return _clock;
  }
}
