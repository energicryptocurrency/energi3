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

'use strict';

const MockProxy = artifacts.require('MockProxy');
const MockContract = artifacts.require('MockContract');
const BlacklistRegistryV1 = artifacts.require('BlacklistRegistryV1');
const IBlacklistRegistry = artifacts.require('IBlacklistRegistry');
const StorageBlacklistRegistryV1 = artifacts.require('StorageBlacklistRegistryV1');

const common = require('./common');

contract("BlacklistRegistryV1", async accounts => {
    const s = {
        artifacts,
        accounts,
        assert,
        it,
        web3,
    };

    before(async () => {
        s.orig = await BlacklistRegistryV1.deployed();
        s.proxy = await MockProxy.at(await s.orig.proxy());
        s.fake = await MockContract.new(s.proxy.address);
        s.proxy_abi = await BlacklistRegistryV1.at(s.proxy.address);
        s.token_abi = await IBlacklistRegistry.at(s.proxy.address);
        await s.proxy.setImpl(s.orig.address);
        s.storage = await StorageBlacklistRegistryV1.at(await s.proxy_abi.v1storage());
        Object.freeze(s);
    });

    after(async () => {
        const impl = await BlacklistRegistryV1.new(s.proxy.address);
        await s.proxy.setImpl(impl.address);
    });

    describe('common pre', () => common.govPreTests(s) );

    //---
    describe('Primary', () => {
    });

    //---
    describe('common post', () => common.govPostTests(s) );
});
