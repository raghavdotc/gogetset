package gogetset

import (
	"context"
	"reflect"
	"testing"
)

var sampleFunc = func(ctx context.Context, args ...string) {}

func TestGet(t *testing.T) {
	type structSample struct {
		Hello string
	}

	type embeddingStructSample struct {
		Embed    structSample
		EmbedMap map[string]string
	}

	level1Ptr := &structSample{Hello: "worldPtr"}
	level2Ptr := &structSample{Hello: "world2Ptr"}

	inputMap := map[string]interface{}{
		"LevelOneIntSlice": []int{1, 2, 3},
		"level1Int":        1,
		"level1String":     "string1",
		"level1Map": map[string]interface{}{
			"level2Int":    2,
			"level2String": "string2",
			"level2Struct": structSample{Hello: "world2"},
			"level2Ptr":    level2Ptr,
			"level2EmbeddingStruct": embeddingStructSample{
				Embed: structSample{Hello: "embeddingWorld"},
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
				"level3Struct": structSample{Hello: "world3"},
			},
			"level2Func": sampleFunc,
		},
		"level1Struct": structSample{Hello: "world"},
		"level1Ptr":    level1Ptr,
		"level1Func":   sampleFunc,
		"level1EmbeddingStruct": embeddingStructSample{
			Embed: structSample{Hello: "embeddingWorld1"},
			EmbedMap: map[string]string{
				"key": "value",
			},
		},
	}
	type args struct {
		path     string
		inputMap map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		hasFunc bool
		wantErr bool
	}{
		{
			name: "level1Int",
			args: args{
				path:     "level1Int",
				inputMap: inputMap,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "level1String",
			args: args{
				path:     "level1String",
				inputMap: inputMap,
			},
			want:    "string1",
			wantErr: false,
		},
		{
			name: "level1Struct",
			args: args{
				path:     "level1Struct",
				inputMap: inputMap,
			},
			want:    structSample{Hello: "world"},
			wantErr: false,
		},
		{
			name: "level1EmbeddingStruct",
			args: args{
				path:     "level1EmbeddingStruct",
				inputMap: inputMap,
			},
			want: embeddingStructSample{
				Embed: structSample{Hello: "embeddingWorld1"},
				EmbedMap: map[string]string{
					"key": "value",
				},
			},
			wantErr: false,
		},
		{
			name: "level1Struct.Hello",
			args: args{
				path:     "level1Struct.Hello",
				inputMap: inputMap,
			},
			want:    "world",
			wantErr: false,
		},
		{
			name: "level1Ptr.Hello",
			args: args{
				path:     "level1Ptr.Hello",
				inputMap: inputMap,
			},
			want:    "worldPtr",
			wantErr: false,
		},
		{
			name: "level1Map",
			args: args{
				path:     "level1Map",
				inputMap: inputMap,
			},
			want: map[string]interface{}{
				"level2Int":    2,
				"level2String": "string2",
				"level2Struct": structSample{Hello: "world2"},
				"level2Ptr":    level2Ptr,
				"level2EmbeddingStruct": embeddingStructSample{
					Embed: structSample{Hello: "embeddingWorld"},
					EmbedMap: map[string]string{
						"key": "value",
					},
				},
				"levelTwoIntSlice": []int{4, 5, 6},
			},
			wantErr: false,
		},
		{
			name: "level1Ptr",
			args: args{
				path:     "level1Ptr",
				inputMap: inputMap,
			},
			want:    level1Ptr,
			wantErr: false,
		},
		{
			name: "level1Func",
			args: args{
				path:     "level1Func",
				inputMap: inputMap,
			},
			hasFunc: true,
			want:    sampleFunc,
			wantErr: false,
		},
		{
			name: "level1Error",
			args: args{
				path:     "level1Error",
				inputMap: inputMap,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "level1Map.level2Int",
			args: args{
				path:     "level1Map.level2Int",
				inputMap: inputMap,
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "level1Map.level2String",
			args: args{
				path:     "level1Map.level2String",
				inputMap: inputMap,
			},
			want:    "string2",
			wantErr: false,
		},
		{
			name: "level1Map.level2Struct",
			args: args{
				path:     "level1Map.level2Struct",
				inputMap: inputMap,
			},
			want:    structSample{Hello: "world2"},
			wantErr: false,
		},
		{
			name: "level1Map.level2Struct.Hello",
			args: args{
				path:     "level1Map.level2Struct.Hello",
				inputMap: inputMap,
			},
			want:    "world2",
			wantErr: false,
		},
		{
			name: "level1Map.level2EmbeddingStruct",
			args: args{
				path:     "level1Map.level2EmbeddingStruct",
				inputMap: inputMap,
			},
			want: embeddingStructSample{
				Embed: structSample{Hello: "embeddingWorld"},
				EmbedMap: map[string]string{
					"key": "value",
				},
			},
			wantErr: false,
		},
		{
			name: "level1Map.level2EmbeddingStruct.Embed",
			args: args{
				path:     "level1Map.level2EmbeddingStruct.Embed",
				inputMap: inputMap,
			},
			want:    structSample{Hello: "embeddingWorld"},
			wantErr: false,
		},
		{
			name: "level1Map.level2EmbeddingStruct.EmbedMap",
			args: args{
				path:     "level1Map.level2EmbeddingStruct.EmbedMap",
				inputMap: inputMap,
			},
			want: map[string]string{
				"key": "value",
			},
			wantErr: false,
		},
		{
			name: "level1Map.level2Ptr.Hello",
			args: args{
				path:     "level1Map.level2Ptr.Hello",
				inputMap: inputMap,
			},
			want:    "world2Ptr",
			wantErr: false,
		},
		{
			name: "level1Map1.level2Map",
			args: args{
				path:     "level1Map1.level2Map",
				inputMap: inputMap,
			},
			want: map[string]interface{}{
				"level3Int":    3,
				"level3String": "string3",
				"level3Struct": structSample{Hello: "world3"},
			},
			wantErr: false,
		},
		{
			name: "level1Map.level2Ptr",
			args: args{
				path:     "level1Map.level2Ptr",
				inputMap: inputMap,
			},
			want:    &structSample{Hello: "world2Ptr"},
			wantErr: false,
		},
		{
			name: "level1Map1.level2Func",
			args: args{
				path:     "level1Map1.level2Func",
				inputMap: inputMap,
			},
			hasFunc: true,
			want:    sampleFunc,
			wantErr: false,
		},
		{
			name: "level1Map.level2Error",
			args: args{
				path:     "level1Map.level2Error",
				inputMap: inputMap,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "LevelOneIntSlice[0]",
			args: args{
				path:     "LevelOneIntSlice[0]",
				inputMap: inputMap,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "LevelOneIntSlice[1]",
			args: args{
				path:     "LevelOneIntSlice[1]",
				inputMap: inputMap,
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "level1Map.levelTwoIntSlice[0]",
			args: args{
				path:     "level1Map.levelTwoIntSlice[0]",
				inputMap: inputMap,
			},
			want:    4,
			wantErr: false,
		},
		{
			name: "level1Map.levelTwoIntSlice[1]",
			args: args{
				path:     "level1Map.levelTwoIntSlice[1]",
				inputMap: inputMap,
			},
			want:    5,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(tt.args.path, tt.args.inputMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() \nerror = %v, \nwantErr = %v\n", err, tt.wantErr)
				return
			}
			if !tt.hasFunc && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() \ngot = %v, \nwant = %v\n", got, tt.want)
			}
		})
	}
}
