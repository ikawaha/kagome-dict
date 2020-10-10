package builder

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_parseMatrix(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    *MatrixDef
		wantErr bool
	}{
		{
			name: "square matrix",
			args: args{
				r: strings.NewReader(strings.Join([]string{
					"2 2",
					"0 0 1",
					"0 1 2",
					"1 0 3",
					"1 1 4",
				}, "\n")),
			},
			want: &MatrixDef{
				rowSize: 2,
				colSize: 2,
				vec: []int16{
					1, 3,
					2, 4,
				},
			},
			wantErr: false,
		},
		{
			name: "not square matrix row < col",
			args: args{
				r: strings.NewReader(strings.Join([]string{
					"2 3",
					"0 0 1",
					"0 1 2",
					"0 2 3",
					"1 0 4",
					"1 1 5",
					"1 2 6",
				}, "\n")),
			},
			want: &MatrixDef{
				rowSize: 2,
				colSize: 3,
				vec: []int16{
					1, 4, //  --> row
					2, 5, //  |
					3, 6, //  ↓ col
				},
			},
			wantErr: false,
		},
		{
			name: "not square matrix row > col",
			args: args{
				r: strings.NewReader(strings.Join([]string{
					"3 1",
					"0 0 1",
					"1 0 2",
					"2 0 3",
				}, "\n")),
			},
			want: &MatrixDef{
				rowSize: 3,
				colSize: 1,
				vec: []int16{
					1, 2, 3,
				},
			},
			wantErr: false,
		},
		{
			name: "not square matrix row < col, invalid right-id",
			args: args{
				r: strings.NewReader(strings.Join([]string{
					"2 3",
					"0 0 1",
					"0 1 2",
					"0 2 3",
					"1 0 4",
					"2 1 5",
					"1 2 6",
				}, "\n")),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "not square matrix row < col, invalid left-id",
			args: args{
				r: strings.NewReader(strings.Join([]string{
					"2 3",
					"0 0 1",
					"0 1 2",
					"0 2 3",
					"1 0 4",
					"1 1 5",
					"1 3 6",
				}, "\n")),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseMatrix(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseMatrix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseMatrix() got = %v, want %v", got, tt.want)
			}
		})
	}
}
