//+build !noasm

/*
 * Minio Cloud Storage, (C) 2017 Minio, Inc.
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

package highwayhash

var (
	useSSE4 = false
	useAVX2 = false
	useNEON = true
)

//go:noescape
func updateArm64(state *[16]uint64, msg []byte)

func initialize(state *[16]uint64, key []byte) {
	initializeGeneric(state, key)
}

func update(state *[16]uint64, msg []byte) {
	if useNEON {
		updateArm64(state, msg)
	} else {
		updateGeneric(state, msg)
	}
}

func finalize(out []byte, state *[16]uint64) {
	finalizeGeneric(out, state)
}
