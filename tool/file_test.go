//Copyright (c) [2023] [JinCanQi]
//[make_data_set_so-vits-svc] is licensed under Mulan PubL v2.
//You can use this software according to the terms and conditions of the Mulan PubL v2.
//You may obtain a copy of Mulan PubL v2 at:
//         http://license.coscl.org.cn/MulanPubL-2.0
//THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
//EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
//MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
//See the Mulan PubL v2 for more details.

package tool

import (
	"testing"
)

func TestFuzzySearchFiles(t *testing.T) {
	type args struct {
		dir      string
		nameList []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				dir: "../data/source/kugou",
				nameList: []string{
					"不要",
					"靠近",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := FuzzySearchFiles(tt.args.dir, tt.args.nameList)
			if (err != nil) != tt.wantErr {
				t.Errorf("FuzzySearchFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(gotResult)
		})
	}
}

//go:generate mkdir -p ../testA ../testB ../testC
//go:generate touch ../testA/a.txt ../testA/b.txt ../testA/c.txt ../testB/a.txt ../testB/b.txt
func TestPlaceFilesWithTheSameNameAsFolderAAndFolderBInFolderC(t *testing.T) {
	type args struct {
		aDir string
		bDir string
		cDir string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				aDir: "../testA",
				bDir: "../testB",
				cDir: "../testC",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PlaceFilesWithTheSameNameAsFolderAAndFolderBInFolderC(tt.args.aDir, tt.args.bDir, tt.args.cDir); (err != nil) != tt.wantErr {
				t.Errorf("PlaceFilesWithTheSameNameAsFolderAAndFolderBInFolderC() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRecoverUltimatevocalremoverFileNameVocals(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{
			args: args{
				fileName: "1_1-13-2023-04-18-再见拜拜_(Vocals).mp3",
			},
			wantResult: "1-13-2023-04-18-再见拜拜.mp3",
		},
		{
			args: args{
				fileName: "4_4-13-2023-01-17-独处_(Instrumental).mp3",
			},
			wantResult: "4-13-2023-01-17-独处.mp3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := RecoverUltimatevocalremoverFileName(tt.args.fileName); gotResult != tt.wantResult {
				t.Errorf("RecoverUltimatevocalremoverFileNameVocals() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
