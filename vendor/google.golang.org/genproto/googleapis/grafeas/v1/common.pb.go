// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grafeas/v1/common.proto

package grafeas

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Kind represents the kinds of notes supported.
type NoteKind int32

const (
	// Unknown.
	NoteKind_NOTE_KIND_UNSPECIFIED NoteKind = 0
	// The note and occurrence represent a package vulnerability.
	NoteKind_VULNERABILITY NoteKind = 1
	// The note and occurrence assert build provenance.
	NoteKind_BUILD NoteKind = 2
	// This represents an image basis relationship.
	NoteKind_IMAGE NoteKind = 3
	// This represents a package installed via a package manager.
	NoteKind_PACKAGE NoteKind = 4
	// The note and occurrence track deployment events.
	NoteKind_DEPLOYMENT NoteKind = 5
	// The note and occurrence track the initial discovery status of a resource.
	NoteKind_DISCOVERY NoteKind = 6
	// This represents a logical "role" that can attest to artifacts.
	NoteKind_ATTESTATION NoteKind = 7
)

var NoteKind_name = map[int32]string{
	0: "NOTE_KIND_UNSPECIFIED",
	1: "VULNERABILITY",
	2: "BUILD",
	3: "IMAGE",
	4: "PACKAGE",
	5: "DEPLOYMENT",
	6: "DISCOVERY",
	7: "ATTESTATION",
}

var NoteKind_value = map[string]int32{
	"NOTE_KIND_UNSPECIFIED": 0,
	"VULNERABILITY":         1,
	"BUILD":                 2,
	"IMAGE":                 3,
	"PACKAGE":               4,
	"DEPLOYMENT":            5,
	"DISCOVERY":             6,
	"ATTESTATION":           7,
}

func (x NoteKind) String() string {
	return proto.EnumName(NoteKind_name, int32(x))
}

func (NoteKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5eb785a4ec53d49e, []int{0}
}

