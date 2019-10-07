package ui

import (
	"game_fly/core/data"
	"game_fly/core/sprite"
)

func AddUi(ui data.UiType, groupName string) {
	flg := false
	for k, v := range sprite.Sprites {
		if v.GroupName == groupName {
			flg = true
			sprite.Sprites[k].Ui = append(v.Ui, ui)
		}
	}

	if flg == false {
		newUi := []data.UiType{ui}
		sprite.Sprites = append(sprite.Sprites, data.SpriteGroup{GroupName: groupName, Ui: newUi})
	}
	ui.OnLoad()
	ui.Start()
}

func GetUi(groupName, name string) (ui data.UiType) {
	for _, v := range sprite.Sprites {
		if v.GroupName == groupName {
			for _, v2 := range v.Ui {
				if v2.GetName() == name {
					return v2
				}
			}
		}
	}
	return nil
}

func ReplaceUi(groupName, name string, newUi data.UiType) {
	if GetUi(groupName, name) == nil {
		AddUi(newUi, groupName)
	}
	for k, v := range sprite.Sprites {
		if v.GroupName == groupName {
			for k2, v2 := range v.Ui {
				if v2.GetName() == name {
					sprite.Sprites[k].Ui[k2] = newUi
				}
			}
		}
	}
}

func GetUiAll(groupName string) (ui []data.UiType) {
	for _, v := range sprite.Sprites {
		if v.GroupName == groupName {
			return v.Ui
		}
	}
	return nil
}
