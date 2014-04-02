package compressing

import (
	"github.com/balzaczyy/golucene/core/store"
)

/*
Efficient index format for block-based Codecs.

This writer generates a file which be loaded into memory using
memory-efficient data structures to quickly locate the block that
contains any document.

In order to have a compact in-memory representation, for every block
of 1024 chunks, this index computes the average number of bytes per
chunk and for every chunk, only stores the difference between

- ${chunk number} * ${average length of a chunk}
- and the actual start offset of the chunk

Data is written as follows:

	- PackedIntsVersion, <Block>^BlockCount, BlocksEndMarker
	- PackedIntsVersion --> VERSION_CURRENT as a vint
	- BlocksEndMarker --> 0 as a vint, this marks the end of blocks since blocks are not allowed to start with 0
	- Block --> BlockChunks, <Docbases>, <StartPointers>
	- BlockChunks --> a vint which is the number of chunks encoded in the block
	- DocBases --> DocBase, AvgChunkDocs, BitsPerDocbaseDelta, DocBaseDeltas
	- DocBase --> first document ID of the block of chunks, as a vint
	- AvgChunkDocs --> average number of documents in a single chunk, as a vint
	- BitsPerDocBaseDelta --> number of bits required to represent a delta from the average using ZigZag encoding
	- DocBaseDeltas --> packed array of BlockChunks elements of BitsPerDocBaseDelta bits each, representing the deltas from the average doc base using ZigZag encoding.
	- StartPointers --> StartointerBase, AveChunkSize, BitsPerStartPointerDelta, StartPointerDeltas
	- StartPointerBase --> the first start ointer of the block, as a vint64
	- AvgChunkSize --> the average size of a chunk of compressed documents, as a vint64
	- BitsPerStartPointerDelta --> number of bits required to represent a delta from the average using ZigZag encoding
	- StartPointerDeltas --> packed array of BlockChunks elements of BitsPerStartPointerDelta bits each, representing the deltas from the average start pointer using ZigZag encoding

Notes

- For any block, the doc base of the n-th chunk can be restored with
DocBase + AvgChunkDocs * n + DOcBsaeDeltas[n].
- For any block, the start pointer of the n-th chunk can be restored
with StartPointerBase + AvgChunkSize * n + StartPointerDeltas[n].
- Once data is loaded into memory, you can lookup the start pointer
of any document by performing two binary searches: a first one based
on the values of DocBase in order to find the right block, and then
inside the block based on DocBaseDeltas (by reconstructing the doc
bases for every chunk).
*/
type StoredFieldsIndexWriter struct {
	fieldsIndexOut store.IndexOutput
}

func NewStoredFieldsIndexWriter(indexOutput store.IndexOutput) (*StoredFieldsIndexWriter, error) {
	panic("not implemented yet")
}

func (w *StoredFieldsIndexWriter) Close() error {
	return w.fieldsIndexOut.Close()
}
