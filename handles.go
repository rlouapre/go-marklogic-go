package goMarklogicGo

// Handle interface
type Handle interface {
	GetFormat() int
	Encode([]byte)
	Decode(interface{})
	Serialized() string
}

// RawHandle returns the raw string results of JSON or XML
type RawHandle struct {
	Format int
	bytes  []byte
}

// GetFormat returns int that represents XML or JSON
func (r *RawHandle) GetFormat() int {
	return r.Format
}

// Encode returns the bytes that represent XML or JSON
func (r *RawHandle) Encode(bytes []byte) {
	r.bytes = bytes
}

// Decode returns the bytes that represent XML or JSON
func (r *RawHandle) Decode(bytes interface{}) {
	r.Encode(bytes.([]byte))
}

// Get returns string of XML or JSON
func (r *RawHandle) Get() string {
	return string(r.bytes)
}

// Serialized returns string of XML or JSON
func (r *RawHandle) Serialized() string {
	return r.Get()
}

// MapHandle returns the raw string results of JSON or XML
type MapHandle struct {
	Format  int
	bytes   []byte
	mapItem *map[string]interface{}
}

// GetFormat returns int that represents XML or JSON
func (m *MapHandle) GetFormat() int {
	return m.Format
}

// Encode returns the bytes that represent XML or JSON
func (m *MapHandle) Encode(bytes []byte) {
	m.bytes = bytes
}

// Decode returns the bytes that represent XML or JSON
func (m *MapHandle) Decode(mapItem interface{}) {
	m.mapItem = mapItem.(*map[string]interface{})

}

// Get returns string of XML or JSON
func (m *MapHandle) Get() *map[string]interface{} {
	return m.mapItem
}

// Serialized returns string of XML or JSON
func (m *MapHandle) Serialized() string {
	return string(m.bytes)
}