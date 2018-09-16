package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
)

type VersionInfo struct {
	Id               int64
	Version          string `orm:size(100)`
	ArtifactId       int64
	UploadFileInfoId int64
}

type PathAndVersion struct {
	PATHFILEID int64
	IDTYPE     int64
	PATHNAME   string
}

func init() {
	orm.RegisterModel(new(VersionInfo))
}

func AddVersionInfo(version string, artifactId int64, uploadFileInfoId int64) (int64, error) {
	o := orm.NewOrm()
	vo, _ := FindVersionInfo(version, artifactId, uploadFileInfoId)
	if vo == nil {
		vo = &VersionInfo{
			Version:          version,
			ArtifactId:       artifactId,
			UploadFileInfoId: uploadFileInfoId,
		}
		id, err := o.Insert(vo)
		return id, err
	} else {
		return vo.Id, nil
	}
}

func FindVersionInfo(version string, artifactId int64, uploadFileInfoId int64) (*VersionInfo, error) {
	o := orm.NewOrm()
	versionInfo := make([]*VersionInfo, 0)
	_, err := o.QueryTable("version_info").Filter("version", version).Filter("artifact_id", artifactId).Filter("upload_file_info_id", uploadFileInfoId).Limit(1).All(&versionInfo)
	if len(versionInfo) > 0 {
		return versionInfo[0], err
	} else {
		return nil, err
	}
}

func FindVersonInfos(pathFileId int64) ([]*VersionInfo, error) {
	o := orm.NewOrm()
	versionInfo := make([]*VersionInfo, 0)
	_, err := o.QueryTable("version_info").Filter("artifact_id", pathFileId).All(&versionInfo)
	return versionInfo, err
}

func FindVersionInfoById(pathFileId int64) (VersionInfo, error) {
	o := orm.NewOrm()
	version := VersionInfo{Id: pathFileId}
	err := o.Read(&version)
	if err == orm.ErrNoRows {
		return VersionInfo{}, errors.New("查询不到相关版本号")
	} else if err == orm.ErrMissPK {
		return VersionInfo{}, errors.New("查询不到相关版本号主键")
	} else if err != nil {
		return VersionInfo{}, err
	} else {
		return version, nil
	}
}
