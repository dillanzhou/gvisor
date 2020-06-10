// automatically generated by stateify.

package kernfs

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (x *DynamicBytesFile) beforeSave() {}
func (x *DynamicBytesFile) save(m state.Map) {
	x.beforeSave()
	m.Save("InodeAttrs", &x.InodeAttrs)
	m.Save("InodeNoopRefCount", &x.InodeNoopRefCount)
	m.Save("InodeNotDirectory", &x.InodeNotDirectory)
	m.Save("InodeNotSymlink", &x.InodeNotSymlink)
	m.Save("locks", &x.locks)
	m.Save("data", &x.data)
}

func (x *DynamicBytesFile) afterLoad() {}
func (x *DynamicBytesFile) load(m state.Map) {
	m.Load("InodeAttrs", &x.InodeAttrs)
	m.Load("InodeNoopRefCount", &x.InodeNoopRefCount)
	m.Load("InodeNotDirectory", &x.InodeNotDirectory)
	m.Load("InodeNotSymlink", &x.InodeNotSymlink)
	m.Load("locks", &x.locks)
	m.Load("data", &x.data)
}

func (x *DynamicBytesFD) beforeSave() {}
func (x *DynamicBytesFD) save(m state.Map) {
	x.beforeSave()
	m.Save("FileDescriptionDefaultImpl", &x.FileDescriptionDefaultImpl)
	m.Save("DynamicBytesFileDescriptionImpl", &x.DynamicBytesFileDescriptionImpl)
	m.Save("LockFD", &x.LockFD)
	m.Save("vfsfd", &x.vfsfd)
	m.Save("inode", &x.inode)
}

func (x *DynamicBytesFD) afterLoad() {}
func (x *DynamicBytesFD) load(m state.Map) {
	m.Load("FileDescriptionDefaultImpl", &x.FileDescriptionDefaultImpl)
	m.Load("DynamicBytesFileDescriptionImpl", &x.DynamicBytesFileDescriptionImpl)
	m.Load("LockFD", &x.LockFD)
	m.Load("vfsfd", &x.vfsfd)
	m.Load("inode", &x.inode)
}

func (x *StaticDirectory) beforeSave() {}
func (x *StaticDirectory) save(m state.Map) {
	x.beforeSave()
	m.Save("InodeNotSymlink", &x.InodeNotSymlink)
	m.Save("InodeDirectoryNoNewChildren", &x.InodeDirectoryNoNewChildren)
	m.Save("InodeAttrs", &x.InodeAttrs)
	m.Save("InodeNoDynamicLookup", &x.InodeNoDynamicLookup)
	m.Save("OrderedChildren", &x.OrderedChildren)
	m.Save("locks", &x.locks)
}

func (x *StaticDirectory) afterLoad() {}
func (x *StaticDirectory) load(m state.Map) {
	m.Load("InodeNotSymlink", &x.InodeNotSymlink)
	m.Load("InodeDirectoryNoNewChildren", &x.InodeDirectoryNoNewChildren)
	m.Load("InodeAttrs", &x.InodeAttrs)
	m.Load("InodeNoDynamicLookup", &x.InodeNoDynamicLookup)
	m.Load("OrderedChildren", &x.OrderedChildren)
	m.Load("locks", &x.locks)
}

func (x *slotList) beforeSave() {}
func (x *slotList) save(m state.Map) {
	x.beforeSave()
	m.Save("head", &x.head)
	m.Save("tail", &x.tail)
}

func (x *slotList) afterLoad() {}
func (x *slotList) load(m state.Map) {
	m.Load("head", &x.head)
	m.Load("tail", &x.tail)
}

func (x *slotEntry) beforeSave() {}
func (x *slotEntry) save(m state.Map) {
	x.beforeSave()
	m.Save("next", &x.next)
	m.Save("prev", &x.prev)
}

func (x *slotEntry) afterLoad() {}
func (x *slotEntry) load(m state.Map) {
	m.Load("next", &x.next)
	m.Load("prev", &x.prev)
}

func init() {
	state.Register("pkg/sentry/fsimpl/kernfs.DynamicBytesFile", (*DynamicBytesFile)(nil), state.Fns{Save: (*DynamicBytesFile).save, Load: (*DynamicBytesFile).load})
	state.Register("pkg/sentry/fsimpl/kernfs.DynamicBytesFD", (*DynamicBytesFD)(nil), state.Fns{Save: (*DynamicBytesFD).save, Load: (*DynamicBytesFD).load})
	state.Register("pkg/sentry/fsimpl/kernfs.StaticDirectory", (*StaticDirectory)(nil), state.Fns{Save: (*StaticDirectory).save, Load: (*StaticDirectory).load})
	state.Register("pkg/sentry/fsimpl/kernfs.slotList", (*slotList)(nil), state.Fns{Save: (*slotList).save, Load: (*slotList).load})
	state.Register("pkg/sentry/fsimpl/kernfs.slotEntry", (*slotEntry)(nil), state.Fns{Save: (*slotEntry).save, Load: (*slotEntry).load})
}
