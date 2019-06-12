// Copyright 2019 The Energi Core Authors
// This file is part of Energi Core.
//
// Energi Core is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Energi Core is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Energi Core. If not, see <http://www.gnu.org/licenses/>.

// Energi Governance system is the fundamental part of Energi Core.

// NOTE: It's not allowed to change the compiler due to byte-to-byte
//       match requirement.
pragma solidity 0.5.9;
//pragma experimental SMTChecker;

import { IGovernedContract, GovernedContract } from "./GovernedContract.sol";
import { IBlacklistRegistry } from "./IBlacklistRegistry.sol";
import { StorageBase }  from "./StorageBase.sol";

/**
 * Permanent storage of Blacklist Registry V1 data.
 */
contract StorageBlacklistRegistryV1 is
    StorageBase
{
    // NOTE: ABIEncoderV2 is not acceptable at the moment of development!
}


/**
 * Genesis hardcoded version of BlacklistRegistry.
 *
 * NOTE: it MUST NOT change after blockchain launch!
 */
contract BlacklistRegistryV1 is
    GovernedContract,
    IBlacklistRegistry
{
    // Data for migration
    //---------------------------------
    StorageBlacklistRegistryV1 public v1storage;
    //---------------------------------

    constructor(address _proxy) public GovernedContract(_proxy) {
        v1storage = new StorageBlacklistRegistryV1();
    }

    // IGovernedContract
    //---------------------------------
    function _destroy(IGovernedContract _newImpl) internal {
        v1storage.setOwner(_newImpl);
    }

    // IBlacklistRegistry
    //---------------------------------

    // Safety
    //---------------------------------
    function () external payable {
        revert("Not supported");
    }
}