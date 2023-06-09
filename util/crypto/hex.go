/*
 * Copyright 2023 veerdone
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/veerdone/gopkg/util"
)

func HexEncodeSha256(s string) string {
	b := HexEncodeSha256B(util.StringToSliceByte(s))

	return util.SliceByteToString(b)
}

func BHexEncodeSha256(b []byte) string {
	bytes := HexEncodeSha256B(b)

	return util.SliceByteToString(bytes)
}

func HexEncodeSha256B(b []byte) []byte {
	h := sha256.New()
	h.Write(b)
	sum := h.Sum(nil)

	dst := make([]byte, hex.EncodedLen(len(sum)))
	hex.Encode(dst, sum)

	return dst
}

func HmacSha256(msg, secret []byte) []byte {
	h := hmac.New(sha256.New, secret)
	h.Write(msg)

	return h.Sum(nil)
}
