package db_test

import (
	"reflect"
	"testing"

	"gopkg.in/nullbio/null.v6"
	"panjiesw.com/mypict/server/model"
)

func TestDatabase_ImageByID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *model.ImageR
		wantErr bool
	}{
		{
			name: "Only ID: Iwr9ILZkRC",
			args: args{id: "Iwr9ILZkRC"},
			want: &model.ImageR{
				ID:            "Iwr9ILZkRC",
				Title:         null.StringFrom("sapien ut nunc vestibulum"),
				UserID:        null.NewString("", false),
				SID:           null.NewString("", false),
				GalleryTitle:  null.NewString("", false),
				GalleryID:     null.NewString("", false),
				ContentPolicy: 1,
			},
			wantErr: false,
		},
		{
			name: "Both, using ID: 4wrrIYZkgm",
			args: args{id: "4wrrIYZkgm"},
			want: &model.ImageR{
				ID:            "4wrrIYZkgm",
				Title:         null.StringFrom("Nisl Nunc Rhoncus Dui Vel"),
				UserID:        null.StringFrom("KXapjMA"),
				SID:           null.StringFrom("sELljSGzg"),
				GalleryTitle:  null.StringFrom("Nisi Eu Orci"),
				GalleryID:     null.StringFrom("LkQzHLWkgZ"),
				ContentPolicy: 1,
			},
			wantErr: false,
		},
		{
			name: "Both, using SID: yPLlCIGkgz",
			args: args{id: "yPLlCIGkgz"},
			want: &model.ImageR{
				ID:            "VQ99SYZzgZ",
				Title:         null.StringFrom("Mauris Morbi Non Lectus"),
				UserID:        null.StringFrom("UpUAY3Ix"),
				SID:           null.StringFrom("yPLlCIGkgz"),
				GalleryTitle:  null.StringFrom("Vel Nulla Eget Eros Elementum"),
				GalleryID:     null.StringFrom("YkQkHYWkRM"),
				ContentPolicy: 0,
			},
			wantErr: false,
		},
		{
			name:    "Not found",
			args:    args{id: "something"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got model.ImageR
			err := d.ImageByID(tt.args.id, &got)
			got.CreatedAt = ""
			if (err != nil) != tt.wantErr {
				t.Errorf("Database.ImageByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			if !reflect.DeepEqual(&got, tt.want) {
				t.Errorf("Database.ImageByID() = %v, want %v", &got, tt.want)
			}
		})
	}
}

func TestDatabase_ImageBSave(t *testing.T) {
	type args struct {
		imgs []model.ImageS
		uid  null.String
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "3 no uid no error",
			args: args{
				imgs: []model.ImageS{
					{ID: "foobar1", Title: null.StringFrom("foobar1"), ContentPolicy: 0},
					{ID: "foobar2", Title: null.NewString("", false), ContentPolicy: 1},
					{ID: "foobar3", Title: null.StringFrom("foobar3"), ContentPolicy: 1},
				},
				uid: null.NewString("", false),
			},
			wantErr: false,
		},
		{
			name: "3 with uid no error",
			args: args{
				imgs: []model.ImageS{
					{ID: "foobar4", Title: null.StringFrom("foobar4")},
					{ID: "foobar5", Title: null.NewString("", false), ContentPolicy: 1},
					{ID: "foobar6", Title: null.StringFrom("foobar6"), ContentPolicy: 1},
				},
				uid: null.StringFrom("user1"),
			},
			wantErr: false,
		},
		{
			name: "3 no uid error conflict id",
			args: args{
				imgs: []model.ImageS{
					{ID: "foobar4", Title: null.StringFrom("foobar4"), ContentPolicy: 0},
					{ID: "foobar5", Title: null.NewString("", false), ContentPolicy: 1},
					{ID: "foobar6", Title: null.StringFrom("foobar6"), ContentPolicy: 1},
				},
				uid: null.NewString("", false),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := d.ImageBSave(tt.args.imgs, tt.args.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Database.ImageBSave() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
