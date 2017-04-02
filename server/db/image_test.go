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
		want    *model.ImageDTO
		wantErr bool
	}{
		{
			name: "Only ID: Iwr9ILZkRC",
			args: args{id: "Iwr9ILZkRC"},
			want: &model.ImageDTO{
				Image: &model.Image{
					ID:            "Iwr9ILZkRC",
					ContentPolicy: 1,
					Title:         null.StringFrom("sapien ut nunc vestibulum"),
					UserID:        null.NewString("", false),
				},
				SID:    "",
				GTitle: "",
				GID:    "",
			},
			wantErr: false,
		},
		{
			name: "Both, using ID: 4wrrIYZkgm",
			args: args{id: "4wrrIYZkgm"},
			want: &model.ImageDTO{
				Image: &model.Image{
					ID:            "4wrrIYZkgm",
					ContentPolicy: 1,
					Title:         null.StringFrom("Nisl Nunc Rhoncus Dui Vel"),
					UserID:        null.StringFrom("KXapjMA"),
				},
				SID:    "sELljSGzg",
				GTitle: "Nisi Eu Orci",
				GID:    "LkQzHLWkgZ",
			},
			wantErr: false,
		},
		{
			name: "Both, using SID: yPLlCIGkgz",
			args: args{id: "yPLlCIGkgz"},
			want: &model.ImageDTO{
				Image: &model.Image{
					ID:            "VQ99SYZzgZ",
					ContentPolicy: 0,
					Title:         null.StringFrom("Mauris Morbi Non Lectus"),
					UserID:        null.StringFrom("UpUAY3Ix"),
				},
				SID:    "yPLlCIGkgz",
				GTitle: "Vel Nulla Eget Eros Elementum",
				GID:    "YkQkHYWkRM",
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
			var got model.ImageDTO
			err := d.ImageByID(tt.args.id, &got)
			if (err != nil) != tt.wantErr {
				t.Errorf("Database.ImageByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			got.Image.CreatedAt = nil
			if !reflect.DeepEqual(&got, tt.want) {
				t.Errorf("Database.ImageByID() = \n%v, want \n%v", &got, tt.want)
			}
		})
	}
}

func TestDatabase_ImageBSave(t *testing.T) {
	type args struct {
		imgs []*model.ImageDTO
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "3 no uid no error",
			args: args{
				imgs: []*model.ImageDTO{
					{
						Image: &model.Image{
							ID:            "foobar1",
							ContentPolicy: 0,
							Title:         null.StringFrom("foobar1"),
						},
					},
					{
						Image: &model.Image{
							ID:            "foobar2",
							ContentPolicy: 1,
							Title:         null.NewString("", false),
						},
					},
					{
						Image: &model.Image{
							ID:            "foobar3",
							ContentPolicy: 1,
							Title:         null.StringFrom("foobar3"),
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "3 with uid no error",
			args: args{
				imgs: []*model.ImageDTO{
					{
						Image: &model.Image{
							ID:     "foobar4",
							Title:  null.StringFrom("foobar4"),
							UserID: null.StringFrom("user1"),
						},
					},
					{
						Image: &model.Image{
							ID:            "foobar5",
							Title:         null.NewString("", false),
							ContentPolicy: 1,
							UserID:        null.StringFrom("user1"),
						},
					},
					{
						Image: &model.Image{
							ID:            "foobar6",
							Title:         null.StringFrom("foobar6"),
							ContentPolicy: 1,
							UserID:        null.StringFrom("user1"),
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := d.ImageBSave(tt.args.imgs)
			if (err != nil) != tt.wantErr {
				t.Errorf("Database.ImageBSave() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			for _, img := range tt.args.imgs {
				if img.Image.ID == "" {
					t.Errorf("Database.ImageBSave() id not generated = %s", img.Image.Title)
				}
			}
		})
	}
}