// Metadata for any related URL information.
type RelatedUrl struct {
	// Specific URL associated with the resource.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// Label to describe usage of the URL.
	Label                string   `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RelatedUrl) Reset()         { *m = RelatedUrl{} }
func (m *RelatedUrl) String() string { return proto.CompactTextString(m) }
func (*RelatedUrl) ProtoMessage()    {}
func (*RelatedUrl) Descriptor() ([]byte, []int) {
	return fileDescriptor_5eb785a4ec53d49e, []int{0}
}

func (m *RelatedUrl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RelatedUrl.Unmarshal(m, b)
}
func (m *RelatedUrl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RelatedUrl.Marshal(b, m, deterministic)
}
func (m *RelatedUrl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RelatedUrl.Merge(m, src)
}
func (m *RelatedUrl) XXX_Size() int {
	return xxx_messageInfo_RelatedUrl.Size(m)
}
func (m *RelatedUrl) XXX_DiscardUnknown() {
	xxx_messageInfo_RelatedUrl.DiscardUnknown(m)
}

var xxx_messageInfo_RelatedUrl proto.InternalMessageInfo

func (m *RelatedUrl) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *RelatedUrl) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

// Verifiers (e.g. Kritis implementations) MUST verify signatures
// with respect to the trust anchors defined in policy (e.g. a Kritis policy).
// Typically this means that the verifier has been configured with a map from
// `public_key_id` to public key material (and any required parameters, e.g.
// signing algorithm).
//
// In particular, verification implementations MUST NOT treat the signature
// `public_key_id` as anything more than a key lookup hint. The `public_key_id`
// DOES NOT validate or authenticate a public key; it only provides a mechanism
// for quickly selecting a public key ALREADY CONFIGURED on the verifier through
// a trusted channel. Verification implementations MUST reject signatures in any
// of the following circumstances:
//   * The `public_key_id` is not recognized by the verifier.
//   * The public key that `public_key_id` refers to does not verify the
//     signature with respect to the payload.
//
// The `signature` contents SHOULD NOT be "attached" (where the payload is
// included with the serialized `signature` bytes). Verifiers MUST ignore any
// "attached" payload and only verify signatures with respect to explicitly
// provided payload (e.g. a `payload` field on the proto message that holds
// this Signature, or the canonical serialization of the proto message that
// holds this signature).
type Signature struct {
	// The content of the signature, an opaque bytestring.
	// The payload that this signature verifies MUST be unambiguously provided
	// with the Signature during verification. A wrapper message might provide
	// the payload explicitly. Alternatively, a message might have a canonical
	// serialization that can always be unambiguously computed to derive the
	// payload.
	Signature []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	// The identifier for the public key that verifies this signature.
	//   * The `public_key_id` is required.
	//   * The `public_key_id` MUST be an RFC3986 conformant URI.
	//   * When possible, the `public_key_id` SHOULD be an immutable reference,
	//     such as a cryptographic digest.
	//
	// Examples of valid `public_key_id`s:
	//
	// OpenPGP V4 public key fingerprint:
	//   * "openpgp4fpr:74FAF3B861BDA0870C7B6DEF607E48D2A663AEEA"
	// See https://www.iana.org/assignments/uri-schemes/prov/openpgp4fpr for more
	// details on this scheme.
	//
	// RFC6920 digest-named SubjectPublicKeyInfo (digest of the DER
	// serialization):
	//   * "ni:///sha-256;cD9o9Cq6LG3jD0iKXqEi_vdjJGecm_iXkbqVoScViaU"
	//   * "nih:///sha-256;703f68f42aba2c6de30f488a5ea122fef76324679c9bf89791ba95a1271589a5"
	PublicKeyId          string   `protobuf:"bytes,2,opt,name=public_key_id,json=publicKeyId,proto3" json:"public_key_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Signature) Reset()         { *m = Signature{} }
func (m *Signature) String() string { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()    {}
func (*Signature) Descriptor() ([]byte, []int) {
	return fileDescriptor_5eb785a4ec53d49e, []int{1}
}

func (m *Signature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Signature.Unmarshal(m, b)
}
func (m *Signature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Signature.Marshal(b, m, deterministic)
}
func (m *Signature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Signature.Merge(m, src)
}
func (m *Signature) XXX_Size() int {
	return xxx_messageInfo_Signature.Size(m)
}
func (m *Signature) XXX_DiscardUnknown() {
	xxx_messageInfo_Signature.DiscardUnknown(m)
}

var xxx_messageInfo_Signature proto.InternalMessageInfo

func (m *Signature) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Signature) GetPublicKeyId() string {
	if m != nil {
		return m.PublicKeyId
	}
	return ""
}

func init() {
	proto.RegisterEnum("grafeas.v1.NoteKind", NoteKind_name, NoteKind_value)
	proto.RegisterType((*RelatedUrl)(nil), "grafeas.v1.RelatedUrl")
	proto.RegisterType((*Signature)(nil), "grafeas.v1.Signature")
}

func init() { proto.RegisterFile("grafeas/v1/common.proto", fileDescriptor_5eb785a4ec53d49e) }

var fileDescriptor_5eb785a4ec53d49e = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xcf, 0x4e, 0xea, 0x40,
	0x18, 0xc5, 0x6f, 0xe1, 0x02, 0xb7, 0x1f, 0xb7, 0x3a, 0x4e, 0x34, 0x62, 0xe2, 0xc2, 0xb0, 0x32,
	0x2e, 0xda, 0x10, 0x5d, 0x98, 0xb8, 0x2a, 0x74, 0x24, 0x93, 0x96, 0xb6, 0xf6, 0x0f, 0x09, 0x6e,
	0x9a, 0x42, 0xc7, 0x49, 0xe3, 0xd0, 0x21, 0xa5, 0x90, 0xf0, 0x0c, 0xbe, 0x85, 0x4f, 0x6a, 0x28,
	0x18, 0x76, 0xe7, 0xfc, 0xce, 0xcc, 0xb7, 0xf8, 0xc1, 0x35, 0x2f, 0xd3, 0x0f, 0x96, 0xae, 0x8d,
	0xed, 0xc0, 0x58, 0xc8, 0xe5, 0x52, 0x16, 0xfa, 0xaa, 0x94, 0x95, 0xc4, 0x70, 0x1c, 0xf4, 0xed,
	0xa0, 0xff, 0x04, 0x10, 0x30, 0x91, 0x56, 0x2c, 0x8b, 0x4b, 0x81, 0x11, 0x34, 0x37, 0xa5, 0xe8,
	0x29, 0x77, 0xca, 0xbd, 0x1a, 0xec, 0x23, 0xbe, 0x84, 0x96, 0x48, 0xe7, 0x4c, 0xf4, 0x1a, 0x35,
	0x3b, 0x94, 0xfe, 0x04, 0xd4, 0x30, 0xe7, 0x45, 0x5a, 0x6d, 0x4a, 0x86, 0x6f, 0x41, 0x5d, 0xff,
	0x96, 0xfa, 0xeb, 0xff, 0xe0, 0x04, 0x70, 0x1f, 0xb4, 0xd5, 0x66, 0x2e, 0xf2, 0x45, 0xf2, 0xc9,
	0x76, 0x49, 0x9e, 0x1d, 0x0f, 0x75, 0x0f, 0xd0, 0x66, 0x3b, 0x9a, 0x3d, 0x7c, 0x29, 0xf0, 0xcf,
	0x95, 0x15, 0xb3, 0xf3, 0x22, 0xc3, 0x37, 0x70, 0xe5, 0x7a, 0x11, 0x49, 0x6c, 0xea, 0x5a, 0x49,
	0xec, 0x86, 0x3e, 0x19, 0xd1, 0x57, 0x4a, 0x2c, 0xf4, 0x07, 0x5f, 0x80, 0x36, 0x8d, 0x1d, 0x97,
	0x04, 0xe6, 0x90, 0x3a, 0x34, 0x9a, 0x21, 0x05, 0xab, 0xd0, 0x1a, 0xc6, 0xd4, 0xb1, 0x50, 0x63,
	0x1f, 0xe9, 0xc4, 0x1c, 0x13, 0xd4, 0xc4, 0x5d, 0xe8, 0xf8, 0xe6, 0xc8, 0xde, 0x97, 0xbf, 0xf8,
	0x0c, 0xc0, 0x22, 0xbe, 0xe3, 0xcd, 0x26, 0xc4, 0x8d, 0x50, 0x0b, 0x6b, 0xa0, 0x5a, 0x34, 0x1c,
	0x79, 0x53, 0x12, 0xcc, 0x50, 0x1b, 0x9f, 0x43, 0xd7, 0x8c, 0x22, 0x12, 0x46, 0x66, 0x44, 0x3d,
	0x17, 0x75, 0x86, 0x6f, 0xa0, 0xe5, 0x52, 0x3f, 0x39, 0xf2, 0x95, 0xf7, 0x67, 0x2e, 0x25, 0x17,
	0x4c, 0xe7, 0x52, 0xa4, 0x05, 0xd7, 0x65, 0xc9, 0x0d, 0xce, 0x8a, 0xda, 0xa6, 0x71, 0x98, 0xd2,
	0x55, 0xbe, 0x36, 0x4e, 0xc2, 0x5f, 0x8e, 0xf1, 0xbb, 0xd1, 0x1c, 0x07, 0xe6, 0xbc, 0x5d, 0x3f,
	0x7d, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0x98, 0x52, 0x07, 0xe2, 0x93, 0x01, 0x00, 0x00,
}