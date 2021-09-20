package gogetset

import (
	"reflect"
	"testing"
)

func TestSet(t *testing.T) {

	type StructSample struct {
		Hello string
		Arr   []interface{}
		Map   map[string]interface{}
	}

	type embeddingStructSample struct {
		Embed    StructSample
		EmbedMap map[string]string
	}

	level1Ptr := &StructSample{Hello: "worldPtr"}
	level2Ptr := &StructSample{Hello: "world2Ptr"}

	inputMap := map[string]interface{}{
		"LevelOneIntSlice": []int{1, 2, 3},
		"level1Int":        1,
		"level1String":     "string1",
		"level1Map": map[string]interface{}{
			"level2Int":    2,
			"level2String": "string2",
			"level2Struct": StructSample{Hello: "world2"},
			"level2Ptr":    level2Ptr,
			"level2EmbeddingStruct": embeddingStructSample{
				Embed: StructSample{Hello: "embeddingWorld"},
				EmbedMap: map[string]string{
					"key": "value",
				},
			},
			"levelTwoIntSlice": []int{4, 5, 6},
		},
		"level1Map1": map[string]interface{}{
			"level2Map": map[string]interface{}{
				"level3Int":    3,
				"level3String": "string3",
				"level3Struct": StructSample{Hello: "world3"},
			},
			"level2Func": sampleFunc,
		},
		"level1Struct": StructSample{Hello: "world"},
		"level1Ptr":    level1Ptr,
		"level1Func":   sampleFunc,
		"level1EmbeddingStruct": embeddingStructSample{
			Embed: StructSample{Hello: "embeddingWorld1"},
			EmbedMap: map[string]string{
				"key": "value",
			},
		},
	}

	type args struct {
		path  string
		input interface{}
		val   interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    interface{}
	}{
		{
			name: "level1Int",
			args: args{
				path:  "level1Int",
				input: inputMap,
				val:   2,
			},
			wantErr: false,
			want:    2,
		},
		{
			name: "setFieldInPtrToStruct",
			args: args{
				path:  "Hello",
				input: &StructSample{},
				val:   "World",
			},
			wantErr: false,
			want:    "World",
		},
		{
			name: "setSliceElementInPtrToStruct",
			args: args{
				path:  "Arr[0]",
				input: &StructSample{},
				val:   "EleWorld",
			},
			wantErr: false,
			want:    "EleWorld",
		},
		{
			name: "setInterfaceSliceElementInPtrToStruct",
			args: args{
				path:  "Arr[1]",
				input: &StructSample{},
				val: map[string]interface{}{
					"arr_one": "val_one",
				},
			},
			wantErr: false,
			want: map[string]interface{}{
				"arr_one": "val_one",
			},
		},
		{
			name: "setInterfaceSliceElementInPtrToStruct",
			args: args{
				path:  "Arr[1].arr_one",
				input: &StructSample{},
				val:   "val_one",
			},
			wantErr: false,
			want:    "val_one",
		},
		{
			name: "setSliceElementInMapError",
			args: args{
				path:  "lOne.lTwo[0]",
				input: nil,
				val:   "MapEleWorld",
			},
			wantErr: true,
			want:    "MapEleWorld",
		},
		{
			name: "setSliceElementInMap",
			args: args{
				path:  "lOne.lTwo[0]",
				input: make(map[string]interface{}),
				val:   "MapEleWorld",
			},
			wantErr: false,
			want:    "MapEleWorld",
		},
		{
			name: "setMapElementInMap",
			args: args{
				path:  "Map.Key",
				input: level1Ptr,
				val:   "MapEleWorld",
			},
			wantErr: false,
			want:    "MapEleWorld",
		},
	}
	var err error
	var got interface{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err = Set(tt.args.path, tt.args.input, tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
			got, err = Get(tt.args.path, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() \nerror = %v, \nwantErr = %v\n", err, tt.wantErr)
			}
			if err != nil {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() \ngot = %v, \nwant = %v\n", got, tt.want)
			}
		})
	}
}
