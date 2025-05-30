// Code generated by fastssz. DO NOT EDIT.
// Hash: 1aadfa27e897a6d615870d7ce32e88af01e893825c6c8e28df52034ba5010ff0
// Version: 0.1.3
package types

import (
	"github.com/attestantio/go-eth2-client/spec"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the Contribution object
func (c *Contribution) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(c)
}

// MarshalSSZTo ssz marshals the Contribution object to a target array
func (c *Contribution) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'SelectionProofSig'
	dst = append(dst, c.SelectionProofSig[:]...)

	// Field (1) 'Contribution'
	if dst, err = c.Contribution.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the Contribution object
func (c *Contribution) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 256 {
		return ssz.ErrSize
	}

	// Field (0) 'SelectionProofSig'
	copy(c.SelectionProofSig[:], buf[0:96])

	// Field (1) 'Contribution'
	if err = c.Contribution.UnmarshalSSZ(buf[96:256]); err != nil {
		return err
	}

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Contribution object
func (c *Contribution) SizeSSZ() (size int) {
	size = 256
	return
}

// HashTreeRoot ssz hashes the Contribution object
func (c *Contribution) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(c)
}

// HashTreeRootWith ssz hashes the Contribution object with a hasher
func (c *Contribution) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'SelectionProofSig'
	hh.PutBytes(c.SelectionProofSig[:])

	// Field (1) 'Contribution'
	if err = c.Contribution.HashTreeRootWith(hh); err != nil {
		return
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Contribution object
func (c *Contribution) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(c)
}

// MarshalSSZ ssz marshals the BeaconVote object
func (b *BeaconVote) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BeaconVote object to a target array
func (b *BeaconVote) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'BlockRoot'
	dst = append(dst, b.BlockRoot[:]...)

	// Field (1) 'Source'
	if b.Source == nil {
		b.Source = new(phase0.Checkpoint)
	}
	if dst, err = b.Source.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (2) 'Target'
	if b.Target == nil {
		b.Target = new(phase0.Checkpoint)
	}
	if dst, err = b.Target.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BeaconVote object
func (b *BeaconVote) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 112 {
		return ssz.ErrSize
	}

	// Field (0) 'BlockRoot'
	copy(b.BlockRoot[:], buf[0:32])

	// Field (1) 'Source'
	if b.Source == nil {
		b.Source = new(phase0.Checkpoint)
	}
	if err = b.Source.UnmarshalSSZ(buf[32:72]); err != nil {
		return err
	}

	// Field (2) 'Target'
	if b.Target == nil {
		b.Target = new(phase0.Checkpoint)
	}
	if err = b.Target.UnmarshalSSZ(buf[72:112]); err != nil {
		return err
	}

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BeaconVote object
func (b *BeaconVote) SizeSSZ() (size int) {
	size = 112
	return
}

// HashTreeRoot ssz hashes the BeaconVote object
func (b *BeaconVote) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BeaconVote object with a hasher
func (b *BeaconVote) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'BlockRoot'
	hh.PutBytes(b.BlockRoot[:])

	// Field (1) 'Source'
	if b.Source == nil {
		b.Source = new(phase0.Checkpoint)
	}
	if err = b.Source.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (2) 'Target'
	if b.Target == nil {
		b.Target = new(phase0.Checkpoint)
	}
	if err = b.Target.HashTreeRootWith(hh); err != nil {
		return
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the BeaconVote object
func (b *BeaconVote) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(b)
}

// MarshalSSZ ssz marshals the ValidatorConsensusData object
func (v *ValidatorConsensusData) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(v)
}

// MarshalSSZTo ssz marshals the ValidatorConsensusData object to a target array
func (v *ValidatorConsensusData) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(16)

	// Offset (0) 'Duty'
	dst = ssz.WriteOffset(dst, offset)
	offset += v.Duty.SizeSSZ()

	// Field (1) 'Version'
	dst = ssz.MarshalUint64(dst, uint64(v.Version))

	// Offset (2) 'DataSSZ'
	dst = ssz.WriteOffset(dst, offset)

	// Field (0) 'Duty'
	if dst, err = v.Duty.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (2) 'DataSSZ'
	if size := len(v.DataSSZ); size > 8388608 {
		err = ssz.ErrBytesLengthFn("ValidatorConsensusData.DataSSZ", size, 8388608)
		return
	}
	dst = append(dst, v.DataSSZ...)

	return
}

// UnmarshalSSZ ssz unmarshals the ValidatorConsensusData object
func (v *ValidatorConsensusData) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 16 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o2 uint64

	// Offset (0) 'Duty'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 != 16 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (1) 'Version'
	v.Version = spec.DataVersion(ssz.UnmarshallUint64(buf[4:12]))

	// Offset (2) 'DataSSZ'
	if o2 = ssz.ReadOffset(buf[12:16]); o2 > size || o0 > o2 {
		return ssz.ErrOffset
	}

	// Field (0) 'Duty'
	{
		buf = tail[o0:o2]
		if err = v.Duty.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (2) 'DataSSZ'
	{
		buf = tail[o2:]
		if len(buf) > 8388608 {
			return ssz.ErrBytesLength
		}
		if cap(v.DataSSZ) == 0 {
			v.DataSSZ = make([]byte, 0, len(buf))
		}
		v.DataSSZ = append(v.DataSSZ, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the ValidatorConsensusData object
func (v *ValidatorConsensusData) SizeSSZ() (size int) {
	size = 16

	// Field (0) 'Duty'
	size += v.Duty.SizeSSZ()

	// Field (2) 'DataSSZ'
	size += len(v.DataSSZ)

	return
}

// HashTreeRoot ssz hashes the ValidatorConsensusData object
func (v *ValidatorConsensusData) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(v)
}

// HashTreeRootWith ssz hashes the ValidatorConsensusData object with a hasher
func (v *ValidatorConsensusData) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Duty'
	if err = v.Duty.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'Version'
	hh.PutUint64(uint64(v.Version))

	// Field (2) 'DataSSZ'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(v.DataSSZ))
		if byteLen > 8388608 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.Append(v.DataSSZ)
		hh.MerkleizeWithMixin(elemIndx, byteLen, (8388608+31)/32)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the ValidatorConsensusData object
func (v *ValidatorConsensusData) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(v)
}
