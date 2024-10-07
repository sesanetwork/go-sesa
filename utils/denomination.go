// Copyright 2023 The go-sesa Authors
// This file is part of the go-sesa library.
//
// The go-sesa library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-sesa library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-sesa library. If not, see <http://www.gnu.org/licenses/>.

package utils

import "math/big"

// These are the multipliers for sesa denominations.
// 1 sesa = 10^18 Wei

const (
	WEI  = 1
	GWEI = 1e9
	sesa  = 1e18
)

func Tosesa(amount uint64) *big.Int {
	return new(big.Int).Mul(new(big.Int).SetUint64(amount), big.NewInt(sesa))
}
