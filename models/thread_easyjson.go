// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	time "time"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson2d00218DecodeForumDBModels(in *jlexer.Lexer, out *ThreadVote) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "nickname":
			out.Nickname = string(in.String())
		case "voice":
			out.Voice = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson2d00218EncodeForumDBModels(out *jwriter.Writer, in ThreadVote) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"nickname\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Nickname))
	}
	{
		const prefix string = ",\"voice\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Voice))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ThreadVote) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2d00218EncodeForumDBModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ThreadVote) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2d00218EncodeForumDBModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ThreadVote) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2d00218DecodeForumDBModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ThreadVote) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2d00218DecodeForumDBModels(l, v)
}
func easyjson2d00218DecodeForumDBModels1(in *jlexer.Lexer, out *ThreadShort) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "author":
			out.Author = string(in.String())
		case "message":
			out.Message = string(in.String())
		case "title":
			out.Title = string(in.String())
		case "slug":
			if in.IsNull() {
				in.Skip()
				out.Slug = nil
			} else {
				if out.Slug == nil {
					out.Slug = new(string)
				}
				*out.Slug = string(in.String())
			}
		case "created":
			if in.IsNull() {
				in.Skip()
				out.Created = nil
			} else {
				if out.Created == nil {
					out.Created = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.Created).UnmarshalJSON(data))
				}
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson2d00218EncodeForumDBModels1(out *jwriter.Writer, in ThreadShort) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"author\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Author))
	}
	{
		const prefix string = ",\"message\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Message))
	}
	{
		const prefix string = ",\"title\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Title))
	}
	if in.Slug != nil {
		const prefix string = ",\"slug\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Slug))
	}
	{
		const prefix string = ",\"created\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Created == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.Created).MarshalJSON())
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ThreadShort) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2d00218EncodeForumDBModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ThreadShort) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2d00218EncodeForumDBModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ThreadShort) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2d00218DecodeForumDBModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ThreadShort) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2d00218DecodeForumDBModels1(l, v)
}
func easyjson2d00218DecodeForumDBModels2(in *jlexer.Lexer, out *ThreadDetailList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(ThreadDetailList, 0, 1)
			} else {
				*out = ThreadDetailList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 ThreadDetail
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson2d00218EncodeForumDBModels2(out *jwriter.Writer, in ThreadDetailList) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v ThreadDetailList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2d00218EncodeForumDBModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ThreadDetailList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2d00218EncodeForumDBModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ThreadDetailList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2d00218DecodeForumDBModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ThreadDetailList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2d00218DecodeForumDBModels2(l, v)
}
func easyjson2d00218DecodeForumDBModels3(in *jlexer.Lexer, out *ThreadDetail) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = uint64(in.Uint64())
		case "forum":
			out.Forum = string(in.String())
		case "votes":
			out.Votes = int(in.Int())
		case "author":
			out.Author = string(in.String())
		case "message":
			out.Message = string(in.String())
		case "title":
			out.Title = string(in.String())
		case "slug":
			if in.IsNull() {
				in.Skip()
				out.Slug = nil
			} else {
				if out.Slug == nil {
					out.Slug = new(string)
				}
				*out.Slug = string(in.String())
			}
		case "created":
			if in.IsNull() {
				in.Skip()
				out.Created = nil
			} else {
				if out.Created == nil {
					out.Created = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.Created).UnmarshalJSON(data))
				}
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson2d00218EncodeForumDBModels3(out *jwriter.Writer, in ThreadDetail) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.Id))
	}
	{
		const prefix string = ",\"forum\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Forum))
	}
	{
		const prefix string = ",\"votes\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Votes))
	}
	{
		const prefix string = ",\"author\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Author))
	}
	{
		const prefix string = ",\"message\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Message))
	}
	{
		const prefix string = ",\"title\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Title))
	}
	if in.Slug != nil {
		const prefix string = ",\"slug\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Slug))
	}
	{
		const prefix string = ",\"created\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Created == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.Created).MarshalJSON())
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ThreadDetail) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2d00218EncodeForumDBModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ThreadDetail) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2d00218EncodeForumDBModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ThreadDetail) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2d00218DecodeForumDBModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ThreadDetail) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2d00218DecodeForumDBModels3(l, v)
}
