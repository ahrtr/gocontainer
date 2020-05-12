// Copyright (c) 2019, Benjamin Wang (benjamin_wang@aliyun.com). All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package collection

// Interface is a type of collection, all containers should implement this interface.
type Interface interface {
	// Size returns the number of elements in the collection.
	Size() int
	// IsEmpty returns true if this container contains no elements.
	IsEmpty() bool
	// Clear removes all of the elements from this container.
	Clear()
}
