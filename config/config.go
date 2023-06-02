//Copyright (c) [2023] [JinCanQi]
//[make_data_set_so-vits-svc] is licensed under Mulan PubL v2.
//You can use this software according to the terms and conditions of the Mulan PubL v2.
//You may obtain a copy of Mulan PubL v2 at:
//         http://license.coscl.org.cn/MulanPubL-2.0
//THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
//EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
//MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
//See the Mulan PubL v2 for more details.

package config

import (
	"errors"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	AuthorName         string   `yaml:"authorName"`         // 作者名字
	KouGouJsonPathList []string `yaml:"kouGouJsonPathList"` // 酷狗json文件路径
	AuthorBirthday     string   `yaml:"authorBirthday"`     // 作者生日
}

var ConfigModel = &Config{}

func init() {
	yamlFile, err := os.ReadFile("./config.yaml")
	if errors.Is(err, os.ErrNotExist) {
		out, err := yaml.Marshal(ConfigModel)
		if err != nil {
			panic(err)
		}
		err = os.WriteFile("./config.yaml", out, os.ModePerm)
		if err != nil {
			panic(err)
		}
		log.Println("config.yaml文件不存在，已经创建，请填写配置")
		return
	}
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &ConfigModel)
	if err != nil {
		panic(err)
	}
}
