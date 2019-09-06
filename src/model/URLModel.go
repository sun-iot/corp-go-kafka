package model

/**
 * Copyright (c) 2018-2028 Corp-ci All Rights Reserved
 *
 * Project:
 * Package: model
 * Version: 1.0
 *
 * Created by Smile-Sun on 2019/9/4 9:19
 */

type URLModel struct {
	appkey   string
	channel  string
	dataware string
	table    string
	data     string
}

func NesURLModel(appkey, channel, dataware, table, data string) *URLModel {
	return &URLModel{
		appkey,
		channel,
		dataware,
		table,
		"",
	}
}

func SetDataURLModel(urlModel *URLModel, data string) *URLModel {
	urlModel.data = data
	return urlModel
}

func GetDataURLModel(urlModel *URLModel) {

}
