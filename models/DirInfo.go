package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"strings"
)

type PathInfo struct {
	Id           int64
	PathName     string `orm:size(1000)`
	PathType     int64  //1、groupId类型，2、artifactId类型
	ParentPathId int64
}

func init() {
	orm.RegisterModel(new(PathInfo))
}

func AddPathInfo(groupId string, artifactId string) (int64, error) {
	if len(groupId) < 1 || len(artifactId) < 1 {
		return int64(0), errors.New("groupId或artifactId不能空！")
	} else {

		//保存groupId
		var paths = strings.Split(groupId, ".")
		o := orm.NewOrm()
		parentId := int64(0)
		for i := 0; i < len(paths); i++ {
			vo, _ := FindPathInfo(paths[i], parentId, int64(1))
			if vo == nil {
				vo = &PathInfo{
					PathName:     paths[i],
					PathType:     int64(1),
					ParentPathId: parentId,
				}
				id, err := o.Insert(vo)
				if err == nil {
					parentId = id
				} else {
					return int64(0), err
				}
			} else {
				parentId = vo.Id
			}
		}

		//保存artifactId
		paths = strings.Split(artifactId, ".")

		for i := 0; i < len(paths); i++ {
			vo, _ := FindPathInfo(paths[i], parentId, int64(2))
			if vo == nil {
				vo = &PathInfo{
					PathName:     paths[i],
					PathType:     int64(2),
					ParentPathId: parentId,
				}
				id, err := o.Insert(vo)
				if err == nil {
					parentId = id
				} else {
					return int64(0), err
				}
			} else {
				parentId = vo.Id
			}
		}
		return parentId, nil
	}
}

func FindPathInfo(pathName string, parentId int64, pathType int64) (*PathInfo, error) {
	o := orm.NewOrm()
	pathInfo := make([]*PathInfo, 0)
	_, err := o.QueryTable("path_info").Filter("path_name", pathName).Filter("parent_path_id", parentId).Filter("path_type", pathType).Limit(1).All(&pathInfo)
	if len(pathInfo) > 0 {
		return pathInfo[0], err
	} else {
		return nil, err
	}
}
