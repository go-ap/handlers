package handlers

import (
	"testing"

	"github.com/go-ap/activitypub"
	pub "github.com/go-ap/activitypub"
)

func TestPathTyper_Type(t *testing.T) {
	t.Skipf("TODO")
}

func TestValidActivityCollection(t *testing.T) {
	t.Skipf("TODO")
}

func TestValidCollection(t *testing.T) {
	t.Skipf("TODO")
}

func TestValidObjectCollection(t *testing.T) {
	t.Skipf("TODO")
}

func TestValidCollectionIRI(t *testing.T) {
	t.Skipf("TODO")
}

func TestSplit(t *testing.T) {
	t.Skipf("TODO")
}

func TestCollectionTypes_Of(t *testing.T) {
	type args struct {
		o activitypub.Item
		t CollectionType
	}
	tests := []struct {
		name string
		args args
		want activitypub.Item
	}{
		{
			name: "nil from nil object",
			args: args{
				o: nil,
				t: "likes",
			},
			want: nil,
		},
		{
			name: "nil from invalid collection type",
			args: args{
				o: activitypub.Object{
					Likes: activitypub.IRI("test"),
				},
				t: "like",
			},
			want: nil,
		},
		{
			name: "nil from nil collection type",
			args: args{
				o: activitypub.Object{
					Likes: nil,
				},
				t: "likes",
			},
			want: nil,
		},
		{
			name: "get likes iri",
			args: args{
				o: activitypub.Object{
					Likes: activitypub.IRI("test"),
				},
				t: "likes",
			},
			want: activitypub.IRI("test"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if ob := test.args.t.Of(test.args.o); ob != test.want {
				t.Errorf("Object received %#v is different, expected #%v", ob, test.want)
			}
		})
	}
}

func TestCollectionType_IRI(t *testing.T) {
	type args struct {
		o activitypub.Item
		t CollectionType
	}
	tests := []struct {
		name string
		args args
		want activitypub.IRI
	}{
		{
			name: "just path from nil object",
			args: args{
				o: nil,
				t: "likes",
			},
			want: pub.IRI("/likes"),
		},
		{
			name: "emptyIRI from invalid collection type",
			args: args{
				o: activitypub.Object{
					Likes: activitypub.IRI("test"),
				},
				t: "like",
			},
			want: "/like",
		},
		{
			name: "just path from object without ID",
			args: args{
				o: activitypub.Object{},
				t: "likes",
			},
			want: pub.IRI("/likes"),
		},
		{
			name: "likes iri on object",
			args: args{
				o: activitypub.Object{
					ID:    "http://example.com",
					Likes: activitypub.IRI("test"),
				},
				t: "likes",
			},
			want: activitypub.IRI("test"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if ob := test.args.t.IRI(test.args.o); ob != test.want {
				t.Errorf("IRI received %q is different, expected %q", ob, test.want)
			}
		})
	}
}

func TestCollectionType_OfActor(t *testing.T) {
	t.Skipf("TODO")
}

func TestCollectionTypes_Contains(t *testing.T) {
	t.Skipf("TODO")
}

func TestIRIf(t *testing.T) {
	type args struct {
		i activitypub.IRI
		t CollectionType
	}
	tests := []struct {
		name string
		args args
		want activitypub.IRI
	}{
		{
			name: "empty iri",
			args: args{
				i: "",
				t: "inbox",
			},
			want: "/inbox",
		},
		{
			name: "plain concat",
			args: args{
				i: "https://example.com",
				t: "inbox",
			},
			want: "https://example.com/inbox",
		},
		{
			name: "strip root from iri",
			args: args{
				i: "https://example.com/",
				t: "inbox",
			},
			want: "https://example.com/inbox",
		},
		{
			name: "invalid iri",
			args: args{
				i: "example.com",
				t: "test",
			},
			want: "example.com/test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IRIf(tt.args.i, tt.args.t); got != tt.want {
				t.Errorf("IRIf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollectionType_AddTo(t *testing.T) {
	type args struct {
		i activitypub.Item
	}
	var i activitypub.Item
	var o *activitypub.Object
	tests := []struct {
		name  string
		t     CollectionType
		args  args
		want  activitypub.IRI
		want1 bool
	}{
		{
			name: "simple",
			t:    "test",
			args: args{
				i: &activitypub.Object{ID: "http://example.com/addTo"},
			},
			want:  "http://example.com/addTo/test",
			want1: false, // this seems to always be false
		},
		{
			name: "on-nil-item",
			t:    "test",
			args: args{
				i: i,
			},
			want:  activitypub.NilIRI,
			want1: false,
		},
		{
			name: "on-nil",
			t:    "test",
			args: args{
				i: nil,
			},
			want:  activitypub.NilIRI,
			want1: false,
		},
		{
			name: "on-nil-object",
			t:    "test",
			args: args{
				i: o,
			},
			want:  activitypub.NilIRI,
			want1: false,
		},
		{
			name: "on-nil-item",
			t:    "test",
			args: args{
				i: i,
			},
			want:  activitypub.NilIRI,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.t.AddTo(tt.args.i)
			if got != tt.want {
				t.Errorf("AddTo() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("AddTo() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCollectionTypes_Split(t *testing.T) {
	t.Skipf("TODO")
}
